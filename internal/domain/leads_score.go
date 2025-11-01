package domain

import (
	"time"
)

type LeadsScore struct {
	ID                      string    `json:"id" db:"i_id"`
	LeadsID                 string    `json:"leads_id" db:"i_leads_id"`
	TAMLeadScore            string    `json:"tam_lead_score" db:"v_tam_lead_score"`
	OutletLeadScore         string    `json:"outlet_lead_score" db:"v_outlet_lead_score"`
	PurchasePlanCriteria    string    `json:"purchase_plan_criteria" db:"c_purchase_plan_criteria"`
	PaymentPreferCriteria   string    `json:"payment_prefer_criteria" db:"c_payment_prefer_criteria"`
	NegotiationCriteria     string    `json:"negotiation_criteria" db:"c_negotiation_criteria"`
	TestDriveCriteria       string    `json:"test_drive_criteria" db:"c_test_drive_criteria"`
	TradeInCriteria         string    `json:"trade_in_criteria" db:"c_trade_in_criteria"`
	BrowsingHistoryCriteria string    `json:"browsing_history_criteria" db:"c_browsing_history_criteria"`
	VehicleAgeCriteria      string    `json:"vehicle_age_criteria" db:"c_vehicle_age_criteria"`
	CreatedAt               time.Time `json:"created_at" db:"d_created_at"`
	CreatedBy               string    `json:"created_by" db:"c_created_by"`
	UpdatedAt               time.Time `json:"updated_at" db:"d_updated_at"`
	UpdatedBy               string    `json:"updated_by" db:"c_updated_by"`
}

// TableName returns the database table name for the User model
func (u *LeadsScore) TableName() string {
	return "dbo.tm_leadscores"
}

// Columns returns the list of database columns for the User model
func (u *LeadsScore) Columns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"v_tam_lead_score",
		"v_outlet_lead_score",
		"c_purchase_plan_criteria",
		"c_payment_prefer_criteria",
		"c_negotiation_criteria",
		"c_test_drive_criteria",
		"c_trade_in_criteria",
		"c_browsing_history_criteria",
		"c_vehicle_age_criteria",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *LeadsScore) SelectColumns() []string {
	return []string{
		"CAST(i_id AS VARCHAR) AS i_id",
		"CAST(i_leads_id AS VARCHAR) AS i_leads_id",
		"v_tam_lead_score",
		"v_outlet_lead_score",
		"c_purchase_plan_criteria",
		"c_payment_prefer_criteria",
		"c_negotiation_criteria",
		"c_test_drive_criteria",
		"c_trade_in_criteria",
		"c_browsing_history_criteria",
		"c_vehicle_age_criteria",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

func (u *LeadsScore) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(u.Columns()))
	values = make([]interface{}, 0, len(u.Columns()))

	if u.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, u.LeadsID)
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
	if u.PaymentPreferCriteria != "" {
		columns = append(columns, "c_payment_prefer_criteria")
		values = append(values, u.PaymentPreferCriteria)
	}
	if u.NegotiationCriteria != "" {
		columns = append(columns, "c_negotiation_criteria")
		values = append(values, u.NegotiationCriteria)
	}
	if u.TestDriveCriteria != "" {
		columns = append(columns, "c_test_drive_criteria")
		values = append(values, u.TestDriveCriteria)
	}
	if u.TradeInCriteria != "" {
		columns = append(columns, "c_trade_in_criteria")
		values = append(values, u.TradeInCriteria)
	}
	if u.BrowsingHistoryCriteria != "" {
		columns = append(columns, "c_browsing_history_criteria")
		values = append(values, u.BrowsingHistoryCriteria)
	}
	if u.VehicleAgeCriteria != "" {
		columns = append(columns, "c_vehicle_age_criteria")
		values = append(values, u.VehicleAgeCriteria)
	}
	columns = append(columns, "c_created_by")
	values = append(values, u.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, u.CreatedBy)
	return columns, values
}

func (u *LeadsScore) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.TAMLeadScore != "" {
		updateMap["v_tam_lead_score"] = u.TAMLeadScore
	}
	if u.OutletLeadScore != "" {
		updateMap["v_outlet_lead_score"] = u.OutletLeadScore
	}
	if u.PurchasePlanCriteria != "" {
		updateMap["c_purchase_plan_criteria"] = u.PurchasePlanCriteria
	}
	if u.PaymentPreferCriteria != "" {
		updateMap["c_payment_prefer_criteria"] = u.PaymentPreferCriteria
	}
	if u.NegotiationCriteria != "" {
		updateMap["c_negotiation_criteria"] = u.NegotiationCriteria
	}
	if u.TestDriveCriteria != "" {
		updateMap["c_test_drive_criteria"] = u.TestDriveCriteria
	}
	if u.TradeInCriteria != "" {
		updateMap["c_trade_in_criteria"] = u.TradeInCriteria
	}
	if u.BrowsingHistoryCriteria != "" {
		updateMap["c_browsing_history_criteria"] = u.BrowsingHistoryCriteria
	}
	if u.VehicleAgeCriteria != "" {
		updateMap["c_vehicle_age_criteria"] = u.VehicleAgeCriteria
	}
	updateMap["d_updated_at"] = time.Now().UTC()
	updateMap["c_updated_by"] = u.UpdatedBy
	return updateMap
}
