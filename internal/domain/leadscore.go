package domain

type LeadScore struct {
	IID                     string `json:"i_id" db:"i_id"`
	TAMLeadScore            string `json:"tam_lead_score" db:"tam_lead_score"`
	OutletLeadScore         string `json:"outlet_lead_score" db:"outlet_lead_score"`
	PurchasePlanCriteria    string `json:"purchase_plan_criteria" db:"purchase_plan_criteria"`
	PaymentPreferCriteria   string `json:"payment_prefer_criteria" db:"payment_prefer_criteria"`
	NegotiationCriteria     string `json:"negotiation_criteria" db:"negotiation_criteria"`
	TestDriveCriteria       string `json:"test_drive_criteria" db:"test_drive_criteria"`
	TradeInCriteria         string `json:"trade_in_criteria" db:"trade_in_criteria"`
	BrowsingHistoryCriteria string `json:"browsing_history_criteria" db:"browsing_history_criteria"`
	VehicleAgeCriteria      string `json:"vehicle_age_criteria" db:"vehicle_age_criteria"`
}

// TableName returns the database table name for the User model
func (u *LeadScore) TableName() string {
	return "dbo.tm_leadscores"
}

// Columns returns the list of database columns for the User model
func (u *LeadScore) Columns() []string {
	return []string{
		"i_id",
		"tam_lead_score",
		"outlet_lead_score",
		"purchase_plan_criteria",
		"payment_prefer_criteria",
		"negotiation_criteria",
		"test_drive_criteria",
		"trade_in_criteria",
		"browsing_history_criteria",
		"vehicle_age_criteria",
	}
}

func (u *LeadScore) ToValues() []interface{} {
	return []interface{}{
		u.IID,
		u.TAMLeadScore,
		u.OutletLeadScore,
		u.PurchasePlanCriteria,
		u.PaymentPreferCriteria,
		u.NegotiationCriteria,
		u.TestDriveCriteria,
		u.TradeInCriteria,
		u.BrowsingHistoryCriteria,
		u.VehicleAgeCriteria,
	}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *LeadScore) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"tam_lead_score",
		"outlet_lead_score",
		"purchase_plan_criteria",
		"payment_prefer_criteria",
		"negotiation_criteria",
		"test_drive_criteria",
		"trade_in_criteria",
		"browsing_history_criteria",
		"vehicle_age_criteria",
	}
}

func (u *LeadScore) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.TAMLeadScore != "" {
		updateMap["tam_lead_score"] = u.TAMLeadScore
	}
	if u.OutletLeadScore != "" {
		updateMap["outlet_lead_score"] = u.OutletLeadScore
	}
	if u.PurchasePlanCriteria != "" {
		updateMap["purchase_plan_criteria"] = u.PurchasePlanCriteria
	}
	if u.PaymentPreferCriteria != "" {
		updateMap["payment_prefer_criteria"] = u.PaymentPreferCriteria
	}
	if u.NegotiationCriteria != "" {
		updateMap["negotiation_criteria"] = u.NegotiationCriteria
	}
	if u.TestDriveCriteria != "" {
		updateMap["test_drive_criteria"] = u.TestDriveCriteria
	}
	if u.TradeInCriteria != "" {
		updateMap["trade_in_criteria"] = u.TradeInCriteria
	}
	if u.BrowsingHistoryCriteria != "" {
		updateMap["browsing_history_criteria"] = u.BrowsingHistoryCriteria
	}
	if u.VehicleAgeCriteria != "" {
		updateMap["vehicle_age_criteria"] = u.VehicleAgeCriteria
	}
	return updateMap
}
