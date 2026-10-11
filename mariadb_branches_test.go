package mariadb_storage

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
)

type takeErrorManager struct{}

func (takeErrorManager) Name() string                           { return "take-error" }
func (takeErrorManager) Take(context.Context) (*sql.DB, error)  { return nil, errors.New("take failed") }
func (takeErrorManager) Return(context.Context, *sql.DB) error  { return nil }
func (takeErrorManager) Shutdown(context.Context) error         { return nil }

func TestMariadbConnectionManagerBranches(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	fromConfig := NewMariadbConnectionManager("localhost", 3306, "user", "pass", "db")
	assert.Equal(t, "db.local", fromConfig.SetHost("db.local").Host)

	fromDB := NewMariadbConnectionManagerFromSqlDb(db)
	got, err := fromDB.Take(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, db, got)
}

func TestNewMariadbStorage_MysqlError(t *testing.T) {
	opts := NewMariadbStorageOptions()
	_, err := NewMariadbStorage(context.Background(), opts)
	assert.ErrorIs(t, err, mysql_storage.ErrMysqlStorageOptionsConnectionManagerNil)
}

func TestNewMariadbStorage_InitError(t *testing.T) {
	opts := NewMariadbStorageOptions().SetConnectionManager(takeErrorManager{})
	_, err := NewMariadbStorage(context.Background(), opts)
	assert.EqualError(t, err, "take failed")
}

func TestNewMariadbStorage_Success(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectExec(".*").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(".*").WillReturnError(errors.New("second init failed"))
	opts := NewMariadbStorageOptions().SetConnectionManager(storage.NewFixedSqlDBConnectionManager(db))
	_, err = NewMariadbStorage(context.Background(), opts)
	assert.EqualError(t, err, "second init failed")
}
