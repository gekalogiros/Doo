package commands

import (
	"reflect"
	"testing"
)

func TestNewTaskListRemoval(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want Command
	}{
		{"Can be Constructed", args{"some date"}, NewTaskListRemoval("some date")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskListRemoval(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskListRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskListRemoval_Execute(t *testing.T) {
	type fields struct {
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Should parse date format 02-01-2006", fields{date:"12-12-2018"}, false},
		{"Should produce error when date is not in 02-01-2006 format", fields{date:"12/12/2018"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := taskListRemoval{
				date: tt.fields.date,
			}
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("taskListRemoval.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
