package commands

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var now = time.Now()

func TestResolveDueDateWithInvalidDate(t *testing.T) {
	assertInvalidDueDate(t, "dsfhdsfiushdf", "failed to parse date")
	assertInvalidDueDate(t, "1235412421241243124213513464562341243124214321D", "failed to parse number of days, period provided is probably too long")
	assertInvalidDueDate(t, "5364367346523545676235345745462353473462356315", "failed to parse number of days, period provided is probably too long")
}

func assertInvalidDueDate(t *testing.T, input string, expected string) {
	_, err := ResolveDueDate(input)

	assert.Error(t, err, expected)
}

func TestResolveDueDate(t *testing.T) {
	type args struct {
		dueDate string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"today", args{dueDate: "today"}, now, false},
		{"tomorrow", args{dueDate: "tomorrow"}, toDate(now.Year(), now.Month(), now.Day()+1), false},
		{"Temporal day", args{dueDate: "134d"}, now.AddDate(0, 0, 134), false},
		{"Temporal day - uppercase", args{dueDate: "134D"}, now.AddDate(0, 0, 134), false},
		{"Temporal day - uppercase with leading zero", args{dueDate: "0134D"}, now.AddDate(0, 0, 134), false},
		{"Temporal month", args{dueDate: "1m"}, now.AddDate(0, 1, 0), false},
		{"Temporal month - uppercase", args{dueDate: "1M"}, now.AddDate(0, 1, 0), false},
		{"Period", args{dueDate: "134"}, now.AddDate(0, 0, 134), false},
		{"Period with leading zero", args{dueDate: "04"}, now.AddDate(0, 0, 4), false},
		{"Period with leading and trailing zero", args{dueDate: "020"}, now.AddDate(0, 0, 20), false},
		{"Date - d/MM/yyyy", args{dueDate: "1/12/2018"}, toDate(2018, 12, 1), false},
		{"Date - dd/MM/yyyy", args{dueDate: "01/12/2018"}, toDate(2018, 12, 1), false},
		{"Date - d/MM/yy", args{dueDate: "1/12/18"}, toDate(2018, 12, 1), false},
		{"Date - d-MM-yyyy", args{dueDate: "1-12-2018"}, toDate(2018, 12, 1), false},
		{"Date - dd-MM-yyyy", args{dueDate: "01-12-2018"}, toDate(2018, 12, 1), false},
		{"Date - d-MM-yy", args{dueDate: "1-12-18"}, toDate(2018, 12, 1), false},
		{"Invalid Due Date - Letters", args{dueDate: "dsfhdsfiushdf"}, now, true},
		{"Invalid Due Date - Temporal", args{dueDate: "1235412421241243124213513464562341243124214321D"}, now, true},
		{"Invalid Due Date - Period", args{dueDate: "5364367346523545676235345745462353473462356315"}, now, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveDueDate(tt.args.dueDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveDueDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotAsString := got.Format("02-01-2006")
			wantAsString := tt.want.Format("02-01-2006")
			if !reflect.DeepEqual(gotAsString, wantAsString) {
				t.Errorf("ResolveDueDate() = %v, want %v", gotAsString, wantAsString)
			}
		})
	}
}

func TestResolveDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"today", args{date: "today"}, now, false},
		{"tomorrow", args{date: "tomorrow"}, toDate(now.Year(), now.Month(), now.Day()+1), false},
		{"yesterday", args{date: "yesterday"}, toDate(now.Year(), now.Month(), now.Day()-1), false},
		{"Temporal day in the past", args{date: "-134d"}, now.AddDate(0, 0, -134), false},
		{"Temporal day", args{date: "134d"}, now.AddDate(0, 0, 134), false},
		{"Temporal day - uppercase", args{date: "134D"}, now.AddDate(0, 0, 134), false},
		{"Temporal day - uppercase with leading zero", args{date: "0134D"}, now.AddDate(0, 0, 134), false},
		{"Temporal month", args{date: "1m"}, now.AddDate(0, 1, 0), false},
		{"Temporal month in the past", args{date: "-1m"}, now.AddDate(0, -1, 0), false},
		{"Temporal month - uppercase", args{date: "1M"}, now.AddDate(0, 1, 0), false},
		{"Period", args{date: "134"}, now.AddDate(0, 0, 134), false},
		{"Period in the past", args{date: "-134"}, now.AddDate(0, 0, -134), false},
		{"Period with leading zero", args{date: "04"}, now.AddDate(0, 0, 4), false},
		{"Period with leading and trailing zero", args{date: "020"}, now.AddDate(0, 0, 20), false},
		{"Date - d/MM/yyyy", args{date: "1/12/2018"}, toDate(2018, 12, 1), false},
		{"Date - dd/MM/yyyy", args{date: "01/12/2018"}, toDate(2018, 12, 1), false},
		{"Date - d/MM/yy", args{date: "1/12/18"}, toDate(2018, 12, 1), false},
		{"Date - d-MM-yyyy", args{date: "1-12-2018"}, toDate(2018, 12, 1), false},
		{"Date - dd-MM-yyyy", args{date: "01-12-2018"}, toDate(2018, 12, 1), false},
		{"Date - d-MM-yy", args{date: "1-12-18"}, toDate(2018, 12, 1), false},
		{"Invalid Due Date - Letters", args{date: "dsfhdsfiushdf"}, now, true},
		{"Invalid Due Date - Temporal", args{date: "1235412421241243124213513464562341243124214321D"}, now, true},
		{"Invalid Due Date - Period", args{date: "5364367346523545676235345745462353473462356315"}, now, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotAsString := got.Format("02-01-2006")
			wantAsString := tt.want.Format("02-01-2006")
			if !reflect.DeepEqual(gotAsString, wantAsString) {
				t.Errorf("ResolveDate() = %v, want %v", gotAsString, wantAsString)
			}
		})
	}
}

func toDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}