package testdrive

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
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
	Leads      leads.LeadsRequest         `json:"leads" validate:"required"`
	Score      leads.Score                `json:"score" validate:"required"`
}

// TestDriveEvent represents the complete webhook payload for test drive booking
type TestDriveEvent struct {
	Process   string             `json:"process" validate:"required"`
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

// ToTestDriveModel converts the TestDriveEvent to the internal TestDrive model
func (be *TestDriveEvent) ToTestDriveModel(customerID string) domain.TestDrive {
	return domain.TestDrive{
		TestDriveID:  be.Data.TestDrive.TestDriveID,
		Model:        be.Data.TestDrive.Model,
		Variant:      be.Data.TestDrive.Variant,
		CreatedAt:    be.Data.TestDrive.GetCreatedTime(),
		StartTime:    be.Data.TestDrive.GetStartTime(),
		EndTime:      be.Data.TestDrive.GetEndTime(),
		Location:     be.Data.TestDrive.Location,
		OutletID:     be.Data.TestDrive.OutletID,
		OutletName:   be.Data.TestDrive.OutletName,
		Status:       be.Data.TestDrive.TestDriveStatus,
		Reason:       utils.ToValue(be.Data.TestDrive.CancellationReason),
		OtherReason:  utils.ToValue(be.Data.TestDrive.OtherCancellationReason),
		Consent:      be.Data.TestDrive.CustomerDrivingConsent,
		OneAccountID: be.Data.OneAccount.OneAccountID,
		CustomerID:   customerID,
		LeadsID:      be.Data.Leads.LeadsID,
	}
}

// ToCustomerModel converts the TestDriveEvent to the internal Customer model
func (be *TestDriveEvent) ToCustomerModel() domain.Customer {
	return domain.Customer{
		OneAccountID: be.Data.OneAccount.OneAccountID,
		FirstName:    be.Data.OneAccount.FirstName,
		LastName:     be.Data.OneAccount.LastName,
		Email:        be.Data.OneAccount.Email,
		PhoneNumber:  be.Data.OneAccount.PhoneNumber,
		Gender:       be.Data.OneAccount.Gender,
	}
}

// ToLeadsModel converts the TestDriveEvent to the internal Leads model
func (be *TestDriveEvent) ToLeadsModel() domain.Leads {
	return domain.Leads{
		LeadsID:                         be.Data.Leads.LeadsID,
		LeadsType:                       be.Data.Leads.LeadsType,
		LeadsFollowUpStatus:             be.Data.Leads.LeadsFollowUpStatus,
		LeadsPreferenceContactTimeStart: be.Data.Leads.LeadsPreferenceContactTimeStart,
		LeadsPreferenceContactTimeEnd:   be.Data.Leads.LeadsPreferenceContactTimeEnd,
		LeadSource:                      be.Data.Leads.LeadSource,
		AdditionalNotes:                 be.Data.Leads.AdditionalNotes,
	}
}

// ToLeadScoreModel converts the TestDriveEvent to the internal LeadScore model
func (be *TestDriveEvent) ToLeadScoreModel() domain.LeadScore {
	return domain.LeadScore{
		IID:                     be.Data.Leads.LeadsID,
		TAMLeadScore:            be.Data.Score.TAMLeadScore,
		OutletLeadScore:         be.Data.Score.OutletLeadScore,
		PurchasePlanCriteria:    be.Data.Score.Parameter.PurchasePlanCriteria,
		PaymentPreferCriteria:   be.Data.Score.Parameter.PaymentPreferCriteria,
		NegotiationCriteria:     be.Data.Score.Parameter.NegotiationCriteria,
		TestDriveCriteria:       be.Data.Score.Parameter.TestDriveCriteria,
		TradeInCriteria:         be.Data.Score.Parameter.TradeInCriteria,
		BrowsingHistoryCriteria: be.Data.Score.Parameter.BrowsingHistoryCriteria,
		VehicleAgeCriteria:      be.Data.Score.Parameter.VehicleAgeCriteria,
	}
}

type GetTestDriveRequest struct {
	IID          *string
	TestDriveID  *string
	CustomerID   *string
	OneAccountID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetTestDriveRequest) Apply(q *sqrl.SelectBuilder) {
	if req.IID != nil {
		q.Where(sqrl.Eq{"i_id": req.IID})
	}
	if req.TestDriveID != nil {
		q.Where(sqrl.Eq{"test_drive_ID": req.TestDriveID})
	}
	if req.CustomerID != nil {
		q.Where(sqrl.Eq{"customer_ID": req.CustomerID})
	}
	if req.OneAccountID != nil {
		q.Where(sqrl.Eq{"one_account_ID": req.OneAccountID})
	}
}
