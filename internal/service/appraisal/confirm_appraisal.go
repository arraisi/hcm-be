package appraisal

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"
	"github.com/arraisi/hcm-be/internal/queue"
)

// ConfirmAppraisal processes appraisal booking confirmation from webhook event
func (s *service) ConfirmAppraisal(ctx context.Context, request appraisal.AppraisalConfirmEvent) error {
	// Start transaction
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = s.transactionRepo.RollbackTransaction(tx)
		}
	}()

	// Upsert customer
	customerID, err := s.customerSvc.UpsertCustomerV2(ctx, tx, request.Data.OneAccount.ToDomain())
	if err != nil {
		return err
	}

	// Get existing appraisal
	existingAppraisal, err := s.appraisalBookingRepo.GetAppraisal(ctx, appraisal.GetAppraisalRequest{
		AppraisalBookingID: request.Data.RequestAppraisal.AppraisalBookingID,
	})
	if err != nil {
		return err
	}

	// Convert to appraisal model and update with new data
	appraisalModel := request.ToAppraisalUpdateModel()
	appraisalModel.ID = existingAppraisal.ID // Preserve existing ID
	appraisalModel.OneAccountID = &customerID
	appraisalModel.VIN = existingAppraisal.VIN
	appraisalModel.LeadsID = existingAppraisal.LeadsID

	// Update appraisal in database
	err = s.appraisalBookingRepo.UpdateAppraisal(ctx, tx, appraisalModel)
	if err != nil {
		return err
	}

	// Upsert leads
	leadsModel := request.Data.Leads.ToDomain(customerID, nil, nil)
	leadsModel.LeadsID = request.Data.Leads.LeadsID

	_, err = s.leadsSvc.UpsertLeads(ctx, tx, leadsModel)
	if err != nil {
		return err
	}

	// Commit transaction
	if err = s.transactionRepo.CommitTransaction(tx); err != nil {
		return err
	}

	// Enqueue task to send confirmation to external API (DIDX)
	err = s.queueClient.EnqueueDIDXAppraisalConfirm(context.Background(), queue.DIDXAppraisalConfirmPayload{
		AppraisalConfirmEvent: request,
	})
	if err != nil {
		return err
	}

	return nil
}
