package domain

import (
	"time"
)

type LeadsScore struct {
	ID                      string    `json:"id" db:"i_id"`
	LeadsID                 string    `json:"leads_id" db:"i_leads_id"`
	TamLeadScore            string    `json:"tam_lead_score" db:"v_tam_lead_score"`
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

//
// ─────────────────────────────────────────────────────────────
//  TABLE NAME
// ─────────────────────────────────────────────────────────────
//

func (l *LeadsScore) TableName() string {
	return "dbo.tm_leads_score"
}

//
// ─────────────────────────────────────────────────────────────
//  COLUMNS (for insert)
// ─────────────────────────────────────────────────────────────
//

func (l *LeadsScore) Columns() []string {
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

//
// ─────────────────────────────────────────────────────────────
//  SELECT COLUMNS (with aliasing as needed)
// ─────────────────────────────────────────────────────────────
//

func (l *LeadsScore) SelectColumns() []string {
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

//
// ─────────────────────────────────────────────────────────────
//  CREATE MAP (for INSERT)
// ─────────────────────────────────────────────────────────────
//

func (l *LeadsScore) ToCreateMap() (columns []string, values []interface{}) {
	columns = []string{}
	values = []interface{}{}

	if l.ID != "" {
		columns = append(columns, "i_id")
		values = append(values, l.ID)
	}
	if l.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, l.LeadsID)
	}
	if l.TamLeadScore != "" {
		columns = append(columns, "v_tam_lead_score")
		values = append(values, l.TamLeadScore)
	}
	if l.OutletLeadScore != "" {
		columns = append(columns, "v_outlet_lead_score")
		values = append(values, l.OutletLeadScore)
	}
	if l.PurchasePlanCriteria != "" {
		columns = append(columns, "c_purchase_plan_criteria")
		values = append(values, l.PurchasePlanCriteria)
	}
	if l.PaymentPreferCriteria != "" {
		columns = append(columns, "c_payment_prefer_criteria")
		values = append(values, l.PaymentPreferCriteria)
	}
	if l.NegotiationCriteria != "" {
		columns = append(columns, "c_negotiation_criteria")
		values = append(values, l.NegotiationCriteria)
	}
	if l.TestDriveCriteria != "" {
		columns = append(columns, "c_test_drive_criteria")
		values = append(values, l.TestDriveCriteria)
	}
	if l.TradeInCriteria != "" {
		columns = append(columns, "c_trade_in_criteria")
		values = append(values, l.TradeInCriteria)
	}
	if l.BrowsingHistoryCriteria != "" {
		columns = append(columns, "c_browsing_history_criteria")
		values = append(values, l.BrowsingHistoryCriteria)
	}
	if l.VehicleAgeCriteria != "" {
		columns = append(columns, "c_vehicle_age_criteria")
		values = append(values, l.VehicleAgeCriteria)
	}

	columns = append(columns, "d_created_at")
	values = append(values, l.CreatedAt)

	columns = append(columns, "c_created_by")
	values = append(values, l.CreatedBy)

	columns = append(columns, "d_updated_at")
	values = append(values, l.UpdatedAt)

	if l.UpdatedBy != "" {
		columns = append(columns, "c_updated_by")
		values = append(values, l.UpdatedBy)
	}

	return
}

//
// ─────────────────────────────────────────────────────────────
//  UPDATE MAP (for UPDATE)
// ─────────────────────────────────────────────────────────────
//

func (l *LeadsScore) ToUpdateMap() map[string]interface{} {
	m := map[string]interface{}{}

	if l.TamLeadScore != "" {
		m["v_tam_lead_score"] = l.TamLeadScore
	}
	if l.OutletLeadScore != "" {
		m["v_outlet_lead_score"] = l.OutletLeadScore
	}
	if l.PurchasePlanCriteria != "" {
		m["c_purchase_plan_criteria"] = l.PurchasePlanCriteria
	}
	if l.PaymentPreferCriteria != "" {
		m["c_payment_prefer_criteria"] = l.PaymentPreferCriteria
	}
	if l.NegotiationCriteria != "" {
		m["c_negotiation_criteria"] = l.NegotiationCriteria
	}
	if l.TestDriveCriteria != "" {
		m["c_test_drive_criteria"] = l.TestDriveCriteria
	}
	if l.TradeInCriteria != "" {
		m["c_trade_in_criteria"] = l.TradeInCriteria
	}
	if l.BrowsingHistoryCriteria != "" {
		m["c_browsing_history_criteria"] = l.BrowsingHistoryCriteria
	}
	if l.VehicleAgeCriteria != "" {
		m["c_vehicle_age_criteria"] = l.VehicleAgeCriteria
	}

	m["d_updated_at"] = l.UpdatedAt

	if l.UpdatedBy != "" {
		m["c_updated_by"] = l.UpdatedBy
	}

	return m
}
