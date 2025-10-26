package domain

type Lead struct {
	IID                             string  `json:"i_id" db:"i_id"`
	LeadsID                         string  `json:"leads_ID" db:"leads_id"`
	LeadsType                       string  `json:"leads_type" db:"leads_type"`
	LeadsFollowUpStatus             string  `json:"leads_follow_up_status" db:"leads_follow_up_status"`
	LeadsPreferenceContactTimeStart string  `json:"leads_preference_contact_time_start" db:"leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   string  `json:"leads_preference_contact_time_end" db:"leads_preference_contact_time_end"`
	LeadSource                      string  `json:"lead_source" db:"lead_source"`
	AdditionalNotes                 *string `json:"additional_notes" db:"additional_notes"`

	// TBD: what the difference between these three scores with table lead_score?
	TAMLeadScore         string `json:"tam_lead_score" db:"tam_lead_score"`
	OutletLeadScore      string `json:"outlet_lead_score" db:"outlet_lead_score"`
	PurchasePlanCriteria string `json:"purchase_plan_criteria" db:"purchase_plan_criteria"`
}

// TableName returns the database table name for the User model
func (u *Lead) TableName() string {
	return "dbo.tm_leads"
}

// Columns returns the list of database columns for the User model
func (u *Lead) Columns() []string {
	return []string{
		"i_id",
		"leads_id",
		"leads_type",
		"leads_follow_up_status",
		"leads_preference_contact_time_start",
		"leads_preference_contact_time_end",
		"lead_source",
		"additional_notes",
		"tam_lead_score",
		"outlet_lead_score",
		"purchase_plan_criteria",
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Lead) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(leads_id AS NVARCHAR(36)) as leads_id",
		"leads_type",
		"leads_follow_up_status",
		"leads_preference_contact_time_start",
		"leads_preference_contact_time_end",
		"lead_source",
		"additional_notes",
		"tam_lead_score",
		"outlet_lead_score",
		"purchase_plan_criteria",
	}
}
