package usedcar

import (
	"context"
	"github.com/arraisi/hcm-be/internal/domain/dto/usedcar"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/elgris/sqrl"
)

func (r *repository) GetUsedCar(ctx context.Context, req usedcar.GetUsedCarRequest) (domain.UsedCar, error) {
	var model domain.UsedCar

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.GetContext(ctx, &model, sqlQuery, args...)
	if err != nil {
		return model, err
	}

	return model, nil
}
