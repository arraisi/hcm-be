package testdrive

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) InsertTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error {
	tx, err := s.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Upsert Customer
	err = s.upsertCustomer(ctx, tx, request)
	if err != nil {
		return err
	}

	// Upsert Leads
	err = s.upsertLeads(ctx, tx, request)
	if err != nil {
		return err
	}

	// Upsert Lead Score
	err = s.upsertLeadScore(ctx, tx, request)
	if err != nil {
		return err
	}

	// Upsert Test Drive
	err = s.upsertTestDrive(ctx, tx, request)
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}

// upsertCustomer checks if a customer exists by OneAccountID. If found, it updates the customer; if not found, it creates a new customer.
func (s *service) upsertCustomer(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	oneAccountID := ev.Data.OneAccount.OneAccountID

	_, err := s.customerRepo.GetCustomer(ctx, customer.GetCustomerRequest{
		OneAccountID: oneAccountID,
	})
	if err == nil {
		// Found → update
		err := s.customerRepo.UpdateCustomer(ctx, tx, ev.ToCustomerModel())
		if err != nil {
			return err
		}
		return nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		if err := s.customerRepo.CreateCustomer(ctx, tx, ev.ToCustomerModel()); err != nil {
			return err
		}
		return nil
	}

	// other error
	return err
}

// upsertTestDrive checks if a test drive exists by TestDriveID. If found, it updates the test drive; if not found, it creates a new test drive.
func (s *service) upsertTestDrive(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	testDriveID := ev.Data.TestDrive.TestDriveID

	_, err := s.repo.GetTestDrive(ctx, testdrive.GetTestDriveRequest{
		TestDriveID: utils.ToPointer(testDriveID),
	})
	if err == nil {
		// Found → update
		err := s.repo.UpdateTestDrive(ctx, tx, ev.ToTestDriveModel())
		if err != nil {
			return err
		}
		return nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		if err := s.repo.CreateTestDrive(ctx, tx, ev.ToTestDriveModel()); err != nil {
			return err
		}
		return nil
	}

	// other error
	return err
}

// upsertLeads checks if a lead exists by LeadsID. If found, it updates the lead; if not found, it creates a new lead.
func (s *service) upsertLeads(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	leadsID := strings.ReplaceAll(ev.Data.Leads.LeadsID, "-", "")

	_, err := s.leadRepo.GetLeads(ctx, leads.GetLeadsRequest{
		LeadsID: utils.ToPointer(leadsID),
	})
	if err == nil {
		// Found → update
		err := s.leadRepo.UpdateLeads(ctx, tx, ev.ToLeadsModel())
		if err != nil {
			return err
		}
		return nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		if err := s.leadRepo.CreateLeads(ctx, tx, ev.ToLeadsModel()); err != nil {
			return err
		}
		return nil
	}

	// other error
	return err
}

// upsertLeadScore checks if a lead score exists by IID. If found, it updates the lead score; if not found, it creates a new lead score.
func (s *service) upsertLeadScore(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	leadsID := strings.ReplaceAll(ev.Data.Leads.LeadsID, "-", "")

	_, err := s.leadScoreRepo.GetLeadScore(ctx, leads.GetLeadScoreRequest{
		IID: utils.ToPointer(leadsID),
	})
	if err == nil {
		// Found → update
		err := s.leadScoreRepo.UpdateLeadScore(ctx, tx, ev.ToLeadScoreModel())
		if err != nil {
			return err
		}
		return nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		if err := s.leadScoreRepo.CreateLeadScore(ctx, tx, ev.ToLeadScoreModel()); err != nil {
			return err
		}
		return nil
	}

	// other error
	return err
}
