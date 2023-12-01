package test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yogavredizon/todolist/pkg/config"
)

func TestDBConn(t *testing.T) {
	config.NewDB()

	t.Log("Database Connected")
}
