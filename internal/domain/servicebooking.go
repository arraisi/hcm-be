package domain

import (
	"time"
)

type ServiceBooking struct {
	ID                           string    `db:"i_id"`
	CustomerID                   string    `db:"i_customer_id"`
	CustomerVehicleID            string    `db:"i_customer_vehicle_id"`
	ServiceBookingID             string    `db:"i_service_booking_id"`
	ServiceBookingNumber         string    `db:"c_service_booking_number"`
	ServiceBookingSource         string    `db:"c_service_booking_source"`
	ServiceBookingStatus         string    `db:"c_service_booking_status"`
	CreatedDatetime              time.Time `db:"d_created_datetime"`
	ServiceCategory              string    `db:"c_service_category"`
	ServiceSequence              int32     `db:"c_service_sequence"`
	SlotDatetimeStart            time.Time `db:"d_slot_datetime_start"`
	SlotDatetimeEnd              time.Time `db:"d_slot_datetime_end"`
	SlotRequestedDatetimeStart   time.Time `db:"d_slot_requested_datetime_start"`
	SlotRequestedDatetimeEnd     time.Time `db:"d_slot_requested_datetime_end"`
	SlotUnavailableFlag          bool      `db:"b_slot_unavailable_flag"`
	CarrierName                  string    `db:"n_carrier_name"`
	CarrierPhoneNumber           string    `db:"c_carrier_phone_number"`
	PreferenceContactPhoneNumber string    `db:"c_preference_contact_phone_number"`
	PreferenceContactTimeStart   string    `db:"t_preference_contact_time_start"`
	PreferenceContactTimeEnd     string    `db:"t_preference_contact_time_end"`
	ServiceLocation              string    `db:"c_service_location"`
	OutletID                     string    `db:"c_outlet_id"`
	OutletName                   string    `db:"n_outlet_name"`
	MobileServiceAddress         string    `db:"e_mobile_service_address"`
	Province                     string    `db:"c_province"`
	City                         string    `db:"c_city"`
	District                     string    `db:"c_district"`
	SubDistrict                  string    `db:"c_subdistrict"`
	PostalCode                   string    `db:"c_postal_code"`
	VehicleProblem               string    `db:"e_vehicle_problem"`
	CancellationReason           string    `db:"e_cancellation_reason"`
	OtherCancellationReason      string    `db:"e_other_cancellation_reason"`
	ServicePricingCallFlag       bool      `db:"b_service_pricing_call_flag"`
	CreatedAt                    time.Time `db:"d_created_at"`
	CreatedBy                    string    `db:"c_created_by"`
	UpdatedAt                    time.Time `db:"d_updated_at"`
	UpdatedBy                    string    `db:"c_updated_by"`
	EventID                      string    `db:"i_event_id"`
	AdditionalVehicleProblem     string    `db:"e_additional_vehicle_problem"`
	InsuranceClaim               string    `db:"c_insurance_claim"`
}

// TableName returns the database table name for the ServiceBooking model
func (sb *ServiceBooking) TableName() string {
	return "dbo.tm_service_booking"
}

