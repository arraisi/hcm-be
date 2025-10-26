package domain

import "time"

// OneAccount represents the one account information from the webhook
type OneAccount struct {
	OneAccountID string `json:"one_account_ID" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Gender       string `json:"gender" validate:"required,oneof=MALE FEMALE"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Email        string `json:"email" validate:"omitempty,email"`
}

// TestDrive represents the test drive information from the webhook
type TestDrive struct {
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

// Leads represents the leads information from the webhook
type Leads struct {
	LeadsID                         string  `json:"leads_ID" validate:"required"`
	LeadsType                       string  `json:"leads_type" validate:"required,eq=TEST_DRIVE_REQUEST"`
	LeadsFollowUpStatus             string  `json:"leads_follow_up_status" validate:"required"`
	LeadsPreferenceContactTimeStart string  `json:"leads_preference_contact_time_start" validate:"required"`
	LeadsPreferenceContactTimeEnd   string  `json:"leads_preference_contact_time_end" validate:"required"`
	LeadsSource                     string  `json:"leads_source" validate:"required"`
	AdditionalNotes                 *string `json:"additional_notes"`
}

// ScoreParameter represents the parameter information in score
type ScoreParameter struct {
	PurchasePlanCriteria    string `json:"purchase_plan_criteria" validate:"required"`
	PaymentPreferCriteria   string `json:"payment_prefer_criteria" validate:"required"`
	NegotiationCriteria     string `json:"negotiation_criteria" validate:"required"`
	TestDriveCriteria       string `json:"test_drive_criteria" validate:"required"`
	TradeInCriteria         string `json:"trade_in_criteria" validate:"required"`
	BrowsingHistoryCriteria string `json:"browsing_history_criteria" validate:"required"`
	VehicleAgeCriteria      string `json:"vehicle_age_criteria" validate:"required"`
}

// Score represents the score information from the webhook
type Score struct {
	IAMLeadScore    string         `json:"iam_lead_score" validate:"required"`
	OutletLeadScore string         `json:"outlet_lead_score" validate:"required"`
	Parameter       ScoreParameter `json:"parameter" validate:"required"`
}

// BookingEventData represents the data payload of the booking event
type BookingEventData struct {
	OneAccount OneAccount `json:"one_account" validate:"required"`
	TestDrive  TestDrive  `json:"test_drive" validate:"required"`
	Leads      Leads      `json:"leads" validate:"required"`
	Score      Score      `json:"score" validate:"required"`
}

// BookingEvent represents the complete webhook payload for test drive booking
type BookingEvent struct {
	Process   string           `json:"process" validate:"required,eq=test drive request"`
	EventID   string           `json:"event_ID" validate:"required,uuid4"`
	Timestamp int64            `json:"timestamp" validate:"required"`
	Data      BookingEventData `json:"data" validate:"required"`
}

// GetEventTimestamp returns the timestamp as time.Time
func (be *BookingEvent) GetEventTimestamp() time.Time {
	return time.Unix(be.Timestamp, 0)
}

// GetTestDriveCreatedTime returns the test drive created datetime as time.Time
func (td *TestDrive) GetCreatedTime() time.Time {
	return time.Unix(td.CreatedDatetime, 0)
}

// GetTestDriveStartTime returns the test drive start datetime as time.Time
func (td *TestDrive) GetStartTime() time.Time {
	return time.Unix(td.TestDriveDatetimeStart, 0)
}

// GetTestDriveEndTime returns the test drive end datetime as time.Time
func (td *TestDrive) GetEndTime() time.Time {
	return time.Unix(td.TestDriveDatetimeEnd, 0)
}
