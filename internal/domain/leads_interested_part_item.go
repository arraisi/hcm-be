package domain

import (
	"time"

	"github.com/elgris/sqrl"
)

// LeadsInterestedPartItem represents the individual items within a package
type LeadsInterestedPartItem struct {
	ID                    string
	LeadsID               string
	LeadsInterestedPartID string
	PartNumber            string
	PartName              string
	CreatedAt             time.Time
}

// TableName returns the table name for LeadsInterestedPartItem
func (LeadsInterestedPartItem) TableName() string {
	return "tm_leads_interested_part_item"
}

// Columns returns the column names for LeadsInterestedPartItem
func (LeadsInterestedPartItem) Columns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"i_leads_interested_part_id",
		"c_part_number",
		"n_part_name",
		"d_created_at",
	}
}

// SelectColumns returns the column selections for LeadsInterestedPartItem
func (LeadsInterestedPartItem) SelectColumns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"i_leads_interested_part_id",
		"c_part_number",
		"n_part_name",
		"d_created_at",
	}
}

// ToCreateMap converts LeadsInterestedPartItem to a map for insertion
func (l *LeadsInterestedPartItem) ToCreateMap() map[string]interface{} {
	return map[string]interface{}{
		"i_id":                       l.ID,
		"i_leads_id":                 l.LeadsID,
		"i_leads_interested_part_id": l.LeadsInterestedPartID,
		"c_part_number":              l.PartNumber,
		"n_part_name":                l.PartName,
		"d_created_at":               l.CreatedAt,
	}
}

// ToUpdateMap converts LeadsInterestedPartItem to a map for update
func (l *LeadsInterestedPartItem) ToUpdateMap() map[string]interface{} {
	updateMap := sqrl.Eq{}
	updateMap["i_leads_id"] = l.LeadsID
	updateMap["i_leads_interested_part_id"] = l.LeadsInterestedPartID
	updateMap["c_part_number"] = l.PartNumber
	updateMap["n_part_name"] = l.PartName
	return updateMap
}
