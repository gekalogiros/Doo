package commands

import "github.com/gekalogiros/Doo/dao"

var tasksDao = dao.NewFileSystemTasksDao() // needs DI

type Command interface {
	Execute() error
}

func setDao(dao dao.TaskDao) {
	tasksDao = dao
}
