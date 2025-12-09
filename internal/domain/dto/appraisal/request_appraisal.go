package appraisal

import (
	"strings"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/shopspring/decimal"
)

// =======================
// Top-level event
// =======================

type EventRequest struct {
	Process   string    `json:"process"`   // "appraisal_booking_request"
	EventID   string    `json:"event_ID"`  // UUID v4
	Timestamp int64     `json:"timestamp"` // UNIX timestamp (seconds)
	Data      EventData `json:"data"`
}

type EventData struct {
	OneAccount       OneAccountDTO       `json:"one_account"`
	RequestAppraisal RequestAppraisalDTO `json:"request_appraisal"`
	Leads            LeadsDTO            `json:"leads"`
	UsedCar          UsedCarDTO          `json:"used_car"`
}

// =======================
// one_account
// =======================

type OneAccountDTO struct {
	OneAccountID            string   `json:"one_account_ID"`            // VARCHAR(32)
	FirstName               string   `json:"first_name"`                // VARCHAR(64)
	LastName                string   `json:"last_name"`                 // VARCHAR(64)
	PhoneNumber             string   `json:"phone_number"`              // VARCHAR(16)
	Email                   string   `json:"email"`                     // VARCHAR(64)
	KTPNumber               string   `json:"ktp_number"`                // VARCHAR(16), conditional
	PreferredContactChannel []string `json:"preferred_contact_channel"` // MTOYOTA, WHATSAPP_OR_SMS, EMAIL, PHONE_CALL
}

func (dto OneAccountDTO) ToCustomerModel() domain.Customer {

	// PreferredContactChannel di DTO adalah []string,
	// sedangkan di domain hanya 1 string → kita gabungkan saja dengan comma.
	contactChannel := strings.Join(dto.PreferredContactChannel, ",")

	// NOTE:
	// Banyak field tidak tersedia di DTO → isi default value.
	// Nanti kalau butuh enrichment dari tempat lain tinggal update.
	empty := ""
	now := time.Now()

	return domain.Customer{
		ID:                      "", // created by DB on upsert
		OneAccountID:            dto.OneAccountID,
		HasjratID:               "", // unknown
		FirstName:               dto.FirstName,
		LastName:                dto.LastName,
		Gender:                  nil, // tidak ada di DTO
		PhoneNumber:             dto.PhoneNumber,
		Email:                   dto.Email,
		IsNew:                   true,  // default: new customer from mTOYOTA
		IsMerge:                 false, // default
		PrimaryUser:             nil,
		DealerCustomerID:        "",
		IsValid:                 true, // dianggap valid karena datang dari mTOYOTA
		IsOmnichannel:           true, // mTOYOTA customer considered omnichannel
		LeadsInID:               "",
		CustomerCategory:        "GENERAL", // default guess
		KTPNumber:               dto.KTPNumber,
		BirthDate:               time.Time{}, // unknown
		ResidenceAddress:        "",
		ResidenceSubdistrict:    "",
		ResidenceDistrict:       "",
		ResidenceCity:           "",
		ResidenceProvince:       "",
		ResidencePostalCode:     "",
		CustomerType:            "INDIVIDUAL", // default; update if needed
		LeadsID:                 "",
		Occupation:              "",
		RegistrationChannel:     "MTOYOTA", // dari mTOYOTA
		RegistrationDatetime:    now,
		ConsentGiven:            false, // unknown
		ConsentGivenAt:          time.Time{},
		ConsentGivenDuring:      "",
		AddressLabel:            "",
		DetailAddress:           "",
		ToyotaIDSingleStatus:    "",
		PreferredContactChannel: contactChannel, // comma-separated string
		CreatedAt:               now,
		CreatedBy:               "system",
		UpdatedAt:               now,
		UpdatedBy:               &empty,
	}
}

// =======================
// request_appraisal
// =======================

