package domain

import "time"

type Leads struct {
	ID                              string    `json:"id" db:"id"`
	LeadsID                         string    `json:"leads_ID" db:"leads_id"`
	LeadsType                       string    `json:"leads_type" db:"leads_type"`
	LeadsFollowUpStatus             string    `json:"leads_follow_up_status" db:"leads_follow_up_status"`
	LeadsPreferenceContactTimeStart string    `json:"leads_preference_contact_time_start" db:"leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   string    `json:"leads_preference_contact_time_end" db:"leads_preference_contact_time_end"`
	LeadSource                      string    `json:"lead_source" db:"lead_source"`
	AdditionalNotes                 *string   `json:"additional_notes" db:"additional_notes"`
	CreatedAt                       time.Time `json:"created_at" db:"created_at"`
	CreatedBy                       string    `json:"created_by" db:"created_by"`
	UpdatedAt                       time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy                       *string   `json:"updated_by" db:"updated_by"`

	// TBD: what the difference between these three scores with table lead_score?
	TAMLeadScore         string `json:"tam_lead_score" db:"tam_lead_score"`
	OutletLeadScore      string `json:"outlet_lead_score" db:"outlet_lead_score"`
	PurchasePlanCriteria string `json:"purchase_plan_criteria" db:"purchase_plan_criteria"`
}

// TableName returns the database table name for the User model
func (u *Leads) TableName() string {
	return "dbo.tm_leads"
}

// Columns returns the list of database columns for the User model
func (u *Leads) Columns() []string {
	return []string{
		"id",
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
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	}
}

func (u *Leads) ToValues() []interface{} {
	return []interface{}{
		u.ID,
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
		u.CreatedAt,
		u.CreatedBy,
		u.UpdatedAt,
		u.UpdatedBy,
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Leads) SelectColumns() []string {
	return []string{
		"CAST(id AS NVARCHAR(36)) as id",
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
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	}
}

func (u *Leads) ToUpdateMap() map[string]interface{} {
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
	updateMap["updated_at"] = u.UpdatedAt
	updateMap["updated_by"] = u.UpdatedBy
	return updateMap
}
