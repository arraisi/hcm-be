package domain

type LeadsFinanceSimulation struct {
	ID                   string `json:"id" db:"i_id"`
	SimulationID         string `json:"simulation_id" db:"i_simulation_id"`
	SimulationNumber     string `json:"simulation_number" db:"c_simulation_number"`
	LeadsID              string `json:"leads_id" db:"i_leads_id"`
	PaymentPreference    string `json:"payment_preference" db:"c_payment_preference"`
	InsurancePeriod      string `json:"insurance_period" db:"c_insurance_period"`
	CreditMethod         string `json:"credit_method" db:"c_credit_method"`
	PackageType          string `json:"package_type" db:"c_package_type"`
	FirstPaymentType     string `json:"first_payment_type" db:"c_first_payment_type"`
	InsuranceType        string `json:"insurance_type" db:"c_insurance_type"`
	InsurancePaymentType string `json:"insurance_payment_type" db:"c_insurance_payment_type"`
	InsuranceCoverage    string `json:"insurance_coverage" db:"c_insurance_coverage"`
	//CreatedAt              time.Time `json:"created_at" db:"d_created_at"`
	//UpdatedAt              time.Time `json:"updated_at" db:"d_updated_at"`
}

// TableName returns the database table name for the LeadsFinanceSimulation model
func (l *LeadsFinanceSimulation) TableName() string {
	return "dbo.tr_leads_finance_simulation"
}

// Columns returns the list of database columns for the LeadsFinanceSimulation model
func (l *LeadsFinanceSimulation) Columns() []string {
	return []string{
		"i_id",
		"i_simulation_id",
		"c_simulation_number",
		"i_leads_id",
		"c_payment_preference",
		"c_insurance_period",
		"c_credit_method",
		"c_package_type",
		"c_first_payment_type",
		"c_insurance_type",
		"c_insurance_payment_type",
		"c_insurance_coverage",
		//"d_created_at",
		//"d_updated_at",
	}
}

// SelectColumns returns the list of columns to select in queries for the LeadsFinanceSimulation model
func (l *LeadsFinanceSimulation) SelectColumns() []string {
	return []string{
		"i_id",
		"i_simulation_id",
		"c_simulation_number",
		"i_leads_id",
		"c_payment_preference",
		"c_insurance_period",
		"c_credit_method",
		"c_package_type",
		"c_first_payment_type",
		"c_insurance_type",
		"c_insurance_payment_type",
		"c_insurance_coverage",
		//"d_created_at",
		//"d_updated_at",
	}
}

// ToCreateMap converts the model to columns and values for insert operation
func (l *LeadsFinanceSimulation) ToCreateMap() (columns []string, values []interface{}) {
	columns = make([]string, 0, len(l.Columns()))
	values = make([]interface{}, 0, len(l.Columns()))

	if l.SimulationID != "" {
		columns = append(columns, "i_simulation_id")
		values = append(values, l.SimulationID)
	}
	if l.SimulationNumber != "" {
		columns = append(columns, "c_simulation_number")
		values = append(values, l.SimulationNumber)
	}
	if l.LeadsID != "" {
		columns = append(columns, "i_leads_id")
		values = append(values, l.LeadsID)
	}
	if l.PaymentPreference != "" {
		columns = append(columns, "c_payment_preference")
		values = append(values, l.PaymentPreference)
	}
	if l.InsurancePeriod != "" {
		columns = append(columns, "c_insurance_period")
		values = append(values, l.InsurancePeriod)
	}
	if l.CreditMethod != "" {
		columns = append(columns, "c_credit_method")
		values = append(values, l.CreditMethod)
	}
	if l.PackageType != "" {
		columns = append(columns, "c_package_type")
		values = append(values, l.PackageType)
	}
	if l.FirstPaymentType != "" {
		columns = append(columns, "c_first_payment_type")
		values = append(values, l.FirstPaymentType)
	}
	if l.InsuranceType != "" {
		columns = append(columns, "c_insurance_type")
		values = append(values, l.InsuranceType)
	}
	if l.InsurancePaymentType != "" {
		columns = append(columns, "c_insurance_payment_type")
		values = append(values, l.InsurancePaymentType)
	}
	if l.InsuranceCoverage != "" {
		columns = append(columns, "c_insurance_coverage")
		values = append(values, l.InsuranceCoverage)
	}
	//if !l.CreatedAt.IsZero() {
	//	columns = append(columns, "d_created_at")
	//	values = append(values, l.CreatedAt.UTC())
	//}
	//if !l.UpdatedAt.IsZero() {
	//	columns = append(columns, "d_updated_at")
	//	values = append(values, l.UpdatedAt.UTC())
	//}

	return columns, values
}

// ToUpdateMap converts the model to a map for update operation
func (l *LeadsFinanceSimulation) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if l.SimulationID != "" {
		updateMap["i_simulation_id"] = l.SimulationID
	}
	if l.SimulationNumber != "" {
		updateMap["c_simulation_number"] = l.SimulationNumber
	}
	if l.PaymentPreference != "" {
		updateMap["c_payment_preference"] = l.PaymentPreference
	}
	if l.InsurancePeriod != "" {
		updateMap["c_insurance_period"] = l.InsurancePeriod
	}
	if l.CreditMethod != "" {
		updateMap["c_credit_method"] = l.CreditMethod
	}
	if l.PackageType != "" {
		updateMap["c_package_type"] = l.PackageType
	}
	if l.FirstPaymentType != "" {
		updateMap["c_first_payment_type"] = l.FirstPaymentType
	}
	if l.InsuranceType != "" {
		updateMap["c_insurance_type"] = l.InsuranceType
	}
	if l.InsurancePaymentType != "" {
		updateMap["c_insurance_payment_type"] = l.InsurancePaymentType
	}
	if l.InsuranceCoverage != "" {
		updateMap["c_insurance_coverage"] = l.InsuranceCoverage
	}
	//if !l.UpdatedAt.IsZero() {
	//	updateMap["d_updated_at"] = l.UpdatedAt.UTC()
	//}
	//if l.UpdatedBy != nil {
	//	updateMap["c_updated_by"] = l.UpdatedBy
	//}

	return updateMap
}
