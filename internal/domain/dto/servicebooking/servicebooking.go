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
	VehicleInsurance      *VehicleInsuranceRequest               `json:"vehicle_insurance"`
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
		BookingType:                  sb.Process,
		EventID:                      sb.EventID,
		CustomerID:                   customerID,
		CustomerVehicleID:            customerVehicleID,
		ServiceBookingID:             sb.Data.ServiceBookingRequest.BookingId,
		ServiceBookingNumber:         sb.Data.ServiceBookingRequest.BookingNumber,
		ServiceBookingSource:         sb.Data.ServiceBookingRequest.BookingSource,
		ServiceBookingStatus:         sb.Data.ServiceBookingRequest.BookingStatus,
		CreatedDatetime:              utils.GetTimeUnix(sb.Data.ServiceBookingRequest.CreatedDatetime).UTC(),
		ServiceCategory:              sb.Data.ServiceBookingRequest.ServiceCategory,
		ServiceSequence:              utils.ToPointer(sb.Data.ServiceBookingRequest.ServiceSequence),
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
		ServiceLocation:              utils.ToPointer(sb.Data.ServiceBookingRequest.ServiceLocation),
		OutletID:                     sb.Data.ServiceBookingRequest.OutletID,
		OutletName:                   sb.Data.ServiceBookingRequest.OutletName,
		MobileServiceAddress:         utils.ToPointer(sb.Data.ServiceBookingRequest.MobileServiceAddress),
		Province:                     utils.ToPointer(sb.Data.ServiceBookingRequest.Province),
		City:                         utils.ToPointer(sb.Data.ServiceBookingRequest.City),
		District:                     utils.ToPointer(sb.Data.ServiceBookingRequest.District),
		SubDistrict:                  utils.ToPointer(sb.Data.ServiceBookingRequest.SubDistrict),
		PostalCode:                   utils.ToPointer(sb.Data.ServiceBookingRequest.PostalCode),
		VehicleProblem:               utils.ToPointer(sb.Data.ServiceBookingRequest.VehicleProblem),
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
	Location         string `json:"location"`
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
		ServiceSequence:              utils.ToValue(sb.ServiceSequence),
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
		ServiceLocation:              utils.ToValue(sb.ServiceLocation),
		OutletID:                     sb.OutletID,
		OutletName:                   sb.OutletName,
		MobileServiceAddress:         utils.ToValue(sb.MobileServiceAddress),
		Province:                     utils.ToValue(sb.Province),
		City:                         utils.ToValue(sb.City),
		District:                     utils.ToValue(sb.District),
		SubDistrict:                  utils.ToValue(sb.SubDistrict),
		PostalCode:                   utils.ToValue(sb.PostalCode),
		VehicleProblem:               utils.ToValue(sb.VehicleProblem),
		CancellationReason:           sb.CancellationReason,
		OtherCancellationReason:      sb.OtherCancellationReason,
		ServicePricingCallFlag:       sb.ServicePricingCallFlag,
	}
}

// DmsAfterSaleServiceBookingEvent represents the structure for DMS After Sales API
type DmsAfterSaleServiceBookingEvent struct {
	Process   string                  `json:"process"`
	EventID   string                  `json:"event_ID"`
	Timestamp int                     `json:"timestamp"`
	Data      DmsAfterSaleDataRequest `json:"data"`
}

type DmsAfterSaleDataRequest struct {
	OneAccount      DmsAfterSaleOneAccount      `json:"one_account"`
	CustomerVehicle DmsAfterSaleCustomerVehicle `json:"customer_vehicle"`
	ServiceBooking  DmsAfterSaleServiceBooking  `json:"service_booking"`
	Job             []DmsAfterSaleJob           `json:"job"`
	Part            []DmsAfterSalePart          `json:"part"`
}

