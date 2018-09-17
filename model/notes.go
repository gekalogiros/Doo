package model

import (
	"time"
)

type Note struct {
	Id          string
	Description string
	Date        time.Time
}

func NewNote(description string, date time.Time) Note {
	n := Note{generateHash(), description, date}
	return n
}
