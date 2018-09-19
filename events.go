package main

import (
	"fmt"
	"github.com/gekalogiros/Doo/dao"
	"github.com/gekalogiros/Doo/model"
	"log"
	"time"
)

const (
	errorMessageTemplate = "%s, Check documentation at github.com/gekalogiros/Doo"
)

var notesDao dao.TaskDao = dao.NewFileSystemTasksDao()

type taskCreation struct {
	dueDate string
	description string
}

type taskListRemoval struct {
	date string
}

func newTaskCreation(dueDate string, description string) taskCreation {
	return taskCreation{dueDate: dueDate, description:description}
}

func (c taskCreation) execute(){
	date, error := ResolveDueDate(c.dueDate)

	if error != nil {
		err := fmt.Sprintf("Invalid due date format provided: %s", date)
		log.Fatal(fmt.Sprintf(errorMessageTemplate, err))
	}

	note := model.NewTask(c.description, date)

	notesDao.Save(&note)
}

func NewTaskListRemoval(date string) taskListRemoval {
	return taskListRemoval{date}
}

func (r taskListRemoval) execute(){

	date, error := time.Parse("02-01-2006", r.date)

	if error != nil {
		err := fmt.Sprintf("Invalid removal date provided: %s", date)
		log.Fatal(fmt.Sprintf(errorMessageTemplate, err))
	}

	notesDao.RemoveAll(date)
}


