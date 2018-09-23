package commands

import (
	"fmt"
	"time"
)

type taskListRemoval struct {
	date string
}

func NewTaskListRemoval(date string) Command {
	return taskListRemoval{date: date}
}

func (r taskListRemoval) Execute() error {

	if date, err := time.Parse("02-01-2006", r.date); err != nil {

		return fmt.Errorf("invalid removal date provided: %s", date)

	} else {

		tasksDao.RemoveAll(date)

		return nil
	}
}
