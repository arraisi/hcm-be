package order

import (
	"strconv"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
)

// OutletTradeRequest represents the outlet trade information
type OutletTradeRequest struct {
	OutletID   string `json:"outlet_ID" validate:"required"`
	OutletName string `json:"outlet_name" validate:"required"`
	PONumber   string `json:"po_number" validate:"required"`
	PODatetime int64  `json:"po_datetime" validate:"required"`
}

// PackagePartRequest represents individual parts within an accessory package
type PackagePartRequest struct {
	AccessoriesNumber string `json:"accessories_number" validate:"required"`
	AccessoriesName   string `json:"accessories_name" validate:"required"`
}

// AccessoryRequest represents accessories information
type AccessoryRequest struct {
	AccessoriesType                 string               `json:"accessories_type" validate:"required,oneof=PACKAGE MERCHANDISE ACCESSORIES"`
	AccessoriesNumber               string               `json:"accessories_number" validate:"required"`
	PackageID                       *string              `json:"package_ID"`
	AccessoriesName                 string               `json:"accessories_name" validate:"required"`
	AccessoriesQty                  int                  `json:"accessories_qty" validate:"required,min=1"`
	PackageParts                    []PackagePartRequest `json:"package_parts"`
	AccessoriesOrderSource          string               `json:"accessories_order_source" validate:"required,oneof=TLS OTHER"`
	AccessoriesAvailabilityStatus   string               `json:"accessories_availability_status" validate:"required,oneof=AVAILABLE NOT_AVAILABLE"`
	AccessoriesItemStatus           string               `json:"accessories_item_status" validate:"required,oneof=PROCESSED PENDING CANCELLED"`
	AccessoriesSize                 *string              `json:"accessories_size"`
	AccessoriesColor                *string              `json:"accessories_color"`
	AccessoriesEstPrice             int64                `json:"accessories_est_price" validate:"required,min=0"`
	AccessoriesInstallationEstPrice int64                `json:"accessories_installation_est_price" validate:"min=0"`
	FlagAccessoriesNeedDownPayment  bool                 `json:"flag_accessories_need_down_payment"`
}

// PaymentRequest represents payment information
type PaymentRequest struct {
	PaymentID       string `json:"payment_ID" validate:"required"`
	PaymentStatus   string `json:"payment_status" validate:"required,oneof=COMPLETED PENDING FAILED CANCELLED"`
	PaymentNumber   string `json:"payment_number" validate:"required"`
	PaymentDatetime int64  `json:"payment_datetime" validate:"required"`
	FundSource      string `json:"fund_source" validate:"required,oneof=CUSTOMER LEASING DEALER"`
	PaymentChannel  string `json:"payment_channel" validate:"required,oneof=BANK_TRANSFER CREDIT_CARD CASH E_WALLET"`
	PaymentStage    int    `json:"payment_stage" validate:"required,min=1"`
	NameOnPayment   string `json:"name_on_payment" validate:"required"`
}

// LeasingApplicationRequest represents leasing application information
type LeasingApplicationRequest struct {
	LeasingID                string `json:"leasing_ID" validate:"required"`
	LeasingCompanyName       string `json:"leasing_company_name" validate:"required"`
	CreatedDatetime          int64  `json:"created_datetime" validate:"required"`
	LeasingApplicationStatus string `json:"leasing_application_status" validate:"required,oneof=PENDING APPROVED REJECTED"`
	LeasingApprovalDate      int64  `json:"leasing_approval_date"`
	LeasingTerms             int    `json:"leasing_terms" validate:"required,min=1"`
}

// InsurancePolicyRequest represents individual insurance policy
type InsurancePolicyRequest struct {
	InsuranceType      string   `json:"insurance_type" validate:"required,oneof=ALL-RISK COMPREHENSIVE TLO"`
	InsuranceCoverage  []string `json:"insurance_coverage"`
	InsuranceStartDate int64    `json:"insurance_start_date" validate:"required"`
	InsuranceEndDate   int64    `json:"insurance_end_date" validate:"required"`
}

