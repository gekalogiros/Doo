package model

import (
	"time"
)

type Task struct {
	Id          string
	Description string
	Date        time.Time
}

func NewTask(description string, date time.Time) Task {
	n := Task{generateHash(), description, date}
	return n
}
