package mariadb_storage

import (
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
)

// MariaDBConnectionManager 创建一个Maria的连接
type MariaDBConnectionManager struct {
	// 其实底层都是基于mysql的
	*mysql_storage.MySQLConnectionManager
}

var _ storage.ConnectionManager[*sql.DB] = &MariaDBConnectionManager{}

// NewMariaStorageConnectionManagerFromDSN 从DSN创建Maria连接
func NewMariaStorageConnectionManagerFromDSN(dsn string) storage.ConnectionManager[*sql.DB] {
	return mysql_storage.NewMySQLConnectionManagerFromDSN(dsn)
}

// NewMariaStorageConnectionManager 从服务器属性创建数据库连接
func NewMariaStorageConnectionManager(host string, port uint, user, passwd, database string) *MariaDBConnectionManager {
	return &MariaDBConnectionManager{
		MySQLConnectionManager: mysql_storage.NewMySQLConnectionProvider(host, port, user, passwd, database),
	}
}
