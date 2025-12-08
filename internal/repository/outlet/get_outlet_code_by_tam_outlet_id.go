package outlet

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
)

func (r *repository) GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error) {
	if tamOutletCode == "" {
		return nil, fmt.Errorf("tam outlet code cannot be empty")
	}

	// Use raw SQL query with TOP 1 for SQL Server
	sqlQuery := `
		SELECT TOP 1 
			i_id,
			i_idbranch,
			c_outlet,
			n_outlet,
			d_createdate,
			c_idhcmcustomer,
			c_tipe,
			c_tamoutlet
		FROM dbo.tr_outlet
		WHERE c_tamoutlet = ?
	`

	sqlQuery = r.db.Rebind(sqlQuery)

	var outlet domain.Outlet
	if err := r.db.GetContext(ctx, &outlet, sqlQuery, tamOutletCode); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &outlet, nil
}
