package customer

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/customer"
	"github.com/elgris/sqrl"
)

func (r repository) GetDetailPenjualanToyota(ctx context.Context, req customer.GetDetailPenjualanToyotaRequest) ([]domain.ViewDetailPenjualanToyota, error) {
	var model domain.ViewDetailPenjualanToyota
	var result []domain.ViewDetailPenjualanToyota

	query := sqrl.Select(model.Columns()...).
		From(model.TableName()).OrderBy("Hasjrat_ID desc")

	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return result, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}
