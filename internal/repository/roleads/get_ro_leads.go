package roleads

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/roleads"
	"github.com/elgris/sqrl"
)

func (r *repository) GetRoLeads(ctx context.Context, req roleads.GetRoLeadsRequest) (domain.RoLeads, error) {
	var model domain.RoLeads

	query := sqrl.Select(model.Columns()...).
		From(model.TableName())

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	err = r.db.GetContext(ctx, &model, r.db.Rebind(sqlQuery), args...)
	if err != nil {
		return model, err
	}

	return model, nil
}
