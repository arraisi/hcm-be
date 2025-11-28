package domain

import (
	"time"
)

// SalesOrder represents a vehicle sales order
type SalesOrder struct {
	ID                               string     `json:"id" db:"i_id"`
	SONumber                         string     `json:"so_number" db:"c_so_number"`
	ColorCode                        string     `json:"color_code" db:"c_color_code"`
	Color                            string     `json:"color" db:"n_color"`
	SROCancelled                     *time.Time `json:"sro_cancelled" db:"d_sro_cancelled"`
	MatchingStatus                   string     `json:"matching_status" db:"c_matching_status"`
	MatchingDate                     *time.Time `json:"matching_date" db:"d_matching_date"`
	VIN                              *string    `json:"vin" db:"c_vin"`
	VINReleaseFlag                   bool       `json:"vin_release_flag" db:"b_vin_release_flag"`
	PlanDeliveryDatetime             *time.Time `json:"plan_delivery_datetime" db:"d_plan_delivery_datetime"`
	RRN                              *string    `json:"rrn" db:"c_rrn"`
	UnitStatus                       string     `json:"unit_status" db:"c_unit_status"`
	MDPDate                          *time.Time `json:"mdp_date" db:"d_mdp_date"`
	OnHandDate                       *time.Time `json:"on_hand_date" db:"d_on_hand_date"`
	VehicleCategory                  string     `json:"vehicle_category" db:"c_vehicle_category"`
	FlagOffTheRoadVehicle            bool       `json:"flag_off_the_road_vehicle" db:"b_flag_off_the_road_vehicle"`
	SettlementStatus                 string     `json:"settlement_status" db:"c_settlement_status"`
	SettlementDatetime               *time.Time `json:"settlement_datetime" db:"d_settlement_datetime"`
	PaymentMethod                    string     `json:"payment_method" db:"c_payment_method"`
	OutletTradeID                    *string    `json:"outlet_trade_id" db:"i_outlet_trade_id"`
	OutletTradeName                  *string    `json:"outlet_trade_name" db:"n_outlet_trade_name"`
	OutletTradePONumber              *string    `json:"outlet_trade_po_number" db:"c_outlet_trade_po_number"`
	OutletTradePODatetime            *time.Time `json:"outlet_trade_po_datetime" db:"d_outlet_trade_po_datetime"`
	DeliveryDONumber                 string     `json:"delivery_do_number" db:"c_delivery_do_number"`
	DeliveryReceivedPlanDatetime     time.Time  `json:"delivery_received_plan_datetime" db:"d_delivery_received_plan_datetime"`
	DeliveryReceivedPlanDatetimeEnd  time.Time  `json:"delivery_received_plan_datetime_end" db:"d_delivery_received_plan_datetime_end"`
	DeliveryGateOutDatetime          *time.Time `json:"delivery_gate_out_datetime" db:"d_delivery_gate_out_datetime"`
	DeliveryActualReceivedDatetime   *time.Time `json:"delivery_actual_received_datetime" db:"d_delivery_actual_received_datetime"`
	DeliveryLocation                 string     `json:"delivery_location" db:"c_delivery_location"`
	DeliveryAddressLabel             string     `json:"delivery_address_label" db:"c_delivery_address_label"`
	DeliveryDestinationAddress       string     `json:"delivery_destination_address" db:"e_delivery_destination_address"`
	DeliveryProvince                 string     `json:"delivery_province" db:"n_delivery_province"`
	DeliveryCity                     string     `json:"delivery_city" db:"n_delivery_city"`
	DeliveryDistrict                 string     `json:"delivery_district" db:"n_delivery_district"`
	DeliverySubdistrict              string     `json:"delivery_subdistrict" db:"n_delivery_subdistrict"`
	DeliveryPostalCode               string     `json:"delivery_postal_code" db:"c_delivery_postal_code"`
	DeliveryFlagBuyerIsRecipient     bool       `json:"delivery_flag_buyer_is_recipient" db:"b_delivery_flag_buyer_is_recipient"`
	DeliveryRecipientName            *string    `json:"delivery_recipient_name" db:"n_delivery_recipient_name"`
	DeliveryRecipientPhoneNumber     *string    `json:"delivery_recipient_phone_number" db:"c_delivery_recipient_phone_number"`
	DeliveryRecipientRelation        *string    `json:"delivery_recipient_relation" db:"c_delivery_recipient_relation"`
	DeliveryRecipientRelationOthers  *string    `json:"delivery_recipient_relation_others" db:"c_delivery_recipient_relation_others"`
	DeliveryOriginLocation           string     `json:"delivery_origin_location" db:"c_delivery_origin_location"`
	DeliveryPreparationStatus        string     `json:"delivery_preparation_status" db:"c_delivery_preparation_status"`
	DeliveryReadyForDeliveryDatetime *time.Time `json:"delivery_ready_for_delivery_datetime" db:"d_delivery_ready_for_delivery_datetime"`
	DeliveryPSPSubmittedDatetime     *time.Time `json:"delivery_psp_submitted_datetime" db:"d_delivery_psp_submitted_datetime"`
	DeliveryPreDECSubmittedDatetime  *time.Time `json:"delivery_pre_dec_submitted_datetime" db:"d_delivery_pre_dec_submitted_datetime"`
	DeliveryDECSubmittedDatetime     *time.Time `json:"delivery_dec_submitted_datetime" db:"d_delivery_dec_submitted_datetime"`
	PDFUStatus                       string     `json:"pdfu_status" db:"c_pdfu_status"`
	PDFUSurveyCompletedDatetime      *time.Time `json:"pdfu_survey_completed_datetime" db:"d_pdfu_survey_completed_datetime"`
	DocSTNKHandoverDatetime          *time.Time `json:"doc_stnk_handover_datetime" db:"d_doc_stnk_handover_datetime"`
	DocBPKBHandoverDatetime          *time.Time `json:"doc_bpkb_handover_datetime" db:"d_doc_bpkb_handover_datetime"`
	DocBPKBReceivedBy                *string    `json:"doc_bpkb_received_by" db:"n_doc_bpkb_received_by"`
	DocSTNKDealerReceivedDatetime    *time.Time `json:"doc_stnk_dealer_received_datetime" db:"d_doc_stnk_dealer_received_datetime"`
	DocBPKBDealerReceivedDatetime    *time.Time `json:"doc_bpkb_dealer_received_datetime" db:"d_doc_bpkb_dealer_received_datetime"`
	DocSTNKRecipientName             *string    `json:"doc_stnk_recipient_name" db:"n_doc_stnk_recipient_name"`
	DocBPKBRecipientName             *string    `json:"doc_bpkb_recipient_name" db:"n_doc_bpkb_recipient_name"`
	DocSTNKStatus                    string     `json:"doc_stnk_status" db:"b_doc_stnk_status"`
	DocBPKBStatus                    string     `json:"doc_bpkb_status" db:"b_doc_bpkb_status"`
	DocCollectionStatus              string     `json:"doc_collection_status" db:"c_doc_collection_status"`
	LeasingID                        *string    `json:"leasing_id" db:"i_leasing_id"`
	LeasingCompanyName               *string    `json:"leasing_company_name" db:"n_leasing_company_name"`
	LeasingCreatedDatetime           *time.Time `json:"leasing_created_datetime" db:"d_leasing_created_datetime"`
	LeasingApplicationStatus         *string    `json:"leasing_application_status" db:"c_leasing_application_status"`
	LeasingApprovalDate              *time.Time `json:"leasing_approval_date" db:"d_leasing_approval_date"`
	LeasingTerms                     *int       `json:"leasing_terms" db:"c_leasing_terms"`
	InsuranceProvider                *string    `json:"insurance_provider" db:"c_insurance_provider"`
	InsuranceProviderOther           *string    `json:"insurance_provider_other" db:"c_insurance_provider_other"`
	InsurancePolicyNumber            *string    `json:"insurance_policy_number" db:"c_insurance_policy_number"`
	CreatedAt                        time.Time  `json:"created_at" db:"d_created_at"`
	UpdatedAt                        time.Time  `json:"updated_at" db:"d_updated_at"`
	EventID                          string     `json:"event_id" db:"i_event_id"`
	SPKID                            string     `json:"spk_id" db:"i_spk_id"`
	CustomerID                       string     `json:"customer_id" db:"i_customer_id"`
}

