package dao

import (
	"fmt"
	"github.com/gekalogiros/Doo/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
)

var now = time.Now()
var configDir = fmt.Sprintf("%s%d", os.TempDir(), now.UnixNano() / int64(time.Millisecond))
var task = model.Task{Id:"xxxx", Description:"I am a test task", Date:now}
var underTest = newFilesystemDao(configDir)

func init(){
	underTest.Save(&task)
}

func TestFilesystem_Save(t *testing.T) {

	tasksFilename := path.Join(configDir, now.Format("02_01_2006"))

	assert.FileExists(t, tasksFilename)

	contentAsBytes, err := ioutil.ReadFile(tasksFilename)

	assert.NoError(t, err)

	assert.Contains(t, string(contentAsBytes), task.Description)
}

func TestFilesystem_RetrieveAllByDate(t *testing.T) {

	tasksFilename := path.Join(configDir, now.Format("02_01_2006"))

	assert.FileExists(t, tasksFilename)

	modelTasks := underTest.RetrieveAllByDate(now)

	assert.Len(t, modelTasks, 1)

	assert.Contains(t, modelTasks, task)
}
