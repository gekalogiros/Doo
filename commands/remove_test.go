package commands

import (
	"reflect"
	"testing"

	"github.com/gekalogiros/Doo/mocks"
	"github.com/golang/mock/gomock"
)

func TestNewTaskListRemoval(t *testing.T) {

	type args struct {
		date    string
		allPast bool
	}
	tests := []struct {
		name string
		args args
		want Command
	}{
		{"Can be Constructed", args{"some date", false}, NewTaskListRemoval("some date", false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskListRemoval(tt.args.date, tt.args.allPast); !reflect.DeepEqual(got, tt.want) {
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
		date    string
		allPast bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"today", fields{date: "today", allPast: false}, false},
		{"future - tomorrow", fields{date: "tomorrow", allPast: false}, false},
		{"future - expression - months", fields{date: "1m", allPast: false}, false},
		{"future - expression - months - case insensitive", fields{date: "1M", allPast: false}, false},
		{"future - expression - days", fields{date: "1d", allPast: false}, false},
		{"future - expression - days - case insensitive", fields{date: "1D", allPast: false}, false},
		{"future - temporal - days - number", fields{date: "1", allPast: false}, false},
		{"future - date - dd/MM/yyyy", fields{date: "02/01/2019", allPast: false}, false},
		{"future - date - dd/MM/yy", fields{date: "02/01/19", allPast: false}, false},
		{"future - date - d/M/yyyy", fields{date: "2/1/19", allPast: false}, false},
		{"future - date - dd-MM-yyyy", fields{date: "02-01-2019", allPast: false}, false},
		{"future - date - dd-MM-yy", fields{date: "02-01-19", allPast: false}, false},
		{"future - date - d-M-yyyy", fields{date: "2-1-19", allPast: false}, false},
		{"past - expression - months", fields{date: "-1m", allPast: false}, false},
		{"past - expression - months - case insensitive", fields{date: "-1M", allPast: false}, false},
		{"past - expression - days", fields{date: "-1d", allPast: false}, false},
		{"past - expression - days - case insensitive", fields{date: "-1D", allPast: false}, false},
		{"past - temporal - days - number", fields{date: "-1", allPast: false}, false},
		{"past", fields{date: "yesterday", allPast: false}, false},
		{"all past", fields{date: "", allPast: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := taskListRemoval{
				date:    tt.fields.date,
				allPast: tt.fields.allPast,
			}
			if r.allPast {
				mockDao.EXPECT().RemovePast().Times(1)
			} else {
				mockDao.EXPECT().RemoveByDate(gomock.Any()).Times(1)
			}
			if err := r.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("taskListRemoval.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
