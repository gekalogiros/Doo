package commands

import (
	"fmt"
	"log"
)

const (
	errorMessageTemplate = "%s, Check documentation at github.com/gekalogiros/Doo"
)

type Command interface {
	Execute()
}

func failIfError(err error, message string){
	if err != nil {
		log.Fatal(fmt.Sprintf(errorMessageTemplate, message))
	}
}
