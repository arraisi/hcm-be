package domain

type TestDrive struct {
	IID         string `json:"i_id" db:"i_id"`
	TestDriveID string `json:"test_drive_ID" db:"test_drive_ID"`
	Katashiki   string `json:"katashiki_code" db:"katashiki_code"`
	Model       string `json:"model" db:"model"`
	Variant     string `json:"variant" db:"variant"`
	StartTime   int64  `json:"test_drive_datetime_start" db:"test_drive_datetime_start"`
	EndTime     int64  `json:"test_drive_datetime_end" db:"test_drive_datetime_end"`
	Location    string `json:"location" db:"location"`
	OutletID    string `json:"outlet_ID" db:"outlet_ID"`
	OutletName  string `json:"outlet_name" db:"outlet_name"`
	Status      string `json:"test_drive_status" db:"test_drive_status"`
	Reason      string `json:"cancellation_reason" db:"cancellation_reason"`
	OtherReason string `json:"other_cancellation_reason" db:"other_cancellation_reason"`
	CreatedAt   int64  `json:"created_datetime" db:"created_datetime"`
	Consent     bool   `json:"customer_driving_consent" db:"customer_driving_consent"`
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
		"outlet_ID",
		"outlet_name",
		"test_drive_status",
		"cancellation_reason",
		"other_cancellation_reason",
		"created_datetime",
		"customer_driving_consent",
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
		"outlet_ID",
		"outlet_name",
		"test_drive_status",
		"cancellation_reason",
		"other_cancellation_reason",
		"created_datetime",
		"customer_driving_consent",
	}
}
