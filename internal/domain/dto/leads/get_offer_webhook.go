package leads

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/google/uuid"
)

// GetOfferWebhookEvent represents the webhook event for get offer request
type GetOfferWebhookEvent struct {
	Process   string              `json:"process" validate:"required"`
	EventID   string              `json:"event_ID" validate:"required"`
	Timestamp int64               `json:"timestamp" validate:"required"`
	Data      GetOfferWebhookData `json:"data" validate:"required"`
}

// GetOfferWebhookData represents the main data in the webhook
type GetOfferWebhookData struct {
	OneAccount     customer.OneAccountRequest `json:"one_account" validate:"required"`
	Leads          GetOfferLeadsRequest       `json:"leads" validate:"required"`
	InterestedPart []InterestedPart           `json:"interested_part"`
	TradeIn        TradeInRequest             `json:"trade_in"`
}

// GetOfferLeadsRequest represents the leads information in get offer webhook
type GetOfferLeadsRequest struct {
	LeadsID                         string  `json:"leads_ID" validate:"required"`
	GetOfferNumber                  string  `json:"get_offer_number" validate:"required"`
	LeadsSource                     string  `json:"leads_source"`
	LeadsType                       string  `json:"leads_type"`
	LeadsFollowUpStatus             string  `json:"leads_follow_up_status"`
	KatashikiSuffix                 string  `json:"katashiki_suffix"`
	ColorCode                       string  `json:"color_code"`
	Model                           string  `json:"model"`
	Variant                         string  `json:"variant"`
	Color                           string  `json:"color"`
	VehicleOTRPrice                 float64 `json:"vehicle_otr_price"`
	OutletID                        string  `json:"outlet_ID"`
	OutletName                      string  `json:"outlet_name"`
	CreatedDatetime                 int64   `json:"created_datetime"`
	ServicePackageID                string  `json:"service_package_ID"`
	ServicePackageName              string  `json:"service_package_name"`
	LeadsPreferenceContactTimeStart string  `json:"leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   string  `json:"leads_preference_contact_time_end"`
	AdditionalNotes                 string  `json:"additional_notes"`
	Score                           Score   `json:"score"`
}

// ToDomain converts GetOfferLeadsRequest to domain.Leads
func (req *GetOfferLeadsRequest) ToDomain(customerID string) domain.Leads {
	createdDatetime := time.Unix(req.CreatedDatetime, 0)

	return domain.Leads{
		ID:                              uuid.New().String(),
		CustomerID:                      customerID,
		LeadsID:                         req.LeadsID,
		LeadsType:                       req.LeadsType,
		LeadsFollowUpStatus:             req.LeadsFollowUpStatus,
		LeadsPreferenceContactTimeStart: utils.ToPointer(req.LeadsPreferenceContactTimeStart),
		LeadsPreferenceContactTimeEnd:   utils.ToPointer(req.LeadsPreferenceContactTimeEnd),
		LeadSource:                      req.LeadsSource,
		AdditionalNotes:                 utils.ToPointer(req.AdditionalNotes),
		TAMLeadScore:                    req.Score.TAMLeadScore,
		OutletLeadScore:                 req.Score.OutletLeadScore,
		PurchasePlanCriteria:            utils.ToPointer(req.Score.Parameter.PurchasePlanCriteria),
		PaymentPreferCriteria:           utils.ToPointer(req.Score.Parameter.PaymentPreferCriteria),
		TestDriveCriteria:               utils.ToPointer(req.Score.Parameter.TestDriveCriteria),
		TradeInCriteria:                 utils.ToPointer(req.Score.Parameter.TradeInCriteria),
		BrowsingHistoryCriteria:         utils.ToPointer(req.Score.Parameter.BrowsingHistoryCriteria),
		VehicleAgeCriteria:              utils.ToPointer(req.Score.Parameter.VehicleAgeCriteria),
		NegotiationCriteria:             utils.ToPointer(req.Score.Parameter.NegotiationCriteria),
		GetOfferNumber:                  utils.ToPointer(req.GetOfferNumber),
		KatashikiSuffix:                 utils.ToPointer(req.KatashikiSuffix),
		ColorCode:                       utils.ToPointer(req.ColorCode),
		Model:                           utils.ToPointer(req.Model),
		Variant:                         utils.ToPointer(req.Variant),
		Color:                           utils.ToPointer(req.Color),
		VehicleOTRPrice:                 utils.ToPointer(req.VehicleOTRPrice),
		OutletID:                        utils.ToPointer(req.OutletID),
		OutletName:                      utils.ToPointer(req.OutletName),
		ServicePackageID:                utils.ToPointer(req.ServicePackageID),
		ServicePackageName:              utils.ToPointer(req.ServicePackageName),
		CreatedDatetime:                 createdDatetime,
	}
}
