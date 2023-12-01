package model

import "time"

type Todo struct {
	Id   int       `json:"id"`
	Todo string    `json:"Todo"`
	Time time.Time `json:"Time"`
}