type DmsAfterSaleOneAccount struct {
	OneAccountID string `json:"one_account_ID"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	PrimaryUser  string `json:"primary_user"`
}

type DmsAfterSaleCustomerVehicle struct {
	VIN             string `json:"vin"`
	KatashikiSuffix string `json:"katashiki_suffix"`
	Model           string `json:"model"`
	Variant         string `json:"variant"`
	PoliceNumber    string `json:"police_number"`
	ActualMileage   int32  `json:"actual_mileage"`
	ColorCode       string `json:"color_code"`
	Color           string `json:"color"`
}

type DmsAfterSaleServiceBooking struct {
	BookingID                    string                 `json:"booking_ID"`
	BookingNumber                string                 `json:"booking_number"`
	BookingSource                string                 `json:"booking_source"`
	BookingStatus                string                 `json:"booking_status"`
	CreatedDatetime              int64                  `json:"created_datetime"`
	ServiceCategory              string                 `json:"service_category"`
	ServiceSequence              int32                  `json:"service_sequence"`
	SlotDatetimeStart            int64                  `json:"slot_datetime_start"`
	SlotDatetimeEnd              int64                  `json:"slot_datetime_end"`
	SlotRequestedDatetimeStart   int64                  `json:"slot_requested_datetime_start"`
	SlotRequestedDatetimeEnd     int64                  `json:"slot_requested_datetime_end"`
	SlotUnavailableFlag          bool                   `json:"slot_unavailable_flag"`
	CarrierName                  string                 `json:"carrier_name"`
	CarrierPhoneNumber           string                 `json:"carrier_phone_number"`
	PreferenceContactPhoneNumber string                 `json:"preference_contact_phone_number"`
	PreferenceContactTimeStart   string                 `json:"preference_contact_time_start"`
	PreferenceContactTimeEnd     string                 `json:"preference_contact_time_end"`
	ServiceLocation              string                 `json:"service_location"`
	OutletID                     string                 `json:"outlet_ID"`
	OutletName                   string                 `json:"outlet_name"`
	MobileServiceAddress         string                 `json:"mobile_service_address"`
	Province                     string                 `json:"province"`
	City                         string                 `json:"city"`
	District                     string                 `json:"district"`
	SubDistrict                  string                 `json:"subdistrict"`
	PostalCode                   string                 `json:"postal_code"`
	VehicleProblem               string                 `json:"vehicle_problem"`
	Warranty                     []DmsAfterSaleWarranty `json:"warranty"`
	Recalls                      []DmsAfterSaleRecall   `json:"recalls"`
	CancellationReason           string                 `json:"cancellation_reason"`
	OtherCancellationReason      string                 `json:"other_cancellation_reason"`
	ServicePricingCallFlag       bool                   `json:"service_pricing_call_flag"`
}

type DmsAfterSaleWarranty struct {
	WarrantyName   string `json:"warranty_name"`
	WarrantyStatus string `json:"warranty_status"`
}

type DmsAfterSaleRecall struct {
	RecallID          string   `json:"recall_ID"`
	RecallDate        string   `json:"recall_date"`
	RecallDescription string   `json:"recall_description"`
	AffectedParts     []string `json:"affected_parts"`
}

type DmsAfterSaleJob struct {
	JobID         string  `json:"job_ID"`
	JobName       string  `json:"job_name"`
	LaborEstPrice float32 `json:"labor_est_price"`
}

type DmsAfterSalePart struct {
	PartType                 string                 `json:"part_type"`
	PackageID                string                 `json:"package_ID"`
	PartNumber               string                 `json:"part_number"`
	PartName                 string                 `json:"part_name"`
	PartQuantity             int32                  `json:"part_quantity"`
	PackageParts             []DmsAfterSalePartItem `json:"package_parts"`
	PartSize                 string                 `json:"part_size"`
	PartColor                string                 `json:"part_color"`
	PartEstPrice             float32                `json:"part_est_price"`
	PartInstallationEstPrice float32                `json:"part_installation_est_price"`
	FlagPartNeedDownPayment  bool                   `json:"flag_part_need_down_payment"`
}

type DmsAfterSalePartItem struct {
	PartNumber string `json:"part_number"`
	PartName   string `json:"part_name"`
}

// ToDmsAfterSaleEvent converts ServiceBookingEvent to DmsAfterSaleServiceBookingEvent
func (sb *ServiceBookingEvent) ToDmsAfterSaleEvent() DmsAfterSaleServiceBookingEvent {
	// Convert warranties
	warranties := make([]DmsAfterSaleWarranty, len(sb.Data.ServiceBookingRequest.Warranty))
	for i, w := range sb.Data.ServiceBookingRequest.Warranty {
		warranties[i] = DmsAfterSaleWarranty(w)
	}

	// Convert recalls
	recalls := make([]DmsAfterSaleRecall, len(sb.Data.ServiceBookingRequest.Recalls))
	for i, r := range sb.Data.ServiceBookingRequest.Recalls {
		recalls[i] = DmsAfterSaleRecall(r)
	}

	// Convert jobs
	jobs := make([]DmsAfterSaleJob, len(sb.Data.Job))
	for i, j := range sb.Data.Job {
		jobs[i] = DmsAfterSaleJob(j)
	}

	// Convert parts
	parts := make([]DmsAfterSalePart, len(sb.Data.Part))
	for i, p := range sb.Data.Part {
		packageParts := make([]DmsAfterSalePartItem, len(p.PackageParts))
		for j, pp := range p.PackageParts {
			packageParts[j] = DmsAfterSalePartItem(pp)
		}

		parts[i] = DmsAfterSalePart{
			PartType:                 p.PartType,
			PackageID:                p.PackageID,
			PartNumber:               p.PartNumber,
			PartName:                 p.PartName,
			PartQuantity:             p.PartQuantity,
			PackageParts:             packageParts,
			PartSize:                 p.PartSize,
			PartColor:                p.PartColor,
			PartEstPrice:             p.PartEstPrice,
			PartInstallationEstPrice: p.PartInstallationEstPrice,
			FlagPartNeedDownPayment:  p.FlagPartNeedDownPayment,
		}
	}

	// Get gender value with default
	gender := ""
	if sb.Data.OneAccount.Gender != nil {
		gender = *sb.Data.OneAccount.Gender
	}

	// Get primary user based on customer type
	primaryUser := "MASTER"
	if sb.Data.OneAccount.CustomerType != nil && *sb.Data.OneAccount.CustomerType != "" {
		primaryUser = *sb.Data.OneAccount.CustomerType
	}

	return DmsAfterSaleServiceBookingEvent{
		Process:   sb.Process,
		EventID:   sb.EventID,
		Timestamp: sb.Timestamp,
		Data: DmsAfterSaleDataRequest{
			OneAccount: DmsAfterSaleOneAccount{
				OneAccountID: sb.Data.OneAccount.OneAccountID,
				FirstName:    sb.Data.OneAccount.FirstName,
				LastName:     sb.Data.OneAccount.LastName,
				Gender:       gender,
				PhoneNumber:  sb.Data.OneAccount.PhoneNumber,
				Email:        sb.Data.OneAccount.Email,
				PrimaryUser:  primaryUser,
			},
			CustomerVehicle: DmsAfterSaleCustomerVehicle{
				VIN:             sb.Data.CustomerVehicle.Vin,
				KatashikiSuffix: sb.Data.CustomerVehicle.KatashikiSuffix,
				Model:           sb.Data.CustomerVehicle.Model,
				Variant:         sb.Data.CustomerVehicle.Variant,
				PoliceNumber:    sb.Data.CustomerVehicle.PoliceNumber,
				ActualMileage:   sb.Data.CustomerVehicle.ActualMileage,
				ColorCode:       sb.Data.CustomerVehicle.ColorCode,
				Color:           sb.Data.CustomerVehicle.Color,
			},
			ServiceBooking: DmsAfterSaleServiceBooking{
				BookingID:                    sb.Data.ServiceBookingRequest.BookingId,
				BookingNumber:                sb.Data.ServiceBookingRequest.BookingNumber,
				BookingSource:                sb.Data.ServiceBookingRequest.BookingSource,
				BookingStatus:                sb.Data.ServiceBookingRequest.BookingStatus,
				CreatedDatetime:              sb.Data.ServiceBookingRequest.CreatedDatetime,
				ServiceCategory:              sb.Data.ServiceBookingRequest.ServiceCategory,
				ServiceSequence:              sb.Data.ServiceBookingRequest.ServiceSequence,
				SlotDatetimeStart:            sb.Data.ServiceBookingRequest.SlotDatetimeStart,
				SlotDatetimeEnd:              sb.Data.ServiceBookingRequest.SlotDatetimeEnd,
				SlotRequestedDatetimeStart:   sb.Data.ServiceBookingRequest.SlotRequestedDatetimeStart,
				SlotRequestedDatetimeEnd:     sb.Data.ServiceBookingRequest.SlotRequestedDatetimeEnd,
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
				Warranty:                     warranties,
				Recalls:                      recalls,
				CancellationReason:           sb.Data.ServiceBookingRequest.CancellationReason,
				OtherCancellationReason:      sb.Data.ServiceBookingRequest.OtherCancellationReason,
				ServicePricingCallFlag:       sb.Data.ServiceBookingRequest.ServicePricingCallFlag,
			},
			Job:  jobs,
			Part: parts,
		},
	}
}
