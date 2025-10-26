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

func (u *Lead) ToValues() []interface{} {
	return []interface{}{
		u.IID,
		u.LeadsID,
		u.LeadsType,
		u.LeadsFollowUpStatus,
		u.LeadsPreferenceContactTimeStart,
		u.LeadsPreferenceContactTimeEnd,
		u.LeadSource,
		u.AdditionalNotes,
		u.TAMLeadScore,
		u.OutletLeadScore,
		u.PurchasePlanCriteria,
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

func (u *Lead) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.LeadsType != "" {
		updateMap["leads_type"] = u.LeadsType
	}
	if u.LeadsFollowUpStatus != "" {
		updateMap["leads_follow_up_status"] = u.LeadsFollowUpStatus
	}
	if u.LeadsPreferenceContactTimeStart != "" {
		updateMap["leads_preference_contact_time_start"] = u.LeadsPreferenceContactTimeStart
	}
	if u.LeadsPreferenceContactTimeEnd != "" {
		updateMap["leads_preference_contact_time_end"] = u.LeadsPreferenceContactTimeEnd
	}
	if u.LeadSource != "" {
		updateMap["lead_source"] = u.LeadSource
	}
	if u.AdditionalNotes != nil {
		updateMap["additional_notes"] = u.AdditionalNotes
	}
	if u.TAMLeadScore != "" {
		updateMap["tam_lead_score"] = u.TAMLeadScore
	}
	if u.OutletLeadScore != "" {
		updateMap["outlet_lead_score"] = u.OutletLeadScore
	}
	if u.PurchasePlanCriteria != "" {
		updateMap["purchase_plan_criteria"] = u.PurchasePlanCriteria
	}
	return updateMap
}
