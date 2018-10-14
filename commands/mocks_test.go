package commands

import (
	"time"

	"github.com/gekalogiros/Doo/dao"
	"github.com/gekalogiros/Doo/model"
)

type tasksDaoMock struct{}

func (d tasksDaoMock) Save(task *model.Task) {}

func (d tasksDaoMock) RemoveByDate(date time.Time) {}

func (d tasksDaoMock) RetrieveByDate(date time.Time) []model.Task {
	return []model.Task{
		{Id: "1111", Description: "first-task"},
		{Id: "2222", Description: "second-task"},
	}
}

func (d tasksDaoMock) RemovePast() {}

func newTasksDaoMock() dao.TaskDao {
	return tasksDaoMock{}
}