// TableName returns the database table name for the SalesOrder model
func (s *SalesOrder) TableName() string {
	return "dbo.tm_sales_order"
}

// Columns returns the list of database columns for the SalesOrder model
func (s *SalesOrder) Columns() []string {
	return []string{
		"c_so_number", "c_color_code", "n_color",
		"d_sro_cancelled", "c_matching_status", "d_matching_date", "c_vin", "b_vin_release_flag",
		"d_plan_delivery_datetime", "c_rrn", "c_unit_status", "d_mdp_date", "d_on_hand_date",
		"c_vehicle_category", "b_flag_off_the_road_vehicle", "c_settlement_status", "d_settlement_datetime", "c_payment_method",
		"i_outlet_trade_id", "n_outlet_trade_name", "c_outlet_trade_po_number", "d_outlet_trade_po_datetime",
		"c_delivery_do_number", "d_delivery_received_plan_datetime", "d_delivery_received_plan_datetime_end",
		"d_delivery_gate_out_datetime", "d_delivery_actual_received_datetime", "c_delivery_location", "c_delivery_address_label",
		"e_delivery_destination_address", "n_delivery_province", "n_delivery_city", "n_delivery_district", "n_delivery_subdistrict",
		"c_delivery_postal_code", "b_delivery_flag_buyer_is_recipient", "n_delivery_recipient_name", "c_delivery_recipient_phone_number",
		"c_delivery_recipient_relation", "c_delivery_recipient_relation_others", "c_delivery_origin_location", "c_delivery_preparation_status",
		"d_delivery_ready_for_delivery_datetime", "d_delivery_psp_submitted_datetime", "d_delivery_pre_dec_submitted_datetime", "d_delivery_dec_submitted_datetime",
		"c_pdfu_status", "d_pdfu_survey_completed_datetime",
		"d_doc_stnk_handover_datetime", "d_doc_bpkb_handover_datetime", "n_doc_bpkb_received_by",
		"d_doc_stnk_dealer_received_datetime", "d_doc_bpkb_dealer_received_datetime", "n_doc_stnk_recipient_name", "n_doc_bpkb_recipient_name",
		"b_doc_stnk_status", "b_doc_bpkb_status", "c_doc_collection_status",
		"i_leasing_id", "n_leasing_company_name", "d_leasing_created_datetime", "c_leasing_application_status", "d_leasing_approval_date", "c_leasing_terms",
		"c_insurance_provider", "c_insurance_provider_other", "c_insurance_policy_number",
		"d_created_at", "d_updated_at",
		"i_event_id", "i_spk_id", "i_customer_id",
	}
}

