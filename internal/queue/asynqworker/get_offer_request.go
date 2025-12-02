package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleDMSGetOfferRequest processes DMS get offer request tasks
func (w *Worker) handleDMSGetOfferRequest(ctx context.Context, t *asynq.Task) error {
	var payload queue.DMSCreateGetOfferPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DMSCreateGetOfferPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DMS GetOfferRequest task (attempt %d) for GetOfferNumber: %s, EventID: %s",
		taskID,
		retried+1,
		payload.GetOfferEvent.Data.Leads.GetOfferNumber,
		payload.GetOfferEvent.EventID,
	)

	// Call the DMS external API
	err := w.dmsSvc.GetOfferRequest(ctx, payload.GetOfferEvent)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DMS GetOfferRequest failed (attempt %d) for GetOfferNumber: %s - Error: %v",
			taskID,
			retried+1,
			payload.GetOfferEvent.Data.Leads.GetOfferNumber,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DMS GetOfferRequest succeeded for GetOfferNumber: %s, EventID: %s",
		taskID,
		payload.GetOfferEvent.Data.Leads.GetOfferNumber,
		payload.GetOfferEvent.EventID,
	)

	return nil
}
