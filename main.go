package main

import (
	"net/http"

	"github.com/yogavredizon/todolist/controller"
)

func main() {
	controller.Handler()
	http.ListenAndServe(":8080", nil)
}
