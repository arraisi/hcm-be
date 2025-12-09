package asynqworker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/hibiken/asynq"
)

// handleAppraisalConfirm processes DIDX appraisal confirm tasks
func (w *Worker) handleAppraisalConfirm(ctx context.Context, t *asynq.Task) error {
	var payload queue.DIDXAppraisalConfirmPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Printf("[ERROR] Failed to unmarshal DIDXAppraisalConfirmPayload: %v", err)
		// Return SkipRetry for unmarshal errors (permanent failure)
		return fmt.Errorf("unmarshal error (will not retry): %w", asynq.SkipRetry)
	}

	taskID, _ := asynq.GetTaskID(ctx)

	// Log task attempt
	retried, _ := asynq.GetRetryCount(ctx)
	log.Printf("[INFO] TaskID: %s Processing DIDX Appraisal Confirm task (attempt %d) for AppraisalBookingID: %s, EventID: %s",
		taskID,
		retried+1,
		payload.AppraisalConfirmEvent.Data.RequestAppraisal.AppraisalBookingID,
		payload.AppraisalConfirmEvent.EventID,
	)

	// Call the DIDX external API
	err := w.didxSvc.ConfirmAppraisal(ctx, payload.AppraisalConfirmEvent)

	// Send callback to DMS
	w.sendDMSCallback(
		ctx,
		taskID,
		payload.AppraisalConfirmEvent.EventID,
		payload.AppraisalConfirmEvent.Data.RequestAppraisal.AppraisalBookingID,
		err,
	)

	if err != nil {
		// Log failure
		log.Printf("[ERROR] TaskID: %s DIDX Appraisal Confirm failed (attempt %d) for AppraisalBookingID: %s - Error: %v",
			taskID,
			retried+1,
			payload.AppraisalConfirmEvent.Data.RequestAppraisal.AppraisalBookingID,
			err,
		)
		return err
	}

	// Log success
	log.Printf("[SUCCESS] TaskID: %s DIDX Appraisal Confirm succeeded for AppraisalBookingID: %s, EventID: %s",
		taskID,
		payload.AppraisalConfirmEvent.Data.RequestAppraisal.AppraisalBookingID,
		payload.AppraisalConfirmEvent.EventID,
	)

	return nil
}
