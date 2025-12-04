package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleTestDriveConfirm processes DIDX test drive confirm tasks
func (w *Worker) handleTestDriveConfirm(ctx context.Context, t *asynq.Task) error {
	var payload queue.DIDXTestDriveConfirmPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DIDXTestDriveConfirmPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DIDX Test Drive Confirm task (attempt %d) for TestDriveID: %s, EventID: %s",
		taskID,
		retried+1,
		payload.TestDriveEvent.Data.TestDrive.TestDriveID,
		payload.TestDriveEvent.EventID,
	)

	// Call the DIDX external API
	err := w.didxSvc.ConfirmTestDrive(ctx, payload.TestDriveEvent)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] TaskID: %s DIDX Test Drive Confirm failed (attempt %d) for TestDriveID: %s - Error: %v",
			taskID,
			retried+1,
			payload.TestDriveEvent.Data.TestDrive.TestDriveID,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DIDX Test Drive Confirm succeeded for TestDriveID: %s, EventID: %s",
		taskID,
		payload.TestDriveEvent.Data.TestDrive.TestDriveID,
		payload.TestDriveEvent.EventID,
	)

	return nil
}
