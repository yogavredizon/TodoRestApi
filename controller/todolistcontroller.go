package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yogavredizon/todolist/helper"
	"github.com/yogavredizon/todolist/model"
	"github.com/yogavredizon/todolist/repository"
)

func fetchTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := repository.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	if r.Method == "GET" {
		web := model.Web{
			Status: "Success Getting Todo",
			Code:   http.StatusOK,
			Data:   todos,
		}
		res, err := json.Marshal(web)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		w.Write(res)
		return
	}

	http.Error(w, "URL Not Found", http.StatusNotFound)
}

func fetchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/todos/"))
	todo := r.FormValue("todo")
	time := r.FormValue("time")

	switch r.Method {
	case "GET":
		if id < 1 {
			http.Error(w, "ID not Found", http.StatusBadRequest)
			return
		}

		todo, err := repository.FindById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		webResponse := &model.Web{
			Status: "Success getting Todo from id",
			Code:   http.StatusOK,
			Data:   todo,
		}
		res, err := json.Marshal(webResponse)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		w.Write(res)
		return
	case "POST":

		_, err := repository.AddTodo(todo, time)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		timeValue := helper.GetHour(time)

		web := model.Web{
			Status: "Success adding todo",
			Code:   http.StatusOK,
			Data: model.Task{
				Todo: todo,
				Time: timeValue,
			},
		}
		res, err := json.Marshal(web)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		w.Write(res)
		return
	case "PUT":
		if todo != "" && time == "" {
			err := repository.UpdateTodo(id, todo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			web := model.Web{
				Status: "Todo updated",
				Code:   http.StatusOK,
				Data:   todo,
			}

			res, err := json.Marshal(web)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}

			w.Write(res)
			return
		}

		if todo == "" && time != "" {
			err := repository.UpdateTime(id, time)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			web := model.Web{
				Status: "Time Updated",
				Code:   http.StatusOK,
				Data:   helper.GetHour(time),
			}

			res, err := json.Marshal(web)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}

			w.Write(res)
			return
		}

	}
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)

}

func Handler() {
	http.HandleFunc("/todos", fetchTodos)
	http.HandleFunc("/todos/", fetchTodo)

	http.ListenAndServe(":8080", nil)
}
