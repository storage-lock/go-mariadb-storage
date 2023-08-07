package mariadb_storage

import (
	"context"
	"github.com/golang-infrastructure/go-iterator"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// MariaDbStorage 基于MariaDb作为Storage
type MariaDbStorage struct {

	// 其实内部就是跟MySQL的实现是一样一样的
	*mysql_storage.MySQLStorage

	options *MariaStorageOptions
}

var _ storage.Storage = &MariaDbStorage{}

// NewMariaDbStorage 创建基于MariaDb的Storage
func NewMariaDbStorage(ctx context.Context, options *MariaStorageOptions) (*MariaDbStorage, error) {

	mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options.MySQLStorageOptions)
	if err != nil {
		return nil, err
	}

	s := &MariaDbStorage{
		options:      options,
		MySQLStorage: mysqlStorage,
	}

	err = s.Init(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

const StorageName = "mariadb-storage"

func (x *MariaDbStorage) GetName() string {
	return StorageName
}

func (x *MariaDbStorage) Init(ctx context.Context) error {
	return x.MySQLStorage.Init(ctx)
}

func (x *MariaDbStorage) UpdateWithVersion(ctx context.Context, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	return x.MySQLStorage.UpdateWithVersion(ctx, lockId, exceptedVersion, newVersion, lockInformation)
}

func (x *MariaDbStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	return x.MySQLStorage.CreateWithVersion(ctx, lockId, version, lockInformation)
}

func (x *MariaDbStorage) DeleteWithVersion(ctx context.Context, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) error {
	return x.MySQLStorage.DeleteWithVersion(ctx, lockId, exceptedVersion, lockInformation)
}

func (x *MariaDbStorage) Get(ctx context.Context, lockId string) (string, error) {
	return x.MySQLStorage.Get(ctx, lockId)
}

func (x *MariaDbStorage) GetTime(ctx context.Context) (time.Time, error) {
	return x.MySQLStorage.GetTime(ctx)
}

func (x *MariaDbStorage) Close(ctx context.Context) error {
	return x.MySQLStorage.Close(ctx)
}

func (x *MariaDbStorage) List(ctx context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	return x.MySQLStorage.List(ctx)
}
