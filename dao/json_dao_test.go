package dao

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"reflect"
	"testing"
	"time"

	"github.com/gekalogiros/Doo/model"
	"github.com/stretchr/testify/assert"
)

func Test_json_Save(t *testing.T) {
	configDir := tempConfigDir()
	dbFile := dbFile()
	task := model.NewTask("description", time.Now())

	type fields struct {
		configDir string
		filename  string
	}
	type args struct {
		n *model.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Should create json with tasks",
			fields{
				configDir: configDir,
				filename:  dbFile,
			},
			args{
				n: &task,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := json{
				configDir: tt.fields.configDir,
				filename:  tt.fields.filename,
			}
			j.Save(tt.args.n)

			file := path.Join(j.configDir, j.filename)

			assert.FileExists(t, file)

			contentAsBytes, err := ioutil.ReadFile(file)

			assert.NoError(t, err)

			assert.Contains(t, string(contentAsBytes), tt.args.n.Description)
		})
	}
}

func Test_json_Move(t *testing.T) {

	from := time.Now()

	to := time.Now().AddDate(0, 0, 1)

	task := model.NewTask("description", from)

	underTest := json{
		configDir: tempConfigDir(),
		filename:  dbFile(),
	}

	underTest.Save(&task)

	type fields struct {
		configDir string
		filename  string
	}
	type args struct {
		taskID string
		from   time.Time
		to     time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Should move tasks between task lists", args{taskID: task.Id, from: from, to: to},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			underTest.Move(tt.args.taskID, tt.args.from, tt.args.to)

			sourceTaskList := underTest.RetrieveByDate(tt.args.from)

			assert.Empty(t, sourceTaskList)

			targetTaskList := underTest.RetrieveByDate(tt.args.to)

			assert.Len(t, targetTaskList, 1)
		})
	}
}

func Test_json_RemoveByDate(t *testing.T) {

	task := model.NewTask("description", time.Now())

	underTest := json{
		configDir: tempConfigDir(),
		filename:  dbFile(),
	}

	underTest.Save(&task)

	type fields struct {
		configDir string
		filename  string
	}
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Should remove tasks for date",
			args{
				date: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			underTest.RemoveByDate(tt.args.date)

			tasks := underTest.RetrieveByDate(tt.args.date)

			assert.Empty(t, tasks)
		})
	}
}

func Test_json_RetrieveByDate(t *testing.T) {

	location := localLocation()
	task := model.NewTask("description", time.Now().In(location).Truncate(24*time.Hour))

	underTest := json{
		configDir: tempConfigDir(),
		filename:  dbFile(),
	}

	underTest.Save(&task)

	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want []model.Task
	}{
		{
			"Should retrieve task from json",
			args{
				date: time.Now(),
			},
			[]model.Task{task},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := underTest.RetrieveByDate(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json.RetrieveByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_json_RemovePast(t *testing.T) {

	location, _ := time.LoadLocation("Local")

	task1 := model.NewTask("desc 1", time.Date(2017, 12, 1, 0, 0, 0, 0, location))
	task2 := model.NewTask("desc 2", time.Date(2016, 12, 1, 0, 0, 0, 0, location))
	task3 := model.NewTask("desc 3", time.Date(2015, 12, 1, 0, 0, 0, 0, location))
	todayTask := model.NewTask("desc 4", time.Now())

	underTest := json{
		configDir: tempConfigDir(),
		filename:  dbFile(),
	}

	underTest.Save(&task1)
	underTest.Save(&task2)
	underTest.Save(&task3)
	underTest.Save(&todayTask)

	underTest.RemovePast()

	assert.Empty(t, underTest.RetrieveByDate(task1.Date))
	assert.Empty(t, underTest.RetrieveByDate(task2.Date))
	assert.Empty(t, underTest.RetrieveByDate(task3.Date))
	assert.NotEmpty(t, underTest.RetrieveByDate(todayTask.Date))
}

func tempConfigDir() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s/test-%d", os.TempDir(), rand.Int())
}

func dbFile() string {
	return "tasks.json"
}

func localLocation() *time.Location {
	loc, _ := time.LoadLocation("Local")
	return loc
}
