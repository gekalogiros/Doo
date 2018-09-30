package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type temporalFunc func(int) int

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

	futureDayLookupIncludingToday = map[string]string{
		"today":    "0",
		"tomorrow": "1",
	}
	pastDayLookupIncludingToday = map[string]string{
		"yesterday": "-1",
		"today":     "0",
	}

	identity temporalFunc = func(period int) int {
		return period
	}
	inverse temporalFunc = func(period int) int {
		return -period
	}
)

func ResolveDueDate(dueDate string) (time.Time, error) {
	switch {
	case futureDayRegex.MatchString(dueDate):
		return resolveByPeriod(futureDayLookupIncludingToday[dueDate], identity)
	case futureTemporalRegex.MatchString(dueDate):
		return resolveByTemporal(dueDate, identity)
	case futurePeriodRegex.MatchString(dueDate):
		return resolveByPeriod(dueDate, identity)
	default:
		return resolveByDate(dueDate)
	}
}

func ResolveDate(date string) (time.Time, error) {
	switch {
	case futureDayRegex.MatchString(date):
		return resolveByPeriod(futureDayLookupIncludingToday[date], identity)
	case futureTemporalRegex.MatchString(date):
		return resolveByTemporal(date, identity)
	case futurePeriodRegex.MatchString(date):
		return resolveByPeriod(date, identity)
	case pastDayRegex.MatchString(date):
		return resolveByPeriod(pastDayLookupIncludingToday[date], identity)
	case pastTemporalRegex.MatchString(date):
		return resolveByTemporal(date, inverse)
	case pastPeriodRegex.MatchString(date):
		return resolveByPeriod(date, inverse)
	default:
		return resolveByDate(date)
	}
}

func resolveByTemporal(expression string, temporalFunc temporalFunc) (time.Time, error) {

	today := time.Now()

	period, err := toInt(expression[:len(expression)-1])
	if err != nil {
		return time.Now(), err
	}

	periodType := strings.ToLower(expression[len(expression)-1:])

	switch periodType {
	case "d":
		return today.AddDate(0, 0, temporalFunc(period)), nil
	case "m":
		return today.AddDate(0, temporalFunc(period), 0), nil
	default:
		return today, fmt.Errorf("failed to resolve expression: %s", expression)
	}
}

func resolveByPeriod(numberOfDays string, temporalFunc temporalFunc) (time.Time, error) {

	today := time.Now()

	numberOfDaysAsInt, err := toInt(numberOfDays)

	if err != nil {
		return today, fmt.Errorf("failed to parse number of days, period provided is probably too long: %s", numberOfDays)
	}

	return today.AddDate(0, 0, temporalFunc(numberOfDaysAsInt)), nil
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
