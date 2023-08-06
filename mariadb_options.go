package mariadb_storage

import mysql_storage "github.com/storage-lock/go-mysql-storage"

// MariaStorageOptions 创建基于Maria为Storage的选项
type MariaStorageOptions struct {
	*mysql_storage.MySQLStorageOptions
}

func NewMariaStorageOptions() *MariaStorageOptions {
	return &MariaStorageOptions{
		MySQLStorageOptions: mysql_storage.NewMySQLStorageOptions(),
	}
}
