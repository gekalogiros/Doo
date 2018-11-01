package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gekalogiros/Doo/commands"
)

const (
	errorMessageTemplate = "%s, Check documentation at github.com/gekalogiros/Doo"

	AddSubCommand        = "add"
	AddDescriptionOption = "t"
	AddDueDateOption     = "d"

	RemoveSubCommand = "rm"
	RemoveDateOption = "d"
	RemovePastOption = "past"

	ListSubCommand = "ls"
	ListDateOption = "d"

	MoveSubCommand     = "mv"
	MoveTaskIDOption   = "id"
	MoveFromDateOption = "f"
	MoveToDateOption   = "t"
)

var (
	version = "unspecified"
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

type MoveCommandOptions struct {
	id   *string
	from *string
	to   *string
}

func (m MoveCommandOptions) valid() bool {
	return *m.id != "" && *m.from != "" && *m.to != ""
}

type RemoveCommandOptions struct {
	date    *string
	allPast *bool
}

func (o RemoveCommandOptions) valid() bool {
	return (*o.date != "" && *o.allPast != true) || (*o.date == "" && *o.allPast == true)
}

type ListCommandOptions struct {
	date *string
}

func (l ListCommandOptions) valid() bool {
	return *l.date != ""
}

func main() {

	addCommand := flag.NewFlagSet(AddSubCommand, flag.ExitOnError)
	todoDescriptionPointer := addCommand.String(AddDescriptionOption, "", "task description")
	todoDatePointer := addCommand.String(AddDueDateOption, "today", "task due date")

	moveCommand := flag.NewFlagSet(MoveSubCommand, flag.ExitOnError)
	moveIDPointer := moveCommand.String(MoveTaskIDOption, "", "task id to move (Required)")
	moveFromDatePointer := moveCommand.String(MoveFromDateOption, "", "date of source task list (Required)")
	moveToDatePointer := moveCommand.String(MoveToDateOption, "", "date of target task list (Required)")

	removeCommand := flag.NewFlagSet(RemoveSubCommand, flag.ExitOnError)
	removeDatePointer := removeCommand.String(RemoveDateOption, "", "Date of the task that you'd like to delete")
	removePastPointer := removeCommand.Bool(RemovePastOption, false, "Remove all past task lists")

	listCommand := flag.NewFlagSet(ListSubCommand, flag.ExitOnError)
	listDatePointer := listCommand.String(ListDateOption, "today", "Date of the task list you'd like to see information for")

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("You need to Provide a command: %s, %s, %s, %s", AddSubCommand, MoveSubCommand, ListSubCommand, RemoveSubCommand))
		os.Exit(1)
	}

	switch os.Args[1] {
	case AddSubCommand:
		addCommand.Parse(os.Args[2:])
		break
	case MoveSubCommand:
		moveCommand.Parse(os.Args[2:])
		break
	case RemoveSubCommand:
		removeCommand.Parse(os.Args[2:])
		break
	case ListSubCommand:
		listCommand.Parse(os.Args[2:])
		break
	case "--version":
		fmt.Println("Doo version " + version)
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

	if moveCommand.Parsed() {

		options := MoveCommandOptions{id: moveIDPointer, from: moveFromDatePointer, to: moveToDatePointer}

		checkArgumentsAndExecute(*moveCommand, options, commands.NewTaskMovement(*options.id, *options.from, *options.to))
	}

	if removeCommand.Parsed() {

		options := RemoveCommandOptions{date: removeDatePointer, allPast: removePastPointer}

		checkArgumentsAndExecute(*removeCommand, options, commands.NewTaskListRemoval(*options.date, *options.allPast))
	}

	if listCommand.Parsed() {

		args := listCommand.Args()

		if len(args) == 1 && flag.Lookup(ListDateOption) == nil {
			listDatePointer = &args[0]
		}

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
