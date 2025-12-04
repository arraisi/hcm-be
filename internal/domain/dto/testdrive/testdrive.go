package testdrive

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

// TestDriveRequest represents the test drive information from the webhook
type TestDriveRequest struct {
	TestDriveID             string  `json:"test_drive_id" validate:"required"`
	TestDriveNumber         string  `json:"test_drive_number" validate:"required"`
	KatashikiCode           string  `json:"katashiki_code" validate:"required"`
	Model                   string  `json:"model" validate:"required"`
	Variant                 string  `json:"variant" validate:"required"`
	CreatedDatetime         int64   `json:"created_datetime" validate:"required"`
	TestDriveDatetimeStart  int64   `json:"test_drive_datetime_start" validate:"required"`
	TestDriveDatetimeEnd    int64   `json:"test_drive_datetime_end" validate:"required"`
	Location                string  `json:"location" validate:"required,oneof=HOME DEALER"`
	OutletID                string  `json:"outlet_ID" validate:"required"`
	OutletName              string  `json:"outlet_name" validate:"required"`
	TestDriveStatus         string  `json:"test_drive_status" validate:"required,oneof=SUBMITTED CHANGE_REQUEST CANCEL_SUBMITTED"`
	CancellationReason      *string `json:"cancellation_reason"`
	OtherCancellationReason *string `json:"other_cancellation_reason"`
	CustomerDrivingConsent  bool    `json:"customer_driving_consent"`
}

func NewTestDriveRequest(td domain.TestDrive) TestDriveRequest {
	return TestDriveRequest{
		TestDriveID:             td.TestDriveID,
		TestDriveNumber:         td.TestDriveNumber,
		KatashikiCode:           td.KatashikiCode,
		Model:                   td.Model,
		Variant:                 td.Variant,
		CreatedDatetime:         td.RequestAt.Unix(),
		TestDriveDatetimeStart:  td.StartTime.Unix(),
		TestDriveDatetimeEnd:    td.EndTime.Unix(),
		Location:                td.Location,
		OutletID:                td.OutletID,
		OutletName:              td.OutletName,
		TestDriveStatus:         td.Status,
		CancellationReason:      utils.ToPointer(td.Reason),
		OtherCancellationReason: utils.ToPointer(td.OtherReason),
		CustomerDrivingConsent:  td.CustomerDrivingConsent,
	}

}

// TestDriveEventData represents the data payload of the booking event
type TestDriveEventData struct {
	OneAccount    customer.OneAccountRequest `json:"one_account" validate:"required"`
	TestDrive     TestDriveRequest           `json:"test_drive" validate:"required"`
	Leads         leads.LeadsRequest         `json:"leads" validate:"required"`
	PICAssignment *PICAssignmentRequest      `json:"pic_assignment,omitempty"`
}

// PICAssignmentRequest represents the PIC assignment information from the webhook
type PICAssignmentRequest struct {
	EmployeeID string `json:"employee_ID"`
	FirstName  string `json:"first_name" `
	LastName   string `json:"last_name"`

	NIK string `json:"nik"`
}

// TestDriveEvent represents the complete webhook payload for test drive booking
type TestDriveEvent struct {
	Process   string             `json:"process" validate:"required"`
	EventID   string             `json:"event_ID" validate:"required,uuid4"`
	Timestamp int64              `json:"timestamp" validate:"required"`
	Data      TestDriveEventData `json:"data" validate:"required"`
}

// ToTestDriveModel converts the TestDriveEvent to the internal TestDrive model
func (be *TestDriveEvent) ToTestDriveModel(customerID string) domain.TestDrive {
	return domain.TestDrive{
		TestDriveID:            be.Data.TestDrive.TestDriveID,
		TestDriveNumber:        be.Data.TestDrive.TestDriveNumber,
		KatashikiCode:          be.Data.TestDrive.KatashikiCode,
		Model:                  be.Data.TestDrive.Model,
		Variant:                be.Data.TestDrive.Variant,
		RequestAt:              utils.GetTimeUnix(be.Data.TestDrive.CreatedDatetime),
		StartTime:              utils.GetTimeUnix(be.Data.TestDrive.TestDriveDatetimeStart),
		EndTime:                utils.GetTimeUnix(be.Data.TestDrive.TestDriveDatetimeEnd),
		Location:               be.Data.TestDrive.Location,
		OutletID:               be.Data.TestDrive.OutletID,
		OutletName:             be.Data.TestDrive.OutletName,
		Status:                 be.Data.TestDrive.TestDriveStatus,
		Reason:                 utils.ToValue(be.Data.TestDrive.CancellationReason),
		OtherReason:            utils.ToValue(be.Data.TestDrive.OtherCancellationReason),
		CustomerDrivingConsent: be.Data.TestDrive.CustomerDrivingConsent,
		CustomerID:             customerID,
		LeadsID:                utils.ToPointer(be.Data.Leads.LeadsID),
		EventID:                be.EventID,
		CreatedAt:              time.Now(),
		CreatedBy:              constants.System,
		UpdatedAt:              time.Now(),
		UpdatedBy:              constants.System, // or fetch from context if available
	}
}

type GetTestDriveRequest struct {
	ID           *string
	TestDriveID  *string
	CustomerID   *string
	OneAccountID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetTestDriveRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.TestDriveID != nil {
		q.Where(sqrl.Eq{"i_test_drive_id": req.TestDriveID})
	}
	if req.CustomerID != nil {
		q.Where(sqrl.Eq{"i_customer_id": req.CustomerID})
	}
	if req.OneAccountID != nil {
		q.Where(sqrl.Eq{"i_one_account_id": req.OneAccountID})
	}
}

type ConfirmTestDriveBookingRequest struct {
	TestDriveID         string `json:"test_drive_id" validate:"required"`
	EmployeeID          string `json:"employee_id" validate:"required"`
	TestDriveStatus     string `json:"test_drive_status" validate:"required,oneof=CONFIRMED CANCELLED COMPLETED NOT_SHOW"`
	LeadsType           string `json:"leads_type,omitempty"`
	LeadsFollowUpStatus string `json:"leads_follow_up_status" validate:"required,oneof=NOT_YET_FOLLOWED_UP ON_CONSIDERATION NO_RESPONSE"`
}
