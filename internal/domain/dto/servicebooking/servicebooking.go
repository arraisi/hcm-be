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
	Process   string        `json:"process"`
	EventID   string        `json:"event_id"`
	Timestamp int           `json:"timestamp"`
	Data      DataRequest   `json:"data"`
	Job       []JobRequest  `json:"job"`
	Part      []PartRequest `json:"part"`
}

type DataRequest struct {
	OneAccount                   customer.OneAccountRequest             `json:"one_account"`
	CustomerVehicle              customervehicle.CustomerVehicleRequest `json:"customer_vehicle"`
	BookingId                    string                                 `json:"booking_id"`
	BookingNumber                string                                 `json:"booking_number"`
	BookingSource                string                                 `json:"booking_source"`
	BookingStatus                string                                 `json:"booking_status"`
	CreatedDatetime              int64                                  `json:"created_datetime"`
	ServiceCategory              string                                 `json:"service_category"`
	ServiceSequence              string                                 `json:"service_sequence"`
	SlotDatetimeStart            int64                                  `json:"slot_datetime_start"`
	SlotDatetimeEnd              int64                                  `json:"slot_datetime_end"`
	SlotRequestedDatetimeStart   int64                                  `json:"slot_requested_datetime_start"`
	SlotRequestedDatetimeEnd     int64                                  `json:"slot_requested_datetime_end"`
	SlotUnavailableFlag          bool                                   `json:"slot_unavailable_flag"`
	CarrierName                  string                                 `json:"carrier_name"`
	CarrierPhoneNumber           string                                 `json:"carrier_phone_number"`
	PreferenceContactPhoneNumber string                                 `json:"preference_contact_phone_number"`
	PreferenceContactTimeStart   string                                 `json:"preference_contact_time_start"`
	PreferenceContactTimeEnd     string                                 `json:"preference_contact_time_end"`
	ServiceLocation              string                                 `json:"service_location"`
	OutletID                     string                                 `json:"outlet_id"`
	OutletName                   string                                 `json:"outlet_name"`
	MobileServiceAddress         string                                 `json:"mobile_service_address"`
	Province                     string                                 `json:"province"`
	City                         string                                 `json:"city"`
	District                     string                                 `json:"district"`
	SubDistrict                  string                                 `json:"sub_district"`
	PostalCode                   string                                 `json:"postal_code"`
	VehicleProblem               string                                 `json:"vehicle_problem"`
	Warranty                     []WarrantyRequest                      `json:"warranty"`
	Recalls                      []RecallRequest                        `json:"recalls"`
	CancellationReason           string                                 `json:"cancellation_reason"`
	OtherCancellationReason      string                                 `json:"other_cancellation_reason"`
	ServicePricingCallFlag       bool                                   `json:"service_pricing_call_flag"`
}

type WarrantyRequest struct {
	WarrantyName   string `json:"warranty_name"`
	WarrantyStatus string `json:"warranty_status"`
}

func (wr WarrantyRequest) ToModel(serviceBookingID string) domain.ServiceBookingWarranty {
	now := time.Now()
	return domain.ServiceBookingWarranty{
		ServiceBookingID: serviceBookingID,
		WarrantyName:     wr.WarrantyName,
		WarrantyStatus:   wr.WarrantyStatus,
		CreatedAt:        now.UTC(),
		CreatedBy:        constants.System,
		UpdatedAt:        now.UTC(),
		UpdatedBy:        constants.System,
	}
}

type RecallRequest struct {
	RecallID          string   `json:"recall_id"`
	RecallDate        string   `json:"recall_date"`
	RecallDescription string   `json:"recall_description"`
	AffectedParts     []string `json:"affected_parts"`
}

func (r *RecallRequest) ToModel(bookingID, part string) domain.ServiceBookingRecall {
	now := time.Now()
	return domain.ServiceBookingRecall{
		ServiceBookingID:  bookingID,
		RecallID:          r.RecallID,
		RecallDate:        r.RecallDate,
		RecallDescription: r.RecallDescription,
		AffectedPart:      part,
		CreatedAt:         now.UTC(),
		CreatedBy:         constants.System,
		UpdatedAt:         now.UTC(),
		UpdatedBy:         constants.System,
	}
}

type JobRequest struct {
	JobID         string `json:"job_ID"`
	JobName       string `json:"job_name"`
	LaborEstPrice string `json:"labor_est_price"`
}

func (j *JobRequest) ToDomain(serviceBookingID string) domain.ServiceBookingJob {
	now := time.Now()
	return domain.ServiceBookingJob{
		ServiceBookingID: serviceBookingID,
		JobID:            j.JobID,
		JobName:          j.JobName,
		LaborEstPrice:    j.LaborEstPrice,
		CreatedAt:        now.UTC(),
		CreatedBy:        constants.System,
		UpdatedAt:        now.UTC(),
		UpdatedBy:        constants.System,
	}
}

type PartRequest struct {
	PartType                 string            `json:"part_type"`
	PackageID                string            `json:"package_ID"`
	PartNumber               string            `json:"part_number"`
	PartName                 string            `json:"part_name"`
	PartQuantity             int32             `json:"part_quantity"`
	PackageParts             []PartItemRequest `json:"package_parts"`
	PartSize                 string            `json:"part_size"`
	PartColor                string            `json:"part_color"`
	PartEstPrice             float32           `json:"part_est_price"`
	PartInstallationEstPrice float32           `json:"part_installation_est_price"`
	FlagPartNeedDownPayment  bool              `json:"flag_part_need_down_payment"`
}

