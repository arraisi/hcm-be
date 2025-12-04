package domain

import (
	"time"

	"github.com/arraisi/hcm-be/pkg/utils"
)

type Leads struct {
	ID                              string     `json:"id" db:"i_id"`
	CustomerID                      string     `json:"customer_id" db:"i_customer_id"`
	LeadsID                         string     `json:"leads_id" db:"i_leads_id"`
	LeadsType                       string     `json:"leads_type" db:"c_leads_type"`
	LeadsFollowUpStatus             string     `json:"leads_follow_up_status" db:"c_leads_follow_up_status"`
	LeadsPreferenceContactTimeStart *string    `json:"leads_preference_contact_time_start" db:"t_leads_preference_contact_time_start"`
	LeadsPreferenceContactTimeEnd   *string    `json:"leads_preference_contact_time_end" db:"t_leads_preference_contact_time_end"`
	LeadSource                      string     `json:"lead_source" db:"c_leads_source"`
	AdditionalNotes                 *string    `json:"additional_notes" db:"e_additional_notes"`
	TAMLeadScore                    string     `json:"tam_lead_score" db:"c_tam_lead_score"`
	OutletLeadScore                 string     `json:"outlet_lead_score" db:"c_outlet_lead_score"`
	PurchasePlanCriteria            *string    `json:"purchase_plan_criteria" db:"c_purchase_plan_criteria"`
	PaymentPreferCriteria           *string    `json:"payment_prefer_criteria" db:"c_payment_prefer_criteria"`
	TestDriveCriteria               *string    `json:"test_drive_criteria" db:"c_test_drive_criteria"`
	TradeInCriteria                 *string    `json:"trade_in_criteria" db:"c_trade_in_criteria"`
	BrowsingHistoryCriteria         *string    `json:"browsing_history_criteria" db:"c_browsing_history_criteria"`
	VehicleAgeCriteria              *string    `json:"vehicle_age_criteria" db:"c_vehicle_age_criteria"`
	CreatedAt                       time.Time  `json:"created_at" db:"d_created_at"`
	CreatedBy                       string     `json:"created_by" db:"c_created_by"`
	UpdatedAt                       time.Time  `json:"updated_at" db:"d_updated_at"`
	UpdatedBy                       *string    `json:"updated_by" db:"c_updated_by"`
	NegotiationCriteria             *string    `json:"negotiation_criteria" db:"c_negotiation_criteria"`
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
	SalesID                         *string    `json:"sales_id" db:"c_sales_id"`
	SalesName                       *string    `json:"sales_name" db:"n_sales_name"`
	GetOfferNumber                  *string    `json:"get_offer_number" db:"c_get_offer_number"`
	FinanceSimulationID             *string    `json:"finance_simulation_id" db:"i_finance_simulation_id"`
	FinanceSimulationNumber         *string    `json:"finance_simulation_number" db:"c_finance_simulation_number"`
	CreatedDatetime                 *time.Time `json:"created_datetime" db:"d_created_datetime"`

	// from test drive
	TestDriveStatus *string `json:"test_drive_status" db:"c_test_drive_status"`
}

type LeadsList []Leads

func (l LeadsList) GetMapBySalesID() map[string]Leads {
	result := make(map[string]Leads)
	for _, leads := range l {
		if leads.SalesID != nil {
			result[utils.ToValue(leads.SalesID)] = leads
		}
	}
	return result
}

// TableName returns the database table name for the User model
func (u *Leads) TableName() string {
	return "dbo.tm_leads"
}

// TableNameAlias returns the database table name for the User model
func (u *Leads) TableNameAlias() string {
	return "dbo.tm_leads AS l"
}

