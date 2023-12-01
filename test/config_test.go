package test

import (
	"testing"

	"github.com/yogavredizon/todolist/config"
)

func TestDBConn(t *testing.T) {
	config.NewDB()

	t.Log("Database Connected")
}
