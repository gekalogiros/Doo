package dao

import (
	"time"

	"github.com/gekalogiros/Doo/model"
)

type TaskDao interface {
	Save(n *model.Task)
	RemoveByDate(date time.Time)
	RetrieveByDate(date time.Time) []model.Task
	RemovePast()
	Move(id string, source time.Time, target time.Time)
}
