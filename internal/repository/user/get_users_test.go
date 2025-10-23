package user

import (
	"testing"
	"time"

	"github.com/arraisi/hcm-be/internal/domain"
	"github.com/arraisi/hcm-be/internal/domain/dto/user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUserRepository_GetUsers(t *testing.T) {
	expectedUsers := []domain.User{
		{
			ID:        "123e4567-e89b-12d3-a456-426614174000",
			Email:     "test1@example.com",
			Name:      "Test User 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "456e7890-e89b-12d3-a456-426614174000",
			Email:     "test2@example.com",
			Name:      "Test User 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("success_no_filter", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{} // No filters

		// given - mock successful retrieval
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users ORDER BY id DESC").Times(1)
		m.mockIDB.EXPECT().SelectContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Do(
			func(ctx interface{}, dest interface{}, query string, args ...interface{}) {
				// Simulate filling the destination with expected users data
				users := dest.(*[]domain.User)
				*users = expectedUsers
			},
		)

		// when
		result, err := m.repository.GetUsers(m.Ctx, req)

		// then
		require.NoError(t, err)
		require.Len(t, result, 2)
		require.Equal(t, expectedUsers[0].ID, result[0].ID)
		require.Equal(t, expectedUsers[1].ID, result[1].ID)
	})

	t.Run("success_with_search_filter", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{
			Search: "test",
		}

		// given - mock successful retrieval with search filter
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users WHERE (name ILIKE @p1 OR email ILIKE @p2) ORDER BY id DESC").Times(1)
		m.mockIDB.EXPECT().SelectContext(gomock.Any(), gomock.Any(), gomock.Any(), "%test%", "%test%").Return(nil).Do(
			func(ctx interface{}, dest interface{}, query string, args ...interface{}) {
				// Simulate filling the destination with filtered users data
				users := dest.(*[]domain.User)
				*users = []domain.User{expectedUsers[0]} // Only one user matches
			},
		)

		// when
		result, err := m.repository.GetUsers(m.Ctx, req)

		// then
		require.NoError(t, err)
		require.Len(t, result, 1)
		require.Equal(t, expectedUsers[0].ID, result[0].ID)
	})

	t.Run("success_with_pagination", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{
			Limit:  10,
			Offset: 0,
		}

		// given - mock successful retrieval with pagination
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users ORDER BY id DESC OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY").Times(1)
		m.mockIDB.EXPECT().SelectContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Do(
			func(ctx interface{}, dest interface{}, query string, args ...interface{}) {
				users := dest.(*[]domain.User)
				*users = expectedUsers
			},
		)

		// when
		result, err := m.repository.GetUsers(m.Ctx, req)

		// then
		require.NoError(t, err)
		require.Len(t, result, 2)
	})

	t.Run("empty_result", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{}

		// given - mock empty result
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users ORDER BY id DESC").Times(1)
		m.mockIDB.EXPECT().SelectContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Do(
			func(ctx interface{}, dest interface{}, query string, args ...interface{}) {
				// Simulate empty result
				users := dest.(*[]domain.User)
				*users = []domain.User{}
			},
		)

		// when
		result, err := m.repository.GetUsers(m.Ctx, req)

		// then
		require.NoError(t, err)
		require.Len(t, result, 0)
	})

	t.Run("database_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.GetUserRequest{}

		// given - mock database error
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("SELECT CAST(id AS NVARCHAR(36)) as id,email,name,created_at,updated_at FROM dbo.users ORDER BY id DESC").Times(1)
		m.mockIDB.EXPECT().SelectContext(gomock.Any(), gomock.Any(), gomock.Any()).Return(sqlxmock.ErrCancelled)

		// when
		result, err := m.repository.GetUsers(m.Ctx, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to query users")
		require.Nil(t, result)
	})
}
