package testdrive

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/lead"
)

// TestDriveRequest represents the test drive information from the webhook
type TestDriveRequest struct {
	TestDriveID             string  `json:"test_drive_ID" validate:"required"`
	TestDriveNumber         string  `json:"test_drive_number" validate:"required"`
	KatashikiCode           string  `json:"katashiki_code" validate:"required"`
	Model                   string  `json:"model" validate:"required"`
	Variant                 string  `json:"variant" validate:"required"`
	CreatedDatetime         int64   `json:"created_datetime" validate:"required"`
	TestDriveDatetimeStart  int64   `json:"test_drive_datetime_start" validate:"required"`
	TestDriveDatetimeEnd    int64   `json:"test_drive_datetime_end" validate:"required"`
	Location                string  `json:"location" validate:"required"`
	OutletID                string  `json:"outlet_ID" validate:"required"`
	OutletName              string  `json:"outlet_name" validate:"required"`
	TestDriveStatus         string  `json:"test_drive_status" validate:"required,oneof=SUBMITTED CHANGE_REQUEST CANCEL_SUBMITTED"`
	CancellationReason      *string `json:"cancellation_reason"`
	OtherCancellationReason *string `json:"other_cancellation_reason"`
	CustomerDrivingConsent  bool    `json:"customer_driving_consent"`
}

// TestDriveEventData represents the data payload of the booking event
type TestDriveEventData struct {
	OneAccount customer.OneAccountRequest `json:"one_account" validate:"required"`
	TestDrive  TestDriveRequest           `json:"test_drive" validate:"required"`
	Leads      lead.LeadsRequest          `json:"leads" validate:"required"`
	Score      lead.Score                 `json:"score" validate:"required"`
}

// TestDriveEvent represents the complete webhook payload for test drive booking
type TestDriveEvent struct {
	Process   string             `json:"process" validate:"required,eq=test drive request"`
	EventID   string             `json:"event_ID" validate:"required,uuid4"`
	Timestamp int64              `json:"timestamp" validate:"required"`
	Data      TestDriveEventData `json:"data" validate:"required"`
}

// GetEventTimestamp returns the timestamp as time.Time
func (be *TestDriveEvent) GetEventTimestamp() time.Time {
	return time.Unix(be.Timestamp, 0)
}

func (td *TestDriveRequest) GetCreatedTime() time.Time {
	return time.Unix(td.CreatedDatetime, 0)
}

func (td *TestDriveRequest) GetStartTime() time.Time {
	return time.Unix(td.TestDriveDatetimeStart, 0)
}

func (td *TestDriveRequest) GetEndTime() time.Time {
	return time.Unix(td.TestDriveDatetimeEnd, 0)
}
