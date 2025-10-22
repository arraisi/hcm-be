package repository

import (
	"database/sql"
	"hcm-be/internal/domain"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"hcm-be/internal/domain/dto/user"
	"github.com/google/uuid"
)

type repository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *sql.DB) UserRepository {
	return &repository{
		db:      db,
	}
}

func (r *repository) GetUsers(ctx context.Context, req user.GetUsersRequest) ([]domain.User, error) {
	// ctx, cancel := context.WithTimeout(ctx, r.timeout)
	// defer cancel()

	// Build the query with optional filters
	query := `SELECT id, email, name, created_at FROM dbo.users WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	// Add search filter if provided
	if req.Search != "" {
		query += fmt.Sprintf(" AND (email LIKE @p%d OR name LIKE @p%d)", argIndex, argIndex+1)
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm)
		argIndex += 2
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

	// Add pagination
	if req.Limit > 0 {
		query += fmt.Sprintf(" OFFSET @p%d ROWS FETCH NEXT @p%d ROWS ONLY", argIndex, argIndex+1)
		args = append(args, req.Offset, req.Limit)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		var id sql.NullString
		if err := rows.Scan(&id, &user.Email, &user.Name, &user.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		user.ID = id.String
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}

func (r *repository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	// ctx, cancel := context.WithTimeout(ctx, r.timeout)
	// defer cancel()

	const query = `SELECT id, email, name, created_at FROM dbo.users WHERE id = @p1`

	var user domain.User
	var uid sql.NullString
	err := r.db.QueryRowContext(ctx, query, id).Scan(&uid, &user.Email, &user.Name, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	user.ID = uid.String
	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, req user.CreateUserRequest) error {
	// ctx, cancel := context.WithTimeout(ctx, r.timeout)
	// defer cancel()

	id := uuid.NewString()
	createdAt := time.Now().UTC()

	const query = `
		INSERT INTO dbo.users (id, email, name, created_at)
		VALUES (@p1, @p2, @p3, @p4)
	`

	_, err := r.db.ExecContext(ctx, query, id, req.Email, req.Name, createdAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *repository) UpdateUser(ctx context.Context, id string, req user.UpdateUserRequest) error {
	// ctx, cancel := context.WithTimeout(ctx, r.timeout)
	// defer cancel()

	// Build dynamic update query
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if req.Email != nil {
		setParts = append(setParts, fmt.Sprintf("email = @p%d", argIndex))
		args = append(args, *req.Email)
		argIndex++
	}

	if req.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = @p%d", argIndex))
		args = append(args, *req.Name)
		argIndex++
	}

	if len(setParts) == 0 {
		return errors.New("no fields to update")
	}

	// Add updated_at
	setParts = append(setParts, fmt.Sprintf("updated_at = @p%d", argIndex))
	args = append(args, time.Now().UTC())
	argIndex++

	// Add WHERE clause
	query := fmt.Sprintf("UPDATE dbo.users SET %s WHERE id = @p%d", strings.Join(setParts, ", "), argIndex)
	args = append(args, id)

	result, err := r.db.ExecContext(ctx, query, args...)
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

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	// ctx, cancel := context.WithTimeout(ctx, r.timeout)
	// defer cancel()

	const query = `DELETE FROM dbo.users WHERE id = @p1`

	result, err := r.db.ExecContext(ctx, query, id)
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
