package test

import (
	"fmt"
	"testing"

	"github.com/yogavredizon/todolist/controller"
)

func TestFetchTodos(t *testing.T) {
	res, err := controller.FetchTodos()

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(res)
}
