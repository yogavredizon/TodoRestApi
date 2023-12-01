package test

import (
	"testing"

	"github.com/yogavredizon/todolist/pkg/config"
	"github.com/yogavredizon/todolist/pkg/repository"
)

func TestQuery(t *testing.T) {
	res, err := repository.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	if res[0].Id != 8 {
		t.Log(res[0].Id)
	}
	t.Log(res[1].Id)
}

func TestFindId(t *testing.T) {
	res, err := repository.FindById(9)

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(res)
}
func TestConfigAddTodo(t *testing.T) {
	config.NewDB()

	repository.AddTodo("Makan", "19:00:00")
	_, err := repository.AddTodo("Makan", "19:00:00")

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestRemove(t *testing.T) {
	// config.NewDB()
	err := repository.RemoveTodo(11)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdateTime(t *testing.T) {
	err := repository.UpdateTime(12, "21:00:00")

	if err != nil {
		t.Fatal(err.Error())
	}
}
