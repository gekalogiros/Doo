package commands

import (
	"fmt"
	"github.com/gekalogiros/Doo/dao"
	"github.com/gekalogiros/Doo/model"
)

type taskCreation struct {
	dueDate     string
	description string
}

func NewTaskCreation(dueDate string, description string) Command {
	return taskCreation{dueDate: dueDate, description: description}
}

func (c taskCreation) Execute() {

	var tasksDao dao.TaskDao = dao.NewFileSystemTasksDao() //DI

	date, error := ResolveDueDate(c.dueDate)

	failIfError(error, fmt.Sprintf("Invalid due date format provided: %s", date))

	note := model.NewTask(c.description, date)

	tasksDao.Save(&note)
}
