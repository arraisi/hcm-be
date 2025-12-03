package domain

import "time"

type Appraisal struct {
	ID                     string `json:"id" db:"i_id"`
	AppraisalBookingID     string `json:"appraisal_booking_id" db:"i_appraisal_booking_id"`
	AppraisalBookingNumber string `json:"appraisal_booking_number" db:"c_appraisal_booking_number"`
	AppraisalType          string `json:"appraisal_type" db:"c_appraisal_type"`

	OutletID   string `json:"outlet_id" db:"i_outlet_id"`
	OutletName string `json:"outlet_name" db:"n_outlet_name"`

	OneAccountID string `json:"one_account_id" db:"i_one_account_id"`
	VIN          string `json:"vin" db:"c_vin"`
	LeadsID      string `json:"leads_id" db:"i_leads_id"`

	AppraisalLocation string `json:"appraisal_location" db:"c_appraisal_location"`
	HomeAddress       string `json:"home_address" db:"e_home_address"`
	Province          string `json:"province" db:"e_province"`
	City              string `json:"city" db:"e_city"`
	District          string `json:"district" db:"e_district"`
	Subdistrict       string `json:"subdistrict" db:"e_subdistrict"`
	PostalCode        string `json:"postal_code" db:"e_postal_code"`

	CreatedDatetime        time.Time  `json:"created_datetime" db:"d_created_datetime"`
	AppraisalStartDatetime *time.Time `json:"appraisal_start_datetime" db:"d_appraisal_start_datetime"`
	AppraisalEndDatetime   *time.Time `json:"appraisal_end_datetime" db:"d_appraisal_end_datetime"`

	ConfirmStartDatetime *time.Time `json:"appraisal_confirmation_start_datetime" db:"d_appraisal_confirmation_start_datetime"`
	ConfirmEndDatetime   *time.Time `json:"appraisal_confirmation_end_datetime" db:"d_appraisal_confirmation_end_datetime"`

	BookingStatus           string  `json:"appraisal_booking_status" db:"c_appraisal_booking_status"`
	CancelledBy             *string `json:"cancelled_by" db:"c_cancelled_by"`
	CancellationReason      *string `json:"cancellation_reason" db:"c_cancellation_reason"`
	OtherCancellationReason string  `json:"other_cancellation_reason" db:"e_other_cancellation_reason"`
	BookingServiceFlag      bool    `json:"booking_service_flag" db:"c_booking_service_flag"`

	// vehicle
	KatashikiSuffix string `json:"katashiki_suffix" db:"c_katashiki_suffix"`
	ColorCode       string `json:"color_code" db:"c_color_code"`
	Model           string `json:"model" db:"n_model"`
	Variant         string `json:"variant" db:"n_variant"`
	Color           string `json:"color" db:"n_color"`

	// trade-in latest summary
	FinalTradeInStatus        *string    `json:"final_trade_in_status" db:"c_final_trade_in_status"`
	LastTradeInStatusDatetime *time.Time `json:"last_trade_in_status_datetime" db:"d_last_trade_in_status_datetime"`

	// sales docs
	SPKNumber string `json:"spk_number" db:"c_spk_number"`
	SONumber  string `json:"so_number" db:"c_so_number"`

	// negotiation (snapshot of latest NEGOTIATION)
	CustomerNegotiationPrice           *float64 `json:"customer_negotiation_price" db:"v_customer_negotiation_price"`
	DealerNegotiationPrice             *float64 `json:"dealer_negotiation_price" db:"v_dealer_negotiation_price"`
	DealPrice                          *float64 `json:"deal_price" db:"v_deal_price"`
	DownPaymentEstimation              *float64 `json:"down_payment_estimation" db:"v_down_payment_estimation"`
	EstimatedRemainingPayment          *float64 `json:"estimated_remaining_payment" db:"v_estimated_remaining_payment"`
	NoDealReason                       string   `json:"no_deal_reason" db:"c_no_deal_reason"`
	NoDealReasonOldVehicleOthers       string   `json:"no_deal_reason_old_vehicle_others" db:"e_no_deal_reason_old_vehicle_others"`
	NoDealReasonOldVehicleExpectedSell *float64 `json:"no_deal_reason_old_vehicle_expected_sell_price" db:"v_no_deal_reason_old_vehicle_expected_sell_price"`
	NoDealReasonOldVehiclePriceSold    *float64 `json:"no_deal_reason_old_vehicle_price_sold" db:"v_no_deal_reason_old_vehicle_price_sold"`
	NoDealReasonNewVehicleOthers       string   `json:"no_deal_reason_new_vehicle_others" db:"e_no_deal_reason_new_vehicle_others"`

	// handover (snapshot of latest HANDOVER)
	TradeInPaymentDatetime  *time.Time `json:"trade_in_payment_datetime" db:"d_trade_in_payment_datetime"`
	TradeInHandoverStatus   *string    `json:"trade_in_handover_status" db:"c_trade_in_handover_status"`
	TradeInHandoverDatetime *time.Time `json:"trade_in_handover_datetime" db:"d_trade_in_handover_datetime"`
	TradeInHandoverLocation *string    `json:"trade_in_handover_location" db:"c_trade_in_handover_location"`
	TradeInHandoverAddress  string     `json:"trade_in_handover_address" db:"e_trade_in_handover_address"`
	HandoverProvince        string     `json:"handover_province" db:"e_handover_province"`
	HandoverCity            string     `json:"handover_city" db:"e_handover_city"`
	HandoverDistrict        string     `json:"handover_district" db:"e_handover_district"`
	HandoverSubdistrict     string     `json:"handover_subdistrict" db:"e_handover_subdistrict"`
	HandoverPostalCode      string     `json:"handover_postal_code" db:"e_handover_postal_code"`

	CreatedDate time.Time  `json:"created_date" db:"d_createdate"`
	UpdatedDate *time.Time `json:"updated_date" db:"d_updatedate"`
}

