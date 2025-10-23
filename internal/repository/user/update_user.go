package user

import (
	"context"
	"fmt"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

func (r *repository) UpdateUser(ctx context.Context, tx *sqlx.Tx, id string, req user.UpdateUserRequest) error {
	model := domain.User{}
	query, args, err := sqrl.Update(model.TableName()).
		SetMap(req.MapToUpdateBuilder()).
		Where(sqrl.Eq{"id": id}).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	query = r.db.Rebind(query)
	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	return nil
}
