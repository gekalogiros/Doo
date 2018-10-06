package commands

import (
	"github.com/gekalogiros/Doo/mocks"
	"github.com/golang/mock/gomock"
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

func TestExecute(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDao := mock_dao.NewMockTaskDao(mockCtrl)

	tasksDao = mockDao

	type fields struct {
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"today", fields{date: "today"}, false},
		{"future - tomorrow", fields{date: "tomorrow"}, false},
		{"future - expression - months", fields{date: "1m"}, false},
		{"future - expression - months - case insensitive", fields{date: "1M"}, false},
		{"future - expression - days", fields{date: "1d"}, false},
		{"future - expression - days - case insensitive", fields{date: "1D"}, false},
		{"future - temporal - days - number", fields{date: "1"}, false},
		{"future - date - dd/MM/yyyy", fields{date: "02/01/2019"}, false},
		{"future - date - dd/MM/yy", fields{date: "02/01/19"}, false},
		{"future - date - d/M/yyyy", fields{date: "2/1/19"}, false},
		{"future - date - dd-MM-yyyy", fields{date: "02-01-2019"}, false},
		{"future - date - dd-MM-yy", fields{date: "02-01-19"}, false},
		{"future - date - d-M-yyyy", fields{date: "2-1-19"}, false},
		{"past - expression - months", fields{date: "-1m"}, false},
		{"past - expression - months - case insensitive", fields{date: "-1M"}, false},
		{"past - expression - days", fields{date: "-1d"}, false},
		{"past - expression - days - case insensitive", fields{date: "-1D"}, false},
		{"past - temporal - days - number", fields{date: "-1"}, false},
		{"past", fields{date: "yesterday"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := taskListRemoval{
				date: tt.fields.date,
			}
			mockDao.EXPECT().RemoveByDate(gomock.Any()).Times(1)
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("taskListRemoval.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
