package mildtg

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	// ErrNotEnoughChars is returned when an invalid date-time-group is provided.
	ErrNotEnoughChars = errors.New("date-time-group too short minimum is ddhhmm")
)

var (
	secondsRegex = regexp.MustCompile(`[0-5][0-9]`)
)

// Time wraps a time.Time to allow for custom
// formatting and parsing of various U.S. military
// date-time-group formats.
type Time struct {
	time.Time
}

// String returns the date-time-group in the format
func (t Time) String() string {
	return t.Format("200601021504")
}

// Format returns the date-time-group in the format
func (t Time) Format(layout string) string {
	return t.Time.Format(layout)
}

// NewTime returns a new Time object.
func NewTime(t time.Time) Time {
	return Time{t}
}

// ParseDTG parses a military date-time-group string in the format
// DDHH[MM]|[MMSS]|(A-Z)[ MMM YY[YY] and returns a Time object.
func ParseDTG(s string) (Time, error) {
	// Remove all spaces from the string.
	s = removeSpaces(s)

	// The minimum length of a date-time-group string is six (DDHHMM).
	if len(s) < 6 {
		return Time{}, ErrNotEnoughChars
	}

	day := int(s[0]-'0')*10 + int(s[1]-'0')
	hour := int(s[2]-'0')*10 + int(s[3]-'0')
	minute := int(s[4]-'0')*10 + int(s[5]-'0')

	second := 0
	tz := ZULU
	i := 6
	// Check for the optional second field (SS).
	if len(s) >= 8 {
		secondsStr := s[i : i+2]
		if secondsRegex.MatchString(secondsStr) {
			second = int(s[i]-'0')*10 + int(s[i+1]-'0')
			i += 2
		} else {
			// No seconds field, check for the time zone designation
			// at the current index.
			second = 0
		}
	}

	_, m, err := parseMonth(s)
	if err != nil {
		return Time{}, err
	}

	year, err := parseYear(s)
	if err != nil {
		return Time{}, err
	}

	return NewTime(time.Date(year, m, day, hour, minute, second, 0, tz.Location())), nil

}

// removeSpaces removes all spaces from a string.
func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