func (a *Appraisal) TableName() string {
	return "dbo.tr_appraisal"
}

func (a *Appraisal) Columns() []string {
	return []string{
		"i_id",
		"i_appraisal_booking_id",
		"c_appraisal_booking_number",
		"c_appraisal_type",
		"i_outlet_id",
		"n_outlet_name",
		"c_appraisal_location",
		"e_home_address",
		"e_province",
		"e_city",
		"e_district",
		"e_subdistrict",
		"e_postal_code",
		"d_created_datetime",
		"d_appraisal_start_datetime",
		"d_appraisal_end_datetime",
		"d_appraisal_confirmation_start_datetime",
		"d_appraisal_confirmation_end_datetime",
		"c_appraisal_booking_status",
		"c_cancelled_by",
		"c_cancellation_reason",
		"e_other_cancellation_reason",
		"c_booking_service_flag",
		"c_katashiki_suffix",
		"c_color_code",
		"n_model",
		"n_variant",
		"n_color",
		"c_final_trade_in_status",
		"d_last_trade_in_status_datetime",
		"c_spk_number",
		"c_so_number",
		"v_customer_negotiation_price",
		"v_dealer_negotiation_price",
		"v_deal_price",
		"v_down_payment_estimation",
		"v_estimated_remaining_payment",
		"c_no_deal_reason",
		"e_no_deal_reason_old_vehicle_others",
		"v_no_deal_reason_old_vehicle_expected_sell_price",
		"v_no_deal_reason_old_vehicle_price_sold",
		"e_no_deal_reason_new_vehicle_others",
		"d_trade_in_payment_datetime",
		"c_trade_in_handover_status",
		"d_trade_in_handover_datetime",
		"c_trade_in_handover_location",
		"e_trade_in_handover_address",
		"e_handover_province",
		"e_handover_city",
		"e_handover_district",
		"e_handover_subdistrict",
		"e_handover_postal_code",
		"d_createdate",
		"d_updatedate",
	}
}

func (a *Appraisal) SelectColumns() []string {
	return []string{
		"i_id",
		"i_appraisal_booking_id",
		"c_appraisal_booking_number",
		"c_appraisal_type",
		"i_outlet_id",
		"n_outlet_name",
		"c_appraisal_location",
		"d_created_datetime",
		"d_appraisal_start_datetime",
		"d_appraisal_end_datetime",
		"c_appraisal_booking_status",
		"c_final_trade_in_status",
		"d_last_trade_in_status_datetime",
		"c_booking_service_flag",
		"d_createdate",
		"d_updatedate",
	}
}

