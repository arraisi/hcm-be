package domain

import (
	"time"

	"github.com/lib/pq"
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
	Deleted                bool      `db:"b_deleted"`
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
		"b_deleted",
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

// ServiceBookingVehicleInsurancePolicy represents individual insurance policy
type ServiceBookingVehicleInsurancePolicy struct {
	ID                     string         `db:"i_id"`
	ServiceBookingID       string         `db:"i_service_booking_id"`
	VehicleInsuranceID     string         `db:"i_vehicle_insurance_id"`
	InsuranceType          string         `db:"c_insurance_type"`
	InsuranceCoverage      pq.StringArray `db:"e_insurance_coverage"`
	InsuranceStartDatetime time.Time      `db:"d_insurance_start_datetime"`
	InsuranceEndDatetime   time.Time      `db:"d_insurance_end_datetime"`
	InsuranceStatus        string         `db:"c_insurance_status"`
	CreatedAt              time.Time      `db:"d_created_at"`
	CreatedBy              string         `db:"c_created_by"`
	UpdatedAt              time.Time      `db:"d_updated_at"`
	UpdatedBy              string         `db:"c_updated_by"`
	Deleted                bool           `db:"b_deleted"`
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
		"d_insurance_start_datetime",
		"d_insurance_end_datetime",
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
		"d_insurance_start_datetime",
		"d_insurance_end_datetime",
		"c_insurance_status",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
		"b_deleted",
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
		columns = append(columns, "d_insurance_start_datetime")
		values = append(values, ip.InsuranceStartDatetime)
	}
	if !ip.InsuranceEndDatetime.IsZero() {
		columns = append(columns, "d_insurance_end_datetime")
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

// ServiceBookingDamageImage represents damage image URLs for service booking
type ServiceBookingDamageImage struct {
	ID               string    `db:"i_id"`
	ServiceBookingID string    `db:"i_service_booking_id"`
	ImageURL         string    `db:"e_image_url"`
	CreatedAt        time.Time `db:"d_created_at"`
	CreatedBy        string    `db:"c_created_by"`
	UpdatedAt        time.Time `db:"d_updated_at"`
	UpdatedBy        string    `db:"c_updated_by"`
	Deleted          bool      `db:"b_deleted"`
}

// TableName returns the database table name
func (di *ServiceBookingDamageImage) TableName() string {
	return "dbo.tr_service_booking_damage_image"
}

// Columns returns the list of database columns
func (di *ServiceBookingDamageImage) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"e_image_url",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select
func (di *ServiceBookingDamageImage) SelectColumns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"e_image_url",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
		"b_deleted",
	}
}

// ToCreateMap prepares the columns and values for inserting
func (di *ServiceBookingDamageImage) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	if di.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, di.ServiceBookingID)
	}
	if di.ImageURL != "" {
		columns = append(columns, "e_image_url")
		values = append(values, di.ImageURL)
	}
	columns = append(columns, "c_created_by")
	values = append(values, di.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, di.CreatedBy)

	return columns, values
}
