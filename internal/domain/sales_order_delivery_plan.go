package domain

import (
	"time"
)

// SalesOrderDeliveryPlan represents delivery amendment/plan information
type SalesOrderDeliveryPlan struct {
	ID                        string    `json:"id" db:"i_id"`
	SalesOrderID              string    `json:"sales_order_id" db:"i_sales_order_id"`
	AmendmentCreatedDatetime  time.Time `json:"amendment_created_datetime" db:"d_amendment_created_datetime"`
	AmendmentStatus           string    `json:"amendment_status" db:"c_amendment_status"`
	AmendmentReason           string    `json:"amendment_reason" db:"c_amendment_reason"`
	AmendmentReasonOthers     *string   `json:"amendment_reason_others" db:"c_amendment_reason_others"`
	AmendmentSource           string    `json:"amendment_source" db:"c_amendment_source"`
	FlagBuyerIsRecipient      bool      `json:"flag_buyer_is_recipient" db:"b_flag_buyer_is_recipient"`
	ReceivedPlanDatetimeStart time.Time `json:"received_plan_datetime_start" db:"d_received_plan_datetime_start"`
	ReceivedPlanDatetimeEnd   time.Time `json:"received_plan_datetime_end" db:"d_received_plan_datetime_end"`
	RecipientName             *string   `json:"recipient_name" db:"n_recipient_name"`
	RecipientPhoneNumber      *string   `json:"recipient_phone_number" db:"c_recipient_phone_number"`
	RecipientRelation         *string   `json:"recipient_relation" db:"c_recipient_relation"`
	RecipientRelationOthers   *string   `json:"recipient_relation_others" db:"c_recipient_relation_others"`
	DeliveryAddressLocation   *string   `json:"delivery_address_location" db:"c_delivery_address_location"`
	DeliveryAddressLabel      string    `json:"delivery_address_label" db:"c_delivery_address_label"`
	DeliveryAddress           string    `json:"delivery_address" db:"e_delivery_address"`
	DeliveryProvince          string    `json:"delivery_province" db:"n_delivery_province"`
	DeliveryCity              string    `json:"delivery_city" db:"n_delivery_city"`
	DeliveryDistrict          string    `json:"delivery_district" db:"n_delivery_district"`
	DeliverySubdistrict       string    `json:"delivery_subdistrict" db:"n_delivery_subdistrict"`
	DeliveryPostalCode        string    `json:"delivery_postal_code" db:"c_delivery_postal_code"`
	CreatedAt                 time.Time `json:"created_at" db:"d_created_at"`
	UpdatedAt                 time.Time `json:"updated_at" db:"d_updated_at"`
}

// TableName returns the database table name
func (s *SalesOrderDeliveryPlan) TableName() string {
	return "dbo.tr_sales_order_delivery_plan"
}

// Columns returns the list of database columns
func (s *SalesOrderDeliveryPlan) Columns() []string {
	return []string{
		"i_sales_order_id",
		"d_amendment_created_datetime",
		"c_amendment_status",
		"c_amendment_reason",
		"c_amendment_reason_others",
		"c_amendment_source",
		"b_flag_buyer_is_recipient",
		"d_received_plan_datetime_start",
		"d_received_plan_datetime_end",
		"n_recipient_name",
		"c_recipient_phone_number",
		"c_recipient_relation",
		"c_recipient_relation_others",
		"c_delivery_address_location",
		"c_delivery_address_label",
		"e_delivery_address",
		"n_delivery_province",
		"n_delivery_city",
		"n_delivery_district",
		"n_delivery_subdistrict",
		"c_delivery_postal_code",
		"d_created_at",
		"d_updated_at",
	}
}

// SelectColumns returns the list of columns to select in queries
func (s *SalesOrderDeliveryPlan) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_sales_order_id AS NVARCHAR(36)) as i_sales_order_id",
		"d_amendment_created_datetime",
		"c_amendment_status",
		"c_amendment_reason",
		"c_amendment_reason_others",
		"c_amendment_source",
		"b_flag_buyer_is_recipient",
		"d_received_plan_datetime_start",
		"d_received_plan_datetime_end",
		"n_recipient_name",
		"c_recipient_phone_number",
		"c_recipient_relation",
		"c_recipient_relation_others",
		"c_delivery_address_location",
		"c_delivery_address_label",
		"e_delivery_address",
		"n_delivery_province",
		"n_delivery_city",
		"n_delivery_district",
		"n_delivery_subdistrict",
		"c_delivery_postal_code",
		"d_created_at",
		"d_updated_at",
	}
}

func (s *SalesOrderDeliveryPlan) ToCreateMap() (columns []string, values []interface{}) {
	columns = s.Columns()
	values = []interface{}{
		s.ID,
		s.SalesOrderID,
		s.AmendmentCreatedDatetime,
		s.AmendmentStatus,
		s.AmendmentReason,
		s.AmendmentReasonOthers,
		s.AmendmentSource,
		s.FlagBuyerIsRecipient,
		s.ReceivedPlanDatetimeStart,
		s.ReceivedPlanDatetimeEnd,
		s.RecipientName,
		s.RecipientPhoneNumber,
		s.RecipientRelation,
		s.RecipientRelationOthers,
		s.DeliveryAddressLocation,
		s.DeliveryAddressLabel,
		s.DeliveryAddress,
		s.DeliveryProvince,
		s.DeliveryCity,
		s.DeliveryDistrict,
		s.DeliverySubdistrict,
		s.DeliveryPostalCode,
		s.CreatedAt,
		s.UpdatedAt,
	}
	return
}