func (a *Appraisal) ToCreateMap() (cols []string, vals []interface{}) {
	cols = make([]string, 0, len(a.Columns()))
	vals = make([]interface{}, 0, len(a.Columns()))

	// mandatory
	cols = append(cols, "i_id")
	vals = append(vals, a.ID)

	cols = append(cols, "i_appraisal_booking_id")
	vals = append(vals, a.AppraisalBookingID)

	cols = append(cols, "c_appraisal_booking_number")
	vals = append(vals, a.AppraisalBookingNumber)

	cols = append(cols, "c_appraisal_type")
	vals = append(vals, a.AppraisalType)

	cols = append(cols, "i_outlet_id")
	vals = append(vals, a.OutletID)

	cols = append(cols, "n_outlet_name")
	vals = append(vals, a.OutletName)

	cols = append(cols, "c_appraisal_location")
	vals = append(vals, a.AppraisalLocation)

	cols = append(cols, "e_home_address")
	vals = append(vals, a.HomeAddress)

	cols = append(cols, "e_province")
	vals = append(vals, a.Province)

	cols = append(cols, "e_city")
	vals = append(vals, a.City)

	cols = append(cols, "e_district")
	vals = append(vals, a.District)

	cols = append(cols, "e_subdistrict")
	vals = append(vals, a.Subdistrict)

	cols = append(cols, "e_postal_code")
	vals = append(vals, a.PostalCode)

	if !a.CreatedDatetime.IsZero() {
		cols = append(cols, "d_created_datetime")
		vals = append(vals, a.CreatedDatetime.UTC())
	}

	if a.AppraisalStartDatetime != nil {
		cols = append(cols, "d_appraisal_start_datetime")
		vals = append(vals, a.AppraisalStartDatetime.UTC())
	}
	if a.AppraisalEndDatetime != nil {
		cols = append(cols, "d_appraisal_end_datetime")
		vals = append(vals, a.AppraisalEndDatetime.UTC())
	}
	if a.ConfirmStartDatetime != nil {
		cols = append(cols, "d_appraisal_confirmation_start_datetime")
		vals = append(vals, a.ConfirmStartDatetime.UTC())
	}
	if a.ConfirmEndDatetime != nil {
		cols = append(cols, "d_appraisal_confirmation_end_datetime")
		vals = append(vals, a.ConfirmEndDatetime.UTC())
	}

	cols = append(cols, "c_appraisal_booking_status")
	vals = append(vals, string(a.BookingStatus))

	if a.CancelledBy != nil {
		cols = append(cols, "c_cancelled_by")
		vals = append(vals, a.CancelledBy)
	}
	if a.CancellationReason != nil {
		cols = append(cols, "c_cancellation_reason")
		vals = append(vals, a.CancellationReason)
	}
	if a.OtherCancellationReason != "" {
		cols = append(cols, "e_other_cancellation_reason")
		vals = append(vals, a.OtherCancellationReason)
	}

	cols = append(cols, "c_booking_service_flag")
	vals = append(vals, a.BookingServiceFlag)

	if a.KatashikiSuffix != "" {
		cols = append(cols, "c_katashiki_suffix")
		vals = append(vals, a.KatashikiSuffix)
	}
	if a.ColorCode != "" {
		cols = append(cols, "c_color_code")
		vals = append(vals, a.ColorCode)
	}
	if a.Model != "" {
		cols = append(cols, "n_model")
		vals = append(vals, a.Model)
	}
	if a.Variant != "" {
		cols = append(cols, "n_variant")
		vals = append(vals, a.Variant)
	}
	if a.Color != "" {
		cols = append(cols, "n_color")
		vals = append(vals, a.Color)
	}

	if a.FinalTradeInStatus != nil {
		cols = append(cols, "c_final_trade_in_status")
		vals = append(vals, a.FinalTradeInStatus)
	}
	if a.LastTradeInStatusDatetime != nil {
		cols = append(cols, "d_last_trade_in_status_datetime")
		vals = append(vals, a.LastTradeInStatusDatetime.UTC())
	}

	if a.SPKNumber != "" {
		cols = append(cols, "c_spk_number")
		vals = append(vals, a.SPKNumber)
	}
	if a.SONumber != "" {
		cols = append(cols, "c_so_number")
		vals = append(vals, a.SONumber)
	}

	// negotiation
	if a.CustomerNegotiationPrice != nil {
		cols = append(cols, "v_customer_negotiation_price")
		vals = append(vals, a.CustomerNegotiationPrice)
	}
	if a.DealerNegotiationPrice != nil {
		cols = append(cols, "v_dealer_negotiation_price")
		vals = append(vals, a.DealerNegotiationPrice)
	}
	if a.DealPrice != nil {
		cols = append(cols, "v_deal_price")
		vals = append(vals, a.DealPrice)
	}
	if a.DownPaymentEstimation != nil {
		cols = append(cols, "v_down_payment_estimation")
		vals = append(vals, a.DownPaymentEstimation)
	}
	if a.EstimatedRemainingPayment != nil {
		cols = append(cols, "v_estimated_remaining_payment")
		vals = append(vals, a.EstimatedRemainingPayment)
	}
	if a.NoDealReason != "" {
		cols = append(cols, "c_no_deal_reason")
		vals = append(vals, a.NoDealReason)
	}
	if a.NoDealReasonOldVehicleOthers != "" {
		cols = append(cols, "e_no_deal_reason_old_vehicle_others")
		vals = append(vals, a.NoDealReasonOldVehicleOthers)
	}
	if a.NoDealReasonOldVehicleExpectedSell != nil {
		cols = append(cols, "v_no_deal_reason_old_vehicle_expected_sell_price")
		vals = append(vals, a.NoDealReasonOldVehicleExpectedSell)
	}
	if a.NoDealReasonOldVehiclePriceSold != nil {
		cols = append(cols, "v_no_deal_reason_old_vehicle_price_sold")
		vals = append(vals, a.NoDealReasonOldVehiclePriceSold)
	}
	if a.NoDealReasonNewVehicleOthers != "" {
		cols = append(cols, "e_no_deal_reason_new_vehicle_others")
		vals = append(vals, a.NoDealReasonNewVehicleOthers)
	}

	// handover
	if a.TradeInPaymentDatetime != nil {
		cols = append(cols, "d_trade_in_payment_datetime")
		vals = append(vals, a.TradeInPaymentDatetime.UTC())
	}
	if a.TradeInHandoverStatus != nil {
		cols = append(cols, "c_trade_in_handover_status")
		vals = append(vals, a.TradeInHandoverStatus)
	}
	if a.TradeInHandoverDatetime != nil {
		cols = append(cols, "d_trade_in_handover_datetime")
		vals = append(vals, a.TradeInHandoverDatetime.UTC())
	}
	if a.TradeInHandoverLocation != nil {
		cols = append(cols, "c_trade_in_handover_location")
		vals = append(vals, a.TradeInHandoverLocation)
	}
	if a.TradeInHandoverAddress != "" {
		cols = append(cols, "e_trade_in_handover_address")
		vals = append(vals, a.TradeInHandoverAddress)
	}
	if a.HandoverProvince != "" {
		cols = append(cols, "e_handover_province")
		vals = append(vals, a.HandoverProvince)
	}
	if a.HandoverCity != "" {
		cols = append(cols, "e_handover_city")
		vals = append(vals, a.HandoverCity)
	}
	if a.HandoverDistrict != "" {
		cols = append(cols, "e_handover_district")
		vals = append(vals, a.HandoverDistrict)
	}
	if a.HandoverSubdistrict != "" {
		cols = append(cols, "e_handover_subdistrict")
		vals = append(vals, a.HandoverSubdistrict)
	}
	if a.HandoverPostalCode != "" {
		cols = append(cols, "e_handover_postal_code")
		vals = append(vals, a.HandoverPostalCode)
	}

	// audit
	if !a.CreatedDate.IsZero() {
		cols = append(cols, "d_createdate")
		vals = append(vals, a.CreatedDate.UTC())
	}
	if a.UpdatedDate != nil {
		cols = append(cols, "d_updatedate")
		vals = append(vals, a.UpdatedDate.UTC())
	}

	return cols, vals
}

