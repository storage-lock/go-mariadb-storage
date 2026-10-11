package mariadb_storage

import (
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
)

// MariadbConnectionManager 创建一个Maria的连接
type MariadbConnectionManager struct {
	// 其实底层都是基于mysql的
	*mysql_storage.MysqlConnectionManager
}

var _ storage.ConnectionManager[*sql.DB] = &MariadbConnectionManager{}

// NewMariadbConnectionManager 从连接属性创建数据库连接
func NewMariadbConnectionManager(host string, port uint, user, passwd, database string) *MariadbConnectionManager {
	return &MariadbConnectionManager{
		MysqlConnectionManager: mysql_storage.NewMysqlConnectionManager(host, port, user, passwd, database),
	}
}

// NewMariadbConnectionManagerFromDsn 从DSN创建MariaDB连接管理器
func NewMariadbConnectionManagerFromDsn(dsn string) *MariadbConnectionManager {
	return &MariadbConnectionManager{
		MysqlConnectionManager: mysql_storage.NewMysqlConnectionManagerFromDsn(dsn),
	}
}

// NewMariadbConnectionManagerFromSqlDb 从*sql.DB创建MariaDB连接管理器
func NewMariadbConnectionManagerFromSqlDb(db *sql.DB) *MariadbConnectionManager {
	return &MariadbConnectionManager{
		MysqlConnectionManager: mysql_storage.NewMysqlConnectionManagerFromSqlDb(db),
	}
}

func (x *MariadbConnectionManager) SetHost(host string) *MariadbConnectionManager {
	x.Host = host
	return x
}