// SelectColumns returns the list of columns to select in queries
func (s *SalesOrder) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id", "c_so_number", "c_color_code", "n_color",
		"d_sro_cancelled", "c_matching_status", "d_matching_date", "c_vin", "b_vin_release_flag",
		"d_plan_delivery_datetime", "c_rrn", "c_unit_status", "d_mdp_date", "d_on_hand_date",
		"c_vehicle_category", "b_flag_off_the_road_vehicle", "c_settlement_status", "d_settlement_datetime", "c_payment_method",
		"i_outlet_trade_id", "n_outlet_trade_name", "c_outlet_trade_po_number", "d_outlet_trade_po_datetime",
		"c_delivery_do_number", "d_delivery_received_plan_datetime", "d_delivery_received_plan_datetime_end",
		"d_delivery_gate_out_datetime", "d_delivery_actual_received_datetime", "c_delivery_location", "c_delivery_address_label",
		"e_delivery_destination_address", "n_delivery_province", "n_delivery_city", "n_delivery_district", "n_delivery_subdistrict",
		"c_delivery_postal_code", "b_delivery_flag_buyer_is_recipient", "n_delivery_recipient_name", "c_delivery_recipient_phone_number",
		"c_delivery_recipient_relation", "c_delivery_recipient_relation_others", "c_delivery_origin_location", "c_delivery_preparation_status",
		"d_delivery_ready_for_delivery_datetime", "d_delivery_psp_submitted_datetime", "d_delivery_pre_dec_submitted_datetime", "d_delivery_dec_submitted_datetime",
		"c_pdfu_status", "d_pdfu_survey_completed_datetime",
		"d_doc_stnk_handover_datetime", "d_doc_bpkb_handover_datetime", "n_doc_bpkb_received_by",
		"d_doc_stnk_dealer_received_datetime", "d_doc_bpkb_dealer_received_datetime", "n_doc_stnk_recipient_name", "n_doc_bpkb_recipient_name",
		"b_doc_stnk_status", "b_doc_bpkb_status", "c_doc_collection_status",
		"CAST(i_leasing_id AS NVARCHAR(36)) as i_leasing_id", "n_leasing_company_name", "d_leasing_created_datetime", "c_leasing_application_status", "d_leasing_approval_date", "c_leasing_terms",
		"c_insurance_provider", "c_insurance_provider_other", "c_insurance_policy_number",
		"d_created_at", "d_updated_at",
		"i_event_id", "i_spk_id", "i_customer_id",
	}
}

