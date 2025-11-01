package servicebooking

import "github.com/elgris/sqrl"

type GetServiceBookingJob struct {
	ID               *string
	JobID            *string
	ServiceBookingID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetServiceBookingJob) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.JobID != nil {
		q.Where(sqrl.Eq{"i_job_id": req.JobID})
	}
	if req.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": req.ServiceBookingID})
	}
}
