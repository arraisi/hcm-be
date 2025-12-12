package customerreminder

import (
	"fmt"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

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
	OneAccount      OneAccount      `json:"one_account" validate:"required"`
	ReminderDetail  ReminderDetail  `json:"reminder" validate:"required"`
	CustomerVehicle CustomerVehicle `json:"customer_vehicle" validate:"required"`
}

// -------- one_account --------

type OneAccount struct {
	OneAccountID            string   `json:"one_account_ID" validate:"required"`     // VARCHAR(32)
	DealerCustomerID        string   `json:"dealer_customer_ID" validate:"required"` // VARCHAR(32)
	FirstName               string   `json:"first_name" validate:"required"`         // VARCHAR(64)
	LastName                string   `json:"last_name" validate:"required"`          // VARCHAR(64)
	PhoneNumber             string   `json:"phone_number" validate:"required"`       // VARCHAR(16)
	Email                   string   `json:"email" validate:"required,email"`        // VARCHAR(64)
	PreferredContactChannel []string `json:"preferred_contact_channel" validate:"required,dive,oneof=MTOYOTA WHATSAPP_OR_SMS EMAIL PHONE_CALL"`
}

func (dto *OneAccount) ToCustomerModel() domain.Customer {
	now := time.Now().UTC()
	entity := domain.Customer{
		OneAccountID:            utils.ToPointer(dto.OneAccountID),
		DealerCustomerID:        dto.DealerCustomerID,
		FirstName:               dto.FirstName,
		LastName:                dto.LastName,
		PhoneNumber:             dto.PhoneNumber,
		Email:                   dto.Email,
		PreferredContactChannel: utils.JoinSCommaSeparatedString(dto.PreferredContactChannel),
		CreatedAt:               now,
		CreatedBy:               constants.System,
		UpdatedAt:               now,
		UpdatedBy:               utils.ToPointer(constants.System),
	}

	return entity
}

// -------- customer_vehicle --------

type CustomerVehicle struct {
	VIN             string `json:"vin" validate:"omitempty"`              // VARCHAR(17) – Y Conditional
	PoliceNumber    string `json:"police_number" validate:"omitempty"`    // VARCHAR(16)
	KatashikiSuffix string `json:"katashiki_suffix" validate:"omitempty"` // VARCHAR(64)
	Model           string `json:"model" validate:"omitempty"`            // VARCHAR(64)
	Variant         string `json:"variant" validate:"omitempty"`          // VARCHAR(128)
	ColorCode       string `json:"color_code" validate:"omitempty"`       // VARCHAR(16)
	Color           string `json:"color" validate:"omitempty"`            // VARCHAR(64)
}

func (dto *CustomerVehicle) ToCustomerVehicleModel() domain.CustomerVehicle {
	now := time.Now().UTC()

	return domain.CustomerVehicle{
		Vin:             dto.VIN,
		KatashikiSuffix: dto.KatashikiSuffix,
		ColorCode:       dto.ColorCode,
		Model:           dto.Model,
		Variant:         dto.Variant,
		Color:           dto.Color,
		PoliceNumber:    dto.PoliceNumber,
		CreatedAt:       now,
		CreatedBy:       constants.System,
		UpdatedAt:       now,
		UpdatedBy:       constants.System,
	}
}

// -------- reminder --------

type ReminderDetail struct {
	ReminderID                string `json:"reminder_ID" validate:"required,uuid4"` // UUID / VARCHAR(36)
	Activity                  string `json:"activity" validate:"required,oneof=SERVICE_BOOKING APPRAISAL_BOOKING TEST_DRIVE_BOOKING DISTRIBUTE_LEADS TRACK_ORDER_STATUS SERVICE_REMINDER DELIVERY_AMEND_REQUEST PRE_DEC TCO ACCOUNT_MANAGEMENT MONITOR_SERVICE_PROGRESS OTHERS"`
	ActivityPlanScheduledDate int64  `json:"activity_plan_scheduled_date" validate:"required"` // Unix timestamp
	AutoReminderStatus        string `json:"auto_reminder_status" validate:"omitempty,oneof=DELIVERED READ NON_MTOYOTA"`
	ReminderMessage           string `json:"reminder_message" validate:"required"` // VARCHAR(256)
	PriorityCall              int    `json:"priority_call" validate:"omitempty"`   // INTEGER (Y Conditional)
	ExtendedWarrantyStatus    string `json:"extended_warranty_status" validate:"omitempty,oneof=ELIGIBLE NOT_ELIGIBLE NOT_APPLICABLE"`
	CustomerHabit             string `json:"customer_habit" validate:"omitempty,oneof=TIME_BASED MILEAGE"`                        // VARCHAR(10)
	LastHabit                 string `json:"last_habit" validate:"omitempty,oneof=PUNCTUAL EARLY LATE INACTIVE PASSIVE"`          // VARCHAR(8)
	NextServiceStatus         string `json:"next_service_status" validate:"omitempty,oneof=PUNCTUAL EARLY LATE INACTIVE PASSIVE"` // VARCHAR(8)
	LastServiceDate           int64  `json:"last_service_date" validate:"omitempty"`                                              // Unix timestamp
	NextServiceDate           int64  `json:"next_service_date" validate:"omitempty"`                                              // Unix timestamp
	NCSStatus                 string `json:"ncs_status" validate:"omitempty,oneof=SAME_OUTLET DIFFERENT"`                         // VARCHAR(11)
	ProgramTab                string `json:"program_tab" validate:"omitempty,oneof=T_CARE GBSB REGULAR"`                          // VARCHAR(7)
	NextServiceStage          int    `json:"next_service_stage" validate:"omitempty"`                                             // INTEGER
}

func (dto *ReminderDetail) ToDomainCustomerReminder() domain.CustomerReminder {
	now := time.Now().UTC()

	entity := domain.CustomerReminder{
		ExternalReminderID:      dto.ReminderID,
		Activity:                dto.Activity,
		ActivityPlanScheduledAt: utils.ToPointer(utils.GetTimeUnix(dto.ActivityPlanScheduledDate)),
		AutoReminderStatus:      &dto.AutoReminderStatus,
		ReminderMessage:         &dto.ReminderMessage,
		PriorityCall:            &dto.PriorityCall,
		ExtendedWarrantyStatus:  &dto.ExtendedWarrantyStatus,
		CustomerHabit:           &dto.CustomerHabit,
		LastHabit:               &dto.LastHabit,
		NextServiceStatus:       &dto.NextServiceStatus,
		NCSStatus:               &dto.NCSStatus,
		ProgramTab:              &dto.ProgramTab,
		NextServiceStage:        &dto.NextServiceStage,
		CreatedAt:               now,
		CreatedBy:               utils.ToPointer(constants.System),
		UpdatedAt:               &now,
		UpdatedBy:               utils.ToPointer(constants.System),
	}

	if dto.LastServiceDate > 0 {
		entity.LastServiceDate = utils.ToPointer(utils.GetTimeUnix(dto.LastServiceDate))
	}

	if dto.NextServiceDate > 0 {
		entity.NextServiceDate = utils.ToPointer(utils.GetTimeUnix(dto.NextServiceDate))

	}

	return entity
}

// GetCustomerRequest represents the request parameters for getting users
type GetCustomerReminderRequest struct {
	Limit              int
	Offset             int
	Search             string
	SortBy             string
	Order              string
	ExternalReminderID string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetCustomerReminderRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ExternalReminderID != "" {
		q.Where(sqrl.Eq{"i_reminder_id": req.ExternalReminderID})
	}

	if req.Limit > 0 {
		q.Suffix(fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", req.Offset, req.Limit))
	}
}
