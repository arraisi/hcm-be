package testdrive

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/errors"
)

func (s service) CreateTestDriveBooking(_ context.Context, request domain.BookingEvent) error {
	// logging the received test drive request
	fmt.Printf("[CreateTestDriveBooking] TestDrive Data: %+v\n", request)

	// Validate test drive data
	if request.Data.TestDrive.TestDriveID == "" {
		return errors.ErrTestDriveInvalidData
	}

	// Validate required fields
	if request.Data.TestDrive.Model == "" || request.Data.TestDrive.OutletID == "" {
		return errors.ErrTestDriveInvalidData
	}

	// Validate event ID
	if request.EventID == "" {
		return errors.ErrTestDriveInvalidData
	}

	// Validate OneAccount information
	if request.Data.OneAccount.OneAccountID == "" || request.Data.OneAccount.FirstName == "" {
		return errors.ErrTestDriveInvalidData
	}

	// Simulate potential database operation failure
	// In real implementation, replace this with actual database operations
	// if dbErr := s.repo.CreateTestDriveBooking(request); dbErr != nil {
	//     return errors.ErrTestDriveCreateFailed
	// }

	return nil
}