// InsuranceApplicationRequest represents insurance application information
type InsuranceApplicationRequest struct {
	InsuranceProvider      string                   `json:"insurance_provider" validate:"required"`
	InsuranceProviderOther *string                  `json:"insurance_provider_other"`
	InsurancePolicyNumber  string                   `json:"insurance_policy_number" validate:"required"`
	Policies               []InsurancePolicyRequest `json:"policies" validate:"required,min=1,dive"`
}

func (i *InsuranceApplicationRequest) ToInsurancePoliciesModel(salesOrderID string) []domain.SalesOrderInsurancePolicy {
	policies := make([]domain.SalesOrderInsurancePolicy, 0, len(i.Policies))
	for _, policyReq := range i.Policies {
		policy := policyReq.ToInsurancePolicyModel(salesOrderID)
		policies = append(policies, policy)

	}
	return policies
}

// PDFURequest represents PDFU (Post Delivery Follow Up) information
type PDFURequest struct {
	Status                  string `json:"status" validate:"required,oneof=PENDING COMPLETED FAILED"`
	SurveyCompletedDatetime int64  `json:"survey_completed_datetime"`
}

// DocumentHandoverRequest represents document handover information
type DocumentHandoverRequest struct {
	STNKHandoverDatetime       int64   `json:"stnk_handover_datetime"`
	BPKBHandoverDatetime       int64   `json:"bpkb_handover_datetime"`
	BPKBReceivedBy             *string `json:"bpkb_received_by" validate:"omitempty,oneof=CUSTOMER LEASING"`
	STNKDealerReceivedDatetime int64   `json:"stnk_dealer_received_datetime"`
	BPKBDealerReceivedDatetime int64   `json:"bpkb_dealer_received_datetime"`
	STNKRecipientName          *string `json:"stnk_recipient_name"`
	BPKBRecipientName          *string `json:"bpkb_recipient_name"`
	STNKStatus                 bool    `json:"stnk_status"`
	BPKBStatus                 bool    `json:"bpkb_status"`
	DocumentCollectionStatus   string  `json:"document_collection_status" validate:"required,oneof=WAITING IN_PROGRESS COMPLETED"`
}

// DeliveryPlanRequest represents delivery amendment/plan information
type DeliveryPlanRequest struct {
	AmendmentCreatedDatetime  int64   `json:"amendment_created_datetime" validate:"required"`
	AmendmentStatus           string  `json:"amendment_status" validate:"required,oneof=PENDING CONFIRMED REJECTED CANCELLED"`
	AmendmentReason           string  `json:"amendment_reason" validate:"required,oneof=TCO CUSTOMER_REQUEST OTHERS"`
	AmendmentReasonOthers     *string `json:"amendment_reason_others"`
	AmendmentSource           string  `json:"amendment_source" validate:"required,oneof=CUSTOMER DEALER OTHERS"`
	FlagBuyerIsRecipient      bool    `json:"flag_buyer_is_recipient"`
	ReceivedPlanDatetimeStart int64   `json:"received_plan_datetime_start" validate:"required"`
	ReceivedPlanDatetimeEnd   int64   `json:"received_plan_datetime_end" validate:"required"`
	RecipientName             *string `json:"recipient_name"`
	RecipientPhoneNumber      *string `json:"recipient_phone_number"`
	RecipientRelation         *string `json:"recipient_relation" validate:"omitempty,oneof=SELF FAMILY_MEMBER FRIEND COLLEAGUE OTHERS"`
	RecipientRelationOthers   *string `json:"recipient_relation_others"`
	DeliveryLocation          *string `json:"delivery_location" validate:"omitempty,oneof=DEALER CUSTOMER_ADDRESS OTHERS"`
	DeliveryAddressLabel      string  `json:"delivery_address_label" validate:"required"`
	DeliveryAddress           string  `json:"delivery_address" validate:"required"`
	DeliveryProvince          string  `json:"delivery_province" validate:"required"`
	DeliveryCity              string  `json:"delivery_city" validate:"required"`
	DeliveryDistrict          string  `json:"delivery_district" validate:"required"`
	DeliverySubdistrict       string  `json:"delivery_subdistrict" validate:"required"`
	DeliveryPostalCode        string  `json:"delivery_postal_code" validate:"required"`
}

