package commands

import (
	"fmt"
	"github.com/gekalogiros/Doo/dao"
	"github.com/gekalogiros/Doo/model"
	"log"
)

type taskCreation struct {
	dueDate     string
	description string
}

func NewTaskCreation(dueDate string, description string) taskCreation {
	return taskCreation{dueDate: dueDate, description: description}
}

func (c taskCreation) Execute() {

	var tasksDao dao.TaskDao = dao.NewFileSystemTasksDao() //DI

	date, error := ResolveDueDate(c.dueDate)

	if error != nil {
		err := fmt.Sprintf("Invalid due date format provided: %s", date)
		log.Fatal(fmt.Sprintf(errorMessageTemplate, err))
	}

	note := model.NewTask(c.description, date)

	tasksDao.Save(&note)
}