type RequestAppraisalDTO struct {
	AppraisalBookingID     string `json:"appraisal_booking_ID"`     // UUID / VARCHAR(36)
	AppraisalBookingNumber string `json:"appraisal_booking_number"` // VARCHAR(32)
	CreatedDatetime        int64  `json:"created_datetime"`         // UNIX TIMESTAMP

	OutletID   string `json:"outlet_ID"`   // VARCHAR(32)
	OutletName string `json:"outlet_name"` // VARCHAR(128)

	AppraisalLocation string `json:"appraisal_location"` // DEALER, HOME_OR_OTHER_ADDRESS
	HomeAddress       string `json:"home_address"`       // conditional
	Province          string `json:"province"`
	City              string `json:"city"`
	District          string `json:"district"`
	Subdistrict       string `json:"subdistrict"`
	PostalCode        string `json:"postal_code"` // VARCHAR(5), conditional

	AppraisalStartDatetime int64 `json:"appraisal_start_datetime"` // UNIX TIMESTAMP, step 15 minutes
	AppraisalEndDatetime   int64 `json:"appraisal_end_datetime"`   // UNIX TIMESTAMP, step 15 minutes

	AppraisalBookingStatus  string `json:"appraisal_booking_status"`  // SUBMITTED, CHANGE_REQUEST, CANCEL_SUBMITTED
	CancelledBy             string `json:"cancelled_by"`              // CUSTOMER, DEALER (conditional)
	CancellationReason      string `json:"cancellation_reason"`       // enum list
	OtherCancellationReason string `json:"other_cancellation_reason"` // only if reason = OTHERS

	BookingServiceFlag bool `json:"booking_service_flag"` // true / false
}

// ToAppraisalModel ToModel converts RequestAppraisalDTO into domain.Appraisal.
func (dto RequestAppraisalDTO) ToAppraisalModel(OneAccountID, vin, leadsID string) domain.Appraisal {
	entity := domain.Appraisal{
		AppraisalBookingID:     dto.AppraisalBookingID,
		AppraisalBookingNumber: dto.AppraisalBookingNumber,

		OutletID:   dto.OutletID,
		OutletName: dto.OutletName,

		OneAccountID: utils.ToPointer(OneAccountID),
		VIN:          utils.ToPointer(vin),
		LeadsID:      utils.ToPointer(leadsID),

		AppraisalLocation: utils.ToPointer(dto.AppraisalLocation),
		HomeAddress:       utils.ToPointer(dto.HomeAddress),
		Province:          utils.ToPointer(dto.Province),
		City:              utils.ToPointer(dto.City),
		District:          utils.ToPointer(dto.District),
		Subdistrict:       utils.ToPointer(dto.Subdistrict),
		PostalCode:        utils.ToPointer(dto.PostalCode),

		CreatedDatetime: utils.GetTimeUnix(dto.CreatedDatetime),

		ConfirmStartDatetime: nil, // belum ada di request
		ConfirmEndDatetime:   nil,

		BookingStatus: dto.AppraisalBookingStatus,

		OtherCancellationReason: utils.ToPointer(dto.OtherCancellationReason),
		BookingServiceFlag:      dto.BookingServiceFlag,

		// vehicle (not present in request, filled on update)
		KatashikiSuffix: nil,
		ColorCode:       nil,
		Model:           nil,
		Variant:         nil,
		Color:           nil,

		// trade-in summary (not present on request)
		FinalTradeInStatus:        nil,
		LastTradeInStatusDatetime: nil,

		// sales docs (not present on request)
		SPKNumber: nil,
		SONumber:  nil,

		// negotiation (not present on request)
		CustomerNegotiationPrice:           nil,
		DealerNegotiationPrice:             nil,
		DealPrice:                          nil,
		DownPaymentEstimation:              nil,
		EstimatedRemainingPayment:          nil,
		NoDealReason:                       nil,
		NoDealReasonOldVehicleOthers:       nil,
		NoDealReasonOldVehicleExpectedSell: nil,
		NoDealReasonOldVehiclePriceSold:    nil,
		NoDealReasonNewVehicleOthers:       nil,

		// handover (not present on request)
		TradeInPaymentDatetime:  nil,
		TradeInHandoverStatus:   nil,
		TradeInHandoverDatetime: nil,
		TradeInHandoverLocation: nil,
		TradeInHandoverAddress:  nil,
		HandoverProvince:        nil,
		HandoverCity:            nil,
		HandoverDistrict:        nil,
		HandoverSubdistrict:     nil,
		HandoverPostalCode:      nil,

		// audit – d_createdate/d_updatedate diisi oleh DB / service layer
		CreatedDate: time.Now(),
		UpdatedDate: nil,
	}

	if dto.AppraisalStartDatetime > 0 {
		entity.AppraisalStartDatetime = utils.ToPointer(utils.GetTimeUnix(dto.AppraisalStartDatetime))
	}

	if dto.AppraisalEndDatetime > 0 {
		entity.AppraisalEndDatetime = utils.ToPointer(utils.GetTimeUnix(dto.AppraisalEndDatetime))
	}

	if dto.CancelledBy != "" {
		entity.CancelledBy = &dto.CancelledBy
	}

	if dto.CancellationReason != "" {
		entity.CancellationReason = &dto.CancellationReason
	}

	return entity
}

