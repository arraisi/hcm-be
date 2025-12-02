package appraisalbooking

// =======================
// Top-level event
// =======================

type EventRequest struct {
	Process   string    `json:"process"`   // "appraisal_booking_request"
	EventID   string    `json:"event_ID"`  // UUID v4
	Timestamp int64     `json:"timestamp"` // UNIX timestamp (seconds)
	Data      EventData `json:"data"`
}

type EventData struct {
	OneAccount       OneAccountDTO       `json:"one_account"`
	RequestAppraisal RequestAppraisalDTO `json:"request_appraisal"`
	Leads            LeadsDTO            `json:"leads"`
	UsedCar          UsedCarDTO          `json:"used_car"`
}

// =======================
// one_account
// =======================

type OneAccountDTO struct {
	OneAccountID            string   `json:"one_account_ID"`            // VARCHAR(32)
	FirstName               string   `json:"first_name"`                // VARCHAR(64)
	LastName                string   `json:"last_name"`                 // VARCHAR(64)
	PhoneNumber             string   `json:"phone_number"`              // VARCHAR(16)
	Email                   string   `json:"email"`                     // VARCHAR(64)
	KTPNumber               string   `json:"ktp_number"`                // VARCHAR(16), conditional
	PreferredContactChannel []string `json:"preferred_contact_channel"` // MTOYOTA, WHATSAPP_OR_SMS, EMAIL, PHONE_CALL
}

// =======================
// request_appraisal
// =======================

type RequestAppraisalDTO struct {
	AppraisalBookingID     string `json:"appraisal_booking_ID"`     // UUID / VARCHAR(36)
	AppraisalBookingNumber string `json:"appraisal_booking_number"` // VARCHAR(32)
	CreatedDatetime        int64  `json:"created_datetime"`         // UNIX TIMESTAMP

	OutletID   string `json:"outlet_ID"`   // VARCHAR(32)
	OutletName string `json:"outlet_name"` // VARCHAR(128)

	AppraisalLocation string `json:"appraisal_location"` // DEALER, HOME_OR_OTHER_ADDRESS
	HomeAddress       string `json:"home_address"`       // conditional
	Province          string `json:"province"`
	City              string `json:"city"`
	District          string `json:"district"`
	Subdistrict       string `json:"subdistrict"`
	PostalCode        string `json:"postal_code"` // VARCHAR(5), conditional

	AppraisalStartDatetime int64 `json:"appraisal_start_datetime"` // UNIX TIMESTAMP, step 15 minutes
	AppraisalEndDatetime   int64 `json:"appraisal_end_datetime"`   // UNIX TIMESTAMP, step 15 minutes

	AppraisalBookingStatus  string `json:"appraisal_booking_status"`  // SUBMITTED, CHANGE_REQUEST, CANCEL_SUBMITTED
	CancelledBy             string `json:"cancelled_by"`              // CUSTOMER, DEALER (conditional)
	CancellationReason      string `json:"cancellation_reason"`       // enum list
	OtherCancellationReason string `json:"other_cancellation_reason"` // only if reason = OTHERS

	BookingServiceFlag bool `json:"booking_service_flag"` // true / false
}

// =======================
// leads (+ score)
// =======================

type LeadsDTO struct {
	LeadsID   string `json:"leads_ID"`         // UUID / VARCHAR(36)
	Katashiki string `json:"katashiki_suffix"` // VARCHAR(64), conditional
	ColorCode string `json:"color_code"`       // VARCHAR(16), conditional
	Model     string `json:"model"`            // VARCHAR(64), conditional
	Variant   string `json:"variant"`          // VARCHAR(128), conditional
	Color     string `json:"color"`            // VARCHAR(64), conditional

	LeadsSource string `json:"leads_source"` // enum values (CUSTOMER_DATABASE, DIGITAL_TOYOTA_WEB, ...)
	LeadsType   string `json:"leads_type"`   // APPRAISAL_REQUEST

	LeadsFollowUpStatus string `json:"leads_follow_up_status"` // NOT_YET_FOLLOWED_UP, ON_CONSIDERATION, NO_RESPONSE

	LeadsPreferenceContactTimeStart string `json:"leads_preference_contact_time_start"` // "HH:mm", step 30 minutes
	LeadsPreferenceContactTimeEnd   string `json:"leads_preference_contact_time_end"`   // "HH:mm", step 30 minutes

	TamLeadScore    string `json:"tam_lead_score"`    // LOW, MEDIUM, HOT
	OutletLeadScore string `json:"outlet_lead_score"` // LOW, MEDIUM, HOT

	ScoreParameter LeadScoreParameterDTO `json:"parameter"`

	// row 72: additional_notes (leads)
	AdditionalNotes string `json:"additional_notes"` // optional: catatan dari customer untuk dealer
}

