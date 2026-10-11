package mariadb_storage

import (
	"context"
	"database/sql"
	"testing"

	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
)

func TestNewMariadbStorageOptions(t *testing.T) {
	opts := NewMariadbStorageOptions()
	assert.NotNil(t, opts)
	assert.NotNil(t, opts.MysqlStorageOptions)
	assert.Equal(t, storage.DefaultStorageTableName, opts.MysqlStorageOptions.TableName)
	assert.Nil(t, opts.MysqlStorageOptions.ConnectionManager)
}

func TestMariadbStorageOptions_SetConnectionManager(t *testing.T) {
	opts := NewMariadbStorageOptions()
	assert.NotNil(t, opts)
	db := &sql.DB{}
	connManager := storage.NewFixedSqlDBConnectionManager(db)
	opts.SetConnectionManager(connManager)
	assert.NotNil(t, opts.ConnectionManager)
}

func TestMariadbStorageOptions_SetTableName(t *testing.T) {
	opts := NewMariadbStorageOptions()
	assert.NotNil(t, opts)
	opts.SetTableName("custom_table")
	assert.Equal(t, "custom_table", opts.TableName)
}

func TestMariadbStorageOptions_Check_NilConnectionManager(t *testing.T) {
	opts := NewMariadbStorageOptions()
	assert.NotNil(t, opts)
	err := opts.Check()
	assert.Error(t, err)
	assert.Equal(t, mysql_storage.ErrMysqlStorageOptionsConnectionManagerNil, err)
}

func TestMariadbStorageOptions_Check_ValidOptions(t *testing.T) {
	opts := NewMariadbStorageOptions()
	assert.NotNil(t, opts)
	db := &sql.DB{}
	connManager := storage.NewFixedSqlDBConnectionManager(db)
	opts.SetConnectionManager(connManager)
	err := opts.Check()
	assert.NoError(t, err)
}

func TestNewMariadbStorage_NilOptions(t *testing.T) {
	assert.Panics(t, func() {
		NewMariadbStorage(context.Background(), nil)
	})
}

func TestMariadbStorage_GetName(t *testing.T) {
	s := &MariadbStorage{}
	name := s.GetName()
	assert.Equal(t, StorageName, name)
}

func TestMariadbStorage_Capabilities(t *testing.T) {
	s := &MariadbStorage{}
	caps := s.Capabilities()
	assert.NotNil(t, caps)
	assert.Len(t, caps, 3)
	assert.Equal(t, storage.CapabilityCAS, caps[0])
	assert.Equal(t, storage.CapabilityReliableTime, caps[1])
	assert.Equal(t, storage.CapabilityWatch, caps[2])
}
