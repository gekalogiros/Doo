package commands

import (
	"fmt"
	"github.com/gekalogiros/Doo/dao"
	"log"
	"time"
)

const (
	errorMessageTemplate = "%s, Check documentation at github.com/gekalogiros/Doo"
)

type taskListRemoval struct {
	date string
}

func NewTaskListRemoval(date string) taskListRemoval {
	return taskListRemoval{date: date}
}

func (r taskListRemoval) Execute() {

	var tasksDao dao.TaskDao = dao.NewFileSystemTasksDao() //DI

	date, error := time.Parse("02-01-2006", r.date)

	if error != nil {
		err := fmt.Sprintf("Invalid removal date provided: %s", date)
		log.Fatal(fmt.Sprintf(errorMessageTemplate, err))
	}

	tasksDao.RemoveAll(date)
}
