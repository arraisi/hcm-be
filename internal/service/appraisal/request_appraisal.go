package appraisal

import (
	"context"
	"fmt"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisal"
	"github.com/arraisi/hcm-be/internal/queue"
)

// ==========================
// Core flow: G01 (request)
// ==========================

// RequestAppraisal
// G01: DI/DX menerima "request appraisal booking" + leads dari mTOYOTA, simpan ke DB, lalu kirim ke DMS.
func (s *service) RequestAppraisal(ctx context.Context, event appraisal.EventRequest) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	// pastikan rollback kalau ada error / panic
	defer func() {
		_ = s.transactionRepo.RollbackTransaction(tx)
	}()

	// ==== 1. Upsert Customer (one_account) ====
	customerModel := event.Data.OneAccount.ToCustomerModel()
	customerID, err := s.customerSvc.UpsertCustomerV2(ctx, tx, customerModel)
	if err != nil {
		return fmt.Errorf("upsert customer: %w", err)
	}

	// ==== 2. Upsert Used Car ====
	usedCarModel := event.Data.UsedCar.ToUsedCarModel(customerID)
	_, err = s.usedCarSvc.UpsertUsedCar(ctx, tx, usedCarModel)
	if err != nil {
		return fmt.Errorf("upsert used car: %w", err)
	}

	// ==== 3. Upsert Leads and Leads Score====
	leadsModel, leadScoreModel := event.Data.Leads.ToLeadsModel(customerID)
	_, err = s.leadsSvc.UpsertLeads(ctx, tx, leadsModel)
	if err != nil {
		return fmt.Errorf("upsert leads: %w", err)
	}
	_, err = s.leadsScoreSvc.UpsertLeadsScore(ctx, tx, leadScoreModel)
	if err != nil {
		return fmt.Errorf("upsert leads: %w", err)
	}

	// ==== 4. Create Appraisal Booking record ====
	appraisalBookingModel := event.Data.RequestAppraisal.ToAppraisalModel(customerModel.OneAccountID, usedCarModel.VIN, leadsModel.LeadsID)
	err = s.appraisalBookingRepo.CreateAppraisal(ctx, tx, &appraisalBookingModel)
	if err != nil {
		return fmt.Errorf("create appraisal booking: %w", err)
	}

	// ==== 5. Commit TX ====
	if err := s.transactionRepo.CommitTransaction(tx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	// ==== 6. Enqueue ke Dealer System (G01 → G02) ====
	payload := queue.DMSAppraisalBookingRequestPayload{
		AppraisalBookingRequest: event,
	}

	if err := s.queueClient.EnqueueDMSAppraisalBookingRequest(context.Background(), payload); err != nil {
		return fmt.Errorf("enqueue appraisal booking request: %w", err)
	}

	return nil
}
