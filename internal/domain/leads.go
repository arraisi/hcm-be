package domain

import (
	"time"
)

type Leads struct {
	ID                              string     `json:"id" db:"i_id"`
	LeadsID                         string     `json:"leads_ID" db:"i_leads_id"`
	LeadsType                       string     `json:"leads_type" db:"c_leads_type"`
	LeadsFollowUpStatus             string     `json:"leads_follow_up_status" db:"c_leads_follow_up_status"`
	LeadsPreferenceContactTimeStart string     `json:"leads_preference_contact_time_start" db:"d_leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   string     `json:"leads_preference_contact_time_end" db:"d_leads_preference_contact_time_end"`
	LeadSource                      string     `json:"lead_source" db:"c_lead_source"`
	TAMLeadScore                    string     `json:"tam_lead_score" db:"v_tam_lead_score"`
	OutletLeadScore                 string     `json:"outlet_lead_score" db:"v_outlet_lead_score"`
	PurchasePlanCriteria            string     `json:"purchase_plan_criteria" db:"c_purchase_plan_criteria"`
	AdditionalNotes                 *string    `json:"additional_notes" db:"e_additional_notes"`
	LastFollowUpDatetime            *time.Time `json:"last_follow_up_datetime" db:"d_last_follow_up_datetime"`
	FollowUpTargetDate              *time.Time `json:"follow_up_target_date" db:"d_follow_up_target_date"`
	GetOfferNumber                  *string    `json:"get_offer_number" db:"c_get_offer_number"`
	KatashikiSuffix                 *string    `json:"katashiki_suffix" db:"c_katashiki_suffix"`
	ColorCode                       *string    `json:"color_code" db:"c_color_code"`
	Model                           *string    `json:"model" db:"c_model"`
	Variant                         *string    `json:"variant" db:"c_variant"`
	Color                           *string    `json:"color" db:"c_color"`
	VehicleOTRPrice                 *float64   `json:"vehicle_otr_price" db:"v_vehicle_otr_price"`
	OutletID                        *string    `json:"outlet_id" db:"i_outlet_id"`
	OutletName                      *string    `json:"outlet_name" db:"n_outlet_name"`
	ServicePackageID                *string    `json:"service_package_id" db:"i_service_package_id"`
	ServicePackageName              *string    `json:"service_package_name" db:"n_service_package_name"`
	CreatedDatetime                 time.Time  `json:"created_datetime" db:"d_created_datetime"`
	LeadsStatus                     *string    `json:"leads_status" db:"c_leads_status"`
	ReasonLeadsStatusUpdate         *string    `json:"reason_leads_status_update" db:"c_reason_leads_status_update"`
	ReasonLeadsStatusUpdateOther    *string    `json:"reason_leads_status_update_other" db:"c_reason_leads_status_update_other"`
	VehicleCategory                 *string    `json:"vehicle_category" db:"c_vehicle_category"`
	DemandStructure                 *string    `json:"demand_structure" db:"d_demand_structure"`
	FinanceSimulationID             *string    `json:"finance_simulation_id" db:"i_finance_simulation_id"`
	FinanceSimulationNumber         *string    `json:"finance_simulation_number" db:"c_finance_simulation_number"`
	CreatedAt                       time.Time  `json:"created_at" db:"d_created_at"`
	CreatedBy                       string     `json:"created_by" db:"c_created_by"`
	UpdatedAt                       time.Time  `json:"updated_at" db:"d_updated_at"`
	UpdatedBy                       *string    `json:"updated_by" db:"c_updated_by"`
}

// TableName returns the database table name for the User model
func (u *Leads) TableName() string {
	return "dbo.tm_leads"
}

// Columns returns the list of database columns for the User model
func (u *Leads) Columns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"c_leads_type",
		"c_leads_follow_up_status",
		"d_leads_preference_contact_time_start",
		"d_leads_preference_contact_time_end",
		"c_lead_source",
		"v_tam_lead_score",
		"v_outlet_lead_score",
		"c_purchase_plan_criteria",
		"e_additional_notes",
		"d_last_follow_up_datetime",
		"d_follow_up_target_date",
		"c_get_offer_number",
		"c_katashiki_suffix",
		"c_color_code",
		"c_model",
		"c_variant",
		"c_color",
		"v_vehicle_otr_price",
		"i_outlet_id",
		"n_outlet_name",
		"i_service_package_id",
		"n_service_package_name",
		"d_created_datetime",
		"c_leads_status",
		"c_reason_leads_status_update",
		"c_reason_leads_status_update_other",
		"c_vehicle_category",
		"d_demand_structure",
		"i_finance_simulation_id",
		"c_finance_simulation_number",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Leads) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_leads_id AS NVARCHAR(36)) as i_leads_id",
		"c_leads_type",
		"c_leads_follow_up_status",
		"d_leads_preference_contact_time_start",
		"d_leads_preference_contact_time_end",
		"c_lead_source",
		"e_additional_notes",
		"v_tam_lead_score",
		"v_outlet_lead_score",
		"c_purchase_plan_criteria",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

func (u *Leads) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(u.Columns()))
	values = make([]interface{}, 0, len(u.Columns()))

	if u.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, u.LeadsID)
	}
	if u.LeadsType != "" {
		columns = append(columns, "c_leads_type")
		values = append(values, u.LeadsType)
	}
	if u.LeadsFollowUpStatus != "" {
		columns = append(columns, "c_leads_follow_up_status")
		values = append(values, u.LeadsFollowUpStatus)
	}
	if u.LeadsPreferenceContactTimeStart != "" {
		columns = append(columns, "d_leads_preference_contact_time_start")
		values = append(values, u.LeadsPreferenceContactTimeStart)
	}
	if u.LeadsPreferenceContactTimeEnd != "" {
		columns = append(columns, "d_leads_preference_contact_time_end")
		values = append(values, u.LeadsPreferenceContactTimeEnd)
	}
	if u.LeadSource != "" {
		columns = append(columns, "c_lead_source")
		values = append(values, u.LeadSource)
	}
	if u.TAMLeadScore != "" {
		columns = append(columns, "v_tam_lead_score")
		values = append(values, u.TAMLeadScore)
	}
	if u.OutletLeadScore != "" {
		columns = append(columns, "v_outlet_lead_score")
		values = append(values, u.OutletLeadScore)
	}
	if u.PurchasePlanCriteria != "" {
		columns = append(columns, "c_purchase_plan_criteria")
		values = append(values, u.PurchasePlanCriteria)
	}
	if u.AdditionalNotes != nil {
		columns = append(columns, "e_additional_notes")
		values = append(values, u.AdditionalNotes)
	}
	if u.LastFollowUpDatetime != nil {
		columns = append(columns, "d_last_follow_up_datetime")
		values = append(values, u.LastFollowUpDatetime.UTC())
	}
	if u.FollowUpTargetDate != nil {
		columns = append(columns, "d_follow_up_target_date")
		values = append(values, u.FollowUpTargetDate.UTC())
	}
	if u.GetOfferNumber != nil {
		columns = append(columns, "c_get_offer_number")
		values = append(values, u.GetOfferNumber)
	}
	if u.KatashikiSuffix != nil {
		columns = append(columns, "c_katashiki_suffix")
		values = append(values, u.KatashikiSuffix)
	}
	if u.ColorCode != nil {
		columns = append(columns, "c_color_code")
		values = append(values, u.ColorCode)
	}
	if u.Model != nil {
		columns = append(columns, "c_model")
		values = append(values, u.Model)
	}
	if u.Variant != nil {
		columns = append(columns, "c_variant")
		values = append(values, u.Variant)
	}
	if u.Color != nil {
		columns = append(columns, "c_color")
		values = append(values, u.Color)
	}
	if u.VehicleOTRPrice != nil {
		columns = append(columns, "v_vehicle_otr_price")
		values = append(values, u.VehicleOTRPrice)
	}
	if u.OutletID != nil {
		columns = append(columns, "i_outlet_id")
		values = append(values, u.OutletID)
	}
	if u.OutletName != nil {
		columns = append(columns, "n_outlet_name")
		values = append(values, u.OutletName)
	}
	if u.ServicePackageID != nil {
		columns = append(columns, "i_service_package_id")
		values = append(values, u.ServicePackageID)
	}
	if u.ServicePackageName != nil {
		columns = append(columns, "n_service_package_name")
		values = append(values, u.ServicePackageName)
	}
	if !u.CreatedDatetime.IsZero() {
		columns = append(columns, "d_created_datetime")
		values = append(values, u.CreatedDatetime.UTC())
	}
	if u.LeadsStatus != nil {
		columns = append(columns, "c_leads_status")
		values = append(values, u.LeadsStatus)
	}
	if u.ReasonLeadsStatusUpdate != nil {
		columns = append(columns, "c_reason_leads_status_update")
		values = append(values, u.ReasonLeadsStatusUpdate)
	}
	if u.ReasonLeadsStatusUpdateOther != nil {
		columns = append(columns, "c_reason_leads_status_update_other")
		values = append(values, u.ReasonLeadsStatusUpdateOther)
	}
	if u.VehicleCategory != nil {
		columns = append(columns, "c_vehicle_category")
		values = append(values, u.VehicleCategory)
	}
	if u.DemandStructure != nil {
		columns = append(columns, "d_demand_structure")
		values = append(values, u.DemandStructure)
	}
	if u.FinanceSimulationID != nil {
		columns = append(columns, "i_finance_simulation_id")
		values = append(values, u.FinanceSimulationID)
	}
	if u.FinanceSimulationNumber != nil {
		columns = append(columns, "c_finance_simulation_number")
		values = append(values, u.FinanceSimulationNumber)
	}
	if !u.CreatedAt.IsZero() {
		columns = append(columns, "d_created_at")
		values = append(values, u.CreatedAt.UTC())
	}
	if u.CreatedBy != "" {
		columns = append(columns, "c_created_by")
		values = append(values, u.CreatedBy)
	}
	if !u.UpdatedAt.IsZero() {
		columns = append(columns, "d_updated_at")
		values = append(values, u.UpdatedAt.UTC())
	}
	if u.UpdatedBy != nil {
		columns = append(columns, "c_updated_by")
		values = append(values, u.UpdatedBy)
	}

	return columns, values
}

func (u *Leads) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.LeadsType != "" {
		updateMap["c_leads_type"] = u.LeadsType
	}
	if u.LeadsFollowUpStatus != "" {
		updateMap["c_leads_follow_up_status"] = u.LeadsFollowUpStatus
	}
	if u.LeadsPreferenceContactTimeStart != "" {
		updateMap["d_leads_preference_contact_time_start"] = u.LeadsPreferenceContactTimeStart
	}
	if u.LeadsPreferenceContactTimeEnd != "" {
		updateMap["d_leads_preference_contact_time_end"] = u.LeadsPreferenceContactTimeEnd
	}
	if u.LeadSource != "" {
		updateMap["c_lead_source"] = u.LeadSource
	}
	if u.AdditionalNotes != nil {
		updateMap["e_additional_notes"] = u.AdditionalNotes
	}
	if u.TAMLeadScore != "" {
		updateMap["v_tam_lead_score"] = u.TAMLeadScore
	}
	if u.OutletLeadScore != "" {
		updateMap["v_outlet_lead_score"] = u.OutletLeadScore
	}
	if u.PurchasePlanCriteria != "" {
		updateMap["c_purchase_plan_criteria"] = u.PurchasePlanCriteria
	}
	updateMap["d_updated_at"] = u.UpdatedAt.UTC()
	updateMap["c_updated_by"] = u.UpdatedBy
	return updateMap
}
