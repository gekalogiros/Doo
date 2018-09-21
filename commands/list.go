package commands

import (
	"fmt"
	"github.com/gekalogiros/Doo/dao"
	"log"
	"time"
)

type taskListRetrieval struct {
	date string
}

func NewTaskListRetrieval(date string) Command {
	return taskListRetrieval{date: date}
}

func (lr taskListRetrieval) Execute() {

	var tasksDao dao.TaskDao = dao.NewFileSystemTasksDao() //DI

	date, error := time.Parse("02-01-2006", lr.date)

	failIfError(error, fmt.Sprintf("Invalid retrieval date provided: %s", date))

	tasks := tasksDao.RetrieveAllByDate(date)

	for _, task := range tasks {
		log.Println(task)
	}
}
