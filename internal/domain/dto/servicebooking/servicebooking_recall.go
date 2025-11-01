package servicebooking

import "github.com/elgris/sqrl"

type GetServiceBookingRecall struct {
	ID               *string
	RecallID         *string
	ServiceBookingID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingRecall) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.RecallID != nil {
		q.Where(sqrl.Eq{"i_recall_id": req.RecallID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
}

type DeleteServiceBookingRecall struct {
	ServiceBookingID *string
}

func (d *DeleteServiceBookingRecall) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}