// DeliveryRequest represents delivery information
type DeliveryRequest struct {
	DONumber                  string  `json:"do_number" validate:"required"`
	ReceivedPlanDatetime      int64   `json:"received_plan_datetime" validate:"required"`
	ReceivedPlanDatetimeEnd   int64   `json:"received_plan_datetime_end" validate:"required"`
	GateOutDatetime           int64   `json:"gate_out_datetime"`
	ActualReceivedDatetime    int64   `json:"actual_received_datetime"`
	DeliveryLocation          string  `json:"delivery_location" validate:"required,oneof=DEALER CUSTOMER_ADDRESS OTHERS"`
	DeliveryAddressLabel      string  `json:"delivery_address_label" validate:"required"`
	DestinationAddress        string  `json:"destination_address" validate:"required"`
	DeliveryProvince          string  `json:"delivery_province" validate:"required"`
	DeliveryCity              string  `json:"delivery_city" validate:"required"`
	DeliveryDistrict          string  `json:"delivery_district" validate:"required"`
	DeliverySubdistrict       string  `json:"delivery_subdistrict" validate:"required"`
	DeliveryPostalCode        string  `json:"delivery_postal_code" validate:"required"`
	FlagBuyerIsRecipient      bool    `json:"flag_buyer_is_recipient"`
	RecipientName             *string `json:"recipient_name"`
	RecipientPhoneNumber      *string `json:"recipient_phone_number"`
	RecipientRelation         *string `json:"recipient_relation" validate:"omitempty,oneof=SELF FAMILY_MEMBER FRIEND COLLEAGUE OTHERS"`
	RecipientRelationOthers   *string `json:"recipient_relation_others"`
	OriginLocation            string  `json:"origin_location" validate:"required"`
	DeliveryPreparationStatus string  `json:"delivery_preparation_status" validate:"required,oneof=PENDING IN_PREPARATION READY DELIVERY_COMPLETED"`
	ReadyForDeliveryDatetime  int64   `json:"ready_for_delivery_datetime"`
	PSPSubmittedDatetime      int64   `json:"psp_submitted_datetime"`
	PreDECSubmittedDatetime   int64   `json:"pre_dec_submitted_datetime"`
	DECSubmittedDatetime      int64   `json:"dec_submitted_datetime"`
}

// SalesOrderRequest represents individual sales order information
type SalesOrderRequest struct {
	SONumber              string                       `json:"so_number" validate:"required"`
	ColorCode             string                       `json:"color_code" validate:"required"`
	Color                 string                       `json:"color" validate:"required"`
	SROCancelledDatetime  int64                        `json:"sro_cancelled_datetime"`
	MatchingStatus        string                       `json:"matching_status" validate:"required,oneof=WAITING VIN_MATCHED UNMATCHED"`
	MatchingDate          int64                        `json:"matching_date"`
	VIN                   string                       `json:"vin"`
	VINReleaseFlag        bool                         `json:"vin_release_flag"`
	PlanDeliveryDatetime  int64                        `json:"plan_delivery_datetime"`
	RRN                   string                       `json:"rrn"`
	UnitStatus            string                       `json:"unit_status" validate:"required"`
	MDPDate               int64                        `json:"mdp_date"`
	OnHandDate            int64                        `json:"on_hand_date"`
	VehicleCategory       string                       `json:"vehicle_category" validate:"required,oneof=RETAIL FLEET"`
	FlagOffTheRoadVehicle bool                         `json:"flag_off_the_road_vehicle"`
	OutletTrade           *OutletTradeRequest          `json:"outlet_trade"`
	Accessories           []AccessoryRequest           `json:"accessories"`
	SettlementStatus      string                       `json:"settlement_status" validate:"required,oneof=WAITING PARTIAL COMPLETED"`
	SettlementDatetime    int64                        `json:"settlement_datetime"`
	PaymentMethod         string                       `json:"payment_method" validate:"required,oneof=CASH CREDIT"`
	DownPayment           []PaymentRequest             `json:"down_payment"`
	Payment               []PaymentRequest             `json:"payment"`
	LeasingApplication    *LeasingApplicationRequest   `json:"leasing_application"`
	InsuranceApplication  *InsuranceApplicationRequest `json:"insurance_application"`
	Delivery              DeliveryRequest              `json:"delivery" validate:"required"`
	PDFU                  PDFURequest                  `json:"pdfu" validate:"required"`
	DocumentHandover      DocumentHandoverRequest      `json:"document_handover" validate:"required"`
	DeliveryPlans         []DeliveryPlanRequest        `json:"delivery_plans"`
}

