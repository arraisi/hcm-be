package domain

import (
	"time"
)

// ServiceBookingVehicleInsurance represents vehicle insurance information for service booking
type ServiceBookingVehicleInsurance struct {
	ID                     string    `db:"i_id"`
	ServiceBookingID       string    `db:"i_service_booking_id"`
	InsuranceProvider      string    `db:"c_insurance_provider"`
	InsuranceProviderOther string    `db:"c_insurance_provider_other"`
	InsurancePolicyNumber  string    `db:"c_insurance_policy_number"`
	CreatedAt              time.Time `db:"d_created_at"`
	CreatedBy              string    `db:"c_created_by"`
	UpdatedAt              time.Time `db:"d_updated_at"`
	UpdatedBy              string    `db:"c_updated_by"`
}

// TableName returns the database table name
func (vi *ServiceBookingVehicleInsurance) TableName() string {
	return "dbo.tr_service_booking_vehicle_insurance"
}

// Columns returns the list of database columns
func (vi *ServiceBookingVehicleInsurance) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"c_insurance_provider",
		"c_insurance_provider_other",
		"c_insurance_policy_number",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select
func (vi *ServiceBookingVehicleInsurance) SelectColumns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"c_insurance_provider",
		"c_insurance_provider_other",
		"c_insurance_policy_number",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting
func (vi *ServiceBookingVehicleInsurance) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	if vi.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, vi.ServiceBookingID)
	}
	if vi.InsuranceProvider != "" {
		columns = append(columns, "c_insurance_provider")
		values = append(values, vi.InsuranceProvider)
	}
	if vi.InsuranceProviderOther != "" {
		columns = append(columns, "c_insurance_provider_other")
		values = append(values, vi.InsuranceProviderOther)
	}
	if vi.InsurancePolicyNumber != "" {
		columns = append(columns, "c_insurance_policy_number")
		values = append(values, vi.InsurancePolicyNumber)
	}
	columns = append(columns, "c_created_by")
	values = append(values, vi.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, vi.CreatedBy)

	return columns, values
}
