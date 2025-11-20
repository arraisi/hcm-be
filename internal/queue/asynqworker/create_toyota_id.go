package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleDMSCreateToyotaID processes DMS create toyota id tasks
func (w *Worker) handleDMSCreateToyotaID(ctx context.Context, t *asynq.Task) error {
	var payload queue.DMSCreateToyotaIDPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DMSCreateToyotaIDPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DMS CreateToyotaID task (attempt %d) for OneAccountID: %s, VIN: %s, EventID: %s",
		taskID,
		retried+1,
		payload.ToyotaIDRequest.Data.OneAccount.OneAccountID,
		payload.ToyotaIDRequest.Data.CustomerVehicle.VIN,
		payload.ToyotaIDRequest.EventID,
	)

	// Call the DMS external API
	err := w.dmsSvc.CreateToyotaID(ctx, payload.ToyotaIDRequest)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DMS CreateToyotaID failed (attempt %d) for OneAccountID: %s, VIN: %s - Error: %v",
			taskID,
			retried+1,
			payload.ToyotaIDRequest.Data.OneAccount.OneAccountID,
			payload.ToyotaIDRequest.Data.CustomerVehicle.VIN,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DMS CreateToyotaID succeeded for OneAccountID: %s, VIN: %s, EventID: %s",
		taskID,
		payload.ToyotaIDRequest.Data.OneAccount.OneAccountID,
		payload.ToyotaIDRequest.Data.CustomerVehicle.VIN,
		payload.ToyotaIDRequest.EventID,
	)

	return nil
}
