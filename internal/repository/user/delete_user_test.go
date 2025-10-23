package user

import (
	"testing"

	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"

	t.Run("success", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		// given - mock successful deletion (note the dbo schema prefix)
		m.mockSqlxDb.ExpectExec(`DELETE FROM dbo\.users WHERE id = \?`).
			WithArgs(userID).
			WillReturnResult(sqlxmock.NewResult(0, 1)) // 1 row affected

		// when
		err := m.repository.DeleteUser(m.Ctx, m.mockDbTx, userID)

		// then
		require.NoError(t, err)
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("user_not_found", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		// given - mock no rows affected (user not found)
		m.mockSqlxDb.ExpectExec(`DELETE FROM dbo\.users WHERE id = \?`).
			WithArgs(userID).
			WillReturnResult(sqlxmock.NewResult(0, 0)) // 0 rows affected

		// when
		err := m.repository.DeleteUser(m.Ctx, m.mockDbTx, userID)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "user with id "+userID+" not found")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("database_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		// given - mock database error
		m.mockSqlxDb.ExpectExec(`DELETE FROM dbo\.users WHERE id = \?`).
			WithArgs(userID).
			WillReturnError(sqlxmock.ErrCancelled)

		// when
		err := m.repository.DeleteUser(m.Ctx, m.mockDbTx, userID)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to delete user")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})

	t.Run("rows_affected_error", func(t *testing.T) {
		m := setupMock(t)
		defer m.Ctrl.Finish()

		// given - mock RowsAffected error
		result := sqlxmock.NewErrorResult(sqlxmock.ErrCancelled) // Result that returns error on RowsAffected
		m.mockSqlxDb.ExpectExec(`DELETE FROM dbo\.users WHERE id = \?`).
			WithArgs(userID).
			WillReturnResult(result)

		// when
		err := m.repository.DeleteUser(m.Ctx, m.mockDbTx, userID)

		// then
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get rows affected")
		require.NoError(t, m.mockSqlxDb.ExpectationsWereMet())
	})
}