// SPKRequest represents SPK (Surat Pesanan Kendaraan) information
type SPKRequest struct {
	SPKNumber               string  `json:"spk_number" validate:"required"`
	LeadsID                 string  `json:"leads_ID" validate:"required"`
	CreatedDatetime         int64   `json:"created_datetime" validate:"required"`
	SPKStatus               string  `json:"spk_status" validate:"required,oneof=PENDING APPROVED REJECTED CANCELLED"`
	Model                   string  `json:"model" validate:"required"`
	Variant                 string  `json:"variant" validate:"required"`
	KatashikiSuffix         string  `json:"katashiki_suffix" validate:"required"`
	Year                    int     `json:"year" validate:"required,min=2000"`
	OutletID                string  `json:"outlet_ID" validate:"required"`
	OutletName              string  `json:"outlet_name" validate:"required"`
	EmployeeID              string  `json:"employee_ID" validate:"required"`
	EmployeeFirstName       string  `json:"employee_first_name" validate:"required"`
	EmployeeLastName        string  `json:"employee_last_name" validate:"required"`
	SPKCustomerConfirmation bool    `json:"spk_customer_confirmation"`
	SPKApprovedDatetime     int64   `json:"spk_approved_datetime"`
	SPKCancelledDatetime    int64   `json:"spk_cancelled_datetime"`
	SPKCancelledReason      *string `json:"spk_cancelled_reason"`
}

// TrackOrderStatusEventData represents the data payload of the track order status event
type TrackOrderStatusEventData struct {
	OneAccount  customer.OneAccountRequest `json:"one_account" validate:"required"`
	SPK         SPKRequest                 `json:"spk" validate:"required"`
	SalesOrders []SalesOrderRequest        `json:"sales_order" validate:"required,min=1,dive"`
}

// TrackOrderStatusEvent represents the complete webhook payload for track order status
type TrackOrderStatusEvent struct {
	Process   string                    `json:"process" validate:"required"`
	EventID   string                    `json:"event_ID" validate:"required,uuid4"`
	Timestamp int64                     `json:"timestamp" validate:"required"`
	Data      TrackOrderStatusEventData `json:"data" validate:"required"`
}

// ToSPKModel converts SPKRequest to domain.SPK
func (r *SPKRequest) ToSPKModel() domain.SPK {
	now := time.Now()
	return domain.SPK{
		SPKNumber:               r.SPKNumber,
		LeadsID:                 r.LeadsID,
		CreatedDatetime:         time.Unix(r.CreatedDatetime, 0),
		SPKStatus:               r.SPKStatus,
		Model:                   r.Model,
		Variant:                 r.Variant,
		KatashikiSuffix:         r.KatashikiSuffix,
		Year:                    r.Year,
		OutletID:                r.OutletID,
		OutletName:              r.OutletName,
		EmployeeID:              r.EmployeeID,
		EmployeeFirstName:       r.EmployeeFirstName,
		EmployeeLastName:        r.EmployeeLastName,
		SPKCustomerConfirmation: r.SPKCustomerConfirmation,
		SPKApprovedDatetime:     utils.UnixToTimePtr(r.SPKApprovedDatetime),
		SPKCancelledDatetime:    utils.UnixToTimePtr(r.SPKCancelledDatetime),
		SPKCancelledReason:      r.SPKCancelledReason,
		CreatedAt:               now,
		UpdatedAt:               now,
	}
}

