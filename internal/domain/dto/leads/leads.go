package leads

import (
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
	IID     *string
	LeadsID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetLeadsRequest) Apply(q *sqrl.SelectBuilder) {
	if req.IID != nil {
		q.Where(sqrl.Eq{"i_id": req.IID})
	}
	if req.LeadsID != nil {
		q.Where(sqrl.Eq{"leads_id": req.LeadsID})
	}
}

type GetLeadScoreRequest struct {
	IID *string
}

// Apply applies the request parameters to the given SelectBuilder
func (req GetLeadScoreRequest) Apply(q *sqrl.SelectBuilder) {
	if req.IID != nil {
		q.Where(sqrl.Eq{"i_id": req.IID})
	}
}
