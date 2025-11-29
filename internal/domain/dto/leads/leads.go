package leads

import (
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/pkg/constants"
	"github.com/arraisi/hcm-be/pkg/utils"
	"github.com/elgris/sqrl"
)

// LeadsRequest represents the leads information from the webhook
type LeadsRequest struct {
	LeadsID                         string  `json:"leads_id" validate:"required"`
	LeadsType                       string  `json:"leads_type" validate:"required"`
	LeadsFollowUpStatus             string  `json:"leads_follow_up_status" validate:"required"`
	LeadsPreferenceContactTimeStart string  `json:"leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   string  `json:"leads_preference_contact_time_end"`
	LeadSource                      string  `json:"leads_source" validate:"required"`
	AdditionalNotes                 *string `json:"additional_notes"`
	LeadsScore                      Score   `json:"score" validate:"required"`
}

func NewLeadsRequest(leads domain.Leads) LeadsRequest {
	return LeadsRequest{
		LeadsID:                         leads.LeadsID,
		LeadsType:                       leads.LeadsType,
		LeadsFollowUpStatus:             leads.LeadsFollowUpStatus,
		LeadsPreferenceContactTimeStart: leads.LeadsPreferenceContactTimeStart,
		LeadsPreferenceContactTimeEnd:   leads.LeadsPreferenceContactTimeEnd,
		LeadSource:                      leads.LeadSource,
		AdditionalNotes:                 leads.AdditionalNotes,
		LeadsScore: Score{
			TAMLeadScore:    leads.TAMLeadScore,
			OutletLeadScore: leads.OutletLeadScore,
			Parameter: ScoreParameter{
				PurchasePlanCriteria:    leads.PurchasePlanCriteria,
				PaymentPreferCriteria:   leads.PaymentPreferCriteria,
				NegotiationCriteria:     leads.NegotiationCriteria,
				TestDriveCriteria:       leads.TestDriveCriteria,
				TradeInCriteria:         leads.TradeInCriteria,
				BrowsingHistoryCriteria: leads.BrowsingHistoryCriteria,
				VehicleAgeCriteria:      leads.VehicleAgeCriteria,
			},
		},
	}
}

// ScoreParameter represents the parameter information in score
type ScoreParameter struct {
	PurchasePlanCriteria    string `json:"purchase_plan_criteria"`
	PaymentPreferCriteria   string `json:"payment_prefer_criteria"`
	NegotiationCriteria     string `json:"negotiation_criteria"`
	TestDriveCriteria       string `json:"test_drive_criteria"`
	TradeInCriteria         string `json:"trade_in_criteria"`
	BrowsingHistoryCriteria string `json:"browsing_history_criteria"`
	VehicleAgeCriteria      string `json:"vehicle_age_criteria"`
}

// Score represents the score information from the webhook
type Score struct {
	TAMLeadScore    string         `json:"tam_lead_score"`
	OutletLeadScore string         `json:"outlet_lead_score"`
	Parameter       ScoreParameter `json:"parameter"`
}

type GetLeadsRequest struct {
	ID      *string
	LeadsID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetLeadsRequest) Apply(q *sqrl.SelectBuilder) {
	if req.ID != nil {
		q.Where(sqrl.Eq{"i_id": req.ID})
	}
	if req.LeadsID != nil {
		q.Where(sqrl.Eq{"i_leads_id": req.LeadsID})
	}
}

// ToDomain converts the RequestTestDrive to the internal Leads model
func (be *LeadsRequest) ToDomain(customerID string) domain.Leads {
	return domain.Leads{
		CustomerID:                      customerID,
		LeadsID:                         be.LeadsID,
		LeadsType:                       be.LeadsType,
		LeadsFollowUpStatus:             be.LeadsFollowUpStatus,
		LeadsPreferenceContactTimeStart: be.LeadsPreferenceContactTimeStart,
		LeadsPreferenceContactTimeEnd:   be.LeadsPreferenceContactTimeEnd,
		LeadSource:                      be.LeadSource,
		AdditionalNotes:                 be.AdditionalNotes,
		TAMLeadScore:                    be.LeadsScore.TAMLeadScore,
		OutletLeadScore:                 be.LeadsScore.OutletLeadScore,
		PurchasePlanCriteria:            be.LeadsScore.Parameter.PurchasePlanCriteria,
		PaymentPreferCriteria:           be.LeadsScore.Parameter.PaymentPreferCriteria,
		TestDriveCriteria:               be.LeadsScore.Parameter.TestDriveCriteria,
		TradeInCriteria:                 be.LeadsScore.Parameter.TradeInCriteria,
		BrowsingHistoryCriteria:         be.LeadsScore.Parameter.BrowsingHistoryCriteria,
		VehicleAgeCriteria:              be.LeadsScore.Parameter.VehicleAgeCriteria,
		NegotiationCriteria:             be.LeadsScore.Parameter.NegotiationCriteria,
		CreatedAt:                       time.Now(),
		CreatedBy:                       constants.System,
		UpdatedAt:                       time.Now(),
		UpdatedBy:                       utils.ToPointer(constants.System),
	}
}
