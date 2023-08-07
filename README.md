# Mariadb Storage

# 一、这是什么
以MariaDB为存储引擎的[Storage](https://github.com/storage-lock/go-storage)实现，当前仓库为比较底层的存储层实现，你可以与[storage-lock](https://github.com/storage-lock/go-storage-lock)结合使用，或者这个项目[mariadb-locks](https://github.com/storage-lock/go-mariadb-locks)里专门封装提供了一些MariaDB锁相关的更易用友好的API。

# 二、安装依赖
```bash
go get -u github.com/storage-lock/go-mariadb-storage
```

# 三、API Examples

## 3.1 从DSN创建MariadbStorage

在Golang的世界中连接数据库最常见的就是DSN，下面的例子演示了如何从一个DSN创建MariadbStorage：

```go
package main

import (
	"context"
	"fmt"
	mariadb_storage "github.com/storage-lock/go-mariadb-storage"
)

func main() {

	// 使用一个DSN形式的数据库连接字符串创建ConnectionManager
	testDsn := "root:UeGqAm8CxYGldMDLoNNt@tcp(127.0.0.1:3306)/storage_lock_test"
	connectionManager := mariadb_storage.NewMariaDBConnectionManagerFromDSN(testDsn)

	// 然后从这个ConnectionManager创建MariadbStorage
	options := mariadb_storage.NewMariaStorageOptions().SetConnectionManager(connectionManager)
	storage, err := mariadb_storage.NewMariaDbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
```

## 3.2 从连接属性（ip、端口、用户名、密码等等）中创建MariadbStorage

或者你的配置文件中存放的并不是DSN，而是零散的几个连接属性，下面是一个创建MariadbStorage的例子：

```go
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
	connectionManager := mariadb_storage.NewMariaDBConnectionManager(host, port, username, passwd, database)

	// 然后从这个连接管理器创建MariadbStorage
	options := mariadb_storage.NewMariaStorageOptions().SetConnectionManager(connectionManager)
	storage, err := mariadb_storage.NewMariaDbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
```

## 3.3 从*sql.DB创建TidbStorage

或者现在你已经有从其它渠道创建的能够连接到MySQL的\*sql.DB，则也可以从这个*sql.DB创建MariadbStorage：

```go
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
	options := mariadb_storage.NewMariaStorageOptions().SetConnectionManager(connectionManager)
	storage, err := mariadb_storage.NewMariaDbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
```







