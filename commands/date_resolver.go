package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	futureDayRegex      = regexp.MustCompile("^today|tomorrow$")
	futureTemporalRegex = regexp.MustCompile("^[0-9]+(d|m|D|M)$")
	futurePeriodRegex   = regexp.MustCompile("^[0-9]+$")
	pastDayRegex        = regexp.MustCompile("^yesterday$")
	pastTemporalRegex   = regexp.MustCompile("^-[0-9]+(d|m|D|M)$")
	pastPeriodRegex     = regexp.MustCompile("^-[0-9]+$")

	allowedDateFormats = [...]string{
		"02/01/2006", "2/1/2006", "02/01/06", "2/1/06",
		"02-01-2006", "2-1-2006", "02-01-06", "2-1-06",
	}

	dayLookup = map[string]string{
		"yesterday": "-1",
		"today":     "0",
		"tomorrow":  "1",
	}
)

func ResolveDueDate(dueDate string) (time.Time, error) {
	switch {
	case futureDayRegex.MatchString(dueDate):
		return resolveByPeriod(dayLookup[dueDate])
	case futureTemporalRegex.MatchString(dueDate):
		return resolveByTemporal(dueDate)
	case futurePeriodRegex.MatchString(dueDate):
		return resolveByPeriod(dueDate)
	default:
		return resolveByDate(dueDate)
	}
}

func ResolveDate(date string) (time.Time, error) {
	switch {
	case futureDayRegex.MatchString(date) || pastDayRegex.MatchString(date):
		return resolveByPeriod(dayLookup[date])
	case futureTemporalRegex.MatchString(date):
		return resolveByTemporal(date)
	case futurePeriodRegex.MatchString(date):
		return resolveByPeriod(date)
	case pastTemporalRegex.MatchString(date):
		return resolveByTemporal(date)
	case pastPeriodRegex.MatchString(date):
		return resolveByPeriod(date)
	default:
		return resolveByDate(date)
	}
}

func resolveByTemporal(expression string) (time.Time, error) {

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
	default:
		return today, fmt.Errorf("failed to resolve expression: %s", expression)
	}
}

func resolveByPeriod(numberOfDays string) (time.Time, error) {

	today := time.Now()

	numberOfDaysAsInt, err := toInt(numberOfDays)

	if err != nil {
		return today, fmt.Errorf("failed to parse number of days, period provided is probably too long: %s", numberOfDays)
	}

	return today.AddDate(0, 0, numberOfDaysAsInt), nil
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

func toInt(time string) (int, error) {
	i, err := strconv.Atoi(time)
	if err != nil {
		return 0, fmt.Errorf("time period %s is too long", time)
	}
	return i, nil
}
