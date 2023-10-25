package test

import (
	"testing"

	repo "github.com/yogavredizon/todolist/pkg/repository"
)

func TestAddTodo(t *testing.T) {

	err := repo.AddTodo("Belajar", "6:00PM")
	repo.AddTodo("Belajar", "6:00PM")
	repo.AddTodo("Makan", "6:00PM")
	repo.AddTodo("Tidur", "6:00PM")

	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestShowTodo(t *testing.T) {
	err := repo.AddTodo("Belajar", "6:00PM")
	// _ = repo.AddTodo("Belajar", "6:00PM")

	if err != nil {
		t.Fatal(err.Error())
	}
	repo.ShowTodo()
}

func TestRemoveTodo(t *testing.T) {
	repo.AddTodo("Belajar", "6:00PM")
	_, err := repo.RemoveTodo(0)

	if err != nil {
		t.Log(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	repo.AddTodo("Belajar", "6:00PM")
	_, err := repo.UpdateTime(1, "7:00PM")
	repo.UpdateTodo(1, "Kerja")

	if err != nil {
		t.Log(err.Error())
	}
}
