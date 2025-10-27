package domain

import "time"

type TestDrive struct {
	IID          string    `json:"i_id" db:"i_id"`
	TestDriveID  string    `json:"test_drive_ID" db:"test_drive_ID"`
	Katashiki    string    `json:"katashiki_code" db:"katashiki_code"`
	Model        string    `json:"model" db:"model"`
	Variant      string    `json:"variant" db:"variant"`
	StartTime    time.Time `json:"test_drive_datetime_start" db:"test_drive_datetime_start"`
	EndTime      time.Time `json:"test_drive_datetime_end" db:"test_drive_datetime_end"`
	Location     string    `json:"location" db:"location"`
	OutletID     string    `json:"outlet_ID" db:"outled_ID"`
	OutletName   string    `json:"outlet_name" db:"outlet_name"`
	Status       string    `json:"test_drive_status" db:"test_drive_status"`
	Reason       string    `json:"cancellation_reason" db:"cancellation_reason"`
	OtherReason  string    `json:"other_cancellation_reason" db:"other_cancellation_reason"`
	CreatedAt    time.Time `json:"created_datetime" db:"created_datetime"`
	Consent      bool      `json:"customer_driving_consent" db:"customer_driving_consent"`
	CustomerID   string    `json:"customer_id" db:"customer_id"`
	OneAccountID string    `json:"one_account_ID" db:"one_account_ID"`
	LeadsID      string    `json:"leads_id" db:"leads_id"`
}

// TableName returns the database table name for the User model
func (u *TestDrive) TableName() string {
	return "dbo.tm_testdrive"
}

// Columns returns the list of database columns for the User model
func (u *TestDrive) Columns() []string {
	return []string{
		"i_id",
		"test_drive_ID",
		"katashiki_code",
		"model",
		"variant",
		"test_drive_datetime_start",
		"test_drive_datetime_end",
		"location",
		"outled_ID",
		"outlet_name",
		"test_drive_status",
		"cancellation_reason",
		"other_cancellation_reason",
		"created_datetime",
		"customer_driving_consent",
		"customer_id",
		"one_account_ID",
		"leads_id",
	}
}

func (u *TestDrive) ToValues() []interface{} {
	return []interface{}{
		u.IID,
		u.TestDriveID,
		u.Katashiki,
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
		u.CreatedAt,
		u.Consent,
		u.CustomerID,
		u.OneAccountID,
		u.LeadsID,
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *TestDrive) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(test_drive_ID AS NVARCHAR(36)) as test_drive_ID",
		"katashiki_code",
		"model",
		"variant",
		"test_drive_datetime_start",
		"test_drive_datetime_end",
		"location",
		"outled_ID",
		"outlet_name",
		"test_drive_status",
		"cancellation_reason",
		"other_cancellation_reason",
		"created_datetime",
		"customer_driving_consent",
		"CAST(customer_id AS NVARCHAR(36)) as customer_id",
		"CAST(one_account_ID AS NVARCHAR(36)) as one_account_ID",
		"CAST(leads_id AS NVARCHAR(36)) as leads_id",
	}
}

func (u *TestDrive) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.Katashiki != "" {
		updateMap["katashiki_code"] = u.Katashiki
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
		updateMap["outled_ID"] = u.OutletID
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
	if !u.CreatedAt.IsZero() {
		updateMap["created_datetime"] = u.CreatedAt
	}
	if u.CustomerID != "" {
		updateMap["customer_id"] = u.CustomerID
	}
	if u.OneAccountID != "" {
		updateMap["one_account_ID"] = u.OneAccountID
	}
	if u.LeadsID != "" {
		updateMap["leads_id"] = u.LeadsID
	}
	updateMap["customer_driving_consent"] = u.Consent
	return updateMap
}