func (p *PartRequest) ToDomain(serviceBookingID string) (domain.ServiceBookingPart, []domain.ServiceBookingPartItem) {
	now := time.Now()

	partItems := make([]domain.ServiceBookingPartItem, 0, len(p.PackageParts))
	for _, item := range p.PackageParts {
		partItems = append(partItems, domain.ServiceBookingPartItem{
			PartNumber: item.PartNumber,
			PartName:   item.PartName,
			CreatedAt:  now.UTC(),
			CreatedBy:  constants.System,
			UpdatedAt:  now.UTC(),
			UpdatedBy:  constants.System,
		})
	}

	return domain.ServiceBookingPart{
		ServiceBookingID:         serviceBookingID,
		PartType:                 p.PartType,
		PackageID:                p.PackageID,
		PartNumber:               p.PartNumber,
		PartName:                 p.PartName,
		PartQuantity:             p.PartQuantity,
		PartSize:                 p.PartSize,
		PartColor:                p.PartColor,
		PartEstPrice:             p.PartEstPrice,
		PartInstallationEstPrice: p.PartInstallationEstPrice,
		FlagPartNeedDownPayment:  p.FlagPartNeedDownPayment,
		CreatedAt:                now.UTC(),
		CreatedBy:                constants.System,
		UpdatedAt:                now.UTC(),
		UpdatedBy:                constants.System,
	}, partItems
}

type PartItemRequest struct {
	PartNumber string `json:"part_number"`
	PartName   string `json:"part_name"`
}

// ToServiceBookingModel converts the DataRequest to the domain.TestDrive model
func (sb *ServiceBookingEvent) ToServiceBookingModel(customerID, customerVehicleID string) domain.ServiceBooking {
	return domain.ServiceBooking{
		EventID:                      sb.EventID,
		CustomerID:                   customerID,
		CustomerVehicleID:            customerVehicleID,
		BookingID:                    sb.Data.BookingId,
		BookingNumber:                sb.Data.BookingNumber,
		BookingSource:                sb.Data.BookingSource,
		BookingStatus:                sb.Data.BookingStatus,
		CreatedDatetime:              utils.GetTimeUnix(sb.Data.CreatedDatetime).UTC(),
		ServiceCategory:              sb.Data.ServiceCategory,
		ServiceSequence:              sb.Data.ServiceSequence,
		SlotDatetimeStart:            utils.GetTimeUnix(sb.Data.SlotDatetimeStart).UTC(),
		SlotDatetimeEnd:              utils.GetTimeUnix(sb.Data.SlotDatetimeEnd).UTC(),
		SlotRequestedDatetimeStart:   utils.GetTimeUnix(sb.Data.SlotRequestedDatetimeStart).UTC(),
		SlotRequestedDatetimeEnd:     utils.GetTimeUnix(sb.Data.SlotRequestedDatetimeEnd).UTC(),
		SlotUnavailableFlag:          sb.Data.SlotUnavailableFlag,
		CarrierName:                  sb.Data.CarrierName,
		CarrierPhoneNumber:           sb.Data.CarrierPhoneNumber,
		PreferenceContactPhoneNumber: sb.Data.PreferenceContactPhoneNumber,
		PreferenceContactTimeStart:   sb.Data.PreferenceContactTimeStart,
		PreferenceContactTimeEnd:     sb.Data.PreferenceContactTimeEnd,
		ServiceLocation:              sb.Data.ServiceLocation,
		OutletID:                     sb.Data.OutletID,
		OutletName:                   sb.Data.OutletName,
		MobileServiceAddress:         sb.Data.MobileServiceAddress,
		Province:                     sb.Data.Province,
		City:                         sb.Data.City,
		District:                     sb.Data.District,
		SubDistrict:                  sb.Data.SubDistrict,
		PostalCode:                   sb.Data.PostalCode,
		VehicleProblem:               sb.Data.VehicleProblem,
		CancellationReason:           sb.Data.CancellationReason,
		OtherCancellationReason:      sb.Data.OtherCancellationReason,
		ServicePricingCallFlag:       sb.Data.ServicePricingCallFlag,
		CreatedAt:                    time.Now().UTC(),
		CreatedBy:                    constants.System, // or fetch from context if available
		UpdatedAt:                    time.Now().UTC(),
		UpdatedBy:                    constants.System, // or fetch from context if available
	}
}

type GetServiceBooking struct {
	ID                   *string
	ServiceBookingID     *string
	ServiceBookingNumber *string
	ServiceBookingSource *string
	EventID              *string
}

func (g *GetServiceBooking) Apply(q *sqrl.SelectBuilder) {
	if g.ID != nil {
		q.Where(sqrl.Eq{"i_id": g.ID})
	}
	if g.ServiceBookingID != nil {
		q.Where(sqrl.Eq{"i_service_booking_id": g.ServiceBookingID})
	}
	if g.ServiceBookingNumber != nil {
		q.Where(sqrl.Eq{"c_service_booking_number": g.ServiceBookingNumber})
	}
	if g.EventID != nil {
		q.Where(sqrl.Eq{"i_event_id": g.EventID})
	}
}
