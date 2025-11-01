package testdrive

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/constants"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) RequestTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error {
	if _, ok := constants.TestDriveStatusMap[request.Data.TestDrive.TestDriveStatus]; !ok {
		return errorx.ErrTestDriveStatusInvalid
	}

	if _, ok := constants.LeadsFollowUpStatusMap[request.Data.Leads.LeadsFollowUpStatus]; !ok {
		return errorx.ErrLeadsFollowUpStatusInvalid
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

	// Upsert Leads
	err = s.upsertLeads(ctx, tx, request)
	if err != nil {
		return err
	}

	// Upsert Test Drive
	_, err = s.UpsertServiceTestDrive(ctx, tx, customerID, request)
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}

// upsertLeads checks if a lead exists by LeadsID. If found, it updates the lead; if not found, it creates a new lead.
func (s *service) upsertLeads(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	leadsID := ev.Data.Leads.LeadsID

	_, err := s.leadRepo.GetLeads(ctx, leads.GetLeadsRequest{
		LeadsID: utils.ToPointer(leadsID),
	})
	if err == nil {
		// Found → update
		err := s.leadRepo.UpdateLeads(ctx, tx, ev.Data.Leads.ToDomain())
		if err != nil {
			return err
		}
		return nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		lds := ev.Data.Leads.ToDomain()
		if err := s.leadRepo.CreateLeads(ctx, tx, &lds); err != nil {
			return err
		}
		return nil
	}

	// other error
	return err
}
