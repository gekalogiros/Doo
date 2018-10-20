package commands

import (
	"fmt"
	"testing"
	"time"

	"github.com/gekalogiros/Doo/dao"
	"github.com/gekalogiros/Doo/mocks"
	"github.com/gekalogiros/Doo/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewTaskListRetrieval_withInvalidDate(t *testing.T) {
	invalidDate := "sfdgsdg"
	err := NewTaskListRetrieval(invalidDate).Execute()
	assert.Equal(t, fmt.Errorf("invalid retrieval date provided: "+invalidDate), err, "")
}

func TestNewTaskListRetrieval_withValidDate(t *testing.T) {

	// golang mock newbie - I'm pretty sure the mock template code below can be done in a better way!!!

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	tasksDao := mock_dao.NewMockTaskDao(mockCtrl)

	setDao(tasksDao)

	defer setDao(dao.NewFileSystemTasksDao())

	taskListDate, _ := time.Parse("02-01-2006", "24-06-2018")

	var actual []model.Task

	printTasks = func(tasks []model.Task) {
		actual = tasks
	}

	expected := []model.Task{
		model.Task{
			Id:          "xxxx",
			Description: "test description",
			Date:        taskListDate,
		},
	}

	tasksDao.EXPECT().RetrieveByDate(taskListDate).Return(expected).Times(1)

	NewTaskListRetrieval("24-06-2018").Execute()

	assert.Equal(t, expected, actual)
}