func (a *Appraisal) ToUpdateMap() map[string]interface{} {
	m := make(map[string]interface{})

	// status & booking
	m["c_appraisal_booking_status"] = string(a.BookingStatus)

	if a.CancelledBy != nil {
		m["c_cancelled_by"] = a.CancelledBy
	}
	if a.CancellationReason != nil {
		m["c_cancellation_reason"] = a.CancellationReason
	}
	if a.OtherCancellationReason != "" {
		m["e_other_cancellation_reason"] = a.OtherCancellationReason
	}

	if a.ConfirmStartDatetime != nil {
		m["d_appraisal_confirmation_start_datetime"] = a.ConfirmStartDatetime.UTC()
	}
	if a.ConfirmEndDatetime != nil {
		m["d_appraisal_confirmation_end_datetime"] = a.ConfirmEndDatetime.UTC()
	}

	// trade-in summary
	if a.FinalTradeInStatus != nil {
		m["c_final_trade_in_status"] = a.FinalTradeInStatus
	}
	if a.LastTradeInStatusDatetime != nil {
		m["d_last_trade_in_status_datetime"] = a.LastTradeInStatusDatetime.UTC()
	}

	// negotiation fields
	if a.CustomerNegotiationPrice != nil {
		m["v_customer_negotiation_price"] = a.CustomerNegotiationPrice
	}
	if a.DealerNegotiationPrice != nil {
		m["v_dealer_negotiation_price"] = a.DealerNegotiationPrice
	}
	if a.DealPrice != nil {
		m["v_deal_price"] = a.DealPrice
	}
	if a.DownPaymentEstimation != nil {
		m["v_down_payment_estimation"] = a.DownPaymentEstimation
	}
	if a.EstimatedRemainingPayment != nil {
		m["v_estimated_remaining_payment"] = a.EstimatedRemainingPayment
	}
	if a.NoDealReason != "" {
		m["c_no_deal_reason"] = a.NoDealReason
	}
	if a.NoDealReasonOldVehicleOthers != "" {
		m["e_no_deal_reason_old_vehicle_others"] = a.NoDealReasonOldVehicleOthers
	}
	if a.NoDealReasonOldVehicleExpectedSell != nil {
		m["v_no_deal_reason_old_vehicle_expected_sell_price"] = a.NoDealReasonOldVehicleExpectedSell
	}
	if a.NoDealReasonOldVehiclePriceSold != nil {
		m["v_no_deal_reason_old_vehicle_price_sold"] = a.NoDealReasonOldVehiclePriceSold
	}
	if a.NoDealReasonNewVehicleOthers != "" {
		m["e_no_deal_reason_new_vehicle_others"] = a.NoDealReasonNewVehicleOthers
	}

	// handover
	if a.TradeInPaymentDatetime != nil {
		m["d_trade_in_payment_datetime"] = a.TradeInPaymentDatetime.UTC()
	}
	if a.TradeInHandoverStatus != nil {
		m["c_trade_in_handover_status"] = a.TradeInHandoverStatus
	}
	if a.TradeInHandoverDatetime != nil {
		m["d_trade_in_handover_datetime"] = a.TradeInHandoverDatetime.UTC()
	}
	if a.TradeInHandoverLocation != nil {
		m["c_trade_in_handover_location"] = a.TradeInHandoverLocation
	}
	if a.TradeInHandoverAddress != "" {
		m["e_trade_in_handover_address"] = a.TradeInHandoverAddress
	}
	if a.HandoverProvince != "" {
		m["e_handover_province"] = a.HandoverProvince
	}
	if a.HandoverCity != "" {
		m["e_handover_city"] = a.HandoverCity
	}
	if a.HandoverDistrict != "" {
		m["e_handover_district"] = a.HandoverDistrict
	}
	if a.HandoverSubdistrict != "" {
		m["e_handover_subdistrict"] = a.HandoverSubdistrict
	}
	if a.HandoverPostalCode != "" {
		m["e_handover_postal_code"] = a.HandoverPostalCode
	}

	// audit
	if a.UpdatedDate != nil {
		m["d_updatedate"] = a.UpdatedDate.UTC()
	}

	return m
}

