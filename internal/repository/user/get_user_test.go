package user

import (
	"database/sql"
	"hcm-be/internal/domain"
	"hcm-be/internal/domain/dto/user"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUserRepository_GetUser(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"
	expectedUser := domain.User{
		ID:        userID,
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{
			ID: userID,
		}

		// given - mock successful retrieval through the database interface
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users WHERE id = @p1").Times(1)
		m.mockIDB.EXPECT().GetContext(gomock.Any(), gomock.Any(), gomock.Any(), userID).Return(nil).Do(
			func(ctx interface{}, dest interface{}, query string, args ...interface{}) {
				// Simulate filling the destination with expected user data
				user := dest.(*domain.User)
				*user = expectedUser
			},
		)

		// when
		result, err := m.repository.GetUser(m.Ctx, req)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedUser.ID, result.ID)
		require.Equal(t, expectedUser.Email, result.Email)
		require.Equal(t, expectedUser.Name, result.Name)
	})

	t.Run("user_not_found", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{
			ID: userID,
		}

		// given - mock no rows found
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users WHERE id = @p1").Times(1)
		m.mockIDB.EXPECT().GetContext(gomock.Any(), gomock.Any(), gomock.Any(), userID).Return(sql.ErrNoRows)

		// when
		result, err := m.repository.GetUser(m.Ctx, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "not found")
		require.Empty(t, result.ID)
	})

	t.Run("database_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{
			ID: userID,
		}

		// given - mock database error
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users WHERE id = @p1").Times(1)
		m.mockIDB.EXPECT().GetContext(gomock.Any(), gomock.Any(), gomock.Any(), userID).Return(sqlxmock.ErrCancelled)

		// when
		result, err := m.repository.GetUser(m.Ctx, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get user by id")
		require.Empty(t, result.ID)
	})
}
