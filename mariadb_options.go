package mariadb_storage

import (
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
)

// MariaStorageOptions 创建基于Maria为Storage的选项
type MariaStorageOptions struct {
	*mysql_storage.MySQLStorageOptions
}

func NewMariaStorageOptions() *MariaStorageOptions {
	return &MariaStorageOptions{
		MySQLStorageOptions: mysql_storage.NewMySQLStorageOptions(),
	}
}

func (x *MariaStorageOptions) SetConnectionManager(connManager storage.ConnectionManager[*sql.DB]) *MariaStorageOptions {
	x.ConnectionManager = connManager
	return x
}

func (x *MariaStorageOptions) SetTableName(tableName string) *MariaStorageOptions {
	x.TableName = tableName
	return x
}
