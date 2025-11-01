package testdrive

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	errorx "github.com/arraisi/hcm-be/pkg/errors"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func (s *service) UpsertServiceTestDrive(ctx context.Context, tx *sqlx.Tx, customerID string, request testdrive.TestDriveEvent) (string, error) {
	testDrives, err := s.repo.GetTestDrives(ctx, testdrive.GetTestDriveRequest{
		TestDriveID: utils.ToPointer(request.Data.TestDrive.TestDriveID),
	})
	if err == nil && len(testDrives) > 0 {
		if len(testDrives) > 1 {
			return testDrives[0].ID, errorx.ErrTestDriveCustomerHasBooking
		}

		// Found → update
		td := request.ToTestDriveModel(customerID)
		td.ID = testDrives[0].ID
		err := s.repo.UpdateTestDrive(ctx, tx, td)
		if err != nil {
			return td.ID, err
		}
	}

	// Not found → create
	if errors.Is(err, sql.ErrNoRows) {
		td := request.ToTestDriveModel(customerID)
		err := s.repo.CreateTestDrive(ctx, tx, &td)
		if err != nil {
			return td.ID, err
		}
		return td.ID, nil
	}

	return "", err
}
