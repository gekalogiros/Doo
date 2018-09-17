package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	AddSubCommand        = "add"
	AddDescriptionOption = "d"
	AddDueDateOption     = "dd"

	RemoveSubCommand  = "rm"
	RemoveDateOption  = "d"
	RemoveForceOption = "f"
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
	addCommand := flag.NewFlagSet(AddSubCommand, flag.ExitOnError)
	todoDescriptionPointer := addCommand.String(AddDescriptionOption, "", "to-do description (Required)")
	todoDatePointer := addCommand.String(AddDueDateOption, "", "to-do due date (Required)")

	removeCommand := flag.NewFlagSet(RemoveSubCommand, flag.ExitOnError)
	removeDatePointer := removeCommand.String(RemoveDateOption, "", "Date of the note that you'd like to delete (Required)")
	removeForcePointer := removeCommand.Bool(RemoveForceOption, false, "Force remove all notes for a certain date")

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("You need to Provide a command: %s, %s", AddSubCommand, RemoveSubCommand))
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
