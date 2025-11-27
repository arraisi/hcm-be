package domain

import (
	"time"
)

// SPK represents Surat Pesanan Kendaraan (Vehicle Order Letter)
type SPK struct {
	ID                      string     `json:"id" db:"i_id"`
	SPKID                   string     `json:"spk_id" db:"i_spk_id"`
	SPKNumber               string     `json:"spk_number" db:"c_spk_number"`
	LeadsID                 string     `json:"leads_id" db:"i_leads_id"`
	CreatedDatetime         time.Time  `json:"created_datetime" db:"d_created_datetime"`
	SPKStatus               string     `json:"spk_status" db:"c_spk_status"`
	Model                   string     `json:"model" db:"c_model"`
	Variant                 string     `json:"variant" db:"c_variant"`
	KatashikiSuffix         string     `json:"katashiki_suffix" db:"c_katashiki_suffix"`
	Year                    int        `json:"year" db:"n_year"`
	OutletID                string     `json:"outlet_id" db:"i_outlet_id"`
	OutletName              string     `json:"outlet_name" db:"n_outlet_name"`
	EmployeeID              string     `json:"employee_id" db:"i_employee_id"`
	EmployeeFirstName       string     `json:"employee_first_name" db:"n_employee_first_name"`
	EmployeeLastName        string     `json:"employee_last_name" db:"n_employee_last_name"`
	SPKCustomerConfirmation bool       `json:"spk_customer_confirmation" db:"b_spk_customer_confirmation"`
	SPKApprovedDatetime     *time.Time `json:"spk_approved_datetime" db:"d_spk_approved_datetime"`
	SPKCancelledDatetime    *time.Time `json:"spk_cancelled_datetime" db:"d_spk_cancelled_datetime"`
	SPKCancelledReason      *string    `json:"spk_cancelled_reason" db:"e_spk_cancelled_reason"`
	CreatedAt               time.Time  `json:"created_at" db:"d_created_at"`
	UpdatedAt               time.Time  `json:"updated_at" db:"d_updated_at"`
}

// TableName returns the database table name for the SPK model
func (s *SPK) TableName() string {
	return "dbo.tm_spk"
}

// Columns returns the list of database columns for the SPK model
func (s *SPK) Columns() []string {
	return []string{
		"i_spk_id",
		"c_spk_number",
		"i_leads_id",
		"d_created_datetime",
		"c_spk_status",
		"c_model",
		"c_variant",
		"c_katashiki_suffix",
		"n_year",
		"i_outlet_ID",
		"n_outlet_name",
		"i_employee_ID",
		"n_employee_first_name",
		"n_employee_last_name",
		"b_spk_customer_confirmation",
		"d_spk_approved_datetime",
		"d_spk_cancelled_datetime",
		"e_spk_cancelled_reason",
		"d_created_at",
		"d_updated_at",
	}
}

// SelectColumns returns the list of columns to select in queries for the SPK model
func (s *SPK) SelectColumns() []string {
	return []string{
		"i_id",
		"i_spk_id",
		"c_spk_number",
		"i_leads_id",
		"d_created_datetime",
		"c_spk_status",
		"c_model",
		"c_variant",
		"c_katashiki_suffix",
		"n_year",
		"i_outlet_ID",
		"n_outlet_name",
		"i_employee_ID",
		"n_employee_first_name",
		"n_employee_last_name",
		"b_spk_customer_confirmation",
		"d_spk_approved_datetime",
		"d_spk_cancelled_datetime",
		"e_spk_cancelled_reason",
		"d_created_at",
		"d_updated_at",
	}
}

func (s *SPK) ToCreateMap() (columns []string, values []interface{}) {
	now := time.Now()
	columns = s.Columns()
	values = []interface{}{
		s.SPKID,
		s.SPKNumber,
		s.LeadsID,
		s.CreatedDatetime,
		s.SPKStatus,
		s.Model,
		s.Variant,
		s.KatashikiSuffix,
		s.Year,
		s.OutletID,
		s.OutletName,
		s.EmployeeID,
		s.EmployeeFirstName,
		s.EmployeeLastName,
		s.SPKCustomerConfirmation,
		s.SPKApprovedDatetime,
		s.SPKCancelledDatetime,
		s.SPKCancelledReason,
		now,
		now,
	}
	return
}

func (s *SPK) ToUpdateMap() map[string]interface{} {
	updateMap := map[string]interface{}{}

	if s.SPKID != "" {
		updateMap["i_spk_id"] = s.SPKID
	}
	if s.SPKNumber != "" {
		updateMap["c_spk_number"] = s.SPKNumber
	}
	if s.LeadsID != "" {
		updateMap["i_leads_id"] = s.LeadsID
	}
	if !s.CreatedDatetime.IsZero() {
		updateMap["d_created_datetime"] = s.CreatedDatetime
	}
	if s.SPKStatus != "" {
		updateMap["c_spk_status"] = s.SPKStatus
	}
	if s.Model != "" {
		updateMap["c_model"] = s.Model
	}
	if s.Variant != "" {
		updateMap["c_variant"] = s.Variant
	}
	if s.KatashikiSuffix != "" {
		updateMap["c_katashiki_suffix"] = s.KatashikiSuffix
	}
	if s.Year != 0 {
		updateMap["n_year"] = s.Year
	}
	if s.OutletID != "" {
		updateMap["i_outlet_ID"] = s.OutletID
	}
	if s.OutletName != "" {
		updateMap["n_outlet_name"] = s.OutletName
	}
	if s.EmployeeID != "" {
		updateMap["i_employee_ID"] = s.EmployeeID
	}
	if s.EmployeeFirstName != "" {
		updateMap["n_employee_first_name"] = s.EmployeeFirstName
	}
	if s.EmployeeLastName != "" {
		updateMap["n_employee_last_name"] = s.EmployeeLastName
	}
	updateMap["b_spk_customer_confirmation"] = s.SPKCustomerConfirmation
	if s.SPKApprovedDatetime != nil {
		updateMap["d_spk_approved_datetime"] = s.SPKApprovedDatetime
	}
	if s.SPKCancelledDatetime != nil {
		updateMap["d_spk_cancelled_datetime"] = s.SPKCancelledDatetime
	}
	if s.SPKCancelledReason != nil {
		updateMap["e_spk_cancelled_reason"] = s.SPKCancelledReason
	}
	updateMap["d_updated_at"] = time.Now()

	return updateMap
}
