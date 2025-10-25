package testdrive

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/errors"
)

func (s service) UpdateTestDriveBooking(_ context.Context, request domain.BookingEvent) error {
	// logging the received test drive request
	fmt.Printf("[UpdateTestDriveBooking] TestDrive Data: %+v\n", request)

	// Validate test drive data
	if request.Data.TestDrive.TestDriveID == "" {
		return errors.ErrTestDriveInvalidData
	}

	// Validate event ID
	if request.EventID == "" {
		return errors.ErrTestDriveInvalidData
	}

	// Validate status for update operations
	status := request.Data.TestDrive.TestDriveStatus
	if status != "CHANGE_REQUEST" && status != "CANCEL_SUBMITTED" {
		return errors.ErrTestDriveInvalidStatus
	}

	// Simulate potential database operation failure
	// In real implementation, replace this with actual database operations
	// if dbErr := s.repo.UpdateTestDriveBooking(request); dbErr != nil {
	//     return errors.ErrTestDriveUpdateFailed
	// }

	return nil
}
