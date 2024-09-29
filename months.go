package mildtg

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	jan = "JAN"
	feb = "FEB"
	mar = "MAR"
	apr = "APR"
	may = "MAY"
	jun = "JUN"
	jul = "JUL"
	aug = "AUG"
	sep = "SEP"
	oct = "OCT"
	nov = "NOV"
	dec = "DEC"
)

var (
	// ErrInvalidMonth is returned when an invalid month is provided.
	ErrInvalidMonth = errors.New("invalid month")
	// ErrInvalidDay is returned when an invalid day is provided.
	ErrInvalidDay = errors.New("invalid day")
)

var (
	// monthRegex is a regular expression to match a three-letter month
	// or the full month name followed by zero or more numbers.
	// No non-numerical characters are allowed after the month (e.g., "JANUARYY").
	monthRegex = regexp.MustCompile(`(?i)` +
		`(JAN(uary)?|FEB(ruary)?|MAR(ch)?|APR(il)?|` +
		`MAY|JUN(e)?|JUL(y)?|AUG(ust)?|SEP(tember)?|` +
		`OCT(ober)?|NOV(ember)?|DEC(ember)?)([0-9\s]*)$`)

	// noMonthRegex is a regular expression to match a date-time-group
	// string that does not contain a month.
	noMonthRegex = regexp.MustCompile(`(?i)^(0[1-9]|2[0-9]|3[01])([0-1][0-9]|2[0-3])([0-5][0-9])([0-5][0-9])?[A-Z]\d{0,4}$`)
)

// monthMap maps the three-letter month abbreviation to the time.Month type.
var months = map[string]time.Month{
	jan: time.January,
	feb: time.February,
	mar: time.March,
	apr: time.April,
	may: time.May,
	jun: time.June,
	jul: time.July,
	aug: time.August,
	sep: time.September,
	oct: time.October,
	nov: time.November,
	dec: time.December,
}

// parseMonth returns the index of the month in the string
// if it exists.
// If there is no month (e.g., "270000Z"), the current month
// is returned.
// If there is an invalid month, (e.g., "270000Z JANUARYY"),
// an error is returned.
func parseMonth(s string) (string, time.Month, error) {
	if noMonthRegex.MatchString(s) {
		// No month in the string.
		return "", time.Now().UTC().Month(), nil
	}

	// Extract the month from the string.
	matches := monthRegex.FindStringSubmatch(s)
	if len(matches) == 0 {
		return "", 0, ErrInvalidMonth
	}

	// Extract the month abbreviation.
	m := strings.ToUpper(matches[1][:3])

	// Check if the month is valid.
	month, ok := months[m]
	if !ok {
		fmt.Printf("month: %s", m)
		return "", 0, ErrInvalidMonth
	}

	return matches[1], month, nil
}

// daysInMonth returns the number of days in a month and year.
func daysInMonth(m time.Month, year int) int {
	switch m {
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29
		}
		return 28
	case time.April, time.June, time.September, time.November:
		return 30
	case time.January, time.March, time.May, time.July,
		time.August, time.October, time.December:
		return 31
	default:
		return 0
	}
}