// Columns returns the list of database columns for the User model
func (u *Leads) Columns() []string {
	return []string{
		"i_id",
		"i_customer_id",
		"i_leads_id",
		"c_leads_type",
		"c_leads_follow_up_status",
		"t_leads_preference_contact_time_start",
		"t_leads_preference_contact_time_end",
		"c_leads_source",
		"e_additional_notes",
		"c_tam_lead_score",
		"c_outlet_lead_score",
		"c_purchase_plan_criteria",
		"c_payment_prefer_criteria",
		"c_test_drive_criteria",
		"c_trade_in_criteria",
		"c_browsing_history_criteria",
		"c_vehicle_age_criteria",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
		"c_negotiation_criteria",
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
		"c_sales_id",
		"n_sales_name",
		"c_get_offer_number",
		"i_finance_simulation_id",
		"c_finance_simulation_number",
		"d_created_datetime",
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *Leads) SelectColumns() []string {
	return []string{
		"i_id",
		"i_customer_id",
		"i_leads_id",
		"c_leads_type",
		"c_leads_follow_up_status",
		"t_leads_preference_contact_time_start",
		"t_leads_preference_contact_time_end",
		"c_leads_source",
		"e_additional_notes",
		"c_tam_lead_score",
		"c_outlet_lead_score",
		"c_purchase_plan_criteria",
		"c_payment_prefer_criteria",
		"c_test_drive_criteria",
		"c_trade_in_criteria",
		"c_browsing_history_criteria",
		"c_vehicle_age_criteria",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
		"c_negotiation_criteria",
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
		"c_sales_id",
		"n_sales_name",
		"c_get_offer_number",
		"i_finance_simulation_id",
		"c_finance_simulation_number",
		"d_created_datetime",
	}
}

func (u *Leads) LeadsTestDriveColumns() []string {
	return []string{
		"l.i_id",
		"l.i_customer_id",
		"l.i_leads_id",
		"l.c_leads_type",
		"l.c_leads_follow_up_status",
		"l.t_leads_preference_contact_time_start",
		"l.t_leads_preference_contact_time_end",
		"l.c_leads_source",
		"l.e_additional_notes",
		"l.c_tam_lead_score",
		"l.c_outlet_lead_score",
		"l.c_purchase_plan_criteria",
		"l.c_payment_prefer_criteria",
		"l.c_test_drive_criteria",
		"l.c_trade_in_criteria",
		"l.c_browsing_history_criteria",
		"l.c_vehicle_age_criteria",
		"l.d_created_at",
		"l.c_created_by",
		"l.d_updated_at",
		"l.c_updated_by",
		"l.c_negotiation_criteria",
		"l.c_katashiki_suffix",
		"l.c_color_code",
		"l.c_model",
		"l.c_variant",
		"l.c_color",
		"l.v_vehicle_otr_price",
		"l.i_outlet_id",
		"l.n_outlet_name",
		"l.i_service_package_id",
		"l.n_service_package_name",
		"l.c_sales_id",
		"l.n_sales_name",
		"l.c_get_offer_number",
		"l.i_finance_simulation_id",
		"l.c_finance_simulation_number",
		"l.d_created_datetime",
		"td.c_test_drive_status",
	}
}

func (u *Leads) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(u.Columns()))
	values = make([]interface{}, 0, len(u.Columns()))

	if u.CustomerID != "" {
		columns = append(columns, "i_customer_id")
		values = append(values, u.CustomerID)
	}
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
	if u.LeadsPreferenceContactTimeStart != nil {
		columns = append(columns, "t_leads_preference_contact_time_start")
		values = append(values, u.LeadsPreferenceContactTimeStart)
	}
	if u.LeadsPreferenceContactTimeEnd != nil {
		columns = append(columns, "t_leads_preference_contact_time_end")
		values = append(values, u.LeadsPreferenceContactTimeEnd)
	}
	if u.LeadSource != "" {
		columns = append(columns, "c_leads_source")
		values = append(values, u.LeadSource)
	}
	if u.AdditionalNotes != nil {
		columns = append(columns, "e_additional_notes")
		values = append(values, u.AdditionalNotes)
	}
	if u.TAMLeadScore != "" {
		columns = append(columns, "c_tam_lead_score")
		values = append(values, u.TAMLeadScore)
	}
	if u.OutletLeadScore != "" {
		columns = append(columns, "c_outlet_lead_score")
		values = append(values, u.OutletLeadScore)
	}
	if u.PurchasePlanCriteria != nil {
		columns = append(columns, "c_purchase_plan_criteria")
		values = append(values, u.PurchasePlanCriteria)
	}
	if u.PaymentPreferCriteria != nil {
		columns = append(columns, "c_payment_prefer_criteria")
		values = append(values, u.PaymentPreferCriteria)
	}
	if u.TestDriveCriteria != nil {
		columns = append(columns, "c_test_drive_criteria")
		values = append(values, u.TestDriveCriteria)
	}
	if u.TradeInCriteria != nil {
		columns = append(columns, "c_trade_in_criteria")
		values = append(values, u.TradeInCriteria)
	}
	if u.BrowsingHistoryCriteria != nil {
		columns = append(columns, "c_browsing_history_criteria")
		values = append(values, u.BrowsingHistoryCriteria)
	}
	if u.VehicleAgeCriteria != nil {
		columns = append(columns, "c_vehicle_age_criteria")
		values = append(values, u.VehicleAgeCriteria)
	}
	if u.NegotiationCriteria != nil {
		columns = append(columns, "c_negotiation_criteria")
		values = append(values, u.NegotiationCriteria)
	}
	if u.GetOfferNumber != nil {
		columns = append(columns, "c_get_offer_number")
		values = append(values, u.GetOfferNumber)
	}
	if u.OutletID != nil {
		columns = append(columns, "i_outlet_id")
		values = append(values, u.OutletID)
	}
	if u.OutletName != nil {
		columns = append(columns, "n_outlet_name")
		values = append(values, u.OutletName)
	}
	if u.KatashikiSuffix != nil {
		columns = append(columns, "c_katashiki_suffix")
		values = append(values, u.KatashikiSuffix)
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
	if u.ColorCode != nil {
		columns = append(columns, "c_color_code")
		values = append(values, u.ColorCode)
	}
	if u.VehicleOTRPrice != nil {
		columns = append(columns, "v_vehicle_otr_price")
		values = append(values, u.VehicleOTRPrice)
	}
	if u.ServicePackageID != nil {
		columns = append(columns, "i_service_package_id")
		values = append(values, u.ServicePackageID)
	}
	if u.ServicePackageName != nil {
		columns = append(columns, "n_service_package_name")
		values = append(values, u.ServicePackageName)
	}
	if u.FinanceSimulationID != nil {
		columns = append(columns, "i_finance_simulation_id")
		values = append(values, u.FinanceSimulationID)
	}
	if u.FinanceSimulationNumber != nil {
		columns = append(columns, "c_finance_simulation_number")
		values = append(values, u.FinanceSimulationNumber)
	}
	if u.SalesID != nil {
		columns = append(columns, "c_sales_id")
		values = append(values, u.SalesID)
	}
	if u.SalesName != nil {
		columns = append(columns, "n_sales_name")
		values = append(values, u.SalesName)
	}
	if u.CreatedDatetime != nil {
		columns = append(columns, "d_created_datetime")
		values = append(values, u.CreatedDatetime)
	}

	columns = append(columns, "c_created_by")
	values = append(values, u.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, u.CreatedBy)

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
	if u.LeadsPreferenceContactTimeStart != nil {
		updateMap["t_leads_preference_contact_time_start"] = u.LeadsPreferenceContactTimeStart
	}
	if u.LeadsPreferenceContactTimeEnd != nil {
		updateMap["t_leads_preference_contact_time_end"] = u.LeadsPreferenceContactTimeEnd
	}
	if u.LeadSource != "" {
		updateMap["c_leads_source"] = u.LeadSource
	}
	if u.AdditionalNotes != nil {
		updateMap["e_additional_notes"] = u.AdditionalNotes
	}
	if u.TAMLeadScore != "" {
		updateMap["c_tam_lead_score"] = u.TAMLeadScore
	}
	if u.OutletLeadScore != "" {
		updateMap["c_outlet_lead_score"] = u.OutletLeadScore
	}
	if u.PurchasePlanCriteria != nil {
		updateMap["c_purchase_plan_criteria"] = u.PurchasePlanCriteria
	}
	if u.PaymentPreferCriteria != nil {
		updateMap["c_payment_prefer_criteria"] = u.PaymentPreferCriteria
	}
	if u.TestDriveCriteria != nil {
		updateMap["c_test_drive_criteria"] = u.TestDriveCriteria
	}
	if u.TradeInCriteria != nil {
		updateMap["c_trade_in_criteria"] = u.TradeInCriteria
	}
	if u.BrowsingHistoryCriteria != nil {
		updateMap["c_browsing_history_criteria"] = u.BrowsingHistoryCriteria
	}
	if u.VehicleAgeCriteria != nil {
		updateMap["c_vehicle_age_criteria"] = u.VehicleAgeCriteria
	}
	if u.NegotiationCriteria != nil {
		updateMap["c_negotiation_criteria"] = u.NegotiationCriteria
	}
	if u.GetOfferNumber != nil {
		updateMap["c_get_offer_number"] = u.GetOfferNumber
	}
	if u.OutletID != nil {
		updateMap["i_outlet_id"] = u.OutletID
	}
	if u.OutletName != nil {
		updateMap["n_outlet_name"] = u.OutletName
	}
	if u.KatashikiSuffix != nil {
		updateMap["c_katashiki_suffix"] = u.KatashikiSuffix
	}
	if u.ColorCode != nil {
		updateMap["c_color_code"] = u.ColorCode
	}
	if u.Model != nil {
		updateMap["c_model"] = u.Model
	}
	if u.Variant != nil {
		updateMap["c_variant"] = u.Variant
	}
	if u.Color != nil {
		updateMap["c_color"] = u.Color
	}
	if u.VehicleOTRPrice != nil {
		updateMap["v_vehicle_otr_price"] = u.VehicleOTRPrice
	}
	if u.ServicePackageID != nil {
		updateMap["i_service_package_id"] = u.ServicePackageID
	}
	if u.ServicePackageName != nil {
		updateMap["n_service_package_name"] = u.ServicePackageName
	}
	if u.FinanceSimulationID != nil {
		updateMap["i_finance_simulation_id"] = u.FinanceSimulationID
	}
	if u.FinanceSimulationNumber != nil {
		updateMap["c_finance_simulation_number"] = u.FinanceSimulationNumber
	}
	if u.SalesID != nil {
		updateMap["c_sales_id"] = u.SalesID
	}
	if u.SalesName != nil {
		updateMap["n_sales_name"] = u.SalesName
	}
	if u.CreatedDatetime != nil {
		updateMap["d_created_datetime"] = u.CreatedDatetime
	}
	updateMap["c_updated_by"] = u.UpdatedBy
	return updateMap
}
