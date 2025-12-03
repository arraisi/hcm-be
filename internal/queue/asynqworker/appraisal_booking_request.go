package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleDMSAppraisalBookingRequest processes DMS appraisal booking request tasks
func (w *Worker) handleDMSAppraisalBookingRequest(ctx context.Context, t *asynq.Task) error {
	var payload queue.DMSAppraisalBookingRequestPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DMSAppraisalBookingRequestPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DMS AppraisalBookingRequest task (attempt %d) for Process: %s, EventID: %s",
		taskID,
		retried+1,
		payload.AppraisalBookingRequest.Process,
		payload.AppraisalBookingRequest.EventID,
	)

	// Call the DMS external API
	err := w.dmsSvc.AppraisalBookingRequest(ctx, payload.AppraisalBookingRequest)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DMS CreateOneAccess failed (attempt %d) for Process: %s, EventID: %s - Error: %v",
			taskID,
			retried+1,
			payload.AppraisalBookingRequest.Process,
			payload.AppraisalBookingRequest.EventID,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DMS CreateOneAccess succeeded for Process: %s, EventID: %s",
		taskID,
		payload.AppraisalBookingRequest.Process,
		payload.AppraisalBookingRequest.EventID,
	)

	return nil
}
