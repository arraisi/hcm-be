package domain

import (
	"database/sql"
	"fmt"
)

// SalesScoring represents the tm_salesscoring table with outlet information
type SalesScoring struct {
	Periode              string         `json:"periode" db:"periode"`
	BranchName           string         `json:"branch_name" db:"BRANCH_NAME"`
	OutletName           string         `json:"outlet_name" db:"OUTLET_NAME"`
	SPVName              string         `json:"spv_name" db:"spv_name"`
	NIK                  string         `json:"nik" db:"NIK"`
	EmpName              string         `json:"emp_name" db:"EmpName"`
	Position             string         `json:"position" db:"Position"`
	Grading              string         `json:"grading" db:"grading"`
	TipeKendaraan        string         `json:"tipe_kendaraan" db:"Tipe_kendaraan"`
	CustomerGroup        string         `json:"customer_group" db:"CustomerGroup"`
	PerformaNilai        sql.NullString `json:"performa_nilai" db:"Performa_nilai"`
	PerformaHuruf        string         `json:"performa_huruf" db:"Performa_Huruf"`
	IDTowas              string         `json:"id_towas" db:"ID_Towas"`
	BranchCode           string         `json:"branch_code" db:"Branch_code"`
	OutletCode           string         `json:"outlet_code" db:"Outlet_code"`
	TAMOutletCode        sql.NullString `json:"tam_outlet_code" db:"c_tamoutlet"`  // from tr_outlet join
	ActiveTestDriveCount int            `json:"active_test_drive_count,omitempty"` // not from DB, computed field
}

type SalesScorings []SalesScoring

func (s SalesScorings) GetUniqueNIKs() []string {
	nikMap := make(map[string]struct{})
	for _, scoring := range s {
		nikMap[scoring.NIK] = struct{}{}
	}

	uniqueNIKs := make([]string, 0, len(nikMap))
	for nik := range nikMap {
		uniqueNIKs = append(uniqueNIKs, nik)
	}

	return uniqueNIKs
}

// GetPerformanceScore safely extracts the numeric performance score from PerformaNilai
// Returns 0 if the value is NULL or cannot be parsed
func (s *SalesScoring) GetPerformanceScore() float64 {
	if !s.PerformaNilai.Valid {
		return 0
	}

	var score float64
	if _, err := fmt.Sscanf(s.PerformaNilai.String, "%f", &score); err != nil {
		return 0
	}
	return score
}

// TableName returns the database table name for the SalesScoring model
func (s *SalesScoring) TableName() string {
	return "dbo.tm_salesscoring"
}

// Columns returns the full list of database columns for the SalesScoring model
func (s *SalesScoring) Columns() []string {
	return []string{
		"periode",
		"BRANCH_NAME",
		"OUTLET_NAME",
		"spv_name",
		"NIK",
		"EmpName",
		"Position",
		"grading",
		"Tipe_kendaraan",
		"CustomerGroup",
		"Performa_nilai",
		"Performa_Huruf",
		"ID_Towas",
		"Branch_code",
		"Outlet_code",
	}
}

// SelectColumns returns columns with table alias for SELECT queries with JOIN
func (s *SalesScoring) SelectColumns() []string {
	return []string{
		"ss.periode",
		"ss.BRANCH_NAME",
		"ss.OUTLET_NAME",
		"ss.spv_name",
		"ss.NIK",
		"ss.EmpName",
		"ss.Position",
		"ss.grading",
		"ss.Tipe_kendaraan",
		"ss.CustomerGroup",
		"ss.Performa_nilai",
		"ss.Performa_Huruf",
		"ss.ID_Towas",
		"ss.Branch_code",
		"ss.Outlet_code",
		"o.c_tamoutlet",
	}
}
