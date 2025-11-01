package servicebooking

import "github.com/elgris/sqrl"

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
