package queue

import (
	"encoding/json"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/hibiken/asynq"
)

// DIDXServiceBookingConfirmPayload represents the payload for DIDX confirm task
type DIDXServiceBookingConfirmPayload struct {
	ServiceBookingEvent servicebooking.ServiceBookingEvent `json:"service_booking_event"`
}

// NewDIDXServiceBookingConfirmTask creates a new Asynq task for DIDX confirm
func NewDIDXServiceBookingConfirmTask(payload DIDXServiceBookingConfirmPayload) (*asynq.Task, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	taskKey := fmt.Sprintf("%s:%s", TaskTypeDIDXConfirm, payload.ServiceBookingEvent.EventID)
	return asynq.NewTask(taskKey, b), nil
}

// DMSTestDriveRequestPayload represents the payload for DMS test drive request task
type DMSTestDriveRequestPayload struct {
	TestDriveEvent testdrive.TestDriveEvent `json:"test_drive_event"`
}

// NewDMSTestDriveRequestTask creates a new Asynq task for DMS test drive request
func NewDMSTestDriveRequestTask(payload DMSTestDriveRequestPayload) (*asynq.Task, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	taskKey := fmt.Sprintf("%s:%s", TaskTypeDMSTestDriveRequest, payload.TestDriveEvent.EventID)
	return asynq.NewTask(taskKey, b), nil
}
