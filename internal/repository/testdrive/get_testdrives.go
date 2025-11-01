package testdrive

import (
	"context"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/testdrive"
	"github.com/elgris/sqrl"
)

func (r *repository) GetTestDrives(ctx context.Context, req testdrive.GetTestDriveRequest) ([]domain.TestDrive, error) {
	var result []domain.TestDrive
	model := domain.TestDrive{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("d_updated_at DESC")

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &result, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
