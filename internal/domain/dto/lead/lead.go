package lead

// LeadsRequest represents the leads information from the webhook
type LeadsRequest struct {
	LeadsID                         string  `json:"leads_ID" validate:"required"`
	LeadsType                       string  `json:"leads_type" validate:"required,eq=TEST_DRIVE_REQUEST"`
	LeadsFollowUpStatus             string  `json:"leads_follow_up_status" validate:"required"`
	LeadsPreferenceContactTimeStart string  `json:"leads_preference_contact_time_start" validate:"required"`
	LeadsPreferenceContactTimeEnd   string  `json:"leads_preference_contact_time_end" validate:"required"`
	LeadsSource                     string  `json:"leads_source" validate:"required"`
	AdditionalNotes                 *string `json:"additional_notes"`
}

// ScoreParameter represents the parameter information in score
type ScoreParameter struct {
	PurchasePlanCriteria    string `json:"purchase_plan_criteria" validate:"required"`
	PaymentPreferCriteria   string `json:"payment_prefer_criteria" validate:"required"`
	NegotiationCriteria     string `json:"negotiation_criteria" validate:"required"`
	TestDriveCriteria       string `json:"test_drive_criteria" validate:"required"`
	TradeInCriteria         string `json:"trade_in_criteria" validate:"required"`
	BrowsingHistoryCriteria string `json:"browsing_history_criteria" validate:"required"`
	VehicleAgeCriteria      string `json:"vehicle_age_criteria" validate:"required"`
}

// Score represents the score information from the webhook
type Score struct {
	IAMLeadScore    string         `json:"iam_lead_score" validate:"required"`
	OutletLeadScore string         `json:"outlet_lead_score" validate:"required"`
	Parameter       ScoreParameter `json:"parameter" validate:"required"`
}
