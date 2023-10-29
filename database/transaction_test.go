package database_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/tyeryan/l-common-util/database"
	logutil "github.com/tyeryan/l-protocol/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

type transactionTestSuite struct {
	suite.Suite

	sqlMock sqlmock.Sqlmock
	db      *gorm.DB

	tx database.Transactional
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(transactionTestSuite))
}

func (suite *transactionTestSuite) SetupSuite() {
	logutil.UseConsoleLogger()
	logutil.EnableDebug()
}

func (suite *transactionTestSuite) SetupTest() {
	var mockDB *sql.DB
	mockDB, suite.sqlMock, _ = sqlmock.New()
	suite.db, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	// if you want to connect to local postgres
	//suite.db, _ = gorm.Open("postgres", "host=localhost dbname=test sslmode=disable")

	suite.tx = database.ProvideTransactional(&mockConnector{db: suite.db})
	suite.db.Logger.LogMode(logger.Info)
}

func (suite *transactionTestSuite) TestWhenNoErrorShouldCommit() {
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectCommit()

	log := logutil.GetLogger("TestWhenNoErrorShouldCommit")

	log.Debugw(context.Background(), "TestWhenNoErrorShouldCommit")

	err := suite.tx.WithTransaction(context.Background(), func(tx *gorm.DB) error {
		return nil
	})

	suite.Nil(err)
}

func (suite *transactionTestSuite) TestWhenErrorShouldRollback() {
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectRollback()

	log := logutil.GetLogger("TestWhenErrorShouldRollback")

	log.Debugw(context.Background(), "TestWhenErrorShouldRollback")

	err := suite.tx.WithTransaction(context.Background(), func(tx *gorm.DB) error {
		return errors.New("oh")
	})

	suite.NotNil(err)
}

func (suite *transactionTestSuite) TestWhenPanicShouldRollback() {
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectRollback()

	log := logutil.GetLogger("TestWhenPanicShouldRollback")

	log.Debugw(context.Background(), "TestWhenPanicShouldRollback")

	wasPanic := false
	defer func() {
		p := recover()
		wasPanic = p != nil
	}()

	err := suite.tx.WithTransaction(context.Background(), func(tx *gorm.DB) error {
		panic("ohohoh")
	})

	suite.NotNil(err)
	suite.True(wasPanic)
}

type mockConnector struct {
	db *gorm.DB
}

func (m *mockConnector) GetDB() *gorm.DB {
	return m.db
}

func (m *mockConnector) GetDBWithContext(ctx context.Context) *gorm.DB {
	return m.db
}
