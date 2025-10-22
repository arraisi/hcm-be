package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"
	"strings"
	"time"

	"github.com/elgris/sqrl"
	"github.com/google/uuid"
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

func (r *repository) GetUsers(ctx context.Context, req user.GetUsersRequest) ([]domain.User, error) {
	var users []domain.User
	var args []interface{}
	argIndex := 0

	// Base query
	query := `SELECT CAST(id AS NVARCHAR(36)) as id, email, name, created_at 
			  FROM dbo.users`

	// Add search filter if provided
	if req.Search != "" {
		searchTerm := "%" + req.Search + "%"
		argIndex++
		query += fmt.Sprintf(" WHERE (email LIKE @p%d OR name LIKE @p%d)", argIndex, argIndex)
		args = append(args, searchTerm)
	}

	// Add sorting
	sortBy := "created_at"
	if req.SortBy != "" {
		// Validate sort field to prevent SQL injection
		validSortFields := map[string]bool{
			"email":      true,
			"name":       true,
			"created_at": true,
		}
		if validSortFields[req.SortBy] {
			sortBy = req.SortBy
		}
	}

	order := "DESC"
	if req.Order != "" && strings.ToUpper(req.Order) == "ASC" {
		order = "ASC"
	}

	query += fmt.Sprintf(" ORDER BY %s %s", sortBy, order)

	// Add SQL Server pagination
	if req.Limit > 0 {
		argIndex++
		offsetParam := argIndex
		argIndex++
		limitParam := argIndex
		query += fmt.Sprintf(" OFFSET @p%d ROWS FETCH NEXT @p%d ROWS ONLY", offsetParam, limitParam)
		args = append(args, req.Offset, req.Limit)
	}

	err := r.db.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}

	return users, nil
}

func (r *repository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	// const query = `SELECT CAST(id AS NVARCHAR(36)) as id, email, name, created_at FROM dbo.users WHERE id = @p1`

	var user domain.User
	query, args, err := sqrl.
		Select(user.TableName()).
		Columns(user.Columns()...).
		Where(sqrl.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, tx *sqlx.Tx, req user.CreateUserRequest) error {
	id := uuid.NewString()
	createdAt := time.Now().UTC()

	model := domain.User{}
	query, args, err := sqrl.Insert(model.TableName()).Columns("id", "email", "name", "created_at").Values(id, req.Email, req.Name, createdAt).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *repository) UpdateUser(ctx context.Context, tx *sqlx.Tx, id string, req user.UpdateUserRequest) error {
	model := domain.User{}
	query, args, err := sqrl.Update(model.TableName()).
		SetMap(map[string]interface{}{
			"email":      req.Email,
			"name":       req.Name,
			"updated_at": time.Now().UTC(),
		}).Where(sqrl.Eq{"id": id}).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

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
