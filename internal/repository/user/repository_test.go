package user

import (
	"context"
	"testing"

	"tabeldata.com/hcm-be/internal/config"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type mock struct {
	Config *config.Config
	Ctrl   *gomock.Controller
	Ctx    context.Context

	mockIDB    *MockiDB
	mockSqlxDb sqlxmock.Sqlmock
	mockDbTx   *sqlx.Tx
	repository *repository
	anyQuery   string
}

func setupMock(t *testing.T) mock {
	m := mock{}
	m.Config = &config.Config{}
	m.Ctrl = gomock.NewController(t)
	m.Ctx = context.Background()

	db, mock, err := sqlxmock.Newx()
	if err != nil {
		panic("Error was not expected while opening a stub safesql connection: " + err.Error())
	}

	m.mockIDB = NewMockiDB(m.Ctrl)
	m.mockSqlxDb = mock
	m.anyQuery = `(.|\s)*\S(.|\s)*`

	// Set up expectations for a transaction
	mock.ExpectBegin()
	tx, _ := db.Beginx()
	m.mockDbTx = tx

	m.repository = NewUserRepository(m.mockIDB)

	return m
}
