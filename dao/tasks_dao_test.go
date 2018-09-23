package dao

import (
	"github.com/gekalogiros/Doo/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
)

var now = time.Now()
var configDir = os.TempDir()
var task = model.NewTask("I am a test task", now)
var underTest = newFilesystemDao(configDir)

func TestFilesystem_Save(t *testing.T) {

	underTest.Save(&task)

	notesFilename := path.Join(configDir, now.Format("02_01_2006"))

	assert.FileExists(t, notesFilename)

	contentAsBytes, err := ioutil.ReadFile(notesFilename)

	assert.NoError(t, err)

	assert.Contains(t, string(contentAsBytes), task.Description)
}
