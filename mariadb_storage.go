package mariadb_storage

import (
	"context"
	"github.com/golang-infrastructure/go-iterator"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"time"

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

func (x *MariadbStorage) Init(ctx context.Context) error {
	return x.MysqlStorage.Init(ctx)
}

func (x *MariadbStorage) UpdateWithVersion(ctx context.Context, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	return x.MysqlStorage.UpdateWithVersion(ctx, lockId, exceptedVersion, newVersion, lockInformation)
}

func (x *MariadbStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	return x.MysqlStorage.CreateWithVersion(ctx, lockId, version, lockInformation)
}

func (x *MariadbStorage) DeleteWithVersion(ctx context.Context, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) error {
	return x.MysqlStorage.DeleteWithVersion(ctx, lockId, exceptedVersion, lockInformation)
}

func (x *MariadbStorage) Get(ctx context.Context, lockId string) (string, error) {
	return x.MysqlStorage.Get(ctx, lockId)
}

func (x *MariadbStorage) GetTime(ctx context.Context) (time.Time, error) {
	return x.MysqlStorage.GetTime(ctx)
}

func (x *MariadbStorage) Close(ctx context.Context) error {
	return x.MysqlStorage.Close(ctx)
}

func (x *MariadbStorage) List(ctx context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	return x.MysqlStorage.List(ctx)
}
