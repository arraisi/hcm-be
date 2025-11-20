package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleDMSCreateOneAccess processes DMS create one access tasks
func (w *Worker) handleDMSCreateOneAccess(ctx context.Context, t *asynq.Task) error {
	var payload queue.DMSCreateOneAccessPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DMSCreateOneAccessPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DMS CreateOneAccess task (attempt %d) for OneAccountID: %s, EventID: %s",
		taskID,
		retried+1,
		payload.OneAccessRequest.Data.OneAccount.OneAccountID,
		payload.OneAccessRequest.EventID,
	)

	// Call the DMS external API
	err := w.dmsSvc.CreateOneAccess(ctx, payload.OneAccessRequest)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DMS CreateOneAccess failed (attempt %d) for OneAccountID: %s - Error: %v",
			taskID,
			retried+1,
			payload.OneAccessRequest.Data.OneAccount.OneAccountID,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DMS CreateOneAccess succeeded for OneAccountID: %s, EventID: %s",
		taskID,
		payload.OneAccessRequest.Data.OneAccount.OneAccountID,
		payload.OneAccessRequest.EventID,
	)

	return nil
}
