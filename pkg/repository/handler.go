package repository

import (
	"errors"
	"fmt"

	"github.com/yogavredizon/todolist/pkg/helper"
	"github.com/yogavredizon/todolist/pkg/model"
)

func AddTodo(todo, timeTodo string) error {
	if todo == "" {
		return errors.New("todo can't blank")
	}
	getHour := helper.GetHour(timeTodo)

	for _, t := range model.Todo {
		if todo == t.Todo || timeTodo == getHour.GoString() {
			return errors.New("todo or Time already attempt")
		}
	}

	var newTodo model.Task

	// newTodo.Id++
	newTodo.Todo = todo
	newTodo.Time = getHour

	model.Todo = append(model.Todo, newTodo)

	return nil
}

func ShowTodo() {
	for i, todo := range model.Todo {
		fmt.Printf("%d. %s : %d:%d \n", i+1, todo.Todo, todo.Time.Hour(), todo.Time.Minute())
	}
}

func RemoveTodo(id int) (int, error) {

	if id < 1 {
		return id, errors.New("id not Found")
	}

	for i := 0; i < len(model.Todo); i++ {
		if id == i {
			model.Todo = append(model.Todo[:id], model.Todo[id+1:]...)
		}
	}

	return id + 1, nil
}

func UpdateTime(id int, timeTodo string) (int, error) {
	id -= 1
	if id > len(model.Todo)-1 {
		return -1, errors.New("id not found")
	}
	model.Todo[id].Time = helper.GetHour(timeTodo)

	return id + 1, nil
}

func UpdateTodo(id int, todo string) (int, error) {
	id -= 1
	if id > len(model.Todo)-1 {
		return -1, errors.New("id not found")
	}
	model.Todo[id].Todo = todo

	return id + 1, nil
}
