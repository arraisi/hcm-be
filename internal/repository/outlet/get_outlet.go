package outlet

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/outlet"
	"github.com/elgris/sqrl"
)

func (r *repository) GetOutlet(ctx context.Context, req outlet.GetOutletRequest) (domain.Outlet, error) {
	var model domain.Outlet
	var result []domain.Outlet

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	err = r.db.SelectContext(ctx, &result, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return model, err
	}

	if len(result) == 0 {
		return model, nil
	}

	return result[0], nil
}
