package domain

import (
	"time"
)

type ServiceBookingJob struct {
	ID               string    `db:"i_id"`
	ServiceBookingID string    `db:"i_service_booking_id"`
	JobID            string    `db:"c_job_id"`
	JobName          string    `db:"n_job_name"`
	LaborEstPrice    float32   `db:"v_labor_est_price"`
	CreatedAt        time.Time `db:"d_created_at"`
	CreatedBy        string    `db:"c_created_by"`
	UpdatedAt        time.Time `db:"d_updated_at"`
	UpdatedBy        string    `db:"c_updated_by"`
	Deleted          bool      `db:"b_deleted"`
}

// TableName returns the database table name for the ServiceBookingJob model
func (sbj *ServiceBookingJob) TableName() string {
	return "dbo.tr_service_booking_job"
}

// Columns returns the list of database columns for the ServiceBookingJob model
func (sbj *ServiceBookingJob) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"c_job_id",
		"n_job_name",
		"n_labor_est_price",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the ServiceBookingJob model
func (sbj *ServiceBookingJob) SelectColumns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"c_job_id",
		"n_job_name",
		"v_labor_est_price",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting a new ServiceBookingJob record
func (sbj *ServiceBookingJob) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0, len(sbj.Columns()))
	values := make([]interface{}, 0, len(sbj.Columns()))

	if sbj.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, sbj.ServiceBookingID)
	}
	if sbj.JobID != "" {
		columns = append(columns, "c_job_id")
		values = append(values, sbj.JobID)
	}
	if sbj.JobName != "" {
		columns = append(columns, "n_job_name")
		values = append(values, sbj.JobName)
	}
	if sbj.LaborEstPrice != 0 {
		columns = append(columns, "v_labor_est_price")
		values = append(values, sbj.LaborEstPrice)
	}
	columns = append(columns, "c_created_by")
	values = append(values, sbj.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, sbj.CreatedBy)

	return columns, values
}

// ToUpdateMap prepares the map of fields to be updated for an existing ServiceBookingJob record
func (sbj *ServiceBookingJob) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if sbj.ServiceBookingID != "" {
		updateMap["i_service_booking_id"] = sbj.ServiceBookingID
	}
	if sbj.JobID != "" {
		updateMap["c_job_id"] = sbj.JobID
	}
	if sbj.JobName != "" {
		updateMap["n_job_name"] = sbj.JobName
	}
	if sbj.LaborEstPrice != 0 {
		updateMap["n_labor_est_price"] = sbj.LaborEstPrice
	}
	updateMap["d_updated_at"] = time.Now().UTC()
	updateMap["c_updated_by"] = sbj.UpdatedBy

	return updateMap
}
