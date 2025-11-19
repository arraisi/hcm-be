package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleServiceBookingConfirm processes DIDX confirm tasks
func (w *Worker) handleServiceBookingConfirm(ctx context.Context, t *asynq.Task) error {
	var payload queue.DIDXServiceBookingConfirmPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DIDXServiceBookingConfirmPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DIDX Confirm task (attempt %d) for ServiceBookingID: %s, EventID: %s",
		taskID,
		retried+1,
		payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
		payload.ServiceBookingEvent.EventID,
	)

	// Call the DIDX external API
	err := w.didxSvc.Confirm(ctx, payload.ServiceBookingEvent)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DIDX Confirm failed (attempt %d) for ServiceBookingID: %s - Error: %v",
			taskID,
			retried+1,
			payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DIDX Confirm succeeded for ServiceBookingID: %s, EventID: %s",
		taskID,
		payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
		payload.ServiceBookingEvent.EventID,
	)

	return nil
}
