package servicebooking

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/elgris/sqrl"
)

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

type DeleteServiceBookingJob struct {
	ServiceBookingID *string
}

func (d *DeleteServiceBookingJob) Apply(q *sqrl.DeleteBuilder) {
	if d.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": d.ServiceBookingID})
	}
}

type JobRequest struct {
	JobID         string  `json:"job_ID"`
	JobName       string  `json:"job_name"`
	LaborEstPrice float32 `json:"labor_est_price"`
}

func (j *JobRequest) ToDomain(serviceBookingID string) domain.ServiceBookingJob {
	now := time.Now()
	return domain.ServiceBookingJob{
		ServiceBookingID: serviceBookingID,
		JobID:            j.JobID,
		JobName:          j.JobName,
		LaborEstPrice:    j.LaborEstPrice,
		CreatedAt:        now.UTC(),
		CreatedBy:        constants.System,
		UpdatedAt:        now.UTC(),
		UpdatedBy:        constants.System,
	}
}

func NewJobsRequest(jobs []domain.ServiceBookingJob) []JobRequest {
	var jobsRequest []JobRequest
	for _, job := range jobs {
		jobsRequest = append(jobsRequest, JobRequest{
			JobID:         job.JobID,
			JobName:       job.JobName,
			LaborEstPrice: job.LaborEstPrice,
		})
	}
	return jobsRequest
}
