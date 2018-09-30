package main

import (
	"flag"
	"fmt"
	"github.com/gekalogiros/Doo/commands"
	"log"
	"os"
)

const (
	errorMessageTemplate = "%s, Check documentation at github.com/gekalogiros/Doo"

	AddSubCommand        = "add"
	AddDescriptionOption = "d"
	AddDueDateOption     = "dd"

	RemoveSubCommand = "rm"
	RemoveDateOption = "dt"

	ListSubCommnad = "ls"
	ListDateOption = "dt"
)

type Options interface {
	valid() bool
}

type AddCommandOptions struct {
	desc *string
	date *string
}

func (o AddCommandOptions) valid() bool {
	return *o.desc != "" && *o.date != ""
}

type RemoveCommandOptions struct {
	date *string
}

func (o RemoveCommandOptions) valid() bool {
	return *o.date != ""
}

type ListCommandOptions struct {
	date *string
}

func (l ListCommandOptions) valid() bool {
	return *l.date != ""
}

func main() {

	addCommand := flag.NewFlagSet(AddSubCommand, flag.ExitOnError)
	todoDescriptionPointer := addCommand.String(AddDescriptionOption, "", "task description (Required)")
	todoDatePointer := addCommand.String(AddDueDateOption, "today", "task due date")

	removeCommand := flag.NewFlagSet(RemoveSubCommand, flag.ExitOnError)
	removeDatePointer := removeCommand.String(RemoveDateOption, "", "Date of the task that you'd like to delete (Required)")

	listCommand := flag.NewFlagSet(ListSubCommnad, flag.ExitOnError)
	listDatePointer := listCommand.String(ListDateOption, "today", "Date of the task list you'd like to see information for (Required)")

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("You need to Provide a command: %s, %s, %s", AddSubCommand, ListSubCommnad, RemoveSubCommand))
		os.Exit(1)
	}

	switch os.Args[1] {
	case AddSubCommand:
		addCommand.Parse(os.Args[2:])
		break
	case RemoveSubCommand:
		removeCommand.Parse(os.Args[2:])
		break
	case ListSubCommnad:
		listCommand.Parse(os.Args[2:])
		break
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCommand.Parsed() {

		args := addCommand.Args()

		if len(args) == 1 && flag.Lookup(AddDescriptionOption) == nil {
			todoDescriptionPointer = &args[0]
		}

		options := AddCommandOptions{date: todoDatePointer, desc: todoDescriptionPointer}

		checkArgumentsAndExecute(*addCommand, options, commands.NewTaskCreation(*options.date, *options.desc))
	}

	if removeCommand.Parsed() {

		options := RemoveCommandOptions{date: removeDatePointer}

		checkArgumentsAndExecute(*removeCommand, options, commands.NewTaskListRemoval(*options.date))
	}

	if listCommand.Parsed() {

		options := ListCommandOptions{date: listDatePointer}

		checkArgumentsAndExecute(*listCommand, options, commands.NewTaskListRetrieval(*options.date))
	}

	os.Exit(0)
}

func checkArgumentsAndExecute(flagSet flag.FlagSet, options Options, command commands.Command) {

	if !options.valid() {
		flagSet.PrintDefaults()
		os.Exit(3)
	}

	executeOrFail(command)
}

func executeOrFail(command commands.Command) {

	if err := command.Execute(); err != nil {
		log.Fatal(fmt.Sprintf(errorMessageTemplate, err.Error()))
	}
}
