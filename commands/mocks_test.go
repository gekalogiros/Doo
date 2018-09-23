package commands

import (
	"github.com/gekalogiros/Doo/dao"
	"github.com/gekalogiros/Doo/model"
	"time"
)

type tasksDaoMock struct {}

func (d tasksDaoMock) Save(task *model.Task){}

func (d tasksDaoMock) RemoveAll(date time.Time){}

func (d tasksDaoMock) RetrieveAllByDate(date time.Time) []string {
	return []string{ "first-task", "second-task" }
}

func newTasksDaoMock() dao.TaskDao {
	return tasksDaoMock{}
}
