package servicebooking

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/elgris/sqrl"
)

type GetServiceBookingWarranty struct {
	ID               *string
	WarrantyID       *string
	ServiceBookingID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingWarranty) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
	if req.WarrantyID != nil {
		q.Where(sqrl.Eq{"i_warranty_id": req.WarrantyID})
	}
}

type DeleteServiceBookingWarranty struct {
	ServiceBookingID *string
}

func (d *DeleteServiceBookingWarranty) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}

func (wr WarrantyRequest) ToModel(serviceBookingID string) domain.ServiceBookingWarranty {
	now := time.Now()
	return domain.ServiceBookingWarranty{
		ServiceBookingID: serviceBookingID,
		WarrantyName:     wr.WarrantyName,
		WarrantyStatus:   wr.WarrantyStatus,
		CreatedAt:        now.UTC(),
		CreatedBy:        constants.System,
		UpdatedAt:        now.UTC(),
		UpdatedBy:        constants.System,
	}
}

type WarrantyRequest struct {
	WarrantyName   string `json:"warranty_name"`
	WarrantyStatus string `json:"warranty_status"`
}

func NewWarrantiesRequest(warranties []domain.ServiceBookingWarranty) []WarrantyRequest {
	var warrantiesRequest []WarrantyRequest
	for _, warranty := range warranties {
		warrantiesRequest = append(warrantiesRequest, WarrantyRequest{
			WarrantyName:   warranty.WarrantyName,
			WarrantyStatus: warranty.WarrantyStatus,
		})
	}
	return warrantiesRequest
}
