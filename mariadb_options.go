package mariadb_storage

import (
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
)

// MariadbStorageOptions 创建基于Maria为Storage的选项
type MariadbStorageOptions struct {
	*mysql_storage.MysqlStorageOptions
}

func NewMariadbStorageOptions() *MariadbStorageOptions {
	return &MariadbStorageOptions{
		MysqlStorageOptions: mysql_storage.NewMysqlStorageOptions(),
	}
}

func (x *MariadbStorageOptions) SetConnectionManager(connManager storage.ConnectionManager[*sql.DB]) *MariadbStorageOptions {
	x.ConnectionManager = connManager
	return x
}

func (x *MariadbStorageOptions) SetTableName(tableName string) *MariadbStorageOptions {
	x.TableName = tableName
	return x
}
func (x *MariadbStorageOptions) Check() error {
	return x.MysqlStorageOptions.Check()
}
