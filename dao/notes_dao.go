package dao

import (
	"fmt"
	"github.com/gekalogiros/todo/model"
	"log"
	"os"
	"path"
	"time"
)

type NotesDao interface {
	Save(n *model.Note)
	RemoveAll(date time.Time)
}

type filesystem struct {
	configDir string
}

func NewFileSystemNotesDao() NotesDao {
	configDir := path.Join(os.Getenv("HOME"), ".doo")
	return newFilesystemDao(configDir)
}

func newFilesystemDao(configFile string) filesystem {
	return filesystem{
		configDir: configFile,
	}
}

func (f filesystem) Save(n *model.Note)  {

	if _, err := os.Stat(f.configDir); os.IsNotExist(err) {
		err = os.MkdirAll(f.configDir, 0755)
		if err != nil {
			log.Fatal("Cannot create or access config directory")
		}
	}

	file := path.Join(f.configDir, n.Date.Format("02_01_2006"))

	fd, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot create directory for storing notes %s", file))
	}

	defer fd.Close()

	if _, err = fd.WriteString(fmt.Sprintf("%s,%s", n.Id, n.Description)); err != nil {
		log.Fatal(fmt.Sprintf("Cannot write note %s", n))
	}
}

func (f filesystem) RemoveAll(date time.Time)  {
	if _, err := os.Stat(f.configDir); os.IsNotExist(err) {

			log.Fatal("Cannot create or access config directory")

	}
}






