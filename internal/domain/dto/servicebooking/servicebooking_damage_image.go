package servicebooking

import "github.com/elgris/sqrl"

// GetServiceBookingDamageImage represents query filters for damage images
type GetServiceBookingDamageImage struct {
	ID               *string
	ServiceBookingID *string
}

// Apply applies the filters to the query builder
func (req *GetServiceBookingDamageImage) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
}

// DeleteServiceBookingDamageImage represents delete filters
type DeleteServiceBookingDamageImage struct {
	ServiceBookingID *string
}

// Apply applies the filters to the delete builder
func (d *DeleteServiceBookingDamageImage) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}
