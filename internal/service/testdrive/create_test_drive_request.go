package testdrive

import (
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (s service) CreateTestDriveRequest(request domain.BookingEvent) error {
	// logging the received test drive request
	fmt.Printf("TestDrive Data: %+v\n", request)

	return nil
}
