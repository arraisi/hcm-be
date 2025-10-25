package testdrive

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (s service) UpdateTestDriveBooking(_ context.Context, request domain.BookingEvent) error {
	// logging the received test drive request
	fmt.Printf("[UpdateTestDriveBooking] TestDrive Data: %+v\n", request)

	return nil
}
