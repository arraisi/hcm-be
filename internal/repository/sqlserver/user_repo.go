package sqlserver

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"hcm-be/internal/domain"
)

type UserRepo struct {
	db *sql.DB
	// optional: timeout per query
	timeout time.Duration
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db, timeout: 5 * time.Second}
}

func (r *UserRepo) FindAll() ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	const q = `
SELECT id, email, name, created_at
FROM dbo.users
ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.User
	for rows.Next() {
		var u domain.User
		var id sql.NullString
		if err := rows.Scan(&id, &u.Email, &u.Name, &u.CreatedAt); err != nil {
			return nil, err
		}
		u.ID = id.String
		out = append(out, u)
	}
	return out, rows.Err()
}

func (r *UserRepo) FindByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	const q = `
SELECT id, email, name, created_at
FROM dbo.users
WHERE id = @p1`

	var u domain.User
	var uid sql.NullString
	err := r.db.QueryRowContext(ctx, q, id).Scan(&uid, &u.Email, &u.Name, &u.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}
	u.ID = uid.String
	return &u, nil
}

func (r *UserRepo) Create(u domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
	}

	// Gunakan OUTPUT untuk mengembalikan kolom jika perlu
	const q = `
INSERT INTO dbo.users (id, email, name, created_at)
VALUES (@p1, @p2, @p3, @p4)`

	_, err := r.db.ExecContext(ctx, q, u.ID, u.Email, u.Name, u.CreatedAt)
	return err
}
