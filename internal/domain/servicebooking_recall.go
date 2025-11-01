package domain

import (
	"time"
)

type ServiceBookingRecall struct {
	ID                string    `db:"i_id"`
	ServiceBookingID  string    `db:"i_service_booking_id"`
	RecallID          string    `db:"i_recall_id"`
	RecallDate        string    `db:"d_recall_date"`
	RecallDescription string    `db:"e_recall_description"`
	AffectedPart      string    `db:"c_affected_part"`
	CreatedAt         time.Time `db:"d_created_at"`
	CreatedBy         string    `db:"c_created_by"`
	UpdatedAt         time.Time `db:"d_updated_at"`
	UpdatedBy         string    `db:"c_updated_by"`
}

// TableName returns the database table name for the ServiceBookingRecall model
func (sbr *ServiceBookingRecall) TableName() string {
	return "dbo.tr_service_booking_recall"
}

// Columns returns the list of database columns for the ServiceBookingRecall model
func (sbr *ServiceBookingRecall) Columns() []string {
	return []string{
		"i_id",
		"i_service_booking_id",
		"i_recall_id",
		"d_recall_date",
		"e_recall_description",
		"c_affected_part",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// SelectColumns returns the list of columns to select in queries for the ServiceBookingRecall model
func (sbr *ServiceBookingRecall) SelectColumns() []string {
	return []string{
		"CAST(i_id AS NVARCHAR(36)) as i_id",
		"CAST(i_service_booking_id AS NVARCHAR(36)) as i_service_booking_id",
		"CAST(i_recall_id AS NVARCHAR(36)) as i_recall_id",
		"d_recall_date",
		"e_recall_description",
		"c_affected_part",
		"d_created_at",
		"c_created_by",
		"d_updated_at",
		"c_updated_by",
	}
}

// ToCreateMap prepares the columns and values for inserting a new ServiceBookingRecall record
func (sbr *ServiceBookingRecall) ToCreateMap() ([]string, []interface{}) {
	columns := make([]string, 0, len(sbr.Columns()))
	values := make([]interface{}, 0, len(sbr.Columns()))

	if sbr.ServiceBookingID != "" {
		columns = append(columns, "i_service_booking_id")
		values = append(values, sbr.ServiceBookingID)
	}
	if sbr.RecallID != "" {
		columns = append(columns, "i_recall_id")
		values = append(values, sbr.RecallID)
	}
	if sbr.RecallDate != "" {
		columns = append(columns, "d_recall_date")
		values = append(values, sbr.RecallDate)
	}
	if sbr.RecallDescription != "" {
		columns = append(columns, "e_recall_description")
		values = append(values, sbr.RecallDescription)
	}
	if sbr.AffectedPart != "" {
		columns = append(columns, "c_affected_part")
		values = append(values, sbr.AffectedPart)
	}
	columns = append(columns, "c_created_by")
	values = append(values, sbr.CreatedBy)
	columns = append(columns, "c_updated_by")
	values = append(values, sbr.CreatedBy)

	return columns, values
}

func (sbr *ServiceBookingRecall) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})

	if sbr.ServiceBookingID != "" {
		updateMap["i_service_booking_id"] = sbr.ServiceBookingID
	}
	if sbr.RecallID != "" {
		updateMap["i_recall_id"] = sbr.RecallID
	}
	if sbr.RecallDate != "" {
		updateMap["d_recall_date"] = sbr.RecallDate
	}
	if sbr.RecallDescription != "" {
		updateMap["e_recall_description"] = sbr.RecallDescription
	}
	if sbr.AffectedPart != "" {
		updateMap["c_affected_part"] = sbr.AffectedPart
	}
	updateMap["d_updated_at"] = time.Now()
	updateMap["c_updated_by"] = sbr.UpdatedBy

	return updateMap
}
