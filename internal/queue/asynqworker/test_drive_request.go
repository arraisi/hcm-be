package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleDMSTestDriveRequest processes DMS test drive request tasks
func (w *Worker) handleDMSTestDriveRequest(ctx context.Context, t *asynq.Task) error {
	var payload queue.DMSTestDriveRequestPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DMSTestDriveRequestPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] Processing DMS TestDriveRequest task (attempt %d) for TestDriveID: %s, EventID: %s",
		retried+1,
		payload.TestDriveEvent.Data.TestDrive.TestDriveID,
		payload.TestDriveEvent.EventID,
	)

	// Call the DMS external API
	err := w.dmsSvc.TestDriveRequest(ctx, payload.TestDriveEvent)
	if err != nil {
		// Log failure and return error to trigger retry
		log.Printf("[ERROR] DMS TestDriveRequest failed (attempt %d) for TestDriveID: %s - Error: %v",
			retried+1,
			payload.TestDriveEvent.Data.TestDrive.TestDriveID,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] DMS TestDriveRequest succeeded for TestDriveID: %s, EventID: %s",
		payload.TestDriveEvent.Data.TestDrive.TestDriveID,
		payload.TestDriveEvent.EventID,
	)

	return nil
}
