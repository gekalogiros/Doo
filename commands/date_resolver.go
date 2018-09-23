package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	timeUnitRegex      = regexp.MustCompile("^[0-9]+(d|m|y|D|M|Y)$")
	periodRegex        = regexp.MustCompile("^[0-9]+$")
	allowedDateFormats = [...]string{
		"02/01/2006", "2/1/2006", "02/01/06", "2/1/06",
		"02-01-2006", "2-1-2006", "02-01-06", "2-1-06"}
)

func ResolveDueDate(dueDate string) (time.Time, error) {
	switch {
	case timeUnitRegex.MatchString(dueDate):
		return resolveByExpression(dueDate)
	case periodRegex.MatchString(dueDate):
		return resolveByNumber(dueDate)
	default:
		return resolveByDate(dueDate)
	}
}

func resolveByExpression(expression string) (time.Time, error) {

	today := time.Now()

	period, err := toInt(expression[:len(expression)-1])
	if err != nil {
		return time.Now(), err
	}

	periodType := strings.ToLower(expression[len(expression)-1:])

	switch periodType {
	case "d":
		return today.AddDate(0, 0, period), nil
	case "m":
		return today.AddDate(0, period, 0), nil
	case "y":
		return today.AddDate(period, 0, 0), nil
	default:
		return today, fmt.Errorf("failed to resolve expression: %s", expression)
	}
}

func resolveByDate(date string) (time.Time, error) {
	for _, element := range allowedDateFormats {
		t, err := time.Parse(element, date)
		if err == nil {
			return t, nil
		}
	}

	return time.Now(), fmt.Errorf("failed to parse date: %s", date)
}

func resolveByNumber(numberOfDays string) (time.Time, error) {

	today := time.Now()

	numberOfDaysAsInt, error := toInt(numberOfDays)

	if error != nil {
		return today, fmt.Errorf("failed to parse number of days, period provided is probably too long: %s", numberOfDays)
	}

	return today.AddDate(0, 0, numberOfDaysAsInt), nil
}

func toInt(time string) (int, error) {
	i, err := strconv.Atoi(time)
	if err != nil {
		return 0, fmt.Errorf("time period %s is too long", time)
	}
	return i, nil
}
