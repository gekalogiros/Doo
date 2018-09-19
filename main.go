package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	AddSubCommand        = "add"
	AddDescriptionOption = "d"
	AddDueDateOption     = "dd"

	RemoveSubCommand = "rm"
	RemoveDateOption = "dt"

	ListSubCommnad = "ls"
	ListDateOption = "dt"
)

type AddCommandOptions struct {
	desc string
	date string
}

func (o AddCommandOptions) valid() bool {
	return o.desc != "" && o.date != ""
}

type RemoveCommandOptions struct {
	date string
}

func (o RemoveCommandOptions) valid() bool {
	return o.date != ""
}

type ListCommandOptions struct {
	date string
}

func (l ListCommandOptions) valid() bool {
	if l.date == "" {
		return false
	}

	if _, err := time.Parse("02-01-2006", l.date); err != nil {
		return false
	}

	return true
}

func main() {
	addCommand := flag.NewFlagSet(AddSubCommand, flag.ExitOnError)
	todoDescriptionPointer := addCommand.String(AddDescriptionOption, "", "task description (Required)")
	todoDatePointer := addCommand.String(AddDueDateOption, "", "task due date (Required)")

	removeCommand := flag.NewFlagSet(RemoveSubCommand, flag.ExitOnError)
	removeDatePointer := removeCommand.String(RemoveDateOption, "", "Date of the task that you'd like to delete (Required)")

	listCommand := flag.NewFlagSet(ListSubCommnad, flag.ExitOnError)
	listDatePointer := listCommand.String(ListDateOption, "", "Date of the task list you'd like to see information for (Required)")

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("You need to Provide a command: %s, %s, %s", AddSubCommand, ListSubCommnad, RemoveSubCommand))
		os.Exit(1)
	}

	switch os.Args[1] {
	case AddSubCommand:
		addCommand.Parse(os.Args[2:])
	case RemoveSubCommand:
		removeCommand.Parse(os.Args[2:])
	case ListSubCommnad:
		listCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCommand.Parsed() {

		options := AddCommandOptions{date: *todoDatePointer, desc: *todoDescriptionPointer}

		if !options.valid() {
			addCommand.PrintDefaults()
			os.Exit(1)
		}

		NewTaskCreation(options.date, options.desc).execute()
	}

	if removeCommand.Parsed() {

		options := RemoveCommandOptions{date: *removeDatePointer}

		if !options.valid() {
			removeCommand.PrintDefaults()
			os.Exit(2)
		}

		NewTaskListRemoval(options.date).execute()
	}

	if listCommand.Parsed() {
		options := ListCommandOptions{date: *listDatePointer}

		if !options.valid() {
			listCommand.PrintDefaults()
			os.Exit(3)
		}

		NewTaskListRetrieval(options.date).execute()
	}
}
