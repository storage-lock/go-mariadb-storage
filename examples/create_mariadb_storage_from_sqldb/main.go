package main

import (
	"context"
	"database/sql"
	"fmt"
	mariadb_storage "github.com/storage-lock/go-mariadb-storage"
	storage "github.com/storage-lock/go-storage"
)

func main() {

	// 假设已经在其它地方初始化数据库连接得到了一个*sql.DB
	testDsn := "root:UeGqAm8CxYGldMDLoNNt@tcp(127.0.0.1:3306)/storage_lock_test"
	db, err := sql.Open("mysql", testDsn)
	if err != nil {
		panic(err)
	}

	// 则可以从这个*sql.DB中创建一个MariadbStorage
	connectionManager := storage.NewFixedSqlDBConnectionManager(db)
	options := mariadb_storage.NewMariadbStorageOptions().SetConnectionManager(connectionManager)
	storage, err := mariadb_storage.NewMariadbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