func (s *SalesOrder) ToCreateMap() (columns []string, values []interface{}) {
	columns = s.Columns()
	values = []interface{}{
		s.SONumber, s.ColorCode, s.Color,
		s.SROCancelled, s.MatchingStatus, s.MatchingDate, s.VIN, s.VINReleaseFlag,
		s.PlanDeliveryDatetime, s.RRN, s.UnitStatus, s.MDPDate, s.OnHandDate,
		s.VehicleCategory, s.FlagOffTheRoadVehicle, s.SettlementStatus, s.SettlementDatetime, s.PaymentMethod,
		s.OutletTradeID, s.OutletTradeName, s.OutletTradePONumber, s.OutletTradePODatetime,
		s.DeliveryDONumber, s.DeliveryReceivedPlanDatetime, s.DeliveryReceivedPlanDatetimeEnd,
		s.DeliveryGateOutDatetime, s.DeliveryActualReceivedDatetime, s.DeliveryLocation, s.DeliveryAddressLabel,
		s.DeliveryDestinationAddress, s.DeliveryProvince, s.DeliveryCity, s.DeliveryDistrict, s.DeliverySubdistrict,
		s.DeliveryPostalCode, s.DeliveryFlagBuyerIsRecipient, s.DeliveryRecipientName, s.DeliveryRecipientPhoneNumber,
		s.DeliveryRecipientRelation, s.DeliveryRecipientRelationOthers, s.DeliveryOriginLocation, s.DeliveryPreparationStatus,
		s.DeliveryReadyForDeliveryDatetime, s.DeliveryPSPSubmittedDatetime, s.DeliveryPreDECSubmittedDatetime, s.DeliveryDECSubmittedDatetime,
		s.PDFUStatus, s.PDFUSurveyCompletedDatetime,
		s.DocSTNKHandoverDatetime, s.DocBPKBHandoverDatetime, s.DocBPKBReceivedBy,
		s.DocSTNKDealerReceivedDatetime, s.DocBPKBDealerReceivedDatetime, s.DocSTNKRecipientName, s.DocBPKBRecipientName,
		s.DocSTNKStatus, s.DocBPKBStatus, s.DocCollectionStatus,
		s.LeasingID, s.LeasingCompanyName, s.LeasingCreatedDatetime, s.LeasingApplicationStatus, s.LeasingApprovalDate, s.LeasingTerms,
		s.InsuranceProvider, s.InsuranceProviderOther, s.InsurancePolicyNumber,
		s.CreatedAt, s.UpdatedAt,
		s.EventID, s.SPKID, s.CustomerID,
	}
	return
}
