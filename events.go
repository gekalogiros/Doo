package main

import (
	"fmt"
	"github.com/gekalogiros/todo/dao"
	"github.com/gekalogiros/todo/model"
	"log"
	"time"
)

const (
	ERROR_MESSAGE_TEMPLATE = "%s, Check documentation at github.com/gekalogiros/doo"
)

var notesDao dao.NotesDao = dao.NewFileSystemNotesDao()

type additionEnquiry struct {
	dueDate string
	description string
}

func NewAdditionEnquiry(dueDate string, description string) additionEnquiry {
	return additionEnquiry{dueDate:dueDate, description:description}
}

func (a additionEnquiry) execute(){
	date, error := ResolveDueDate(a.dueDate)

	if error != nil {
		err := fmt.Sprintf("Invalid due date format provided: %s", date)
		log.Fatal(fmt.Sprintf(ERROR_MESSAGE_TEMPLATE, err))
	}

	note := model.NewNote(a.description, date)

	notesDao.Save(&note)
}

type pastDateRemovalEnquiry struct {
	date string
}

func NewPastDateRemovalEnquiry(date string) pastDateRemovalEnquiry {
	return pastDateRemovalEnquiry{date}
}

func (r pastDateRemovalEnquiry) execute(){

	date, error := time.Parse("02-01-2006", r.date)

	if error != nil {
		err := fmt.Sprintf("Invalid removal date provided: %s", date)
		log.Fatal(fmt.Sprintf(ERROR_MESSAGE_TEMPLATE, err))
	}

	notesDao.RemoveAll(date)
}


