package mariadb_storage

import (
	"context"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewMariadbStorage(t *testing.T) {
	envName := "STORAGE_LOCK_MARIA_DSN"
	dsn := os.Getenv(envName)
	assert.NotEmpty(t, dsn)
	options := NewMariadbStorageOptions().SetConnectionManager(NewMariadbConnectionManagerFromDsn(dsn))
	s, err := NewMariadbStorage(context.Background(), options)
	assert.Nil(t, err)
	storage_test_helper.TestStorage(t, s)
}
