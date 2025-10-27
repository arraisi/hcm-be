package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/user"

	"github.com/elgris/sqrl"
)

func (r *repository) GetUser(ctx context.Context, req user.GetUserRequest) (domain.User, error) {
	var model domain.User

	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName())
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return model, fmt.Errorf("failed to build query: %w", err)
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.GetContext(ctx, &model, sqlQuery, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model, fmt.Errorf("user with id %v not found", req)
		}
		return model, fmt.Errorf("failed to get user by id: %w", err)
	}

	return model, nil
}
