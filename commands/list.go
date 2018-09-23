package commands

import (
	"fmt"
	"time"
)

type taskListRetrieval struct {
	date string
}

func NewTaskListRetrieval(date string) Command {
	return taskListRetrieval{date: date}
}

func (lr taskListRetrieval) Execute() error {

	if date, err := time.Parse("02-01-2006", lr.date); err != nil {

		return fmt.Errorf("invalid retrieval date provided: %s", lr.date)

	} else {

		tasks := tasksDao.RetrieveAllByDate(date)

		for _, task := range tasks {
			fmt.Println(task)
		}

		return nil
	}
}
