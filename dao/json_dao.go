package dao

import (
	js "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/gekalogiros/Doo/model"
)

type json struct {
	configDir string
	filename  string
}

type taskJSON struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewJSONDao() TaskDao {

	configDir := path.Join(os.Getenv("HOME"), ".doo")

	return newJSONDao(configDir)
}

func newJSONDao(configFile string) json {
	return json{
		configDir: configFile,
		filename:  "tasks.json",
	}
}

func (db json) Save(task *model.Task) {

	err := db.init()

	if err != nil {
		log.Fatal("Failed to create task list db")
	}

	db.update(db.dateToString(task.Date), db.adaptFromTask(*task))
}

func (db json) Move(taskID string, from time.Time, to time.Time) {

	if db.exists() {

		fromDate := db.dateToString(from)

		toDate := db.dateToString(to)

		db.moveByID(fromDate, toDate, taskID)
	}
}

func (db json) RemoveByDate(date time.Time) {

	if db.exists() {

		targetDate := db.dateToString(date)

		db.remove(targetDate)
	}
}

func (db json) RetrieveByDate(date time.Time) []model.Task {

	if db.exists() {

		targetDate := db.dateToString(date)

		return db.adaptToTaskList(db.findAll()[targetDate], date)
	}

	return make([]model.Task, 0)
}

func (db json) RemovePast() {

	if db.exists() {

		location, _ := time.LoadLocation("Local")

		now := time.Now().In(location).Truncate(24 * time.Hour)

		all := db.findAll()

		for dateAsString := range all {

			date := db.dateFromString(dateAsString)

			if date.Before(now) {
				delete(all, dateAsString)
			}
		}

		db.persistAll(all)
	}
}

func (db json) findAll() map[string][]taskJSON {

	dbPath := db.path()

	dbFile, _ := os.Open(dbPath)

	defer dbFile.Close()

	byteValue, _ := ioutil.ReadAll(dbFile)

	var tasksPerDate map[string][]taskJSON

	js.Unmarshal(byteValue, &tasksPerDate)

	return tasksPerDate
}

func (db json) update(date string, task taskJSON) {

	taskList := db.findAll()

	tasksForDate := taskList[date]

	if tasksForDate == nil {
		tasksForDate = make([]taskJSON, 0)
	}

	tasksForDate = append(tasksForDate, task)

	taskList[date] = tasksForDate

	db.persistAll(taskList)
}

func (db json) remove(date string) {

	taskList := db.findAll()

	delete(taskList, date)

	db.persistAll(taskList)
}

func (db json) moveByID(fromDate string, toDate string, id string) {

	taskList := db.findAll()

	var index int
	var taskJSON taskJSON

	for i, task := range taskList[fromDate] {
		if task.ID == id {
			index = i
			taskJSON = task
			break
		}
	}

	taskList[fromDate] = remove(taskList[fromDate], index)

	taskList[toDate] = append(taskList[toDate], taskJSON)

	db.persistAll(taskList)
}

func (db json) persistAll(tasks map[string][]taskJSON) {

	updatedJSON, err := js.Marshal(tasks)

	if isError(err) {
		return
	}

	ioutil.WriteFile(db.path(), updatedJSON, 0644)
}

func remove(slice []taskJSON, s int) []taskJSON {
	return append(slice[:s], slice[s+1:]...)
}

func (db json) adaptFromTask(task model.Task) taskJSON {
	return taskJSON{
		ID:          task.Id,
		Description: task.Description,
	}
}

func (db json) adaptToTask(task taskJSON, date time.Time) model.Task {
	location, _ := time.LoadLocation("Local")
	return model.Task{
		Id:          task.ID,
		Description: task.Description,
		Date:        date.In(location).Truncate(24 * time.Hour),
	}
}

func (db json) adaptToTaskList(tasks []taskJSON, date time.Time) []model.Task {

	if tasks == nil {
		tasks = make([]taskJSON, 0)
	}

	data := make([]model.Task, 0)

	for _, element := range tasks {

		data = append(data, db.adaptToTask(element, date))

	}

	return data
}

func (db json) init() error {

	if !db.configDirExists() {
		err := os.MkdirAll(db.configDir, os.ModePerm)
		if isError(err) {
			return err
		}
	}

	if !db.exists() {

		file, err := os.Create(db.path())

		if isError(err) {
			fmt.Println(err)
			return err
		}

		defer file.Close()

		_, err = file.WriteString("{}")
		if isError(err) {
			return err
		}

		err = file.Sync()
		if isError(err) {
			return err
		}
	}

	return nil
}

func (db json) dateToString(date time.Time) string {
	return date.Format("20060102")
}

func (db json) dateFromString(dateAsString string) time.Time {

	date, _ := time.Parse("20060102", dateAsString)

	location, _ := time.LoadLocation("Local")

	return date.In(location).Truncate(24 * time.Hour)
}

func (db json) path() string {
	return path.Join(db.configDir, db.filename)
}

func (db json) exists() bool {
	return resourceExists(db.path())
}

func (db json) configDirExists() bool {
	return resourceExists(db.configDir)
}

func resourceExists(path string) bool {

	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func isError(err error) bool {

	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
