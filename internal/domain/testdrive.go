package domain

import (
	"time"

	"github.com/google/uuid"
)

type TestDrive struct {
	ID                     string    `json:"id" db:"i_id"`
	TestDriveID            string    `json:"test_drive_id" db:"i_test_drive_id"`
	KatashikiCode          string    `json:"katashiki_code" db:"c_katashiki_code"`
	Model                  string    `json:"model" db:"c_model"`
	Variant                string    `json:"variant" db:"c_variant"`
	StartTime              time.Time `json:"start_time" db:"d_test_drive_datetime_start"`
	EndTime                time.Time `json:"end_time" db:"d_test_drive_datetime_end"`
	Location               string    `json:"location" db:"c_location"`
	OutletID               string    `json:"outlet_id" db:"i_outlet_id"`
	OutletName             string    `json:"outlet_name" db:"n_outlet_name"`
	Status                 string    `json:"status" db:"c_test_drive_status"`
	Reason                 string    `json:"reason" db:"e_cancellation_reason"`
	OtherReason            string    `json:"other_reason" db:"e_other_cancellation_reason"`
	CustomerDrivingConsent bool      `json:"customer_driving_consent" db:"b_customer_driving_consent"`
	TestDriveNumber        string    `json:"test_drive_number" db:"c_test_drive_number"`
	CustomerID             string    `json:"customer_id" db:"i_customer_id"`
	LeadsID                string    `json:"leads_id" db:"i_leads_id"`
	EventID                string    `json:"event_id" db:"i_event_id"`
	RequestAt              time.Time `json:"request_at" db:"d_request_at"`
	CreatedAt              time.Time `json:"created_at" db:"d_created_at"`
	CreatedBy              string    `json:"created_by" db:"c_created_by"`
	UpdatedAt              time.Time `json:"updated_at" db:"d_updated_at"`
	UpdatedBy              string    `json:"updated_by" db:"c_updated_by"`
}

// TableName returns the database table name for the User model
func (u *TestDrive) TableName() string {
	return "dbo.tm_testdrive"
}

// Columns returns the list of database columns for the User model
func (u *TestDrive) Columns() []string {
	return []string{
		"i_id",
		"i_test_drive_id",
		"c_katashiki_code",
		"c_model",
		"c_variant",
		"d_test_drive_datetime_start",
		"d_test_drive_datetime_end",
		"c_location",
		"i_outlet_id",
		"n_outlet_name",
		"c_test_drive_status",
		"e_cancellation_reason",
		"e_other_cancellation_reason",
		"b_customer_driving_consent",
		"c_test_drive_number",
		"i_customer_id",
		"i_leads_id",
		"i_event_id",
		"d_request_at",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *TestDrive) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_test_drive_id AS NVARCHAR(36)) as i_test_drive_id",
		"c_katashiki_code",
		"c_model",
		"c_variant",
		"d_test_drive_datetime_start",
		"d_test_drive_datetime_end",
		"c_location",
		"i_outlet_id",
		"n_outlet_name",
		"c_test_drive_status",
		"e_cancellation_reason",
		"e_other_cancellation_reason",
		"b_customer_driving_consent",
		"c_test_drive_number",
		"CAST(i_customer_id AS NVARCHAR(36)) as i_customer_id",
		"CAST(i_leads_id AS NVARCHAR(36)) as i_leads_id",
		"CAST(i_event_id AS NVARCHAR(36)) as i_event_id",
		"d_request_at",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

