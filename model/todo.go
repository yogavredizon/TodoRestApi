package model

import "time"

type Task struct {
	Id   int
	Todo string
	Time time.Time
}
