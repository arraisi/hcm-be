package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

type iDB interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Rebind(query string) string
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type repository struct {
	db iDB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db iDB) *repository {
	return &repository{
		db: db,
	}
}

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

func (r *repository) DeleteUser(ctx context.Context, tx *sqlx.Tx, id string) error {
	model := domain.User{}
	query, args, err := sqrl.Delete(model.TableName()).Where(sqrl.Eq{"id": id}).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
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
