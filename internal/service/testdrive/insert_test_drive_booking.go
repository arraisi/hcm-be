package testdrive

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/arraisi/hcm-be/pkg/constants"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) InsertTestDriveBooking(ctx context.Context, request testdrive.TestDriveEvent) error {
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
	customerID, err := s.upsertCustomer(ctx, tx, request)
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
	err = s.upsertTestDrive(ctx, tx, request, customerID)
	if err != nil {
		return err
	}

	return s.transactionRepo.CommitTransaction(tx)
}

// upsertCustomer checks if a customer exists by OneAccountID. If found, it updates the customer; if not found, it creates a new customer.
func (s *service) upsertCustomer(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) (string, error) {
	oneAccountID := ev.Data.OneAccount.OneAccountID

	_, err := s.customerRepo.GetCustomer(ctx, customer.GetCustomerRequest{
		OneAccountID: oneAccountID,
	})
	if err == nil {
		// Found → update
		customerID, err := s.customerRepo.UpdateCustomer(ctx, tx, ev.ToCustomerModel())
		if err != nil {
			return customerID, err
		}
		return customerID, nil
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		customerID, err := s.customerRepo.CreateCustomer(ctx, tx, ev.ToCustomerModel())
		if err != nil {
			return customerID, err
		}
		return customerID, nil
	}

	// other error
	return "", err
}

// upsertTestDrive checks if a test drive exists by TestDriveID. If found, it updates the test drive; if not found, it creates a new test drive.
func (s *service) upsertTestDrive(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent, customerID string) error {
	testDriveID := ev.Data.TestDrive.TestDriveID

	testDrives, err := s.repo.GetTestDrives(ctx, testdrive.GetTestDriveRequest{
		CustomerID: utils.ToPointer(customerID),
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// Ensure only one active test drive per customer
	if len(testDrives) > 1 {
		return errorx.ErrTestDriveCustomerHasBooking
	}

	// Update existing test drive if found
	if len(testDrives) == 1 && testDrives[0].TestDriveID == testDriveID {
		return s.repo.UpdateTestDrive(ctx, tx, ev.ToTestDriveModel(customerID))
	}

	// Create new test drive if not found
	return s.repo.CreateTestDrive(ctx, tx, ev.ToTestDriveModel(customerID))
}

// upsertLeads checks if a lead exists by LeadsID. If found, it updates the lead; if not found, it creates a new lead.
func (s *service) upsertLeads(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	leadsID := ev.Data.Leads.LeadsID

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

// upsertLeadScore checks if a lead score exists by ID. If found, it updates the lead score; if not found, it creates a new lead score.
func (s *service) upsertLeadScore(ctx context.Context, tx *sqlx.Tx, ev testdrive.TestDriveEvent) error {
	leadsID := ev.Data.Leads.LeadsID

	_, err := s.leadScoreRepo.GetLeadScore(ctx, leads.GetLeadScoreRequest{
		ID: utils.ToPointer(leadsID),
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
