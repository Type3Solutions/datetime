package mildtg

import (
	"errors"
	"time"
)

const (
	jan          = "JAN"
	feb          = "FEB"
	mar          = "MAR"
	apr          = "APR"
	may          = "MAY"
	jun          = "JUN"
	jul          = "JUL"
	aug          = "AUG"
	sep          = "SEP"
	oct          = "OCT"
	nov          = "NOV"
	dec          = "DEC"
	january      = "JANUARY"
	february     = "FEBRUARY"
	march        = "MARCH"
	april        = "APRIL"
	june         = "JUNE"
	july         = "JULY"
	august       = "AUGUST"
	september    = "SEPTEMBER"
	october      = "OCTOBER"
	novemberLong = "NOVEMBER"
	december     = "DECEMBER"
)

var (
	// ErrInvalidMonth is returned when an invalid month is provided.
	ErrInvalidMonth = errors.New("invalid month")
	// ErrInvalidDay is returned when an invalid day is provided.
	ErrInvalidDay = errors.New("invalid day")
)

// monthMap maps the three-letter month abbreviation to the time.Month type.
var months = map[string]time.Month{
	jan:          time.January,
	feb:          time.February,
	mar:          time.March,
	apr:          time.April,
	may:          time.May,
	jun:          time.June,
	jul:          time.July,
	aug:          time.August,
	sep:          time.September,
	oct:          time.October,
	nov:          time.November,
	dec:          time.December,
	january:      time.January,
	february:     time.February,
	march:        time.March,
	april:        time.April,
	june:         time.June,
	july:         time.July,
	august:       time.August,
	september:    time.September,
	october:      time.October,
	novemberLong: time.November,
	december:     time.December,
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
