package mariadb_storage

import (
	"context"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewMariaDbStorage(t *testing.T) {
	envName := "STORAGE_LOCK_MARIA_DSN"
	dsn := os.Getenv(envName)
	assert.NotEmpty(t, dsn)
	s, err := NewMariaDbStorage(context.Background(), &MariaStorageOptions{
		MySQLStorageOptions: &mysql_storage.MySQLStorageOptions{
			ConnectionManager: NewMariaStorageConnectionManagerFromDSN(dsn),
			TableName:         storage_test_helper.TestTableName,
		},
	})
	assert.Nil(t, err)
	storage_test_helper.TestStorage(t, s)
}
