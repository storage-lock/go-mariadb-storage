package mariadb_storage

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"sync"
)

// MariaDBConnectionManager 创建一个Maria的连接
type MariaDBConnectionManager struct {
	// 其实底层都是基于mysql的
	once *sync.Once
	*mysql_storage.MySQLConnectionManager
}

var _ storage.ConnectionManager[*sql.DB] = &MariaDBConnectionManager{}

// NewMariaDBConnectionManagerFromDSN 从DSN创建MariaDB连接管理器
func NewMariaDBConnectionManagerFromDSN(dsn string) *MariaDBConnectionManager {
	return &MariaDBConnectionManager{
		once:                   &sync.Once{},
		MySQLConnectionManager: mysql_storage.NewMySQLConnectionManagerFromDSN(dsn),
	}
}

// NewMariaDBConnectionManager 从连接属性创建数据库连接
func NewMariaDBConnectionManager(host string, port uint, user, passwd, database string) *MariaDBConnectionManager {
	return &MariaDBConnectionManager{
		once:                   &sync.Once{},
		MySQLConnectionManager: mysql_storage.NewMySQLConnectionManager(host, port, user, passwd, database),
	}
}

func (x *MariaDBConnectionManager) SetHost(host string) *MariaDBConnectionManager {
	x.Host = host
	return x
}

func (x *MariaDBConnectionManager) SetPort(port uint) *MariaDBConnectionManager {
	x.Port = port
	return x
}

func (x *MariaDBConnectionManager) SetUser(user string) *MariaDBConnectionManager {
	x.User = user
	return x
}

func (x *MariaDBConnectionManager) SetPasswd(passwd string) *MariaDBConnectionManager {
	x.Passwd = passwd
	return x
}

func (x *MariaDBConnectionManager) SetDatabaseName(databaseName string) *MariaDBConnectionManager {
	x.DatabaseName = databaseName
	return x
}

const MariaDBConnectionManagerName = "mariadb-connection-manager"

func (x *MariaDBConnectionManager) Name() string {
	return MariaDBConnectionManagerName
}

// Take 获取到数据库的连接
func (x *MariaDBConnectionManager) Take(ctx context.Context) (*sql.DB, error) {
	return x.MySQLConnectionManager.Take(ctx)
}

func (x *MariaDBConnectionManager) GetDSN() string {
	return x.MySQLConnectionManager.GetDSN()
}

func (x *MariaDBConnectionManager) Return(ctx context.Context, db *sql.DB) error {
	return x.MySQLConnectionManager.Return(ctx, db)
}

func (x *MariaDBConnectionManager) Shutdown(ctx context.Context) error {
	return x.MySQLConnectionManager.Shutdown(ctx)
}
