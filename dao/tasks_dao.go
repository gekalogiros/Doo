package dao

import (
	"fmt"
	"github.com/gekalogiros/Doo/model"
	"log"
	"os"
	"path"
	"time"
)

type TaskDao interface {
	Save(n *model.Task)
	RemoveAll(date time.Time)
}

type filesystem struct {
	configDir  string
	fileFormat string
}

func NewFileSystemTasksDao() TaskDao {
	configDir := path.Join(os.Getenv("HOME"), ".doo")
	return newFilesystemDao(configDir)
}

func newFilesystemDao(configFile string) filesystem {
	return filesystem{
		configDir:  configFile,
		fileFormat: "02_01_2006",
	}
}

func (f filesystem) ensureConfigDirectoryIsPresent() {
	if _, err := os.Stat(f.configDir); os.IsNotExist(err) {
		err = os.MkdirAll(f.configDir, 0755)
		if err != nil {
			log.Fatal("Cannot create or access config directory")
		}
	}
}

func (f filesystem) Save(n *model.Task) {

	f.ensureConfigDirectoryIsPresent()

	file := path.Join(f.configDir, n.Date.Format(f.fileFormat))

	fd, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot create directory for storing task %s", file))
	}

	defer fd.Close()

	if _, err = fd.WriteString(fmt.Sprintf("%s,%s", n.Id, n.Description)); err != nil {
		log.Fatal(fmt.Sprintf("Cannot write task %s", n))
	}
}

func (f filesystem) RemoveAll(date time.Time) {
	// TODO
}
