package user

import (
	"context"
	"fmt"

	"tabeldata.com/hcm-be/internal/domain"
	"tabeldata.com/hcm-be/internal/domain/dto/user"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) CreateUser(ctx context.Context, tx *sqlx.Tx, req user.CreateUserRequest) error {
	model := domain.User{}
	query, args, err := sqrl.Insert(model.TableName()).
		Columns("id", "email", "name", "created_at", "updated_at").
		Values(req.ID, req.Email, req.Name, req.CreatedAt, req.CreatedAt).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	query = r.db.Rebind(query)
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