// ToSalesOrderModel converts SalesOrderRequest to domain.SalesOrder
func (r *SalesOrderRequest) ToSalesOrderModel(spkID, customerID, eventID string) domain.SalesOrder {
	now := time.Now()
	return domain.SalesOrder{
		SONumber:                         r.SONumber,
		ColorCode:                        r.ColorCode,
		Color:                            r.Color,
		SROCancelled:                     utils.UnixToTimePtr(r.SROCancelledDatetime),
		MatchingStatus:                   r.MatchingStatus,
		MatchingDate:                     utils.UnixToTimePtr(r.MatchingDate),
		VIN:                              utils.ToPointer(r.VIN),
		VINReleaseFlag:                   r.VINReleaseFlag,
		PlanDeliveryDatetime:             utils.UnixToTimePtr(r.PlanDeliveryDatetime),
		RRN:                              utils.ToPointer(r.RRN),
		UnitStatus:                       r.UnitStatus,
		MDPDate:                          utils.UnixToTimePtr(r.MDPDate),
		OnHandDate:                       utils.UnixToTimePtr(r.OnHandDate),
		VehicleCategory:                  r.VehicleCategory,
		FlagOffTheRoadVehicle:            r.FlagOffTheRoadVehicle,
		SettlementStatus:                 r.SettlementStatus,
		SettlementDatetime:               utils.UnixToTimePtr(r.SettlementDatetime),
		PaymentMethod:                    r.PaymentMethod,
		OutletTradeID:                    utils.ToPointerIf(r.OutletTrade != nil, r.OutletTrade.OutletID),
		OutletTradeName:                  utils.ToPointerIf(r.OutletTrade != nil, r.OutletTrade.OutletName),
		OutletTradePONumber:              utils.ToPointerIf(r.OutletTrade != nil, r.OutletTrade.PONumber),
		OutletTradePODatetime:            utils.ToPointerIf(r.OutletTrade != nil, time.Unix(r.OutletTrade.PODatetime, 0)),
		DeliveryDONumber:                 r.Delivery.DONumber,
		DeliveryReceivedPlanDatetime:     time.Unix(r.Delivery.ReceivedPlanDatetime, 0),
		DeliveryReceivedPlanDatetimeEnd:  time.Unix(r.Delivery.ReceivedPlanDatetimeEnd, 0),
		DeliveryGateOutDatetime:          utils.UnixToTimePtr(r.Delivery.GateOutDatetime),
		DeliveryActualReceivedDatetime:   utils.UnixToTimePtr(r.Delivery.ActualReceivedDatetime),
		DeliveryLocation:                 r.Delivery.DeliveryLocation,
		DeliveryAddressLabel:             r.Delivery.DeliveryAddressLabel,
		DeliveryDestinationAddress:       r.Delivery.DestinationAddress,
		DeliveryProvince:                 r.Delivery.DeliveryProvince,
		DeliveryCity:                     r.Delivery.DeliveryCity,
		DeliveryDistrict:                 r.Delivery.DeliveryDistrict,
		DeliverySubdistrict:              r.Delivery.DeliverySubdistrict,
		DeliveryPostalCode:               r.Delivery.DeliveryPostalCode,
		DeliveryFlagBuyerIsRecipient:     r.Delivery.FlagBuyerIsRecipient,
		DeliveryRecipientName:            r.Delivery.RecipientName,
		DeliveryRecipientPhoneNumber:     r.Delivery.RecipientPhoneNumber,
		DeliveryRecipientRelation:        r.Delivery.RecipientRelation,
		DeliveryRecipientRelationOthers:  r.Delivery.RecipientRelationOthers,
		DeliveryOriginLocation:           r.Delivery.OriginLocation,
		DeliveryPreparationStatus:        r.Delivery.DeliveryPreparationStatus,
		DeliveryReadyForDeliveryDatetime: utils.UnixToTimePtr(r.Delivery.ReadyForDeliveryDatetime),
		DeliveryPSPSubmittedDatetime:     utils.UnixToTimePtr(r.Delivery.PSPSubmittedDatetime),
		DeliveryPreDECSubmittedDatetime:  utils.UnixToTimePtr(r.Delivery.PreDECSubmittedDatetime),
		DeliveryDECSubmittedDatetime:     utils.UnixToTimePtr(r.Delivery.DECSubmittedDatetime),
		PDFUStatus:                       r.PDFU.Status,
		PDFUSurveyCompletedDatetime:      utils.UnixToTimePtr(r.PDFU.SurveyCompletedDatetime),
		DocSTNKHandoverDatetime:          utils.UnixToTimePtr(r.DocumentHandover.STNKHandoverDatetime),
		DocBPKBHandoverDatetime:          utils.UnixToTimePtr(r.DocumentHandover.BPKBHandoverDatetime),
		DocBPKBReceivedBy:                r.DocumentHandover.BPKBReceivedBy,
		DocSTNKDealerReceivedDatetime:    utils.UnixToTimePtr(r.DocumentHandover.STNKDealerReceivedDatetime),
		DocBPKBDealerReceivedDatetime:    utils.UnixToTimePtr(r.DocumentHandover.BPKBDealerReceivedDatetime),
		DocSTNKRecipientName:             r.DocumentHandover.STNKRecipientName,
		DocBPKBRecipientName:             r.DocumentHandover.BPKBRecipientName,
		DocSTNKStatus:                    r.DocumentHandover.STNKStatus,
		DocBPKBStatus:                    r.DocumentHandover.BPKBStatus,
		DocCollectionStatus:              r.DocumentHandover.DocumentCollectionStatus,
		LeasingID:                        utils.ToPointerIf(r.LeasingApplication != nil, r.LeasingApplication.LeasingID),
		LeasingCompanyName:               utils.ToPointerIf(r.LeasingApplication != nil, r.LeasingApplication.LeasingCompanyName),
		LeasingCreatedDatetime:           utils.ToPointerIf(r.LeasingApplication != nil, time.Unix(r.LeasingApplication.CreatedDatetime, 0)),
		LeasingApplicationStatus:         utils.ToPointerIf(r.LeasingApplication != nil, r.LeasingApplication.LeasingApplicationStatus),
		LeasingApprovalDate:              utils.ToPointerIf(r.LeasingApplication != nil && r.LeasingApplication.LeasingApprovalDate > 0, time.Unix(r.LeasingApplication.LeasingApprovalDate, 0)),
		LeasingTerms:                     utils.ToPointerIf(r.LeasingApplication != nil, r.LeasingApplication.LeasingTerms),
		InsuranceProvider:                utils.ToPointerIf(r.InsuranceApplication != nil, r.InsuranceApplication.InsuranceProvider),
		InsuranceProviderOther:           r.InsuranceApplication.InsuranceProviderOther,
		InsurancePolicyNumber:            utils.ToPointerIf(r.InsuranceApplication != nil, r.InsuranceApplication.InsurancePolicyNumber),
		CreatedAt:                        now,
		UpdatedAt:                        now,
		EventID:                          eventID,
		SPKID:                            spkID,
		CustomerID:                       customerID,
	}
}

