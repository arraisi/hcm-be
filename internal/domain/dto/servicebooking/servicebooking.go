package servicebooking

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/internal/domain/dto/customervehicle"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

type ServiceBookingEvent struct {
	Process   string      `json:"process"`
	EventID   string      `json:"event_ID"`
	Timestamp int         `json:"timestamp"`
	Data      DataRequest `json:"data"`
}

type DataRequest struct {
	OneAccount            customer.OneAccountRequest             `json:"one_account"`
	CustomerVehicle       customervehicle.CustomerVehicleRequest `json:"customer_vehicle"`
	Job                   []JobRequest                           `json:"job"`
	Part                  []PartRequest                          `json:"part"`
	ServiceBookingRequest ServiceBookingRequest                  `json:"service_booking"`
	VehicleInsurance      VehicleInsuranceRequest                `json:"vehicle_insurance"`
}

type ServiceBookingRequest struct {
	Warranty                     []WarrantyRequest `json:"warranty"`
	Recalls                      []RecallRequest   `json:"recalls"`
	BookingId                    string            `json:"booking_ID" validate:"required"`
	BookingNumber                string            `json:"booking_number" validate:"required"`
	BookingSource                string            `json:"booking_source" validate:"required"`
	BookingStatus                string            `json:"booking_status" validate:"required"`
	CreatedDatetime              int64             `json:"created_datetime" validate:"required"`
	ServiceCategory              string            `json:"service_category" validate:"required"`
	ServiceSequence              int32             `json:"service_sequence"`
	SlotDatetimeStart            int64             `json:"slot_datetime_start"`
	SlotDatetimeEnd              int64             `json:"slot_datetime_end"`
	SlotRequestedDatetimeStart   int64             `json:"slot_requested_datetime_start"`
	SlotRequestedDatetimeEnd     int64             `json:"slot_requested_datetime_end"`
	SlotUnavailableFlag          bool              `json:"slot_unavailable_flag"`
	CarrierName                  string            `json:"carrier_name"`
	CarrierPhoneNumber           string            `json:"carrier_phone_number"`
	PreferenceContactPhoneNumber string            `json:"preference_contact_phone_number"`
	PreferenceContactTimeStart   string            `json:"preference_contact_time_start"`
	PreferenceContactTimeEnd     string            `json:"preference_contact_time_end"`
	ServiceLocation              string            `json:"service_location"`
	OutletID                     string            `json:"outlet_ID" validate:"required"`
	OutletName                   string            `json:"outlet_name" validate:"required"`
	MobileServiceAddress         string            `json:"mobile_service_address"`
	Province                     string            `json:"province"`
	City                         string            `json:"city"`
	District                     string            `json:"district"`
	SubDistrict                  string            `json:"subdistrict"`
	PostalCode                   string            `json:"postal_code"`
	VehicleProblem               string            `json:"vehicle_problem"`
	CancellationReason           string            `json:"cancellation_reason"`
	OtherCancellationReason      string            `json:"other_cancellation_reason"`
	ServicePricingCallFlag       bool              `json:"service_pricing_call_flag"`

	// from service booking bp
	AppointmentDatetimeStart          int64    `json:"appointment_datetime_start"`
	AppointmentDatetimeEnd            int64    `json:"appointment_datetime_end"`
	AppointmentRequestedDatetimeStart int64    `json:"appointment_requested_datetime_start"`
	AppointmentRequestedDatetimeEnd   int64    `json:"appointment_requested_datetime_end"`
	AdditionalVehicleProblem          string   `json:"additional_vehicle_problem"`
	DamageImage                       []string `json:"damage_image"`
	InsuranceClaim                    string   `json:"insurance_claim"`
}

