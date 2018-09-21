package commands

import (
	"fmt"
	"github.com/gekalogiros/Doo/dao"
	"time"
)

type taskListRemoval struct {
	date string
}

func NewTaskListRemoval(date string) Command {
	return taskListRemoval{date: date}
}

func (r taskListRemoval) Execute() {

	var tasksDao dao.TaskDao = dao.NewFileSystemTasksDao() //DI

	date, error := time.Parse("02-01-2006", r.date)

	failIfError(error, fmt.Sprintf("Invalid removal date provided: %s", date))

	tasksDao.RemoveAll(date)
}
