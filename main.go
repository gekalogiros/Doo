package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ADD_SUB_COMMAND        = "add"
	ADD_DESCRIPTION_OPTION = "d"
	ADD_DUE_DATE_OPTION    = "dd"

	REMOVE_SUB_COMMAND  = "rm"
	REMOVE_DATE_OPTION  = "d"
	REMOVE_FORCE_OPTION = "f"
)

type AddCommandOptions struct {
	desc string
	date string
}

func (o AddCommandOptions) valid() bool {
	return o.desc != "" && o.date != ""
}

type RemoveCommandOptions struct {
	date  string
	force bool
}

func (o RemoveCommandOptions) valid() bool {
	return o.date != ""
}

func main() {
	addCommand := flag.NewFlagSet(ADD_SUB_COMMAND, flag.ExitOnError)
	todoDescriptionPointer := addCommand.String(ADD_DESCRIPTION_OPTION, "", "to-do description (Required)")
	todoDatePointer := addCommand.String(ADD_DUE_DATE_OPTION, "", "to-do due date (Required)")

	removeCommand := flag.NewFlagSet(REMOVE_SUB_COMMAND, flag.ExitOnError)
	removeDatePointer := removeCommand.String(REMOVE_DATE_OPTION, "", "Date of the note that you'd like to delete (Required)")
	removeForcePointer := removeCommand.Bool(REMOVE_FORCE_OPTION, false, "Force remove all notes for a certain date")

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("You need to Provide a command: %s, %s", ADD_SUB_COMMAND, REMOVE_SUB_COMMAND))
		os.Exit(1)
	}

	fmt.Println(os.Args)

	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
	case "rm":
		removeCommand.Parse(os.Args[2:])
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

		NewAdditionEnquiry(options.date, options.desc).execute()
	}

	if removeCommand.Parsed() {

		options := RemoveCommandOptions{date: *removeDatePointer, force: *removeForcePointer}

		if !options.valid() {
			removeCommand.PrintDefaults()
			os.Exit(2)
		}

		if options.force {
			NewPastDateRemovalEnquiry(options.date).execute()
		}
	}
}
