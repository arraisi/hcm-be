package appraisal

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/leads"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// AppraisalConfirmEventData represents the data payload of the confirmation event
type AppraisalConfirmEventData struct {
	OneAccount       customer.OneAccountRequest `json:"one_account" validate:"required"`
	RequestAppraisal AppraisalConfirmRequest    `json:"request_appraisal" validate:"required"`
	Leads            leads.LeadsRequest         `json:"leads" validate:"required"`
	UsedCar          UsedCarDTO                 `json:"used_car" validate:"required"`
	PICAssignment    *PICAssignmentRequest      `json:"pic_assignment,omitempty"`
}

// AppraisalConfirmRequest represents the appraisal confirmation information
type AppraisalConfirmRequest struct {
	AppraisalBookingID                 string  `json:"appraisal_booking_ID" validate:"required"`
	AppraisalBookingNumber             string  `json:"appraisal_booking_number" validate:"required"`
	CreatedDatetime                    int64   `json:"created_datetime" validate:"required"`
	OutletID                           string  `json:"outlet_ID" validate:"required"`
	OutletName                         string  `json:"outlet_name" validate:"required"`
	AppraisalLocation                  string  `json:"appraisal_location" validate:"required,oneof=DEALER HOME_OR_OTHER_ADDRESS"`
	HomeAddress                        string  `json:"home_address,omitempty"`
	Province                           string  `json:"province,omitempty"`
	City                               string  `json:"city,omitempty"`
	District                           string  `json:"district,omitempty"`
	Subdistrict                        string  `json:"subdistrict,omitempty"`
	PostalCode                         string  `json:"postal_code,omitempty"`
	AppraisalStartDatetime             int64   `json:"appraisal_start_datetime" validate:"required"`
	AppraisalEndDatetime               int64   `json:"appraisal_end_datetime" validate:"required"`
	AppraisalConfirmationStartDatetime int64   `json:"appraisal_confirmation_start_datetime" validate:"required"`
	AppraisalConfirmationEndDatetime   int64   `json:"appraisal_confirmation_end_datetime" validate:"required"`
	AppraisalBookingStatus             string  `json:"appraisal_booking_status" validate:"required,oneof=CONFIRMED CANCEL_CONFIRMED CHANGE_REQUEST_CONFIRMED"`
	CancelledBy                        *string `json:"cancelled_by,omitempty"`
	CancellationReason                 *string `json:"cancellation_reason,omitempty"`
	OtherCancellationReason            *string `json:"other_cancellation_reason,omitempty"`
}

// PICAssignmentRequest represents the PIC assignment information
type PICAssignmentRequest struct {
	EmployeeID string `json:"employee_ID"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NIK        string `json:"nik"`
}

// AppraisalConfirmEvent represents the complete webhook payload for appraisal confirmation
type AppraisalConfirmEvent struct {
	Process   string                    `json:"process" validate:"required"`
	EventID   string                    `json:"event_ID" validate:"required,uuid4"`
	Timestamp int64                     `json:"timestamp" validate:"required"`
	Data      AppraisalConfirmEventData `json:"data" validate:"required"`
}

// NewAppraisalConfirmRequest creates a new AppraisalConfirmRequest from domain model
func NewAppraisalConfirmRequest(appraisal domain.Appraisal) AppraisalConfirmRequest {
	return AppraisalConfirmRequest{
		AppraisalBookingID:      appraisal.AppraisalBookingID,
		AppraisalBookingNumber:  appraisal.AppraisalBookingNumber,
		AppraisalBookingStatus:  appraisal.BookingStatus,
		OutletID:                appraisal.OutletID,
		OutletName:              appraisal.OutletName,
		AppraisalLocation:       utils.ToValue(appraisal.AppraisalLocation),
		AppraisalStartDatetime:  appraisal.ConfirmStartDatetime.Unix(),
		AppraisalEndDatetime:    appraisal.ConfirmEndDatetime.Unix(),
		CancelledBy:             appraisal.CancelledBy,
		CancellationReason:      appraisal.CancellationReason,
		OtherCancellationReason: appraisal.OtherCancellationReason,
	}
}

// ToAppraisalUpdateModel converts the AppraisalConfirmEvent to update the appraisal model
func (e *AppraisalConfirmEvent) ToAppraisalUpdateModel() domain.Appraisal {
	var confirmStart, confirmEnd, appraisalStart, appraisalEnd *time.Time
	var createdDatetime time.Time

	if e.Data.RequestAppraisal.AppraisalConfirmationStartDatetime > 0 {
		t := utils.GetTimeUnix(e.Data.RequestAppraisal.AppraisalConfirmationStartDatetime)
		confirmStart = &t
	}
	if e.Data.RequestAppraisal.AppraisalConfirmationEndDatetime > 0 {
		t := utils.GetTimeUnix(e.Data.RequestAppraisal.AppraisalConfirmationEndDatetime)
		confirmEnd = &t
	}
	if e.Data.RequestAppraisal.AppraisalStartDatetime > 0 {
		t := utils.GetTimeUnix(e.Data.RequestAppraisal.AppraisalStartDatetime)
		appraisalStart = &t
	}
	if e.Data.RequestAppraisal.AppraisalEndDatetime > 0 {
		t := utils.GetTimeUnix(e.Data.RequestAppraisal.AppraisalEndDatetime)
		appraisalEnd = &t
	}
	if e.Data.RequestAppraisal.CreatedDatetime > 0 {
		createdDatetime = utils.GetTimeUnix(e.Data.RequestAppraisal.CreatedDatetime)
	}

	return domain.Appraisal{
		AppraisalBookingID:      e.Data.RequestAppraisal.AppraisalBookingID,
		AppraisalBookingNumber:  e.Data.RequestAppraisal.AppraisalBookingNumber,
		OutletID:                e.Data.RequestAppraisal.OutletID,
		OutletName:              e.Data.RequestAppraisal.OutletName,
		AppraisalLocation:       utils.ToPointer(e.Data.RequestAppraisal.AppraisalLocation),
		HomeAddress:             utils.ToPointer(e.Data.RequestAppraisal.HomeAddress),
		Province:                utils.ToPointer(e.Data.RequestAppraisal.Province),
		City:                    utils.ToPointer(e.Data.RequestAppraisal.City),
		District:                utils.ToPointer(e.Data.RequestAppraisal.District),
		Subdistrict:             utils.ToPointer(e.Data.RequestAppraisal.Subdistrict),
		PostalCode:              utils.ToPointer(e.Data.RequestAppraisal.PostalCode),
		CreatedDatetime:         createdDatetime,
		AppraisalStartDatetime:  appraisalStart,
		AppraisalEndDatetime:    appraisalEnd,
		ConfirmStartDatetime:    confirmStart,
		ConfirmEndDatetime:      confirmEnd,
		BookingStatus:           e.Data.RequestAppraisal.AppraisalBookingStatus,
		CancelledBy:             e.Data.RequestAppraisal.CancelledBy,
		CancellationReason:      e.Data.RequestAppraisal.CancellationReason,
		OtherCancellationReason: e.Data.RequestAppraisal.OtherCancellationReason,
	}
}
