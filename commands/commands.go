package commands

import "github.com/gekalogiros/Doo/dao"

var tasksDao = dao.NewJSONDao() // needs DI

type Command interface {
	Execute() error
}

func setDao(dao dao.TaskDao) {
	tasksDao = dao
}
