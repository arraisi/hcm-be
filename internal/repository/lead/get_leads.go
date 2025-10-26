package lead

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/lead"
	"github.com/elgris/sqrl"
)

func (r *repository) GetLeads(ctx context.Context, req lead.GetLeadRequest) (domain.Lead, error) {
	var model domain.Lead

	query := sqrl.Select(model.SelectColumns()...).
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