// =======================
// leads (+ score)
// =======================

type LeadsDTO struct {
	LeadsID   string `json:"leads_ID"`         // UUID / VARCHAR(36)
	Katashiki string `json:"katashiki_suffix"` // VARCHAR(64), conditional
	ColorCode string `json:"color_code"`       // VARCHAR(16), conditional
	Model     string `json:"model"`            // VARCHAR(64), conditional
	Variant   string `json:"variant"`          // VARCHAR(128), conditional
	Color     string `json:"color"`            // VARCHAR(64), conditional

	LeadsSource string `json:"leads_source"` // enum values (CUSTOMER_DATABASE, DIGITAL_TOYOTA_WEB, ...)
	LeadsType   string `json:"leads_type"`   // APPRAISAL_REQUEST

	LeadsFollowUpStatus string `json:"leads_follow_up_status"` // NOT_YET_FOLLOWED_UP, ON_CONSIDERATION, NO_RESPONSE

	LeadsPreferenceContactTimeStart string `json:"leads_preference_contact_time_start"` // "HH:mm", step 30 minutes
	LeadsPreferenceContactTimeEnd   string `json:"leads_preference_contact_time_end"`   // "HH:mm", step 30 minutes
	AdditionalNotes                 string `json:"additional_notes"`                    // optional: catatan dari customer untuk dealer

	Score LeadScoreDTO `json:"score"`
}

type LeadScoreDTO struct {
	TamLeadScore    string                `json:"tam_lead_score"`    // LOW, MEDIUM, HOT
	OutletLeadScore string                `json:"outlet_lead_score"` // LOW, MEDIUM, HOT
	Parameter       LeadScoreParameterDTO `json:"parameter"`
}

type LeadScoreParameterDTO struct {
	PurchasePlanCriteria    string `json:"purchase_plan_criteria"`    // 1_TO_7_DAYS, 8_TO_30_DAYS, 31_DAYS_TO_INFINITE
	PaymentPreferCriteria   string `json:"payment_prefer_criteria"`   // CASH, CREDIT_APPLIED, CREDIT_NOT_YET_APPLIED, NO_PREFERENCE
	NegotiationCriteria     string `json:"negotiation_criteria"`      // HAVE_STARTED_NEGOTIATIONS, HAVE_NOT_STARTED_NEGOTIATIONS
	TestDriveCriteria       string `json:"test_drive_criteria"`       // COMPLETED, CONFIRMED, SUBMITTED, CANCELLED, UNSPECIFIED
	TradeInCriteria         string `json:"trade_in_criteria"`         // DELIVERY, HANDOVER, PAYMENT, NEGOTIATION, CONFIRMED, SUBMITTED, CANCELLED, UNSPECIFIED
	BrowsingHistoryCriteria string `json:"browsing_history_criteria"` // MORE_THAN_5_PAGES, LESS_THAN_5_PAGES, UNSPECIFIED
	VehicleAgeCriteria      string `json:"vehicle_age_criteria"`      // MORE_THAN_2.5_YEARS, LESS_THAN_2.5_YEARS, UNSPECIFIED
}

