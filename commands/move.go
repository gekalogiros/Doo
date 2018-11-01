package commands

import (
	"fmt"
)

type taskMovement struct {
	id   string
	from string
	to   string
}

func NewTaskMovement(taskId string, from string, to string) Command {
	return taskMovement{id: taskId, from: from, to: to}
}

func (m taskMovement) Execute() error {

	fromDate, err := ResolveDate(m.from)

	if err != nil {
		return fmt.Errorf("invalid due date format provided: %s", fromDate)
	}

	toDate, err := ResolveDueDate(m.to)

	if err != nil {
		return fmt.Errorf("invalid target date format provided: %s", toDate)
	}

	tasksDao.Move(m.id, fromDate, toDate)

	return nil
}
