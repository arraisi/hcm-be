package customerreminder

// Root request payload
type Request struct {
	Process   string `json:"process" validate:"required"`        // "customer_reminder"
	EventID   string `json:"event_ID" validate:"required,uuid4"` // UUID v4
	Timestamp int64  `json:"timestamp" validate:"required"`      // Unix timestamp
	Data      Data   `json:"data" validate:"required"`
}

// Data holds outlet and reminder list
type Data struct {
	OutletID  string     `json:"outlet_ID" validate:"required"`      // VARCHAR(32)
	Reminders []Reminder `json:"reminders" validate:"required,dive"` // at least 1 reminder
}

// One reminder entry (one_account + reminder + vehicle)
type Reminder struct {
	OneAccount      ReminderOneAccount      `json:"one_account" validate:"required"`
	ReminderDetail  ReminderDetail          `json:"reminder" validate:"required"`
	CustomerVehicle ReminderCustomerVehicle `json:"customer_vehicle" validate:"required"`
}

// -------- one_account --------

type ReminderOneAccount struct {
	OneAccountID            string   `json:"one_account_ID" validate:"required"`     // VARCHAR(32)
	DealerCustomerID        string   `json:"dealer_customer_ID" validate:"required"` // VARCHAR(32)
	FirstName               string   `json:"first_name" validate:"required"`         // VARCHAR(64)
	LastName                string   `json:"last_name" validate:"required"`          // VARCHAR(64)
	PhoneNumber             string   `json:"phone_number" validate:"required"`       // VARCHAR(16)
	Email                   string   `json:"email" validate:"required,email"`        // VARCHAR(64)
	PreferredContactChannel []string `json:"preferred_contact_channel" validate:"required,dive,oneof=MTOYOTA WHATSAPP_OR_SMS EMAIL PHONE_CALL"`
}

// -------- reminder --------

type ReminderDetail struct {
	ReminderID string `json:"reminder_ID" validate:"required,uuid4"` // UUID / VARCHAR(36)

	// Activity being reminded
	Activity string `json:"activity" validate:"required,oneof=SERVICE_BOOKING APPRAISAL_BOOKING TEST_DRIVE_BOOKING DISTRIBUTE_LEADS TRACK_ORDER_STATUS SERVICE_REMINDER DELIVERY_AMEND_REQUEST PRE_DEC TCO ACCOUNT_MANAGEMENT MONITOR_SERVICE_PROGRESS OTHERS"`

	ActivityPlanScheduledDate int64  `json:"activity_plan_scheduled_date" validate:"required"` // Unix timestamp
	AutoReminderStatus        string `json:"auto_reminder_status" validate:"omitempty,oneof=DELIVERED READ NON_MTOYOTA"`
	ReminderMessage           string `json:"reminder_message" validate:"required"` // VARCHAR(256)
	PriorityCall              int    `json:"priority_call" validate:"omitempty"`   // INTEGER (Y Conditional)
	ExtendedWarrantyStatus    string `json:"extended_warranty_status" validate:"omitempty,oneof=ELIGIBLE NOT_ELIGIBLE NOT_APPLICABLE"`

	CustomerHabit     string `json:"customer_habit" validate:"omitempty,oneof=TIME_BASED MILEAGE"`                        // VARCHAR(10)
	LastHabit         string `json:"last_habit" validate:"omitempty,oneof=PUNCTUAL EARLY LATE INACTIVE PASSIVE"`          // VARCHAR(8)
	NextServiceStatus string `json:"next_service_status" validate:"omitempty,oneof=PUNCTUAL EARLY LATE INACTIVE PASSIVE"` // VARCHAR(8)

	LastServiceDate int64 `json:"last_service_date" validate:"omitempty"` // Unix timestamp
	NextServiceDate int64 `json:"next_service_date" validate:"omitempty"` // Unix timestamp

	NCSStatus  string `json:"ncs_status" validate:"omitempty,oneof=SAME_OUTLET DIFFERENT"` // VARCHAR(11)
	ProgramTab string `json:"program_tab" validate:"omitempty,oneof=T_CARE GBSB REGULAR"`  // VARCHAR(7)

	NextServiceStage int `json:"next_service_stage" validate:"omitempty"` // INTEGER
}

// -------- customer_vehicle --------

type ReminderCustomerVehicle struct {
	VIN             string `json:"vin" validate:"omitempty"`              // VARCHAR(17) – Y Conditional
	PoliceNumber    string `json:"police_number" validate:"omitempty"`    // VARCHAR(16)
	KatashikiSuffix string `json:"katashiki_suffix" validate:"omitempty"` // VARCHAR(64)
	Model           string `json:"model" validate:"omitempty"`            // VARCHAR(64)
	Variant         string `json:"variant" validate:"omitempty"`          // VARCHAR(128)
	ColorCode       string `json:"color_code" validate:"omitempty"`       // VARCHAR(16)
	Color           string `json:"color" validate:"omitempty"`            // VARCHAR(64)
}
