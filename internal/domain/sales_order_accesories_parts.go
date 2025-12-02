package domain

import (
	"time"
)

// SalesOrderAccessoriesPart represents accessories attached to a sales order
type SalesOrderAccessoriesPart struct {
	ID                string    `json:"id" db:"i_id"`
	AccessoriesID     string    `json:"accessories_id" db:"i_accessories_id"`
	AccessoriesNumber string    `json:"accessories_number" db:"c_accessories_number"`
	AccessoriesName   string    `json:"accessories_name" db:"n_accessories_name"`
	CreatedAt         time.Time `json:"created_at" db:"d_created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"d_updated_at"`
}

// TableName returns the database table name
func (s *SalesOrderAccessoriesPart) TableName() string {
	return "dbo.tr_sales_order_accessories_parts"
}

// Columns returns the list of database columns
func (s *SalesOrderAccessoriesPart) Columns() []string {
	return []string{
		"i_accessories_id",
		"c_accessories_number",
		"n_accessories_name",
		"d_created_at",
		"d_updated_at",
	}
}

// SelectColumns returns the list of columns to select in queries
func (s *SalesOrderAccessoriesPart) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"i_accessories_id",
		"c_accessories_number",
		"n_accessories_name",
		"d_created_at",
		"d_updated_at",
	}
}

func (s *SalesOrderAccessoriesPart) ToCreateMap() (columns []string, values []interface{}) {
	columns = s.Columns()
	values = []interface{}{
		s.AccessoriesID,
		s.AccessoriesNumber,
		s.AccessoriesName,
		s.CreatedAt,
		s.UpdatedAt,
	}
	return
}
