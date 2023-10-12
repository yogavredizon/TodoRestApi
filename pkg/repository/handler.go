package repository

import (
	"errors"
	"fmt"

	"github.com/yogavredizon/todolist/pkg/helper"
	"github.com/yogavredizon/todolist/pkg/model"
)

type Todo model.Todo

func (t *Todo) AddTodo(task, timeTodo string) (Todo, error) {
	if task == "" {
		return *t, errors.New("task can't blank")
	}

	for _, todo := range t.Todo {
		if todo == task {
			return *t, errors.New("task already taken")
		}
	}

	t.Todo = append(t.Todo, task)
	t.Time = append(t.Time, helper.GetHour(timeTodo))

	return *t, nil
}

func (t *Todo) ShowTodo() {

	for i, todo := range t.Todo {
		fmt.Printf("%d. %s : %d:%d \n", i+1, todo, t.Time[i].Hour(), t.Time[i].Minute())
	}
}

func (t *Todo) RemoveTodo(id int) (int, error) {
	id = id - 1

	if id > len(t.Todo)-1 {
		return -1, errors.New("id not found")
	}

	t.Todo = append(t.Todo[:id], t.Todo[id+1:]...)
	t.Time = append(t.Time[:id], t.Time[id+1:]...)

	return id + 1, nil
}

func (t *Todo) UpdateTime(id int, timeTodo string) (int, error) {
	id -= 1
	if id > len(t.Time)-1 {
		return -1, errors.New("id not found")
	}
	t.Time[id] = helper.GetHour(timeTodo)

	return id + 1, nil
}
func (t *Todo) UpdateTodo(id int, todo string) (int, error) {
	id -= 1
	if id > len(t.Todo)-1 {
		return -1, errors.New("id not found")
	}
	t.Todo[id] = todo

	return id + 1, nil
}
