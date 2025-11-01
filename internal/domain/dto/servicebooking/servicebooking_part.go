package servicebooking

import "github.com/elgris/sqrl"

type GetServiceBookingPart struct {
	ID               *string
	ServiceBookingID *string
	PartType         *string
	PackageID        *string
	PartNumber       *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingPart) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
	if req.PartType != nil {
		q.Where(sqrl.Eq{"c_part_type": req.PartType})
	}
	if req.PackageID != nil {
		q.Where(sqrl.Eq{"i_package_id": req.PackageID})
	}
	if req.PartNumber != nil {
		q.Where(sqrl.Eq{"c_part_number": req.PartNumber})
	}
}

type DeleteServiceBookingPart struct {
	ServiceBookingID *string
}

func (d *DeleteServiceBookingPart) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}
