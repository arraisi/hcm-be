package queue

import (
	"encoding/json"
	"fmt"
	"github.com/arraisi/hcm-be/internal/domain/dto/appraisalbooking"
	"github.com/arraisi/hcm-be/internal/domain/dto/oneaccess"
	"github.com/arraisi/hcm-be/internal/domain/dto/toyotaid"

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

// DMSCreateOneAccessPayload represents the payload for DMS create one access task
type DMSCreateOneAccessPayload struct {
	OneAccessRequest oneaccess.Request `json:"one_access_request"`
}

// NewDMSCreateOneAccessTask creates a new Asynq task for DMS toyota id
func NewDMSCreateOneAccessTask(payload DMSCreateOneAccessPayload) (*asynq.Task, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	taskKey := fmt.Sprintf("%s:%s", TaskTypeDMSCreateOneAccess, payload.OneAccessRequest.EventID)
	return asynq.NewTask(taskKey, b), nil
}

// DMSCreateToyotaIDPayload represents the payload for DMS create toyota id task
type DMSCreateToyotaIDPayload struct {
	ToyotaIDRequest toyotaid.Request `json:"toyota_id_request"`
}

// NewDMSCreateToyotaIDTask creates a new Asynq task for DMS toyota id
func NewDMSCreateToyotaIDTask(payload DMSCreateToyotaIDPayload) (*asynq.Task, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	taskKey := fmt.Sprintf("%s:%s", TaskTypeDMSCreateToyotaID, payload.ToyotaIDRequest.EventID)
	return asynq.NewTask(taskKey, b), nil
}

// DMSAppraisalBookingRequestPayload represents the payload for the DMS appraisal booking request task
type DMSAppraisalBookingRequestPayload struct {
	AppraisalBookingRequest appraisalbooking.EventRequest `json:"appraisal_booking_request"`
}

// NewDMSAppraisalBookingRequestTask creates a new Asynq task for DMS appraisal booking request
func NewDMSAppraisalBookingRequestTask(payload DMSAppraisalBookingRequestPayload) (*asynq.Task, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	taskKey := fmt.Sprintf("%s:%s", TaskTypeDMSAppraisalBookingRequest, payload.AppraisalBookingRequest.EventID)
	return asynq.NewTask(taskKey, b), nil
}
