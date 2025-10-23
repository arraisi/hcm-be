package user

import (
	"hcm-be/internal/domain/dto/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"
	name := "Updated User"
	email := "updated@example.com"

	t.Run("success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{
			Name:  &name,
			Email: &email,
		}

		// given - mock successful update (note: sqrl field order may vary)
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("UPDATE dbo.users SET email = @p1, name = @p2 WHERE id = @p3").Times(1)
		m.mockSqlxDb.ExpectExec(`UPDATE dbo\.users SET email = @p1, name = @p2 WHERE id = @p3`).
			WithArgs(email, name, userID).
			WillReturnResult(sqlxmock.NewResult(0, 1)) // 1 row affected

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then
		require.NoError(t, err)
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("update_name_only", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{
			Name: &name,
		}

		// given - mock successful update with only name
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("UPDATE dbo.users SET name = @p1 WHERE id = @p2").Times(1)
		m.mockSqlxDb.ExpectExec(`UPDATE dbo\.users SET name = @p1 WHERE id = @p2`).
			WithArgs(name, userID).
			WillReturnResult(sqlxmock.NewResult(0, 1)) // 1 row affected

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then
		require.NoError(t, err)
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("update_email_only", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{
			Email: &email,
		}

		// given - mock successful update with only email
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("UPDATE dbo.users SET email = @p1 WHERE id = @p2").Times(1)
		m.mockSqlxDb.ExpectExec(`UPDATE dbo\.users SET email = @p1 WHERE id = @p2`).
			WithArgs(email, userID).
			WillReturnResult(sqlxmock.NewResult(0, 1)) // 1 row affected

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then
		require.NoError(t, err)
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("user_not_found", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{
			Name:  &name,
			Email: &email,
		}

		// given - mock no rows affected (user not found)
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("UPDATE dbo.users SET email = @p1, name = @p2 WHERE id = @p3").Times(1)
		m.mockSqlxDb.ExpectExec(`UPDATE dbo\.users SET email = @p1, name = @p2 WHERE id = @p3`).
			WithArgs(email, name, userID).
			WillReturnResult(sqlxmock.NewResult(0, 0)) // 0 rows affected

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "user with id "+userID+" not found")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("database_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{
			Name:  &name,
			Email: &email,
		}

		// given - mock database error
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("UPDATE dbo.users SET email = @p1, name = @p2 WHERE id = @p3").Times(1)
		m.mockSqlxDb.ExpectExec(`UPDATE dbo\.users SET email = @p1, name = @p2 WHERE id = @p3`).
			WithArgs(email, name, userID).
			WillReturnError(sqlxmock.ErrCancelled)

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to update user")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("rows_affected_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{
			Name:  &name,
			Email: &email,
		}

		// given - mock RowsAffected error
		result := sqlxmock.NewErrorResult(sqlxmock.ErrCancelled)
		m.mockIDB.EXPECT().Rebind(gomock.Any()).Return("UPDATE dbo.users SET email = @p1, name = @p2 WHERE id = @p3").Times(1)
		m.mockSqlxDb.ExpectExec(`UPDATE dbo\.users SET email = @p1, name = @p2 WHERE id = @p3`).
			WithArgs(email, name, userID).
			WillReturnResult(result)

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get rows affected")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("empty_update_request", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		req := user.UpdateUserRequest{} // No fields to update

		// when
		err := m.repository.UpdateUser(m.Ctx, m.mockDbTx, userID, req)

		// then - this should return an error since sqrl requires at least one Set clause
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to build query")
	})
}
