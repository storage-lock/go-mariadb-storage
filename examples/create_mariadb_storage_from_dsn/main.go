package main

import (
	"context"
	"fmt"
	mariadb_storage "github.com/storage-lock/go-mariadb-storage"
)

func main() {

	// 使用一个DSN形式的数据库连接字符串创建ConnectionManager
	testDsn := "root:UeGqAm8CxYGldMDLoNNt@tcp(127.0.0.1:3306)/storage_lock_test"
	connectionManager := mariadb_storage.NewMariadbConnectionManagerFromDsn(testDsn)

	// 然后从这个ConnectionManager创建MariadbStorage
	options := mariadb_storage.NewMariadbStorageOptions().SetConnectionManager(connectionManager)
	storage, err := mariadb_storage.NewMariadbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
