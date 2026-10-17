package mariadb_storage

import (
	"context"
	"os"
	"testing"

	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
)

// TestMariadbRealSemantics 连接真实 MariaDB 介质跑接口级语义套件。
// 需要环境变量 STORAGE_LOCK_MARIADB_DSN，指向真实 MariaDB 服务。
func TestMariadbRealSemantics(t *testing.T) {
	dsn := os.Getenv("STORAGE_LOCK_MARIADB_DSN")
	if dsn == "" {
		t.Skip("STORAGE_LOCK_MARIADB_DSN not set, skipping real-mariadb semantics test")
	}

	ctx := context.Background()
	cm := NewMariadbConnectionManagerFromDsn(dsn)
	s, err := NewMariadbStorage(ctx, NewMariadbStorageOptions().SetConnectionManager(cm))
	if err != nil {
		t.Fatalf("NewMariadbStorage: %v", err)
	}
	storage_test_helper.TestStorage(t, s)
}