package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Todo []string
	Time []time.Time
}

func getHour(hour string) time.Time {

	// if hour

	getHour, _ := time.Parse(time.Kitchen, hour)

	return getHour

}

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
	t.Time = append(t.Time, getHour(timeTodo))

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
	t.Time[id] = getHour(timeTodo)

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

func main() {

	run := 1
	var todo Todo
	scanner := bufio.NewScanner(os.Stdin)
	todo.AddTodo("p", ":00PM")

	for run > 0 {
		fmt.Print("\nTodolist\n")
		todo.ShowTodo()

		fmt.Println("\nMenu Todolist")
		fmt.Print(" 1. Tambah Todo \n 2. Update Todo \n 3. Update Time \n 4. Hapus Todo \n 5. Keluar \n")

		// Get input from terminal dan convert into integer
		fmt.Print("Pilih Menu : ")
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())

		switch num {
		case 1:
			fmt.Println("Tambah Todo : ")

			fmt.Print("Todo : ")
			scanner.Scan()
			scanner.Text()
			task := scanner.Text()

			fmt.Print("Waktu : ")
			scanner.Scan()
			scanner.Text()
			time := scanner.Text()

			_, err := todo.AddTodo(task, time)
			if err != nil {
				fmt.Printf("\nError : %v \n", err)
			}
		case 2:
			fmt.Println("Update Todo dari Id")
			fmt.Print("Masukan ID : ")
			scanner.Scan()
			scanner.Text()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukan Todo : ")
			scanner.Scan()
			scanner.Text()
			task := scanner.Text()
			res, err := todo.UpdateTodo(id, task)

			if err != nil {
				fmt.Printf("\nError : %v \n", err)
			} else {
				fmt.Printf("\n %v Updated \n", res)
			}

		case 3:
			fmt.Println("Update Waktu dari Id")
			fmt.Print("Masukan ID : ")
			scanner.Scan()
			scanner.Text()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Masukan Waktu : ")
			scanner.Scan()
			scanner.Text()
			task := scanner.Text()
			res, err := todo.UpdateTime(id, task)

			if err != nil {
				fmt.Printf("\nError : %v \n", err)
			} else {
				fmt.Printf("\n %v :Updated \n", res)
			}
		case 4:
			scanner.Scan()
			scanner.Text()
			id, _ := strconv.Atoi(scanner.Text())

			res, err := todo.RemoveTodo(id)

			if err != nil {
				fmt.Printf("\nError %v : \n", err)
			} else {
				fmt.Printf("\n %v :Updated \n", res)
			}
		case 5:
			run = 0
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}

}