// Columns returns the list of database columns for the ServiceBooking model
func (sb *ServiceBooking) Columns() []string {
	return []string{
		"i_id",
		"i_event_id",
		"i_customer_id",
		"i_customer_vehicle_id",
		"i_service_booking_id",
		"c_service_booking_number",
		"c_service_booking_source",
		"c_service_booking_status",
		"d_created_datetime",
		"c_service_category",
		"c_service_sequence",
		"d_slot_datetime_start",
		"d_slot_datetime_end",
		"d_slot_requested_datetime_start",
		"d_slot_requested_datetime_end",
		"b_slot_unavailable_flag",
		"n_carrier_name",
		"c_carrier_phone_number",
		"c_preference_contact_phone_number",
		"t_preference_contact_time_start",
		"t_preference_contact_time_end",
		"c_service_location",
		"c_outlet_id",
		"n_outlet_name",
		"e_mobile_service_address",
		"c_province",
		"c_city",
		"c_district",
		"c_subdistrict",
		"c_postal_code",
		"e_vehicle_problem",
		"e_cancellation_reason",
		"e_other_cancellation_reason",
		"b_service_pricing_call_flag",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the ServiceBooking model
func (sb *ServiceBooking) SelectColumns() []string {
	return []string{
		"i_id",
		"i_customer_id",
		"i_customer_vehicle_id",
		"i_service_booking_id",
		"i_event_id",
		"c_service_booking_number",
		"c_service_booking_source",
		"c_service_booking_status",
		"d_created_datetime",
		"c_service_category",
		"c_service_sequence",
		"d_slot_datetime_start",
		"d_slot_datetime_end",
		"d_slot_requested_datetime_start",
		"d_slot_requested_datetime_end",
		"b_slot_unavailable_flag",
		"n_carrier_name",
		"c_carrier_phone_number",
		"c_preference_contact_phone_number",
		"t_preference_contact_time_start",
		"t_preference_contact_time_end",
		"c_service_location",
		"c_outlet_id",
		"n_outlet_name",
		"e_mobile_service_address",
		"c_province",
		"c_city",
		"c_district",
		"c_subdistrict",
		"c_postal_code",
		"e_vehicle_problem",
		"e_cancellation_reason",
		"e_other_cancellation_reason",
		"b_service_pricing_call_flag",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap converts the ServiceBooking struct to a map of columns and values for insertion
func (sb *ServiceBooking) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0, len(sb.Columns()))
	values := make([]interface{}, 0, len(sb.Columns()))

	if sb.EventID != "" {
		columns = append(columns, "i_event_id")
		values = append(values, sb.EventID)
	}
	if sb.CustomerID != "" {
		columns = append(columns, "i_customer_id")
		values = append(values, sb.CustomerID)
	}
	if sb.CustomerVehicleID != "" {
		columns = append(columns, "i_customer_vehicle_id")
		values = append(values, sb.CustomerVehicleID)
	}
	if sb.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, sb.ServiceBookingID)
	}
	if sb.ServiceBookingNumber != "" {
		columns = append(columns, "c_service_booking_number")
		values = append(values, sb.ServiceBookingNumber)
	}
	if sb.ServiceBookingSource != "" {
		columns = append(columns, "c_service_booking_source")
		values = append(values, sb.ServiceBookingSource)
	}
	if sb.ServiceBookingStatus != "" {
		columns = append(columns, "c_service_booking_status")
		values = append(values, sb.ServiceBookingStatus)
	}
	if !sb.CreatedDatetime.IsZero() {
		columns = append(columns, "d_created_datetime")
		values = append(values, sb.CreatedDatetime.UTC())
	}
	if sb.ServiceCategory != "" {
		columns = append(columns, "c_service_category")
		values = append(values, sb.ServiceCategory)
	}
	if sb.ServiceSequence != 0 {
		columns = append(columns, "c_service_sequence")
		values = append(values, sb.ServiceSequence)
	}
	if !sb.SlotDatetimeStart.IsZero() {
		columns = append(columns, "d_slot_datetime_start")
		values = append(values, sb.SlotDatetimeStart.UTC())
	}
	if !sb.SlotDatetimeEnd.IsZero() {
		columns = append(columns, "d_slot_datetime_end")
		values = append(values, sb.SlotDatetimeEnd.UTC())
	}
	if !sb.SlotRequestedDatetimeStart.IsZero() {
		columns = append(columns, "d_slot_requested_datetime_start")
		values = append(values, sb.SlotRequestedDatetimeStart.UTC())
	}
	if !sb.SlotRequestedDatetimeEnd.IsZero() {
		columns = append(columns, "d_slot_requested_datetime_end")
		values = append(values, sb.SlotRequestedDatetimeEnd.UTC())
	}
	columns = append(columns, "b_slot_unavailable_flag")
	values = append(values, sb.SlotUnavailableFlag)
	if sb.CarrierName != "" {
		columns = append(columns, "n_carrier_name")
		values = append(values, sb.CarrierName)
	}
	if sb.CarrierPhoneNumber != "" {
		columns = append(columns, "c_carrier_phone_number")
		values = append(values, sb.CarrierPhoneNumber)
	}
	if sb.PreferenceContactPhoneNumber != "" {
		columns = append(columns, "c_preference_contact_phone_number")
		values = append(values, sb.PreferenceContactPhoneNumber)
	}
	if sb.PreferenceContactTimeStart != "" {
		columns = append(columns, "t_preference_contact_time_start")
		values = append(values, sb.PreferenceContactTimeStart)
	}
	if sb.PreferenceContactTimeEnd != "" {
		columns = append(columns, "t_preference_contact_time_end")
		values = append(values, sb.PreferenceContactTimeEnd)
	}
	if sb.ServiceLocation != "" {
		columns = append(columns, "c_service_location")
		values = append(values, sb.ServiceLocation)
	}
	if sb.OutletID != "" {
		columns = append(columns, "c_outlet_id")
		values = append(values, sb.OutletID)
	}
	if sb.OutletName != "" {
		columns = append(columns, "n_outlet_name")
		values = append(values, sb.OutletName)
	}
	if sb.MobileServiceAddress != "" {
		columns = append(columns, "e_mobile_service_address")
		values = append(values, sb.MobileServiceAddress)
	}
	if sb.Province != "" {
		columns = append(columns, "c_province")
		values = append(values, sb.Province)
	}
	if sb.City != "" {
		columns = append(columns, "c_city")
		values = append(values, sb.City)
	}
	if sb.District != "" {
		columns = append(columns, "c_district")
		values = append(values, sb.District)
	}
	if sb.SubDistrict != "" {
		columns = append(columns, "c_subdistrict")
		values = append(values, sb.SubDistrict)
	}
	if sb.PostalCode != "" {
		columns = append(columns, "c_postal_code")
		values = append(values, sb.PostalCode)
	}
	if sb.VehicleProblem != "" {
		columns = append(columns, "e_vehicle_problem")
		values = append(values, sb.VehicleProblem)
	}
	if sb.CancellationReason != "" {
		columns = append(columns, "e_cancellation_reason")
		values = append(values, sb.CancellationReason)
	}
	if sb.OtherCancellationReason != "" {
		columns = append(columns, "e_other_cancellation_reason")
		values = append(values, sb.OtherCancellationReason)
	}
	columns = append(columns, "b_service_pricing_call_flag")
	values = append(values, sb.ServicePricingCallFlag)

	columns = append(columns, "c_created_by")
	values = append(values, sb.CreatedBy)

	columns = append(columns, "c_updated_by")
	values = append(values, sb.CreatedBy)

	return columns, values
}