func (u *TestDrive) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(u.Columns()))
	values = make([]interface{}, 0, len(u.Columns()))

	id := uuid.NewString()
	columns = append(columns, "i_id")
	values = append(values, id)

	if u.TestDriveID != "" {
		columns = append(columns, "i_test_drive_id")
		values = append(values, u.TestDriveID)
	}
	if u.KatashikiCode != "" {
		columns = append(columns, "c_katashiki_code")
		values = append(values, u.KatashikiCode)
	}
	if u.Model != "" {
		columns = append(columns, "c_model")
		values = append(values, u.Model)
	}
	if u.Variant != "" {
		columns = append(columns, "c_variant")
		values = append(values, u.Variant)
	}
	if !u.StartTime.IsZero() {
		columns = append(columns, "d_test_drive_datetime_start")
		values = append(values, u.StartTime)
	}
	if !u.EndTime.IsZero() {
		columns = append(columns, "d_test_drive_datetime_end")
		values = append(values, u.EndTime)
	}
	if u.Location != "" {
		columns = append(columns, "c_location")
		values = append(values, u.Location)
	}
	if u.OutletID != "" {
		columns = append(columns, "i_outlet_id")
		values = append(values, u.OutletID)
	}
	if u.OutletName != "" {
		columns = append(columns, "n_outlet_name")
		values = append(values, u.OutletName)
	}
	if u.Status != "" {
		columns = append(columns, "c_test_drive_status")
		values = append(values, u.Status)
	}
	if u.Reason != "" {
		columns = append(columns, "e_cancellation_reason")
		values = append(values, u.Reason)
	}
	if u.OtherReason != "" {
		columns = append(columns, "e_other_cancellation_reason")
		values = append(values, u.OtherReason)
	}
	columns = append(columns, "b_customer_driving_consent")
	values = append(values, u.CustomerDrivingConsent)
	if u.TestDriveNumber != "" {
		columns = append(columns, "c_test_drive_number")
		values = append(values, u.TestDriveNumber)
	}
	if u.CustomerID != "" {
		columns = append(columns, "i_customer_id")
		values = append(values, u.CustomerID)
	}
	if u.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, u.LeadsID)
	}
	if u.EventID != "" {
		columns = append(columns, "i_event_id")
		values = append(values, u.EventID)
	}
	if !u.RequestAt.IsZero() {
		columns = append(columns, "d_request_at")
		values = append(values, u.RequestAt)
	}
	if !u.CreatedAt.IsZero() {
		columns = append(columns, "d_created_at")
		values = append(values, u.CreatedAt)
	}
	if u.CreatedBy != "" {
		columns = append(columns, "c_created_by")
		values = append(values, u.CreatedBy)
	}
	if !u.UpdatedAt.IsZero() {
		columns = append(columns, "d_updated_at")
		values = append(values, u.UpdatedAt)
	}
	if u.UpdatedBy != "" {
		columns = append(columns, "c_updated_by")
		values = append(values, u.UpdatedBy)
	}
	return columns, values
}

func (u *TestDrive) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if u.KatashikiCode != "" {
		updateMap["c_katashiki_code"] = u.KatashikiCode
	}
	if u.Model != "" {
		updateMap["c_model"] = u.Model
	}
	if u.Variant != "" {
		updateMap["c_variant"] = u.Variant
	}
	if !u.StartTime.IsZero() {
		updateMap["d_test_drive_datetime_start"] = u.StartTime
	}
	if !u.EndTime.IsZero() {
		updateMap["d_test_drive_datetime_end"] = u.EndTime
	}
	if u.Location != "" {
		updateMap["c_location"] = u.Location
	}
	if u.OutletID != "" {
		updateMap["i_outlet_id"] = u.OutletID
	}
	if u.OutletName != "" {
		updateMap["n_outlet_name"] = u.OutletName
	}
	if u.Status != "" {
		updateMap["c_test_drive_status"] = u.Status
	}
	if u.Reason != "" {
		updateMap["e_cancellation_reason"] = u.Reason
	}
	if u.OtherReason != "" {
		updateMap["e_other_cancellation_reason"] = u.OtherReason
	}
	if u.CustomerID != "" {
		updateMap["i_customer_id"] = u.CustomerID
	}
	if u.LeadsID != "" {
		updateMap["i_leads_id"] = u.LeadsID
	}
	if !u.RequestAt.IsZero() {
		updateMap["d_request_at"] = u.RequestAt
	}
	if u.TestDriveNumber != "" {
		updateMap["c_test_drive_number"] = u.TestDriveNumber
	}
	updateMap["b_customer_driving_consent"] = u.CustomerDrivingConsent
	updateMap["d_updated_at"] = u.UpdatedAt
	updateMap["c_updated_by"] = u.UpdatedBy
	return updateMap
}
