package domain

import (
	"time"
)

// ServiceBookingVehicleInsurancePolicy represents individual insurance policy
type ServiceBookingVehicleInsurancePolicy struct {
	ID                     string    `db:"i_id"`
	ServiceBookingID       string    `db:"i_service_booking_id"`
	VehicleInsuranceID     string    `db:"i_vehicle_insurance_id"`
	InsuranceType          string    `db:"c_insurance_type"`
	InsuranceCoverage      string    `db:"e_insurance_coverage"`
	InsuranceStartDatetime time.Time `db:"d_insurance_datetime_start"`
	InsuranceEndDatetime   time.Time `db:"d_insurance_datetime_end"`
	InsuranceStatus        string    `db:"c_insurance_status"`
	CreatedAt              time.Time `db:"d_created_at"`
	CreatedBy              string    `db:"c_created_by"`
	UpdatedAt              time.Time `db:"d_updated_at"`
	UpdatedBy              string    `db:"c_updated_by"`
}

// TableName returns the database table name
func (ip *ServiceBookingVehicleInsurancePolicy) TableName() string {
	return "dbo.tr_service_booking_vehicle_insurance_policy"
}

// Columns returns the list of database columns
func (ip *ServiceBookingVehicleInsurancePolicy) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"i_vehicle_insurance_id",
		"c_insurance_type",
		"e_insurance_coverage",
		"d_insurance_datetime_start",
		"d_insurance_datetime_end",
		"c_insurance_status",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select
func (ip *ServiceBookingVehicleInsurancePolicy) SelectColumns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"i_vehicle_insurance_id",
		"c_insurance_type",
		"e_insurance_coverage",
		"d_insurance_datetime_start",
		"d_insurance_datetime_end",
		"c_insurance_status",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting
func (ip *ServiceBookingVehicleInsurancePolicy) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	if ip.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, ip.ServiceBookingID)
	}
	if ip.VehicleInsuranceID != "" {
		columns = append(columns, "i_vehicle_insurance_id")
		values = append(values, ip.VehicleInsuranceID)
	}
	if ip.InsuranceType != "" {
		columns = append(columns, "c_insurance_type")
		values = append(values, ip.InsuranceType)
	}
	if len(ip.InsuranceCoverage) > 0 {
		columns = append(columns, "e_insurance_coverage")
		values = append(values, ip.InsuranceCoverage)
	}
	if !ip.InsuranceStartDatetime.IsZero() {
		columns = append(columns, "d_insurance_datetime_start")
		values = append(values, ip.InsuranceStartDatetime)
	}
	if !ip.InsuranceEndDatetime.IsZero() {
		columns = append(columns, "d_insurance_datetime_end")
		values = append(values, ip.InsuranceEndDatetime)
	}
	if ip.InsuranceStatus != "" {
		columns = append(columns, "c_insurance_status")
		values = append(values, ip.InsuranceStatus)
	}
	columns = append(columns, "c_created_by")
	values = append(values, ip.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, ip.CreatedBy)

	return columns, values
}
