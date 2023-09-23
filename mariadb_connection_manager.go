package mariadb_storage

import (
	"context"
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

// NewMariaDBConnectionManagerFromDsn 从DSN创建MariaDB连接管理器
func NewMariaDBConnectionManagerFromDsn(dsn string) *MariadbConnectionManager {
	return &MariadbConnectionManager{
		MysqlConnectionManager: mysql_storage.NewMysqlConnectionManagerFromDsn(dsn),
	}
}

// NewMariaDBConnectionManager 从连接属性创建数据库连接
func NewMariaDBConnectionManager(host string, port uint, user, passwd, database string) *MariadbConnectionManager {
	return &MariadbConnectionManager{
		MysqlConnectionManager: mysql_storage.NewMysqlConnectionManager(host, port, user, passwd, database),
	}
}

func (x *MariadbConnectionManager) SetHost(host string) *MariadbConnectionManager {
	x.Host = host
	return x
}

func (x *MariadbConnectionManager) SetPort(port uint) *MariadbConnectionManager {
	x.Port = port
	return x
}

func (x *MariadbConnectionManager) SetUser(user string) *MariadbConnectionManager {
	x.User = user
	return x
}

func (x *MariadbConnectionManager) SetPasswd(passwd string) *MariadbConnectionManager {
	x.Passwd = passwd
	return x
}

func (x *MariadbConnectionManager) SetDatabaseName(databaseName string) *MariadbConnectionManager {
	x.DatabaseName = databaseName
	return x
}

const MariaDBConnectionManagerName = "mariadb-connection-manager"

func (x *MariadbConnectionManager) Name() string {
	return MariaDBConnectionManagerName
}

// Take 获取到数据库的连接
func (x *MariadbConnectionManager) Take(ctx context.Context) (*sql.DB, error) {
	return x.MysqlConnectionManager.Take(ctx)
}

func (x *MariadbConnectionManager) GetDSN() string {
	return x.MysqlConnectionManager.GetDSN()
}

func (x *MariadbConnectionManager) Return(ctx context.Context, db *sql.DB) error {
	return x.MysqlConnectionManager.Return(ctx, db)
}

func (x *MariadbConnectionManager) Shutdown(ctx context.Context) error {
	return x.MysqlConnectionManager.Shutdown(ctx)
}
