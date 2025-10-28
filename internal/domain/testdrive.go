package domain

import "time"

type TestDrive struct {
	ID              string    `json:"id" db:"id"`
	TestDriveID     string    `json:"test_drive_ID" db:"test_drive_id"`
	TestDriveNumber string    `json:"test_drive_number" db:"test_drive_number"`
	KatashikiCode   string    `json:"katashiki_code" db:"katashiki_code"`
	Model           string    `json:"model" db:"model"`
	Variant         string    `json:"variant" db:"variant"`
	RequestAt       time.Time `json:"created_datetime" db:"request_at"`
	StartTime       time.Time `json:"test_drive_datetime_start" db:"test_drive_datetime_start"`
	EndTime         time.Time `json:"test_drive_datetime_end" db:"test_drive_datetime_end"`
	Location        string    `json:"location" db:"location"`
	OutletID        string    `json:"outlet_ID" db:"outlet_id"`
	OutletName      string    `json:"outlet_name" db:"outlet_name"`
	Status          string    `json:"test_drive_status" db:"test_drive_status"`
	Reason          string    `json:"cancellation_reason" db:"cancellation_reason"`
	OtherReason     string    `json:"other_cancellation_reason" db:"other_cancellation_reason"`
	Consent         bool      `json:"customer_driving_consent" db:"customer_driving_consent"`
	CustomerID      string    `json:"customer_id" db:"customer_id"`
	LeadsID         string    `json:"leads_id" db:"leads_id"`
	EventID         string    `json:"event_id" db:"event_id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	CreatedBy       string    `json:"created_by" db:"created_by"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy       *string   `json:"updated_by" db:"updated_by"`
}

// TableName returns the database table name for the User model
func (u *TestDrive) TableName() string {
	return "dbo.tm_testdrive"
}

// Columns returns the list of database columns for the User model
func (u *TestDrive) Columns() []string {
	return []string{
		"id",
		"test_drive_id",
		"test_drive_number",
		"katashiki_code",
		"model",
		"variant",
		"test_drive_datetime_start",
		"test_drive_datetime_end",
		"location",
		"outlet_id",
		"outlet_name",
		"test_drive_status",
		"cancellation_reason",
		"other_cancellation_reason",
		"request_at",
		"customer_driving_consent",
		"customer_id",
		"leads_id",
		"event_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	}
}

func (u *TestDrive) ToValues() []interface{} {
	return []interface{}{
		u.ID,
		u.TestDriveID,
		u.TestDriveNumber,
		u.KatashikiCode,
		u.Model,
		u.Variant,
		u.StartTime,
		u.EndTime,
		u.Location,
		u.OutletID,
		u.OutletName,
		u.Status,
		u.Reason,
		u.OtherReason,
		u.RequestAt,
		u.Consent,
		u.CustomerID,
		u.LeadsID,
		u.EventID,
		u.CreatedAt,
		u.CreatedBy,
		u.UpdatedAt,
		u.UpdatedBy,
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *TestDrive) SelectColumns() []string {
	return []string{
		"CAST(id AS NVARCHAR(36)) as id",
		"CAST(test_drive_id AS NVARCHAR(36)) as test_drive_id",
		"katashiki_code",
		"model",
		"variant",
		"test_drive_datetime_start",
		"test_drive_datetime_end",
		"location",
		"outlet_id",
		"outlet_name",
		"test_drive_status",
		"cancellation_reason",
		"other_cancellation_reason",
		"customer_driving_consent",
		"test_drive_number",
		"CAST(customer_id AS NVARCHAR(36)) as customer_id",
		"CAST(leads_id AS NVARCHAR(36)) as leads_id",
		"CAST(event_id AS NVARCHAR(36)) as event_id",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"request_at",
	}
}

func (u *TestDrive) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.KatashikiCode != "" {
		updateMap["katashiki_code"] = u.KatashikiCode
	}
	if u.Model != "" {
		updateMap["model"] = u.Model
	}
	if u.Variant != "" {
		updateMap["variant"] = u.Variant
	}
	if !u.StartTime.IsZero() {
		updateMap["test_drive_datetime_start"] = u.StartTime
	}
	if !u.EndTime.IsZero() {
		updateMap["test_drive_datetime_end"] = u.EndTime
	}
	if u.Location != "" {
		updateMap["location"] = u.Location
	}
	if u.OutletID != "" {
		updateMap["outlet_id"] = u.OutletID
	}
	if u.OutletName != "" {
		updateMap["outlet_name"] = u.OutletName
	}
	if u.Status != "" {
		updateMap["test_drive_status"] = u.Status
	}
	if u.Reason != "" {
		updateMap["cancellation_reason"] = u.Reason
	}
	if u.OtherReason != "" {
		updateMap["other_cancellation_reason"] = u.OtherReason
	}
	if u.CustomerID != "" {
		updateMap["customer_id"] = u.CustomerID
	}
	if u.LeadsID != "" {
		updateMap["leads_id"] = u.LeadsID
	}
	if !u.RequestAt.IsZero() {
		updateMap["request_at"] = u.RequestAt
	}
	if u.TestDriveNumber != "" {
		updateMap["test_drive_number"] = u.TestDriveNumber
	}
	updateMap["customer_driving_consent"] = u.Consent
	updateMap["update_at"] = u.UpdatedAt
	updateMap["updated_by"] = u.UpdatedBy
	return updateMap
}
