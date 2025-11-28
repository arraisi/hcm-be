package domain

import (
	"time"

	"github.com/lib/pq"
)

// SalesOrderInsurancePolicy represents individual insurance policy coverage
type SalesOrderInsurancePolicy struct {
	ID                 string         `json:"id" db:"i_id"`
	SalesOrderID       string         `json:"sales_order_id" db:"i_sales_order_id"`
	InsuranceType      string         `json:"insurance_type" db:"c_insurance_type"`
	InsuranceStartDate time.Time      `json:"insurance_start_date" db:"d_insurance_start_date"`
	InsuranceEndDate   time.Time      `json:"insurance_end_date" db:"d_insurance_end_date"`
	CreatedAt          time.Time      `json:"created_at" db:"d_created_at"`
	UpdatedAt          time.Time      `json:"updated_at" db:"d_updated_at"`
	InsuranceCoverage  pq.StringArray `json:"insurance_coverage" db:"e_insurance_coverage"`
}

// TableName returns the database table name
func (s *SalesOrderInsurancePolicy) TableName() string {
	return "dbo.tr_sales_order_insurance_policy"
}

// Columns returns the list of database columns
func (s *SalesOrderInsurancePolicy) Columns() []string {
	return []string{
		"i_sales_order_id",
		"c_insurance_type",
		"d_insurance_start_date",
		"d_insurance_end_date",
		"e_insurance_coverage",
		"d_created_at",
		"d_updated_at",
	}
}

// SelectColumns returns the list of columns to select in queries
func (s *SalesOrderInsurancePolicy) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"i_sales_order_id",
		"c_insurance_type",
		"e_insurance_coverage",
		"d_insurance_start_date",
		"d_insurance_end_date",
		"d_created_at",
		"d_updated_at",
	}
}

func (s *SalesOrderInsurancePolicy) ToCreateMap() (columns []string, values []interface{}) {
	columns = s.Columns()
	values = []interface{}{
		s.SalesOrderID,
		s.InsuranceType,
		s.InsuranceCoverage,
		s.InsuranceStartDate,
		s.InsuranceEndDate,
		s.CreatedAt,
		s.UpdatedAt,
	}
	return
}
