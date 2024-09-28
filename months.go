package mildtg

import (
	"errors"
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
)

var (
	janRegex = regexp.MustCompile(`JAN|JANUARY`)
	febRegex = regexp.MustCompile(`FEB|FEBRUARY`)
	marRegex = regexp.MustCompile(`MAR|MARCH`)
	aprRegex = regexp.MustCompile(`APR|APRIL`)
	mayRegex = regexp.MustCompile(`MAY`)
	junRegex = regexp.MustCompile(`JUN|JUNE`)
	julRegex = regexp.MustCompile(`JUL|JULY`)
	augRegex = regexp.MustCompile(`AUG|AUGUST`)
	sepRegex = regexp.MustCompile(`SEP|SEPTEMBER`)
	octRegex = regexp.MustCompile(`OCT|OCTOBER`)
	novRegex = regexp.MustCompile(`NOV|NOVEMBER`)
	decRegex = regexp.MustCompile(`DEC|DECEMBER`)
)

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

// parseMonth returns the month from a date-time-group string.
func parseMonth(s string) (time.Month, error) {
	s = strings.ToUpper(removeSpaces(s))

	switch {
	case janRegex.MatchString(s):
		return time.January, nil
	case febRegex.MatchString(s):
		return time.February, nil
	case marRegex.MatchString(s):
		return time.March, nil
	case aprRegex.MatchString(s):
		return time.April, nil
	case mayRegex.MatchString(s):
		return time.May, nil
	case junRegex.MatchString(s):
		return time.June, nil
	case julRegex.MatchString(s):
		return time.July, nil
	case augRegex.MatchString(s):
		return time.August, nil
	case sepRegex.MatchString(s):
		return time.September, nil
	case octRegex.MatchString(s):
		return time.October, nil
	case novRegex.MatchString(s):
		return time.November, nil
	case decRegex.MatchString(s):
		return time.December, nil
	default:
		return 0, ErrInvalidMonth
	}
}

// month returns the month from a date-time-group string.
func month(s string) (time.Month, error) {
	s = removeSpaces(strings.ToUpper(s))
	return parseMonth(s)
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