type LeadScoreParameterDTO struct {
	PurchasePlanCriteria    string `json:"purchase_plan_criteria"`    // 1_TO_7_DAYS, 8_TO_30_DAYS, 31_DAYS_TO_INFINITE
	PaymentPreferCriteria   string `json:"payment_prefer_criteria"`   // CASH, CREDIT_APPLIED, CREDIT_NOT_YET_APPLIED, NO_PREFERENCE
	NegotiationCriteria     string `json:"negotiation_criteria"`      // HAVE_STARTED_NEGOTIATIONS, HAVE_NOT_STARTED_NEGOTIATIONS
	TestDriveCriteria       string `json:"test_drive_criteria"`       // COMPLETED, CONFIRMED, SUBMITTED, CANCELLED, UNSPECIFIED
	TradeInCriteria         string `json:"trade_in_criteria"`         // DELIVERY, HANDOVER, PAYMENT, NEGOTIATION, CONFIRMED, SUBMITTED, CANCELLED, UNSPECIFIED
	BrowsingHistoryCriteria string `json:"browsing_history_criteria"` // MORE_THAN_5_PAGES, LESS_THAN_5_PAGES, UNSPECIFIED
	VehicleAgeCriteria      string `json:"vehicle_age_criteria"`      // MORE_THAN_2.5_YEARS, LESS_THAN_2.5_YEARS, UNSPECIFIED
}

// =======================
// used_car
// =======================

type UsedCarDTO struct {
	UsedCarBrand string `json:"used_car_brand"` // VARCHAR(64), Y

	VIN          string `json:"vin,omitempty"` // VARCHAR(17), conditional
	PoliceNumber string `json:"police_number"` // VARCHAR(16), conditional

	KatashikiSuffix string `json:"katashiki_suffix"` // VARCHAR(64), conditional
	ColorCode       string `json:"color_code"`       // VARCHAR(16), conditional

	UsedCarModel   string `json:"used_car_model"`   // VARCHAR(64), conditional
	UsedCarVariant string `json:"used_car_variant"` // VARCHAR(64), conditional
	UsedCarColor   string `json:"used_car_color"`   // VARCHAR(64), conditional

	InitialEstimatedUsedCarPrice float64 `json:"initial_estimated_used_car_price"` // FLOAT, Y

	Mileage        int64  `json:"mileage"`                  // INTEGER, Y (KM)
	Province       string `json:"province"`                 // VARCHAR(64), Y
	UsedCarYear    int    `json:"used_car_year"`            // INTEGER, N
	Transmission   string `json:"used_car_transmission"`    // VARCHAR(64), N (M_T, A_T, CVT)
	Fuel           string `json:"used_car_fuel"`            // VARCHAR(64), N (GASOLINE, DIESEL, EV, HYBRID, HYDROGEN)
	EngineCapacity int64  `json:"used_car_engine_capacity"` // INTEGER, N
	Mover          string `json:"mover"`                    // VARCHAR(64), N (4x2, 4x4, FWD, etc.)

	StnkExpiryDate int64 `json:"stnk_expiry_date"` // UNIX TIMESTAMP, Y

	Stnk   bool `json:"stnk"`   // BOOLEAN, N
	Bpkb   bool `json:"bpkb"`   // BOOLEAN, N
	Faktur bool `json:"faktur"` // BOOLEAN, N

	ServiceBook          bool `json:"service_book"`           // BOOLEAN, N
	AvailabilitySpareKey bool `json:"availability_spare_key"` // BOOLEAN, N
	SpareTireCheck       bool `json:"spare_tire_check"`       // BOOLEAN, N

	AdditionalNotes        string  `json:"additional_notes"`         // VARCHAR(256), N
	CustomerRequestedPrice float64 `json:"customer_requested_price"` // FLOAT, N
}
