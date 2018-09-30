package formatter

import (
	"strconv"
)

const (
	reset   = "\033[0m"
	normal  = 0
	Bold    = 1
	Bolder  = 2
	Boldest = 3
	red     = 31
	green   = 32
	yellow  = 33
)

var display = func(colorCode int, emphasis int, content string) string {
	return "\033[" + strconv.Itoa(emphasis) + ";" + strconv.Itoa(colorCode) + "m" + content + reset
}

func Red(text string) string {
	return display(red, normal, text)
}

func Green(text string) string {
	return display(green, normal, text)
}

func Yellow(text string) string {
	return display(yellow, normal, text)
}

func BRed(text string, strength int) string {
	return display(red, strength, text)
}

func BGreen(text string, strength int) string {
	return display(green, strength, text)
}

func BYellow(text string, strength int) string {
	return display(yellow, strength, text)
}
