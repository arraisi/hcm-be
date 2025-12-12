package asynqworker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	webhookdto "github.com/arraisi/hcm-be/internal/domain/dto/webhook"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

type DIDXSvc interface {
	Confirm(ctx context.Context, body any) error
	ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error
	ConfirmAppraisal(ctx context.Context, request appraisal.AppraisalConfirmEvent) error
}

type DMSSvc interface {
	TestDriveRequest(ctx context.Context, body any) error
	CreateOneAccess(ctx context.Context, body any) error
	CreateToyotaID(ctx context.Context, body any) error
	AppraisalBookingRequest(ctx context.Context, body any) error
	GetOfferRequest(ctx context.Context, body any) error
	CallbackTamResponse(ctx context.Context, body any) error
}

type DMSAfterSalesSvc interface {
	ServiceBookingRequest(ctx context.Context, body any) error
}

// Worker handles Asynq task processing
type Worker struct {
	srv              *asynq.Server
	mux              *asynq.ServeMux
	didxSvc          DIDXSvc
	dmsSvc           DMSSvc
	dmsAfterSalesSvc DMSAfterSalesSvc
}

// New creates a new Asynq worker instance
func New(cfg *config.Config, didxSvc DIDXSvc, dmsSvc DMSSvc, dmsAfterSalesSvc DMSAfterSalesSvc) *Worker {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.Asynq.RedisAddr,
			DB:       cfg.Asynq.RedisDB,
			Password: cfg.Asynq.RedisPassword,
		},
		asynq.Config{
			Concurrency: cfg.Asynq.Concurrency,
			Queues: map[string]int{
				cfg.Asynq.Queue: 1,
			},
			// Custom retry delay function: 1m, 5m, 10m
			// n is the number of times the task has been retried (0-indexed for next retry)
			RetryDelayFunc: func(n int, err error, t *asynq.Task) time.Duration {
				switch n {
				case 0:
					return time.Minute // 1st retry after 1 minute
				case 1:
					return 5 * time.Minute // 2nd retry after 5 minutes
				case 2:
					return 10 * time.Minute // 3rd retry after 10 minutes
				default:
					return 10 * time.Minute
				}
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				retried, _ := asynq.GetRetryCount(ctx)
				maxRetry, _ := asynq.GetMaxRetry(ctx)
				if retried >= maxRetry {
					log.Printf("[ERROR] Task %s failed after %d retries: %v", task.Type(), retried, err)
				}
			}),
		},
	)

	mux := asynq.NewServeMux()

	w := &Worker{
		srv:              srv,
		mux:              mux,
		didxSvc:          didxSvc,
		dmsSvc:           dmsSvc,
		dmsAfterSalesSvc: dmsAfterSalesSvc,
	}

	// Register task handlers
	mux.HandleFunc(queue.TaskTypeDIDXServiceBookingConfirm, w.handleServiceBookingConfirm)
	mux.HandleFunc(queue.TaskTypeDIDXTestDriveConfirm, w.handleTestDriveConfirm)
	mux.HandleFunc(queue.TaskTypeDIDXAppraisalConfirm, w.handleAppraisalConfirm)
	mux.HandleFunc(queue.TaskTypeDMSTestDriveRequest, w.handleDMSTestDriveRequest)
	mux.HandleFunc(queue.TaskTypeDMSServiceBookingRequest, w.handleDMSServiceBookingRequest)
	mux.HandleFunc(queue.TaskTypeDMSCreateOneAccess, w.handleDMSCreateOneAccess)
	mux.HandleFunc(queue.TaskTypeDMSCreateToyotaID, w.handleDMSCreateToyotaID)
	mux.HandleFunc(queue.TaskTypeDMSAppraisalBookingRequest, w.handleDMSAppraisalBookingRequest)
	mux.HandleFunc(queue.TaskTypeDMSCreateGetOffer, w.handleDMSGetOfferRequest)

	return w
}

// Run starts the Asynq worker server
func (w *Worker) Run() error {
	log.Printf("Starting Asynq worker...")
	return w.srv.Run(w.mux)
}

// Shutdown gracefully shuts down the worker
func (w *Worker) Shutdown() {
	log.Printf("Shutting down Asynq worker...")
	w.srv.Shutdown()
}

// sendDMSCallback sends callback response to DMS TAM
func (w *Worker) sendDMSCallback(ctx context.Context, taskID, eventID, documentID string, err error) {
	var callback webhookdto.CallbackTamResponse

	if err != nil {
		callback = webhookdto.NewFailureCallback(eventID, documentID, err)
	} else {
		callback = webhookdto.NewSuccessCallback(eventID, documentID, "Test drive confirmed successfully")
	}

	// Convert struct to map[string]interface{} using JSON marshal
	var callbackPayload map[string]interface{}
	jsonData, _ := json.Marshal(callback)
	if err := json.Unmarshal(jsonData, &callbackPayload); err != nil {
		log.Printf("[ERROR] TaskID: %s Failed to convert callback to map for DocumentID: %s - Error: %v",
			taskID,
			documentID,
			err,
		)
		return
	}

	if callbackErr := w.dmsSvc.CallbackTamResponse(ctx, callbackPayload); callbackErr != nil {
		log.Printf("[ERROR] TaskID: %s Failed to send %s callback to DMS for DocumentID: %s - Error: %v",
			taskID,
			callback.Status,
			documentID,
			callbackErr,
		)
	}

	log.Printf("[INFO] TaskID: %s Sent %s callback to DMS for DocumentID: %s",
		taskID,
		callback.Status,
		documentID,
	)
}
