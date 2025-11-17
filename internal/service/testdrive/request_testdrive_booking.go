package testdrive

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/internal/queue"
	"github.com/arraisi/hcm-be/pkg/constants"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) RequestTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error {
	if _, ok := constants.TestDriveStatusMap[request.Data.TestDrive.TestDriveStatus]; !ok {
		return errorx.ErrTestDriveStatusInvalid
	}

	if _, ok := constants.TestDriveLocationMap[request.Data.TestDrive.Location]; !ok {
		return errorx.ErrTestDriveLocationInvalid
	}

	if _, ok := constants.LeadsFollowUpStatusMap[request.Data.Leads.LeadsFollowUpStatus]; !ok {
		return errorx.ErrLeadsFollowUpStatusInvalid
	}

	if request.Data.Leads.LeadsType != constants.LeadsTypeTestDriveRequest {
		return errorx.ErrLeadsTypeInvalid
	}

	if _, ok := constants.LeadsSourceMap[request.Data.Leads.LeadSource]; !ok {
		return errorx.ErrLeadsSourceInvalid
	}

	// Validate cancellation reason when test drive status is CANCEL_SUBMITTED
	if request.Data.TestDrive.TestDriveStatus == constants.TestDriveBookingStatusCancelSubmitted {
		// Cancellation reason is required
		if request.Data.TestDrive.CancellationReason == nil || *request.Data.TestDrive.CancellationReason == "" {
			return errorx.ErrTestDriveCancellationReasonRequired
		}

		// Validate cancellation reason is one of the allowed values
		if _, ok := constants.CancellationReasonMap[*request.Data.TestDrive.CancellationReason]; !ok {
			return errorx.ErrTestDriveCancellationReasonInvalid
		}

		// If cancellation reason is OTHERS, other_cancellation_reason is required
		if *request.Data.TestDrive.CancellationReason == constants.CancellationReasonOthers {
			if request.Data.TestDrive.OtherCancellationReason == nil || *request.Data.TestDrive.OtherCancellationReason == "" {
				return errorx.ErrTestDriveOtherCancellationReasonRequired
			}
		}
	}

	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Upsert Customer
	customerID, err := s.customerSvc.UpsertCustomer(ctx, tx, request.Data.OneAccount)
	if err != nil {
		return err
	}

	// Upsert Test Drive
	testDriveID, err := s.UpsertServiceTestDrive(ctx, tx, customerID, request)
	if err != nil {
		return err
	}

	// Upsert Leads
	err = s.upsertLeads(ctx, tx, customerID, testDriveID, request)
	if err != nil {
		return err
	}

	err = s.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return err
	}

	err = s.queueClient.EnqueueDMSTestDriveRequest(context.Background(), queue.DMSTestDriveRequestPayload{
		TestDriveEvent: request,
	})
	if err != nil {
		return fmt.Errorf("failed to enqueue DMS test drive request: %w", err)
	}

	return nil
}

// upsertLeads checks if a lead exists by LeadsID. If found, it updates the lead; if not found, it creates a new lead.
func (s *service) upsertLeads(ctx context.Context, tx *sqlx.Tx, customerID, testDriveID string, ev testdrive.TestDriveEvent) error {
	leadsID := ev.Data.Leads.LeadsID

	lead, err := s.leadRepo.GetLeads(ctx, leads.GetLeadsRequest{
		LeadsID: utils.ToPointer(leadsID),
	})
	if err == nil {
		// Found → update
		lds := ev.Data.Leads.ToDomain(customerID, testDriveID)
		lds.ID = lead.ID
		err := s.leadRepo.UpdateLeads(ctx, tx, lds)
		if err != nil {
			return err
		}
		return nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		lds := ev.Data.Leads.ToDomain(customerID, testDriveID)
		if err := s.leadRepo.CreateLeads(ctx, tx, &lds); err != nil {
			return err
		}
		return nil
	}

	// other error
	return err
}
