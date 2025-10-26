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
