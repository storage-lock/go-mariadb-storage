package mariadb_storage

import (
	"context"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"

	_ "github.com/go-sql-driver/mysql"
)

// MariadbStorage 基于MariaDb作为Storage
type MariadbStorage struct {

	// 其实内部就是跟MySQL的实现是一样一样的
	*mysql_storage.MysqlStorage

	options *MariadbStorageOptions
}

var _ storage.Storage = &MariadbStorage{}

// NewMariadbStorage 创建基于MariaDb的Storage
func NewMariadbStorage(ctx context.Context, options *MariadbStorageOptions) (*MariadbStorage, error) {

	mysqlStorage, err := mysql_storage.NewMysqlStorage(ctx, options.MysqlStorageOptions)
	if err != nil {
		return nil, err
	}

	s := &MariadbStorage{
		options:      options,
		MysqlStorage: mysqlStorage,
	}

	err = s.Init(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

const StorageName = "mariadb-storage"

func (x *MariadbStorage) GetName() string {
	return StorageName
}

// Capabilities 声明 MariaDB 存储支持的能力。
// MariaDB 兼容 MySQL 协议，同样支持 CAS 和可靠时间源。
// 未初始化的零值不委托给内部的 MysqlStorage，否则空指针。
func (x *MariadbStorage) Capabilities() []storage.StorageCapability {
	if x.MysqlStorage == nil {
		return []storage.StorageCapability{storage.CapabilityCAS, storage.CapabilityReliableTime, storage.CapabilityWatch}
	}
	return x.MysqlStorage.Capabilities()
}
