package dao

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/gekalogiros/Doo/model"
	"github.com/stretchr/testify/assert"
)

const (
	dateFormat = "02_01_2006"
)

var location, _ = time.LoadLocation("Local")
var now = time.Now().In(location)
var configDir = fmt.Sprintf("%stest-%d", os.TempDir(), rand.Int())
var task = model.Task{Id: "xxx0", Description: "I am a test task", Date: now}
var oneYearInThePastDate = now.AddDate(-1, 0, 0)
var pastTask = model.Task{Id: "xxx1", Description: "I am a test past task", Date: oneYearInThePastDate}
var oneMonthInThePastDate = now.AddDate(0, -1, 0)
var anotherPastTask = model.Task{Id: "xxx2", Description: "I am a test past task 2", Date: oneMonthInThePastDate}
var underTest = newFilesystemDao(configDir)

func setUp() {
	cleanUp()
	underTest.Save(&task)
	underTest.Save(&pastTask)
	underTest.Save(&anotherPastTask)
}

func cleanUp() {
	os.RemoveAll(path.Join(configDir, now.Format(dateFormat)))
	os.RemoveAll(path.Join(configDir, oneYearInThePastDate.Format(dateFormat)))
	os.RemoveAll(path.Join(configDir, oneMonthInThePastDate.Format(dateFormat)))
}

func TestFilesystem_Save(t *testing.T) {

	setUp()

	tasksFilename := path.Join(configDir, now.Format(dateFormat))

	assert.FileExists(t, tasksFilename)

	contentAsBytes, err := ioutil.ReadFile(tasksFilename)

	assert.NoError(t, err)

	assert.Contains(t, string(contentAsBytes), task.Description)
}

func TestFilesystem_RetrieveAllByDate(t *testing.T) {

	setUp()

	tasksFilename := path.Join(configDir, now.Format(dateFormat))

	assert.FileExists(t, tasksFilename)

	modelTasks := underTest.RetrieveByDate(now)

	assert.Len(t, modelTasks, 1)

	assert.Contains(t, modelTasks, task)
}

func TestFilesystem_RemoveByDate(t *testing.T) {

	setUp()

	tasksFilename := path.Join(configDir, now.Format(dateFormat))

	assert.FileExists(t, tasksFilename)

	underTest.RemoveByDate(now)

	assert.False(t, underTest.directoryExists(tasksFilename))
}

func TestFilesystem_RemovePast(t *testing.T) {

	setUp()

	tasksFilename1 := path.Join(configDir, now.Format(dateFormat))

	tasksFilename2 := path.Join(configDir, oneYearInThePastDate.Format(dateFormat))

	tasksFilename3 := path.Join(configDir, oneMonthInThePastDate.Format(dateFormat))

	underTest.RemovePast()

	assert.True(t, underTest.directoryExists(tasksFilename1))

	assert.False(t, underTest.directoryExists(tasksFilename2))

	assert.False(t, underTest.directoryExists(tasksFilename3))
}
