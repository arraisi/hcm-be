package domain

import (
	"time"

	"github.com/elgris/sqrl"
)

// LeadsInterestedPart represents the interested part/package/merchandise for a lead
type LeadsInterestedPart struct {
	ID                       string
	LeadsID                  string
	PartType                 string
	PackageID                *string
	PartNumber               string
	PartName                 string
	PartQuantity             int
	PartSize                 *string
	PartColor                *string
	PartEstPrice             *float64
	PartInstallationEstPrice *float64
	FlagPartNeedDownPayment  bool
	CreatedAt                time.Time
	UpdatedAt                *time.Time
}

// TableName returns the table name for LeadsInterestedPart
func (LeadsInterestedPart) TableName() string {
	return "tm_leads_interested_part"
}

// Columns returns the column names for LeadsInterestedPart
func (LeadsInterestedPart) Columns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"c_part_type",
		"i_package_id",
		"c_part_number",
		"n_part_name",
		"v_part_quantity",
		"v_part_size",
		"c_part_color",
		"v_part_est_price",
		"v_part_installation_est_price",
		"b_flag_part_need_down_payment",
		"d_created_at",
		"d_updated_at",
	}
}

// SelectColumns returns the column selections for LeadsInterestedPart
func (LeadsInterestedPart) SelectColumns() []string {
	return []string{
		"i_id",
		"i_leads_id",
		"c_part_type",
		"i_package_id",
		"c_part_number",
		"n_part_name",
		"v_part_quantity",
		"v_part_size",
		"c_part_color",
		"v_part_est_price",
		"v_part_installation_est_price",
		"b_flag_part_need_down_payment",
		"d_created_at",
		"d_updated_at",
	}
}

// ToCreateMap converts LeadsInterestedPart to a map for insertion
func (l *LeadsInterestedPart) ToCreateMap() map[string]interface{} {
	return map[string]interface{}{
		"i_id":                          l.ID,
		"i_leads_id":                    l.LeadsID,
		"c_part_type":                   l.PartType,
		"i_package_id":                  l.PackageID,
		"c_part_number":                 l.PartNumber,
		"n_part_name":                   l.PartName,
		"v_part_quantity":               l.PartQuantity,
		"v_part_size":                   l.PartSize,
		"c_part_color":                  l.PartColor,
		"v_part_est_price":              l.PartEstPrice,
		"v_part_installation_est_price": l.PartInstallationEstPrice,
		"b_flag_part_need_down_payment": l.FlagPartNeedDownPayment,
		"d_created_at":                  l.CreatedAt,
		"d_updated_at":                  l.UpdatedAt,
	}
}

// ToUpdateMap converts LeadsInterestedPart to a map for update
func (l *LeadsInterestedPart) ToUpdateMap() map[string]interface{} {
	updateMap := sqrl.Eq{}
	updateMap["i_leads_id"] = l.LeadsID
	updateMap["c_part_type"] = l.PartType
	updateMap["i_package_id"] = l.PackageID
	updateMap["c_part_number"] = l.PartNumber
	updateMap["n_part_name"] = l.PartName
	updateMap["v_part_quantity"] = l.PartQuantity
	updateMap["v_part_size"] = l.PartSize
	updateMap["c_part_color"] = l.PartColor
	updateMap["v_part_est_price"] = l.PartEstPrice
	updateMap["v_part_installation_est_price"] = l.PartInstallationEstPrice
	updateMap["b_flag_part_need_down_payment"] = l.FlagPartNeedDownPayment
	updateMap["d_updated_at"] = l.UpdatedAt
	return updateMap
}
