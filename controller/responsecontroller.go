package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yogavredizon/todolist/model"
)

var BaseURL = "localhost:8080"

func FetchTodos() ([]model.Todo, error) {
	client := http.Client{}

	var data []model.Todo

	request, errReq := http.NewRequest(http.MethodGet, BaseURL+"/", nil)

	if errReq != nil {
		return data, nil
	}

	response, errResponse := client.Do(request)

	if errResponse != nil {
		return data, nil
	}
	result, _ := io.ReadAll(response.Body)
	fmt.Println(result)

	json.NewDecoder(response.Body).Decode(&data)

	return data, nil

}
