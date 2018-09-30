package commands

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	tasksDao = newTasksDaoMock()
}

func TestNewTaskListRetrieval_withInvalidDate(t *testing.T) {
	invalidDate := "sfdgsdg"
	err := NewTaskListRetrieval(invalidDate).Execute()
	assert.Equal(t, fmt.Errorf("invalid retrieval date provided: "+invalidDate), err, "")
}

func ExampleNewTaskListRetrieval() {
	NewTaskListRetrieval("24-06-2018").Execute()
	// Output:
	// 1111	first-task
	// 2222	second-task
}
