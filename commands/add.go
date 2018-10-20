package commands

import (
	"fmt"

	"github.com/gekalogiros/Doo/model"
)

type taskCreation struct {
	dueDate     string
	description string
}

func NewTaskCreation(dueDate string, description string) Command {
	return taskCreation{dueDate: dueDate, description: description}
}

func (c taskCreation) Execute() error {

	if date, err := ResolveDueDate(c.dueDate); err != nil {

		return fmt.Errorf("invalid due date format provided: %s", c.dueDate)

	} else {

		task := model.NewTask(c.description, date)

		tasksDao.Save(&task)

		return nil
	}
}
