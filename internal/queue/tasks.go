package queue

import (
	"encoding/json"

	"github.com/arraisi/hcm-be/internal/domain/dto/servicebooking"
	"github.com/hibiken/asynq"
)

// DIDXConfirmPayload represents the payload for DIDX confirm task
type DIDXConfirmPayload struct {
	ServiceBookingEvent servicebooking.ServiceBookingEvent `json:"service_booking_event"`
}

// NewDIDXConfirmTask creates a new Asynq task for DIDX confirm
func NewDIDXConfirmTask(payload DIDXConfirmPayload) (*asynq.Task, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TaskTypeDIDXConfirm, b), nil
}
