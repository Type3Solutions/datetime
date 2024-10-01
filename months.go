package mildtg

import (
	"errors"
	"strings"
	"time"
)

var (
	// ErrInvalidMonth is returned when an invalid month is provided.
	ErrInvalidMonth = errors.New("invalid month")
	// ErrInvalidDay is returned when an invalid day is provided.
	ErrInvalidDay = errors.New("invalid day")
)

// monthMap maps the three-letter month abbreviation to the time.Month type.
var (
	months map[string]time.Month
)

func init() {
	months = make(map[string]time.Month)

	for _, m := range []time.Month{
		time.January, time.February, time.March, time.April, time.May, time.June,
		time.July, time.August, time.September, time.October, time.November, time.December,
	} {
		months[strings.ToUpper(m.String()[:3])] = m
		months[strings.ToUpper(m.String())] = m
	}
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
