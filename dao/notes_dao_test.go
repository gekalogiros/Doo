package dao

import (
	"github.com/gekalogiros/todo/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
)

var now = time.Now()

func TestFilesystem_Save(t *testing.T) {

	configDir := os.TempDir()

	dao := newFilesystemDao(configDir)

	task := model.NewTask("I am a test task", now)

	dao.Save(&task)

	notesFilename := path.Join(configDir, now.Format("02_01_2006"))

	assert.FileExists(t, notesFilename)

	contentAsBytes, error := ioutil.ReadFile(notesFilename)

	assert.NoError(t, error)

	assert.Contains(t, string(contentAsBytes), task.Description)
}
