package dao

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gekalogiros/Doo/model"
)

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
	if !f.configDirectoryExists() {
		err := os.MkdirAll(f.configDir, 0755)
		if err != nil {
			log.Fatal("Cannot create or access config directory. ", err)
		}
	}
}

func (f filesystem) configDirectoryExists() bool {
	return f.directoryExists(f.configDir)
}

func (f filesystem) directoryExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
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

	if _, err = fd.WriteString(fmt.Sprintf("%s,%s\n", n.Id, n.Description)); err != nil {
		log.Fatal(fmt.Sprintf("Cannot write task %s", n))
	}
}

func (f filesystem) RemoveByDate(date time.Time) {
	taskListPath := path.Join(f.configDir, date.Format(f.fileFormat))
	os.RemoveAll(taskListPath)
}

func (f filesystem) RemovePast() {

	if files, err := filepath.Glob(f.configDir + "/*"); err != nil {

		log.Fatal(err)

	} else {

		for _, element := range files {

			base := filepath.Base(element)

			if taskListDate, e := time.Parse(f.fileFormat, base); e != nil {

				log.Fatal(e)

			} else {

				now := time.Now().Truncate(24 * time.Hour)

				location, _ := time.LoadLocation("Local")

				taskListDate = taskListDate.In(location)

				if taskListDate.Before(now) {
					os.RemoveAll(element)
				}
			}
		}
	}
}

func (f filesystem) RetrieveByDate(date time.Time) []model.Task {
	taskListPath := path.Join(f.configDir, date.Format(f.fileFormat))
	if f.configDirectoryExists() && f.directoryExists(taskListPath) {
		if lines, err := readLines(taskListPath); err == nil {
			tasks := make([]model.Task, len(lines))
			for i, element := range lines {
				lineSplit := strings.Split(element, ",")
				task := model.Task{Id: lineSplit[0], Description: lineSplit[1], Date: date}
				tasks[i] = task
			}
			return tasks
		}
	}
	return []model.Task{}
}

func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
