package commands

import (
	"fmt"
	"github.com/gekalogiros/Doo/formatter"
)

type taskListRetrieval struct {
	date string
}

func NewTaskListRetrieval(date string) Command {
	return taskListRetrieval{date: date}
}

func (lr taskListRetrieval) Execute() error {

	if retrievalDate, err := ResolveDate(lr.date); err != nil {

		return fmt.Errorf("invalid retrieval date provided: %s", lr.date)

	} else {

		tasks := tasksDao.RetrieveAllByDate(retrievalDate)

		if len(tasks) == 0 {

			fmt.Printf("%s %s",
				formatter.Red("Cannot find task list for date provided: "),
				formatter.BRed(retrievalDate.Format("02-01-2006"), formatter.Boldest))

		} else {
			for _, task := range tasks {
				fmt.Println(task)
			}
		}

		return nil
	}
}