// ToServiceBookingModel converts the DataRequest to the domain.TestDrive model
func (sb *ServiceBookingEvent) ToServiceBookingModel(customerID, customerVehicleID string) domain.ServiceBooking {
	serviceBooking := domain.ServiceBooking{
		EventID:                      sb.EventID,
		CustomerID:                   customerID,
		CustomerVehicleID:            customerVehicleID,
		ServiceBookingID:             sb.Data.ServiceBookingRequest.BookingId,
		ServiceBookingNumber:         sb.Data.ServiceBookingRequest.BookingNumber,
		ServiceBookingSource:         sb.Data.ServiceBookingRequest.BookingSource,
		ServiceBookingStatus:         sb.Data.ServiceBookingRequest.BookingStatus,
		CreatedDatetime:              utils.GetTimeUnix(sb.Data.ServiceBookingRequest.CreatedDatetime).UTC(),
		ServiceCategory:              sb.Data.ServiceBookingRequest.ServiceCategory,
		ServiceSequence:              sb.Data.ServiceBookingRequest.ServiceSequence,
		SlotDatetimeStart:            utils.GetTimeUnix(sb.Data.ServiceBookingRequest.SlotDatetimeStart).UTC(),
		SlotDatetimeEnd:              utils.GetTimeUnix(sb.Data.ServiceBookingRequest.SlotDatetimeEnd).UTC(),
		SlotRequestedDatetimeStart:   utils.GetTimeUnix(sb.Data.ServiceBookingRequest.SlotRequestedDatetimeStart).UTC(),
		SlotRequestedDatetimeEnd:     utils.GetTimeUnix(sb.Data.ServiceBookingRequest.SlotRequestedDatetimeEnd).UTC(),
		SlotUnavailableFlag:          sb.Data.ServiceBookingRequest.SlotUnavailableFlag,
		CarrierName:                  sb.Data.ServiceBookingRequest.CarrierName,
		CarrierPhoneNumber:           sb.Data.ServiceBookingRequest.CarrierPhoneNumber,
		PreferenceContactPhoneNumber: sb.Data.ServiceBookingRequest.PreferenceContactPhoneNumber,
		PreferenceContactTimeStart:   sb.Data.ServiceBookingRequest.PreferenceContactTimeStart,
		PreferenceContactTimeEnd:     sb.Data.ServiceBookingRequest.PreferenceContactTimeEnd,
		ServiceLocation:              sb.Data.ServiceBookingRequest.ServiceLocation,
		OutletID:                     sb.Data.ServiceBookingRequest.OutletID,
		OutletName:                   sb.Data.ServiceBookingRequest.OutletName,
		MobileServiceAddress:         sb.Data.ServiceBookingRequest.MobileServiceAddress,
		Province:                     sb.Data.ServiceBookingRequest.Province,
		City:                         sb.Data.ServiceBookingRequest.City,
		District:                     sb.Data.ServiceBookingRequest.District,
		SubDistrict:                  sb.Data.ServiceBookingRequest.SubDistrict,
		PostalCode:                   sb.Data.ServiceBookingRequest.PostalCode,
		VehicleProblem:               sb.Data.ServiceBookingRequest.VehicleProblem,
		CancellationReason:           sb.Data.ServiceBookingRequest.CancellationReason,
		OtherCancellationReason:      sb.Data.ServiceBookingRequest.OtherCancellationReason,
		ServicePricingCallFlag:       sb.Data.ServiceBookingRequest.ServicePricingCallFlag,
		CreatedAt:                    time.Now().UTC(),
		CreatedBy:                    constants.System, // or fetch from context if available
		UpdatedAt:                    time.Now().UTC(),
		UpdatedBy:                    constants.System, // or fetch from context if available
		AdditionalVehicleProblem:     sb.Data.ServiceBookingRequest.AdditionalVehicleProblem,
		InsuranceClaim:               sb.Data.ServiceBookingRequest.InsuranceClaim,
	}
	if sb.Data.ServiceBookingRequest.AppointmentDatetimeStart != 0 {
		serviceBooking.SlotDatetimeStart = utils.GetTimeUnix(sb.Data.ServiceBookingRequest.AppointmentDatetimeStart).UTC()
	}
	if sb.Data.ServiceBookingRequest.AppointmentDatetimeEnd != 0 {
		serviceBooking.SlotDatetimeEnd = utils.GetTimeUnix(sb.Data.ServiceBookingRequest.AppointmentDatetimeEnd).UTC()
	}
	if sb.Data.ServiceBookingRequest.AppointmentRequestedDatetimeStart != 0 {
		serviceBooking.SlotRequestedDatetimeStart = utils.GetTimeUnix(sb.Data.ServiceBookingRequest.AppointmentRequestedDatetimeStart).UTC()
	}
	if sb.Data.ServiceBookingRequest.AppointmentRequestedDatetimeEnd != 0 {
		serviceBooking.SlotRequestedDatetimeEnd = utils.GetTimeUnix(sb.Data.ServiceBookingRequest.AppointmentRequestedDatetimeEnd).UTC()
	}
	return serviceBooking
}

