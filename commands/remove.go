package commands

import (
	"fmt"
)

type taskListRemoval struct {
	date string
}

func NewTaskListRemoval(date string) Command {
	return taskListRemoval{date: date}
}

func (r taskListRemoval) Execute() error {

	if removalDate, err := ResolveDate(r.date); err != nil {

		return fmt.Errorf("invalid removal date provided: %s", r.date)

	} else {

		tasksDao.RemoveByDate(removalDate)

		return nil
	}
}