// ToAccessoryModel converts AccessoryRequest to domain.SalesOrderAccessory
func (r *AccessoryRequest) ToAccessoryModel(salesOrderID string) domain.SalesOrderAccessory {
	now := time.Now()
	return domain.SalesOrderAccessory{
		AccessoriesType:                 r.AccessoriesType,
		PackageID:                       utils.ToValue(r.PackageID),
		AccessoriesNumber:               r.AccessoriesNumber,
		AccessoriesName:                 r.AccessoriesName,
		AccessoriesQty:                  strconv.Itoa(r.AccessoriesQty),
		AccessoriesOrderSource:          r.AccessoriesOrderSource,
		AccessoriesAvailabilityStatus:   r.AccessoriesAvailabilityStatus,
		AccessoriesItemStatus:           r.AccessoriesItemStatus,
		AccessoriesSize:                 utils.ToValue(r.AccessoriesSize),
		AccessoriesColor:                utils.ToValue(r.AccessoriesColor),
		AccessoriesEstPrice:             strconv.FormatInt(r.AccessoriesEstPrice, 10),
		FlagAccessoriesNeedDownPayment:  r.FlagAccessoriesNeedDownPayment,
		AccessoriesInstallationEstPrice: strconv.FormatInt(r.AccessoriesInstallationEstPrice, 10),
		CreatedAt:                       now,
		UpdatedAt:                       now,
		SalesOrderID:                    salesOrderID,
	}
}

