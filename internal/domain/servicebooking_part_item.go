package domain

import (
	"time"
)

type ServiceBookingPartItem struct {
	ID                   string    `db:"i_id"`
	ServiceBookingID     string    `db:"i_service_booking_id"`
	ServiceBookingPartID string    `db:"i_service_booking_part_id"`
	PartNumber           string    `db:"c_part_number"`
	PartName             string    `db:"n_part_name"`
	CreatedAt            time.Time `db:"d_created_at"`
	CreatedBy            string    `db:"c_created_by"`
	UpdatedAt            time.Time `db:"d_updated_at"`
	UpdatedBy            string    `db:"c_updated_by"`
	Deleted              bool      `db:"b_deleted"`
}

// TableName returns the database table name for the ServiceBookingPartItem model
func (pi *ServiceBookingPartItem) TableName() string {
	return "dbo.tr_service_booking_part_item"
}

// Columns returns the list of database columns for the ServiceBookingPartItem model
func (pi *ServiceBookingPartItem) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"i_service_booking_part_id",
		"c_part_number",
		"n_part_name",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the ServiceBookingPartItem model
func (pi *ServiceBookingPartItem) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_service_booking_id AS NVARCHAR(36)) as i_service_booking_id",
		"CAST(i_service_booking_part_id AS NVARCHAR(36)) as i_service_booking_part_id",
		"c_part_number",
		"n_part_name",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting a new ServiceBookingPartItem record
func (pi *ServiceBookingPartItem) ToCreateMap(serviceBookingID, packageID string) ([]string, []interface{}) {
	columns := make([]string, 0, len(pi.Columns()))
	values := make([]interface{}, 0, len(pi.Columns()))

	if serviceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, serviceBookingID)
	}

	if packageID != "" {
		columns = append(columns, "i_service_booking_part_id")
		values = append(values, packageID)
	}

	if pi.PartNumber != "" {
		columns = append(columns, "c_part_number")
		values = append(values, pi.PartNumber)
	}

	if pi.PartName != "" {
		columns = append(columns, "n_part_name")
		values = append(values, pi.PartName)
	}
	columns = append(columns, "c_created_by")
	values = append(values, pi.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, pi.CreatedBy)

	return columns, values
}

// ToUpdateMap prepares the columns and values for updating an existing ServiceBookingPartItem record
func (pi *ServiceBookingPartItem) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if pi.PartNumber != "" {
		updateMap["c_part_number"] = pi.PartNumber
	}

	if pi.PartName != "" {
		updateMap["n_part_name"] = pi.PartName
	}

	updateMap["d_updated_at"] = time.Now()
	updateMap["c_updated_by"] = pi.UpdatedBy

	return updateMap
}
