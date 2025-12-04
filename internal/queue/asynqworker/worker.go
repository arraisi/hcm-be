package asynqworker

import (
	"context"
	"log"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

type DIDXSvc interface {
	Confirm(ctx context.Context, body any) error
	ConfirmTestDrive(ctx context.Context, request testdrive.TestDriveEvent) error
}

type DMSSvc interface {
	TestDriveRequest(ctx context.Context, body any) error
	CreateOneAccess(ctx context.Context, body any) error
	CreateToyotaID(ctx context.Context, body any) error
	AppraisalBookingRequest(ctx context.Context, body any) error
}

// Worker handles Asynq task processing
type Worker struct {
	srv     *asynq.Server
	mux     *asynq.ServeMux
	didxSvc DIDXSvc
	dmsSvc  DMSSvc
}

// New creates a new Asynq worker instance
func New(cfg config.AsynqConfig, didxSvc DIDXSvc, dmsSvc DMSSvc) *Worker {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     cfg.RedisAddr,
			DB:       cfg.RedisDB,
			Password: cfg.RedisPassword,
		},
		asynq.Config{
			Concurrency: cfg.Concurrency,
			Queues: map[string]int{
				cfg.Queue: 1,
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
		srv:     srv,
		mux:     mux,
		didxSvc: didxSvc,
		dmsSvc:  dmsSvc,
	}

	// Register task handlers
	mux.HandleFunc(queue.TaskTypeDIDXConfirm, w.handleServiceBookingConfirm)
	mux.HandleFunc(queue.TaskTypeDMSTestDriveRequest, w.handleDMSTestDriveRequest)
	mux.HandleFunc(queue.TaskTypeDMSCreateOneAccess, w.handleDMSCreateOneAccess)
	mux.HandleFunc(queue.TaskTypeDMSCreateToyotaID, w.handleDMSCreateToyotaID)
	mux.HandleFunc(queue.TaskTypeDMSAppraisalBookingRequest, w.handleDMSAppraisalBookingRequest)

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