// ToUpdateMap converts the ServiceBooking struct to a map of columns and values for updating
func (sb *ServiceBooking) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if sb.CustomerID != "" {
		updateMap["i_customer_id"] = sb.CustomerID
	}
	if sb.CustomerVehicleID != "" {
		updateMap["i_customer_vehicle_id"] = sb.CustomerVehicleID
	}
	if sb.ServiceBookingID != "" {
		updateMap["i_service_booking_id"] = sb.ServiceBookingID
	}
	if sb.ServiceBookingNumber != "" {
		updateMap["c_service_booking_number"] = sb.ServiceBookingNumber
	}
	if sb.ServiceBookingSource != "" {
		updateMap["c_service_booking_source"] = sb.ServiceBookingSource
	}
	if sb.ServiceBookingStatus != "" {
		updateMap["c_service_booking_status"] = sb.ServiceBookingStatus
	}
	if !sb.CreatedDatetime.IsZero() {
		updateMap["d_created_datetime"] = sb.CreatedDatetime.UTC()
	}
	if sb.ServiceCategory != "" {
		updateMap["c_service_category"] = sb.ServiceCategory
	}
	if sb.ServiceSequence != 0 {
		updateMap["c_service_sequence"] = sb.ServiceSequence
	}
	if !sb.SlotDatetimeStart.IsZero() {
		updateMap["d_slot_datetime_start"] = sb.SlotDatetimeStart.UTC()
	}
	if !sb.SlotDatetimeEnd.IsZero() {
		updateMap["d_slot_datetime_end"] = sb.SlotDatetimeEnd.UTC()
	}
	if !sb.SlotRequestedDatetimeStart.IsZero() {
		updateMap["d_slot_requested_datetime_start"] = sb.SlotRequestedDatetimeStart.UTC()
	}
	if !sb.SlotRequestedDatetimeEnd.IsZero() {
		updateMap["d_slot_requested_datetime_end"] = sb.SlotRequestedDatetimeEnd.UTC()
	}
	if sb.CarrierName != "" {
		updateMap["n_carrier_name"] = sb.CarrierName
	}
	if sb.CarrierPhoneNumber != "" {
		updateMap["c_carrier_phone_number"] = sb.CarrierPhoneNumber
	}
	if sb.PreferenceContactPhoneNumber != "" {
		updateMap["c_preference_contact_phone_number"] = sb.PreferenceContactPhoneNumber
	}
	if sb.PreferenceContactTimeStart != "" {
		updateMap["t_preference_contact_time_start"] = sb.PreferenceContactTimeStart
	}
	if sb.PreferenceContactTimeEnd != "" {
		updateMap["t_preference_contact_time_end"] = sb.PreferenceContactTimeEnd
	}
	if sb.ServiceLocation != "" {
		updateMap["c_service_location"] = sb.ServiceLocation
	}
	if sb.OutletID != "" {
		updateMap["c_outlet_id"] = sb.OutletID
	}
	if sb.OutletName != "" {
		updateMap["n_outlet_name"] = sb.OutletName
	}
	if sb.MobileServiceAddress != "" {
		updateMap["e_mobile_service_address"] = sb.MobileServiceAddress
	}
	if sb.Province != "" {
		updateMap["c_province"] = sb.Province
	}
	if sb.City != "" {
		updateMap["c_city"] = sb.City
	}
	if sb.District != "" {
		updateMap["c_district"] = sb.District
	}
	if sb.SubDistrict != "" {
		updateMap["c_subdistrict"] = sb.SubDistrict
	}
	if sb.PostalCode != "" {
		updateMap["c_postal_code"] = sb.PostalCode
	}
	if sb.VehicleProblem != "" {
		updateMap["e_vehicle_problem"] = sb.VehicleProblem
	}
	if sb.CancellationReason != "" {
		updateMap["e_cancellation_reason"] = sb.CancellationReason
	}
	if sb.OtherCancellationReason != "" {
		updateMap["e_other_cancellation_reason"] = sb.OtherCancellationReason
	}

	updateMap["b_slot_unavailable_flag"] = sb.SlotUnavailableFlag
	updateMap["b_service_pricing_call_flag"] = sb.ServicePricingCallFlag
	updateMap["d_updated_at"] = time.Now().UTC()
	updateMap["c_updated_by"] = sb.UpdatedBy

	return updateMap
}
