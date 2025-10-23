package user

import (
	"testing"
	"time"

	"github.com/arraisi/hcm-be/internal/domain/dto/user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUserRepository_CreateUser(t *testing.T) {
	req := user.CreateUserRequest{
		ID:        "123e4567-e89b-12d3-a456-426614174000",
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		// The actual SQL after rebinding (what gets executed)
		expectedSQLServerQuery := `INSERT INTO users \(id,email,name,created_at,updated_at\) VALUES \(@p1,@p2,@p3,@p4,@p5\)`

		// given - mock the database operations
		// First, mock the query building (Rebind converts ? to @p1, @p2, etc for SQL Server)
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("INSERT INTO users (id,email,name,created_at,updated_at) VALUES (@p1,@p2,@p3,@p4,@p5)").Times(1)

		// Then mock the actual execution with the SQL Server syntax
		m.mockSqlxDb.ExpectExec(expectedSQLServerQuery).
			WithArgs(req.ID, req.Email, req.Name, req.CreatedAt, req.CreatedAt).
			WillReturnResult(sqlxmock.NewResult(1, 1))

		// when
		err := m.repository.CreateUser(m.Ctx, m.mockDbTx, req)

		// then
		require.NoError(t, err)
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("database_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		expectedSQLServerQuery := `INSERT INTO users \(id,email,name,created_at,updated_at\) VALUES \(@p1,@p2,@p3,@p4,@p5\)`

		// given - mock database error
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("INSERT INTO users (id,email,name,created_at,updated_at) VALUES (@p1,@p2,@p3,@p4,@p5)").Times(1)
		m.mockSqlxDb.ExpectExec(expectedSQLServerQuery).
			WithArgs(req.ID, req.Email, req.Name, req.CreatedAt, req.CreatedAt).
			WillReturnError(sqlxmock.ErrCancelled)

		// when
		err := m.repository.CreateUser(m.Ctx, m.mockDbTx, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to create user")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})
}