func (dto LeadsDTO) ToLeadsModel(customerID string) (domain.Leads, domain.LeadsScore) {
	timeNow := time.Now()

	lead := domain.Leads{
		ID:         "", // generated by DB
		CustomerID: customerID,

		LeadsID:             dto.LeadsID,
		LeadsType:           dto.LeadsType,
		LeadsFollowUpStatus: dto.LeadsFollowUpStatus,
		LeadSource:          dto.LeadsSource,

		LeadsPreferenceContactTimeStart: utils.ToPointer(dto.LeadsPreferenceContactTimeStart),
		LeadsPreferenceContactTimeEnd:   utils.ToPointer(dto.LeadsPreferenceContactTimeEnd),
		AdditionalNotes:                 utils.ToPointer(dto.AdditionalNotes),

		TAMLeadScore:    dto.Score.TamLeadScore,
		OutletLeadScore: dto.Score.OutletLeadScore,

		PurchasePlanCriteria:    utils.ToPointer(dto.Score.Parameter.PurchasePlanCriteria),
		PaymentPreferCriteria:   utils.ToPointer(dto.Score.Parameter.PaymentPreferCriteria),
		TestDriveCriteria:       utils.ToPointer(dto.Score.Parameter.TestDriveCriteria),
		TradeInCriteria:         utils.ToPointer(dto.Score.Parameter.TradeInCriteria),
		BrowsingHistoryCriteria: utils.ToPointer(dto.Score.Parameter.BrowsingHistoryCriteria),
		VehicleAgeCriteria:      utils.ToPointer(dto.Score.Parameter.VehicleAgeCriteria),
		NegotiationCriteria:     utils.ToPointer(dto.Score.Parameter.NegotiationCriteria),

		CreatedAt: timeNow,
		CreatedBy: constants.System,
		UpdatedAt: timeNow,
		UpdatedBy: nil,

		// field-field “legacy/old table” yang belum ada di DTO → biarkan kosong
		GetOfferNumber:          nil,
		KatashikiSuffix:         utils.ToPointer(dto.Katashiki),
		ColorCode:               utils.ToPointer(dto.ColorCode),
		Model:                   utils.ToPointer(dto.Model),
		Variant:                 utils.ToPointer(dto.Variant),
		Color:                   utils.ToPointer(dto.Color),
		VehicleOTRPrice:         nil,
		OutletID:                nil,
		OutletName:              nil,
		ServicePackageID:        nil,
		ServicePackageName:      nil,
		CreatedDatetime:         utils.ToPointer(timeNow),
		FinanceSimulationID:     nil,
		FinanceSimulationNumber: nil,
	}

	leadScore := domain.LeadsScore{
		ID:              "", // generated by DB
		LeadsID:         dto.LeadsID,
		TamLeadScore:    dto.Score.TamLeadScore,
		OutletLeadScore: dto.Score.OutletLeadScore,

		PurchasePlanCriteria:    dto.Score.Parameter.PurchasePlanCriteria,
		PaymentPreferCriteria:   dto.Score.Parameter.PaymentPreferCriteria,
		NegotiationCriteria:     dto.Score.Parameter.NegotiationCriteria,
		TestDriveCriteria:       dto.Score.Parameter.TestDriveCriteria,
		TradeInCriteria:         dto.Score.Parameter.TradeInCriteria,
		BrowsingHistoryCriteria: dto.Score.Parameter.BrowsingHistoryCriteria,
		VehicleAgeCriteria:      dto.Score.Parameter.VehicleAgeCriteria,

		CreatedAt: timeNow,
		CreatedBy: constants.System,
		UpdatedAt: timeNow,
		UpdatedBy: constants.System,
	}

	return lead, leadScore
}

// =======================
// used_car
// =======================

