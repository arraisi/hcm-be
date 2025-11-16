package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/arraisi/hcm-be/internal/config"
	"github.com/arraisi/hcm-be/internal/external/didx"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// Worker handles Asynq task processing
type Worker struct {
	srv        *asynq.Server
	mux        *asynq.ServeMux
	didxClient didx.ClientInterface
}

// New creates a new Asynq worker instance
func New(cfg config.AsynqConfig, didxClient didx.ClientInterface) *Worker {
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
			RetryDelayFunc: func(n int, err error, t *asynq.Task) time.Duration {
				switch n {
				case 1:
					return time.Minute
				case 2:
					return 5 * time.Minute
				case 3:
					return 10 * time.Minute
				default:
					return 0
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
		srv:        srv,
		mux:        mux,
		didxClient: didxClient,
	}

	// Register task handlers
	mux.HandleFunc(queue.TaskTypeDIDXConfirm, w.handleDIDXConfirm)

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

// handleDIDXConfirm processes DIDX confirm tasks
func (w *Worker) handleDIDXConfirm(ctx context.Context, t *asynq.Task) error {
	var payload queue.DIDXConfirmPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DIDXConfirmPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] Processing DIDX Confirm task (attempt %d) for ServiceBookingID: %s, EventID: %s",
		retried+1,
		payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
		payload.ServiceBookingEvent.EventID,
	)

	// Call the DIDX external API
	err := w.didxClient.Confirm(ctx, payload.ServiceBookingEvent)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] DIDX Confirm failed (attempt %d) for ServiceBookingID: %s - Error: %v",
			retried+1,
			payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] DIDX Confirm succeeded for ServiceBookingID: %s, EventID: %s",
		payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
		payload.ServiceBookingEvent.EventID,
	)

	return nil
}
