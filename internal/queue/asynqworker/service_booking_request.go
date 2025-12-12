package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleDMSServiceBookingRequest processes DMS service booking request tasks
func (w *Worker) handleDMSServiceBookingRequest(ctx context.Context, t *asynq.Task) error {
	var payload queue.DMSServiceBookingRequestPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DMSServiceBookingRequestPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DMS ServiceBookingRequest task (attempt %d) for BookingID: %s, EventID: %s",
		taskID,
		retried+1,
		payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
		payload.ServiceBookingEvent.EventID,
	)

	// Call the DMS After Sales external API
	err := w.dmsAfterSalesSvc.ServiceBookingRequest(ctx, payload.ServiceBookingEvent)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DMS ServiceBookingRequest failed (attempt %d) for BookingID: %s - Error: %v",
			taskID,
			retried+1,
			payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DMS ServiceBookingRequest succeeded for BookingID: %s, EventID: %s",
		taskID,
		payload.ServiceBookingEvent.Data.ServiceBookingRequest.BookingId,
		payload.ServiceBookingEvent.EventID,
	)

	return nil
}