type UsedCarDTO struct {
	UsedCarBrand string `json:"used_car_brand"` // VARCHAR(64), Y

	VIN          string `json:"vin,omitempty"` // VARCHAR(17), conditional
	PoliceNumber string `json:"police_number"` // VARCHAR(16), conditional

	KatashikiSuffix string `json:"katashiki_suffix"` // VARCHAR(64), conditional
	ColorCode       string `json:"color_code"`       // VARCHAR(16), conditional

	UsedCarModel   string `json:"used_car_model"`   // VARCHAR(64), conditional
	UsedCarVariant string `json:"used_car_variant"` // VARCHAR(64), conditional
	UsedCarColor   string `json:"used_car_color"`   // VARCHAR(64), conditional

	InitialEstimatedUsedCarPrice decimal.Decimal `json:"initial_estimated_used_car_price"` // FLOAT, Y

	Mileage        int64  `json:"mileage"`                  // INTEGER, Y (KM)
	Province       string `json:"province"`                 // VARCHAR(64), Y
	UsedCarYear    int    `json:"used_car_year"`            // INTEGER, N
	Transmission   string `json:"used_car_transmission"`    // VARCHAR(64), N (M_T, A_T, CVT)
	Fuel           string `json:"used_car_fuel"`            // VARCHAR(64), N (GASOLINE, DIESEL, EV, HYBRID, HYDROGEN)
	EngineCapacity int64  `json:"used_car_engine_capacity"` // INTEGER, N
	Mover          string `json:"mover"`                    // VARCHAR(64), N (4x2, 4x4, FWD, etc.)

	StnkExpiryDate int64 `json:"stnk_expiry_date"` // UNIX TIMESTAMP, Y

	Stnk   bool `json:"stnk"`   // BOOLEAN, N
	Bpkb   bool `json:"bpkb"`   // BOOLEAN, N
	Faktur bool `json:"faktur"` // BOOLEAN, N

	ServiceBook          bool `json:"service_book"`           // BOOLEAN, N
	AvailabilitySpareKey bool `json:"availability_spare_key"` // BOOLEAN, N
	SpareTireCheck       bool `json:"spare_tire_check"`       // BOOLEAN, N

	AdditionalNotes        string          `json:"additional_notes"`         // VARCHAR(256), N
	CustomerRequestedPrice decimal.Decimal `json:"customer_requested_price"` // FLOAT, N
}

func (dto UsedCarDTO) ToUsedCarModel(customerID string) domain.UsedCar {
	now := time.Now()

	var stnkExpiry time.Time
	if dto.StnkExpiryDate > 0 {
		// asumsi timestamp dalam detik
		stnkExpiry = time.Unix(dto.StnkExpiryDate, 0)
	}

	model := domain.UsedCar{
		ID:                           0, // identity di DB
		CustomerID:                   customerID,
		UsedCarBrand:                 dto.UsedCarBrand,
		VIN:                          dto.VIN,
		PoliceNumber:                 dto.PoliceNumber,
		KatashikiSuffix:              dto.KatashikiSuffix,
		ColorCode:                    dto.ColorCode,
		UsedCarModel:                 dto.UsedCarModel,
		UsedCarVariant:               dto.UsedCarVariant,
		UsedCarColor:                 dto.UsedCarColor,
		InitialEstimatedUsedCarPrice: dto.InitialEstimatedUsedCarPrice,
		UsedCarYear:                  dto.UsedCarYear,
		UsedCarFuel:                  dto.Fuel,
		UsedCarEngineCapacity:        int(dto.EngineCapacity),
		UsedCarTransmission:          dto.Transmission,
		Mileage:                      int(dto.Mileage),
		Province:                     dto.Province,
		Mover:                        dto.Mover,
		StnkExpiryDate:               stnkExpiry,
		Stnk:                         dto.Stnk,
		Bpkb:                         dto.Bpkb,
		Faktur:                       dto.Faktur,
		ServiceBook:                  dto.ServiceBook,
		AvailabilitySpareKey:         dto.AvailabilitySpareKey,
		SpareTireCheck:               dto.SpareTireCheck,
		CustomerRequestedPrice:       dto.CustomerRequestedPrice,
		AdditionalNotes:              dto.AdditionalNotes,
		CreatedAt:                    now,
	}

	return model
}
