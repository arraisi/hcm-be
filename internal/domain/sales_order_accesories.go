package domain

import (
	"time"
)

// SalesOrderAccessory represents accessories attached to a sales order
type SalesOrderAccessory struct {
	ID                              string    `json:"id" db:"i_id"`
	AccessoriesType                 string    `json:"accessories_type" db:"c_accessories_type"`
	AccessoriesNumber               string    `json:"accessories_number" db:"c_accessories_number"`
	AccessoriesName                 string    `json:"accessories_name" db:"n_accessories_name"`
	AccessoriesQty                  string    `json:"accessories_qty" db:"v_accessories_qty"`
	AccessoriesOrderSource          string    `json:"accessories_order_source" db:"c_accessories_order_source"`
	AccessoriesAvailabilityStatus   string    `json:"accessories_availability_status" db:"c_accessories_availability_status"`
	AccessoriesItemStatus           string    `json:"accessories_item_status" db:"c_accessories_item_status"`
	AccessoriesSize                 string    `json:"accessories_size" db:"v_accessories_size"`
	AccessoriesColor                string    `json:"accessories_color" db:"c_accessories_color"`
	AccessoriesEstPrice             string    `json:"accessories_est_price" db:"v_accessories_est_price"`
	FlagAccessoriesNeedDownPayment  string    `json:"flag_accessories_need_down_payment" db:"c_flag_accessories_need_down_payment"`
	AccessoriesInstallationEstPrice string    `json:"accessories_installation_est_price" db:"v_accessories_installation_est_price"`
	CreatedAt                       time.Time `json:"created_at" db:"d_created_at"`
	UpdatedAt                       time.Time `json:"updated_at" db:"d_updated_at"`
	SalesOrderID                    string    `json:"sales_order_id" db:"sales_order_id"`
}

// TableName returns the database table name
func (s *SalesOrderAccessory) TableName() string {
	return "dbo.tr_sales_order_accessories"
}

// Columns returns the list of database columns
func (s *SalesOrderAccessory) Columns() []string {
	return []string{
		"c_accessories_type",
		"c_accessories_number",
		"n_accessories_name",
		"v_accessories_qty",
		"c_accessories_order_source",
		"c_accessories_availability_status",
		"c_accessories_item_status",
		"v_accessories_size",
		"c_accessories_color",
		"v_accessories_est_price",
		"c_flag_accessories_need_down_payment",
		"v_accessories_installation_est_price",
		"d_created_at",
		"d_updated_at",
		"i_sales_order_id",
	}
}

// SelectColumns returns the list of columns to select in queries
func (s *SalesOrderAccessory) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"c_accessories_type",
		"c_accessories_number",
		"n_accessories_name",
		"v_accessories_qty",
		"c_accessories_order_source",
		"c_accessories_availability_status",
		"c_accessories_item_status",
		"v_accessories_size",
		"c_accessories_color",
		"v_accessories_est_price",
		"c_flag_accessories_need_down_payment",
		"v_accessories_installation_est_price",
		"d_created_at",
		"d_updated_at",
		"i_sales_order_id",
	}
}

func (s *SalesOrderAccessory) ToCreateMap() (columns []string, values []interface{}) {
	columns = s.Columns()
	values = []interface{}{
		s.AccessoriesType,
		s.AccessoriesNumber,
		s.AccessoriesName,
		s.AccessoriesQty,
		s.AccessoriesOrderSource,
		s.AccessoriesAvailabilityStatus,
		s.AccessoriesItemStatus,
		s.AccessoriesSize,
		s.AccessoriesColor,
		s.AccessoriesEstPrice,
		s.FlagAccessoriesNeedDownPayment,
		s.AccessoriesInstallationEstPrice,
		s.CreatedAt,
		s.UpdatedAt,
		s.SalesOrderID,
	}
	return
}