func (r *PackagePartRequest) ToAccessoryPartModel(accessoriesID string) domain.SalesOrderAccessoriesPart {
	now := time.Now()
	return domain.SalesOrderAccessoriesPart{
		AccessoriesID:     accessoriesID,
		AccessoriesNumber: r.AccessoriesNumber,
		AccessoriesName:   r.AccessoriesName,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

// ToPaymentModel converts PaymentRequest to domain.SalesOrderPayment
func (r *PaymentRequest) ToPaymentModel(salesOrderID string, isDownPayment bool) domain.SalesOrderPayment {
	// Determine payment type based on isDownPayment flag
	paymentType := constants.PaymentTypePayment
	if isDownPayment {
		paymentType = constants.PaymentTypeDownPayment
	}

	now := time.Now()

	return domain.SalesOrderPayment{
		SalesOrderID:    salesOrderID,
		PaymentID:       r.PaymentID,
		PaymentNumber:   r.PaymentNumber,
		PaymentDatetime: time.Unix(r.PaymentDatetime, 0),
		NameOnPayment:   r.NameOnPayment,
		PaymentStatus:   r.PaymentStatus,
		FundSource:      r.FundSource,
		PaymentChannel:  r.PaymentChannel,
		PaymentStage:    strconv.Itoa(r.PaymentStage),
		PaymentType:     paymentType,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// ToInsurancePolicyModel converts InsurancePolicyRequest to domain.SalesOrderInsurancePolicy
func (r *InsurancePolicyRequest) ToInsurancePolicyModel(salesOrderID string) domain.SalesOrderInsurancePolicy {
	now := time.Now()
	return domain.SalesOrderInsurancePolicy{
		SalesOrderID:       salesOrderID,
		InsuranceType:      r.InsuranceType,
		InsuranceCoverage:  r.InsuranceCoverage,
		InsuranceStartDate: time.Unix(r.InsuranceStartDate, 0),
		InsuranceEndDate:   time.Unix(r.InsuranceEndDate, 0),
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

// ToDeliveryPlanModel converts DeliveryPlanRequest to domain.SalesOrderDeliveryPlan
func (r *DeliveryPlanRequest) ToDeliveryPlanModel(salesOrderID string) domain.SalesOrderDeliveryPlan {
	now := time.Now()
	return domain.SalesOrderDeliveryPlan{
		SalesOrderID:              salesOrderID,
		AmendmentCreatedDatetime:  time.Unix(r.AmendmentCreatedDatetime, 0),
		AmendmentStatus:           r.AmendmentStatus,
		AmendmentReason:           r.AmendmentReason,
		AmendmentReasonOthers:     r.AmendmentReasonOthers,
		AmendmentSource:           r.AmendmentSource,
		FlagBuyerIsRecipient:      r.FlagBuyerIsRecipient,
		ReceivedPlanDatetimeStart: time.Unix(r.ReceivedPlanDatetimeStart, 0),
		ReceivedPlanDatetimeEnd:   time.Unix(r.ReceivedPlanDatetimeEnd, 0),
		RecipientName:             r.RecipientName,
		RecipientPhoneNumber:      r.RecipientPhoneNumber,
		RecipientRelation:         r.RecipientRelation,
		RecipientRelationOthers:   r.RecipientRelationOthers,
		DeliveryAddressLocation:   r.DeliveryLocation,
		DeliveryAddressLabel:      r.DeliveryAddressLabel,
		DeliveryAddress:           r.DeliveryAddress,
		DeliveryProvince:          r.DeliveryProvince,
		DeliveryCity:              r.DeliveryCity,
		DeliveryDistrict:          r.DeliveryDistrict,
		DeliverySubdistrict:       r.DeliverySubdistrict,
		DeliveryPostalCode:        r.DeliveryPostalCode,
		CreatedAt:                 now,
		UpdatedAt:                 now,
	}
}
