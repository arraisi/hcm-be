package servicebooking

import "github.com/elgris/sqrl"

type GetServiceBookingPartItem struct {
	ID                   *string
	ServiceBookingPartID *string
	PartNumber           *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingPartItem) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingPartID != nil {
		q.Where(sqrl.Eq{"i_service_booking_part_id": req.ServiceBookingPartID})
	}
	if req.PartNumber != nil {
		q.Where(sqrl.Eq{"c_part_number": req.PartNumber})
	}
}
