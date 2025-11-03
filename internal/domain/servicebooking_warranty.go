package domain

import (
	"time"
)

type ServiceBookingWarranty struct {
	ID               string    `db:"i_id"`
	ServiceBookingID string    `db:"i_service_booking_id"`
	WarrantyName     string    `db:"c_warranty_name"`
	WarrantyStatus   string    `db:"c_warranty_status"`
	CreatedAt        time.Time `db:"d_created_at"`
	CreatedBy        string    `db:"c_created_by"`
	UpdatedAt        time.Time `db:"d_updated_at"`
	UpdatedBy        string    `db:"c_updated_by"`
}

type ServiceBookingWarranties []ServiceBookingWarranty

// TableName returns the database table name for the ServiceBookingWarranty model
func (sbw *ServiceBookingWarranty) TableName() string {
	return "dbo.tr_service_booking_warranty"
}

// Columns returns the list of database columns for the ServiceBookingWarranty model
func (sbw *ServiceBookingWarranty) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"c_warranty_name",
		"c_warranty_status",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the ServiceBookingWarranty model
func (sbw *ServiceBookingWarranty) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_service_booking_id AS NVARCHAR(36)) as i_service_booking_id",
		"c_warranty_name",
		"c_warranty_status",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

func (sbw *ServiceBookingWarranty) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0, len(sbw.Columns()))
	values := make([]interface{}, 0, len(sbw.Columns()))

	if sbw.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, sbw.ServiceBookingID)
	}
	if sbw.WarrantyName != "" {
		columns = append(columns, "c_warranty_name")
		values = append(values, sbw.WarrantyName)
	}
	if sbw.WarrantyStatus != "" {
		columns = append(columns, "c_warranty_status")
		values = append(values, sbw.WarrantyStatus)
	}
	columns = append(columns, "c_created_by")
	values = append(values, sbw.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, sbw.CreatedBy)

	return columns, values
}

func (sbw *ServiceBookingWarranty) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if sbw.ServiceBookingID != "" {
		updateMap["i_service_booking_id"] = sbw.ServiceBookingID
	}
	if sbw.WarrantyName != "" {
		updateMap["c_warranty_name"] = sbw.WarrantyName
	}
	if sbw.WarrantyStatus != "" {
		updateMap["c_warranty_status"] = sbw.WarrantyStatus
	}
	updateMap["d_updated_at"] = time.Now()
	updateMap["c_updated_by"] = sbw.UpdatedBy

	return updateMap
}
