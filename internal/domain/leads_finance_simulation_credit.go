package domain

import "github.com/elgris/sqrl"

// LeadsFinanceSimulationCredit represents the credit simulation results for finance simulation
type LeadsFinanceSimulationCredit struct {
	ID                       string
	LeadsID                  string
	LeadsFinanceSimulationID string
	Tenor                    int
	DownPayment              float64
	TotalFirstPayment        float64
	MonthlyInstallment       float64
	IsActive                 bool
}

// TableName returns the table name for LeadsFinanceSimulationCredit
func (LeadsFinanceSimulationCredit) TableName() string {
	return "tr_leads_finance_simulation_credit"
}

// Columns returns the column names for LeadsFinanceSimulationCredit
func (LeadsFinanceSimulationCredit) Columns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"i_leads_finance_simulation_id",
		"v_tenor",
		"v_down_payment",
		"v_total_first_payment",
		"v_monthly_installment",
		"b_is_active",
	}
}

// SelectColumns returns the column selections for LeadsFinanceSimulationCredit
func (LeadsFinanceSimulationCredit) SelectColumns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"i_leads_finance_simulation_id",
		"v_tenor",
		"v_down_payment",
		"v_total_first_payment",
		"v_monthly_installment",
		"b_is_active",
	}
}

// ToCreateMap converts LeadsFinanceSimulationCredit to a map for insertion
func (l *LeadsFinanceSimulationCredit) ToCreateMap() map[string]interface{} {
	return map[string]interface{}{
		"i_id":                          l.ID,
		"i_leads_id":                    l.LeadsID,
		"i_leads_finance_simulation_id": l.LeadsFinanceSimulationID,
		"v_tenor":                       l.Tenor,
		"v_down_payment":                l.DownPayment,
		"v_total_first_payment":         l.TotalFirstPayment,
		"v_monthly_installment":         l.MonthlyInstallment,
		"b_is_active":                   l.IsActive,
	}
}

// ToUpdateMap converts LeadsFinanceSimulationCredit to a map for update
func (l *LeadsFinanceSimulationCredit) ToUpdateMap() map[string]interface{} {
	updateMap := sqrl.Eq{}
	updateMap["i_leads_id"] = l.LeadsID
	updateMap["i_leads_finance_simulation_id"] = l.LeadsFinanceSimulationID
	updateMap["v_tenor"] = l.Tenor
	updateMap["v_down_payment"] = l.DownPayment
	updateMap["v_total_first_payment"] = l.TotalFirstPayment
	updateMap["v_monthly_installment"] = l.MonthlyInstallment
	updateMap["b_is_active"] = l.IsActive
	return updateMap
}
