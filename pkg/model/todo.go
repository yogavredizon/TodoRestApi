package model

import "time"

type Task struct {
	Todo string
	Time time.Time
}

var Todo = []Task{}
