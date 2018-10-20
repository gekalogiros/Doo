package commands

import (
	"fmt"
)

type taskListRemoval struct {
	date    string
	allPast bool
}

func NewTaskListRemoval(date string, allPast bool) Command {
	return taskListRemoval{date: date, allPast: allPast}
}

func (r taskListRemoval) Execute() error {

	if r.allPast {

		return removePast()

	}

	return removeDate(r.date)
}

func removePast() error {

	tasksDao.RemovePast()

	return nil
}

func removeDate(date string) error {

	if removalDate, err := ResolveDate(date); err != nil {

		return fmt.Errorf("invalid removal date provided: %s", date)

	} else {

		tasksDao.RemoveByDate(removalDate)

		return nil
	}
}