//
// STATUS UPDATE HISTORY MODEL – matches tr_appraisal_status_update
//

type AppraisalStatusUpdate struct {
	ID                    string    `json:"id" db:"i_id"`
	AppraisalID           string    `json:"appraisal_id" db:"i_appraisal_id"`
	TradeInStatus         string    `json:"trade_in_status" db:"c_trade_in_status"`
	TradeInStatusDatetime time.Time `json:"trade_in_status_datetime" db:"d_trade_in_status_datetime"`
	CreatedDate           time.Time `json:"created_date" db:"d_createdate"`
}

func (s *AppraisalStatusUpdate) TableName() string {
	return "dbo.tr_appraisal_status_update"
}

func (s *AppraisalStatusUpdate) Columns() []string {
	return []string{
		"i_id",
		"i_appraisal_id",
		"c_trade_in_status",
		"d_trade_in_status_datetime",
		"d_createdate",
	}
}

func (s *AppraisalStatusUpdate) SelectColumns() []string {
	return []string{
		"i_id",
		"i_appraisal_id",
		"c_trade_in_status",
		"d_trade_in_status_datetime",
		"d_createdate",
	}
}

func (s *AppraisalStatusUpdate) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(s.Columns()))
	values = make([]interface{}, 0, len(s.Columns()))

	// mandatory
	columns = append(columns, "i_id")
	values = append(values, s.ID)

	columns = append(columns, "i_appraisal_id")
	values = append(values, s.AppraisalID)

	columns = append(columns, "c_trade_in_status")
	values = append(values, s.TradeInStatus)

	if !s.TradeInStatusDatetime.IsZero() {
		columns = append(columns, "d_trade_in_status_datetime")
		values = append(values, s.TradeInStatusDatetime.UTC())
	}

	return columns, values
}
