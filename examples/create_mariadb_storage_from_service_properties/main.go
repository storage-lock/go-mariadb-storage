package main

import (
	"context"
	"fmt"
	mariadb_storage "github.com/storage-lock/go-mariadb-storage"
)

func main() {

	// 数据库连接不是DSN的形式，就是一堆零散的属性，则依次设置，可以得到一个连接管理器
	host := "127.0.0.1"
	port := uint(3306)
	username := "root"
	passwd := "UeGqAm8CxYGldMDLoNNt"
	database := "storage_lock_test"
	connectionManager := mariadb_storage.NewMariadbConnectionManager(host, port, username, passwd, database)

	// 然后从这个连接管理器创建MariadbStorage
	options := mariadb_storage.NewMariadbStorageOptions().SetConnectionManager(connectionManager)
	storage, err := mariadb_storage.NewMariadbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
