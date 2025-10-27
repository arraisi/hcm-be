package user

import (
	"context"
	"fmt"

	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/user"

	"github.com/elgris/sqrl"
)

func (r *repository) GetUsers(ctx context.Context, req user.GetUserRequest) ([]domain.User, error) {
	var users []domain.User
	model := domain.User{}

	// Base query
	query := sqrl.Select(model.SelectColumns()...).
		From(model.TableName()).
		OrderBy("id DESC")

	// Add search filter if provided
	req.Apply(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	sqlQuery = r.db.Rebind(sqlQuery)
	err = r.db.SelectContext(ctx, &users, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}

	return users, nil
}
