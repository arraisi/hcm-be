package testdrive

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
)

type Service interface {
	ConfirmTestDriveBooking(ctx context.Context, request testdrive.ConfirmTestDriveBookingRequest) error
}

// Handler handles HTTP requests for user operations
type Handler struct {
	svc Service
}

// New creates a new CustomerHandler instance
func New(svc Service) Handler {
	return Handler{svc: svc}
}
