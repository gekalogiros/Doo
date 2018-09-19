package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var now time.Time = time.Now()

func TestResolveDueDateWithInvalidDate(t *testing.T) {
	assertInvalidDueDate(t, "dsfhdsfiushdf", "failed to parse date")
	assertInvalidDueDate(t, "1235412421241243124213513464562341243124214321D", "failed to parse number of days, period provided is probably too long")
	assertInvalidDueDate(t, "5364367346523545676235345745462353473462356315", "failed to parse number of days, period provided is probably too long")
}

func TestResolveDueDateWithExpressionResolver(t *testing.T) {
	assertDueDate(t, "134d", now.AddDate(0, 0, 134))
	assertDueDate(t, "134D", now.AddDate(0, 0, 134))
	assertDueDate(t, "0134D", now.AddDate(0, 0, 134))
	assertDueDate(t, "1m", now.AddDate(0, 1, 0))
	assertDueDate(t, "1M", now.AddDate(0, 1, 0))
	assertDueDate(t, "10y", now.AddDate(10, 0, 0))
	assertDueDate(t, "10Y", now.AddDate(10, 0, 0))
}

func TestResolveDueDateWithDayResolver(t *testing.T) {
	assertDueDate(t, "134", now.AddDate(0, 0, 134))
	assertDueDate(t, "04", now.AddDate(0, 0, 4))
	assertDueDate(t, "020", now.AddDate(0, 0, 20))
}

func TestResolveDueDateWithDateResolver(t *testing.T) {
	expected := toDate(2018, 12, 1)
	assertDueDate(t, "1/12/2018", expected)
	assertDueDate(t, "01/12/2018", expected)
	assertDueDate(t, "1/12/18", expected)
	assertDueDate(t, "1-12-2018", expected)
	assertDueDate(t, "01-12-2018", expected)
	assertDueDate(t, "1-12-18", expected)
}

func assertInvalidDueDate(t *testing.T, input string, expected string){
	_, error := ResolveDueDate(input)

	assert.Error(t, error, expected)
}

func assertDueDate(t *testing.T, input string, expected time.Time){
	date, error := ResolveDueDate(input)

	assert.NoError(t, error)

	format := "2006-01-02"
	assert.Equal(t, date.Format(format), expected.Format(format))
}

func toDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