type GetServiceBooking struct {
	ID                   *string
	CustomerID           *string
	ServiceBookingID     *string
	ServiceBookingNumber *string
	ServiceBookingSource *string
	ServiceBookingStatus *string
	ServiceCategory      *string
	EventID              *string
}

func (g *GetServiceBooking) Apply(q *sqrl.SelectBuilder) {
	if g.ID != nil {
		q.Where(sqrl.Eq{"i_id": g.ID})
	}
	if g.CustomerID != nil {
		q.Where(sqrl.Eq{"i_customer_id": g.CustomerID})
	}
	if g.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": g.ServiceBookingID})
	}
	if g.ServiceBookingNumber != nil {
		q.Where(sqrl.Eq{"c_service_booking_number": g.ServiceBookingNumber})
	}
	if g.ServiceBookingStatus != nil {
		q.Where(sqrl.Eq{"c_service_booking_status": g.ServiceBookingStatus})
	}
	if g.ServiceCategory != nil {
		q.Where(sqrl.Eq{"c_service_category": g.ServiceCategory})
	}
	if g.EventID != nil {
		q.Where(sqrl.Eq{"i_event_id": g.EventID})
	}
}

type ConfirmServiceBookingRequest struct {
	ServiceBookingID string `json:"service_booking_id" validate:"required"`
	EmployeeID       string `json:"employee_id" validate:"required"`
	Status           string `json:"status" validate:"required,oneof=MANUALLY_CONFIRMED CANCELLED COMPLETED NOT_SHOW SYSTEM_CONFIRMED"`
	Location         string `json:"location" validate:"required,oneof=WORKSHOP MOBILE_SERVICES"`
}

// ServiceBookingEventData represents the data payload for confirm event
type ServiceBookingEventData struct {
	OneAccount     customer.OneAccountRequest `json:"one_account" validate:"required"`
	ServiceBooking *ServiceBookingRequest     `json:"service_booking" validate:"required"`
	PICAssignment  *PICAssignmentRequest      `json:"pic_assignment,omitempty"`
}

// PICAssignmentRequest represents the PIC assignment information
type PICAssignmentRequest struct {
	EmployeeID string `json:"employee_id" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
}

// NewServiceBookingRequest creates a ServiceBookingRequest from domain model
func NewServiceBookingRequest(sb domain.ServiceBooking, warranties []WarrantyRequest, recalls []RecallRequest) ServiceBookingRequest {
	return ServiceBookingRequest{
		Warranty:                     warranties,
		Recalls:                      recalls,
		BookingId:                    sb.ServiceBookingID,
		BookingNumber:                sb.ServiceBookingNumber,
		BookingSource:                sb.ServiceBookingSource,
		BookingStatus:                sb.ServiceBookingStatus,
		CreatedDatetime:              sb.CreatedDatetime.Unix(),
		ServiceCategory:              sb.ServiceCategory,
		ServiceSequence:              sb.ServiceSequence,
		SlotDatetimeStart:            sb.SlotDatetimeStart.Unix(),
		SlotDatetimeEnd:              sb.SlotDatetimeEnd.Unix(),
		SlotRequestedDatetimeStart:   sb.SlotRequestedDatetimeStart.Unix(),
		SlotRequestedDatetimeEnd:     sb.SlotRequestedDatetimeEnd.Unix(),
		SlotUnavailableFlag:          sb.SlotUnavailableFlag,
		CarrierName:                  sb.CarrierName,
		CarrierPhoneNumber:           sb.CarrierPhoneNumber,
		PreferenceContactPhoneNumber: sb.PreferenceContactPhoneNumber,
		PreferenceContactTimeStart:   sb.PreferenceContactTimeStart,
		PreferenceContactTimeEnd:     sb.PreferenceContactTimeEnd,
		ServiceLocation:              sb.ServiceLocation,
		OutletID:                     sb.OutletID,
		OutletName:                   sb.OutletName,
		MobileServiceAddress:         sb.MobileServiceAddress,
		Province:                     sb.Province,
		City:                         sb.City,
		District:                     sb.District,
		SubDistrict:                  sb.SubDistrict,
		PostalCode:                   sb.PostalCode,
		VehicleProblem:               sb.VehicleProblem,
		CancellationReason:           sb.CancellationReason,
		OtherCancellationReason:      sb.OtherCancellationReason,
		ServicePricingCallFlag:       sb.ServicePricingCallFlag,
	}
}
