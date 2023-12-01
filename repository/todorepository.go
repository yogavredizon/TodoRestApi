package repository

import (
	"errors"

	"github.com/yogavredizon/todolist/config"
	"github.com/yogavredizon/todolist/helper"
	"github.com/yogavredizon/todolist/model"
)

func AddTodo(todo, timeTodo string) (model.Task, error) {
	conn := config.NewDB()
	newTodo := model.Task{}

	if todo == "" || timeTodo == "" {
		return newTodo, errors.New("please Fill all field")
	}

	defer conn.Close()

	rows, _ := FindAll()

	for i := 0; i < len(rows); i++ {
		if helper.GetHour(timeTodo) == rows[i].Time {
			return newTodo, errors.New("time is already set")
		}
	}

	stmt, err := conn.Prepare("INSERT INTO todo(todo, time) VALUES (?,?)")

	if err != nil {
		return newTodo, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(todo, timeTodo)

	rowaffect, _ := res.RowsAffected()
	if rowaffect == 0 {
		return newTodo, errors.New("unable to add todo")
	}

	if err != nil {
		return newTodo, err
	}

	return newTodo, nil
}

func FindAll() ([]model.Task, error) {

	// create array of model.Task
	var arrOfTask []model.Task

	// connection to database
	conn := config.NewDB()
	defer conn.Close()

	// calling database to select colom and send value into statement
	stmt, err := conn.Prepare("SELECT id, todo, time from todo")

	if err != nil {
		panic(err)
	}

	// close statement when accessing database has done
	defer stmt.Close()

	// excute prepare statment for querying value
	rows, _ := stmt.Query()

	// rows.Next() will stop when rows.Scan() have false value
	// all value will replace into struct of Task
	for rows.Next() {
		var hour string

		var todo model.Task
		rows.Scan(&todo.Id, &todo.Todo, &hour)
		getHour := helper.GetHour(hour)
		todo.Time = getHour
		arrOfTask = append(arrOfTask, todo)
	}

	// every task which get from databse will add/append into array of Task

	return arrOfTask, nil
}

func FindById(id int) (model.Task, error) {
	conn := config.NewDB()
	defer conn.Close()

	var todo = model.Task{}

	stmt, err := conn.Prepare("SELECT id, todo, time FROM todo where id = ?")
	if err != nil {
		return todo, err
	}

	hour := ""
	err = stmt.QueryRow(id).Scan(&todo.Id, &todo.Todo, &hour)

	todo.Time = helper.GetHour(hour)

	if err != nil {
		return todo, errors.New("id not found")
	}

	return todo, nil
}

func RemoveTodo(id int) error {

	conn := config.NewDB()

	defer conn.Close()

	stmt, err := conn.Prepare("DELETE FROM todo WHERE id=?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	rowAffected, _ := res.RowsAffected()

	if rowAffected == 0 {
		return errors.New("id not found")
	}
	if err != nil {
		return err
	}
	return nil
}

func UpdateTime(id int, timeTodo string) error {

	conn := config.NewDB()
	defer conn.Close()

	todo, _ := FindAll()

	for _, t := range todo {
		if helper.GetHour(timeTodo) == t.Time {
			return errors.New("time already set, change to another time")
		}
	}

	stmt, err := conn.Prepare("UPDATE todo set time = ? where id = ?")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(timeTodo, id)

	if err != nil {
		return err
	}

	defer stmt.Close()

	rowAffect, _ := result.RowsAffected()
	if rowAffect == 0 {
		return errors.New("id not found")
	}
	return nil
}

func UpdateTodo(id int, todo string) error {
	conn := config.NewDB()
	defer conn.Close()

	getId, _ := FindById(id)

	if getId.Todo == todo {
		return errors.New("todo can't same with older")
	}
	stmt, err := conn.Prepare("UPDATE todo set todo = ? where id = ?")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(todo, id)

	if err != nil {
		return err
	}

	defer stmt.Close()

	rowAffect, _ := result.RowsAffected()
	if rowAffect == 0 {
		return errors.New("id not found")
	}
	return nil
}
