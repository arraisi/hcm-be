package domain

import "time"

// ServiceBookingImage represents damage image URLs for service booking
type ServiceBookingImage struct {
	ID               string    `db:"i_id"`
	ServiceBookingID string    `db:"i_service_booking_id"`
	ImageURL         string    `db:"n_image_url"`
	CreatedAt        time.Time `db:"d_created_at"`
	CreatedBy        string    `db:"c_created_by"`
	UpdatedAt        time.Time `db:"d_updated_at"`
	UpdatedBy        string    `db:"c_updated_by"`
}

// TableName returns the database table name
func (di *ServiceBookingImage) TableName() string {
	return "dbo.tr_service_booking_image"
}

// Columns returns the list of database columns
func (di *ServiceBookingImage) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"n_image_url",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select
func (di *ServiceBookingImage) SelectColumns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"n_image_url",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting
func (di *ServiceBookingImage) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	if di.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, di.ServiceBookingID)
	}
	if di.ImageURL != "" {
		columns = append(columns, "n_image_url")
		values = append(values, di.ImageURL)
	}
	columns = append(columns, "c_created_by")
	values = append(values, di.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, di.CreatedBy)

	return columns, values
}
