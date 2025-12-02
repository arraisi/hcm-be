package outlet

import (
	"context"
	"fmt"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

func (r *repository) GetOutletCodeByTAMOutletID(ctx context.Context, tamOutletCode string) (*domain.Outlet, error) {
	if tamOutletCode == "" {
		return nil, fmt.Errorf("tam outlet code cannot be empty")
	}

	model := domain.Outlet{}

	query := sqrl.
		Select(model.SelectColumns()...).
		From(model.TableName()).
		Where(sqrl.Eq{"c_tamoutlet": tamOutletCode}).
		Limit(1)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)

	var outlet domain.Outlet
	if err := r.db.GetContext(ctx, &outlet, sqlQuery, args...); err != nil {
		return nil, err
	}

	return &outlet, nil
}
