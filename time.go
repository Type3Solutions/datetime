package mildtg

import (
	"strings"
	"time"
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

// ParseDTG parses a date-time-group string in the format
func ParseDTG(s string) (Time, error) {
	t, err := time.Parse("200601021504", s)
	if err != nil {
		return Time{}, err
	}
	return NewTime(t), nil
}

// removeSpaces removes all spaces from a string.
func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
