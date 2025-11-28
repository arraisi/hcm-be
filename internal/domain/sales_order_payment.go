package domain

import (
	"time"
)

// SalesOrderPayment represents payment information for a sales order
type SalesOrderPayment struct {
	ID              string    `json:"id" db:"i_id"`
	PaymentID       string    `json:"payment_id" db:"i_payment_id"`
	PaymentNumber   string    `json:"payment_number" db:"c_payment_number"`
	PaymentDatetime time.Time `json:"payment_datetime" db:"d_payment_datetime"`
	NameOnPayment   string    `json:"name_on_payment" db:"n_name_on_payment"`
	PaymentStatus   string    `json:"payment_status" db:"c_payment_status"`
	FundSource      string    `json:"fund_source" db:"c_fund_source"`
	PaymentChannel  string    `json:"payment_channel" db:"c_payment_channel"`
	PaymentStage    string    `json:"payment_stage" db:"v_payment_stage"`
	PaymentType     string    `json:"payment_type" db:"c_payment_type"`
	CreatedAt       time.Time `json:"created_at" db:"d_created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"d_updated_at"`
	SalesOrderID    string    `json:"sales_order_id" db:"i_sales_order_id"`
}

// TableName returns the database table name
func (s *SalesOrderPayment) TableName() string {
	return "dbo.tr_sales_order_payment"
}

// Columns returns the list of database columns
func (s *SalesOrderPayment) Columns() []string {
	return []string{
		"i_payment_id",
		"c_payment_number",
		"d_payment_datetime",
		"n_name_on_payment",
		"c_payment_status",
		"c_fund_source",
		"c_payment_channel",
		"v_payment_stage",
		"c_payment_type",
		"d_created_at",
		"d_updated_at",
		"i_sales_order_id",
	}
}

// SelectColumns returns the list of columns to select in queries
func (s *SalesOrderPayment) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"i_payment_id",
		"c_payment_number",
		"d_payment_datetime",
		"n_name_on_payment",
		"c_payment_status",
		"c_fund_source",
		"c_payment_channel",
		"v_payment_stage",
		"c_payment_type",
		"d_created_at",
		"d_updated_at",
		"i_sales_order_id",
	}
}

func (s *SalesOrderPayment) ToCreateMap() (columns []string, values []interface{}) {
	columns = s.Columns()
	values = []interface{}{
		s.PaymentID,
		s.PaymentNumber,
		s.PaymentDatetime,
		s.NameOnPayment,
		s.PaymentStatus,
		s.FundSource,
		s.PaymentChannel,
		s.PaymentStage,
		s.PaymentType,
		s.CreatedAt,
		s.UpdatedAt,
		s.SalesOrderID,
	}
	return
}
