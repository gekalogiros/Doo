package commands

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	tasksDao = newTasksDaoMock()
}

func TestNewTaskCreation(t *testing.T) {
	type fields struct {
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"expression - months", fields{ date: "1m"}, false},
		{"expression - months - case insensitive", fields{ date: "1M"}, false},
		{"expression - days", fields{ date: "1d"}, false},
		{"expression - days - case insensitive", fields{ date: "1D"}, false},
		{"temporal - days - number", fields{ date: "1"}, false},
		{"date - dd/MM/yyyy", fields{ date: "02/01/2019"}, false},
		{"date - dd/MM/yy", fields{ date: "02/01/19"}, false},
		{"date - d/M/yyyy", fields{ date: "2/1/19"}, false},
		{"date - dd-MM-yyyy", fields{ date: "02-01-2019"}, false},
		{"date - dd-MM-yy", fields{ date: "02-01-19"}, false},
		{"date - d-M-yyyy", fields{ date: "2-1-19"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := taskCreation{
				dueDate: tt.fields.date,
			}
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("taskCreation.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTaskCreation_withInvalidDate(t *testing.T) {
	invalidDate := "invalid-date"
	description := "description"
	err := NewTaskCreation(invalidDate, description).Execute()
	assert.Equal(t, fmt.Errorf("invalid due date format provided: "+invalidDate), err, "")
}
