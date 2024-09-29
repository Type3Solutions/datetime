package mildtg

import (
	"errors"
	"regexp"
	"time"
)

const (
	minYear = 1941 // The U.S. entered World War II in 1941.
	maxYear = 9999 // The maximum year allowed.
)

var (
	// ErrInvalidYear is returned when an invalid hour is provided.
	ErrInvalidYear = errors.New("invalid year")

	// ErrYearOutOfRange is returned when the year is out of range.
	ErrYearOutOfRange = errors.New("year out of range")

	// yearRegex is a regular expression to match a two or four-digit year
	// at the end of a date-time-group string.
	// A one, three, or more than four-digit year is not allowed and will
	// return an error.
	yearRegex = regexp.MustCompile(`\d+$`)

	// nonAlphaNumericRegex is a regular expression to match any non-alphanumeric
	// characters in a string.
	nonAlphaNumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

// parseYear returns the year from a date-time-group string.
//
// The year is optional and may not be present in the string.
// If the year is not present, the current year is returned.
//
// The year is expected to be in the format YY or YYYY and
// occurs after the month and day in the date-time-group string.
// If the year is not in the expected format or out of the range,
// an error is returned.
//
// The year is expected to be in the range 00-99 or 0000-9999.
// If the two-digit year is less than 50, we assume the year to be
// in the 21st century.
// If the two-digit year is greater than or equal to 50, the year
// is assumed to be in the 20th century.
//
// There is a chance that the string may contain an operational
// date-time-group format (DDHHMMZ) which does not contain a year.
// If this is the case, the current year is returned.
// All other formats are expected to contain a two or four-digit year.
//
// Examples:
//
//	"270000Z JAN 20" → 2020
//	"270000Z JAN 2020" → 2020
//	"270000Z" → current year
func parseYear(s string) (int, error) {
	// Remove all spaces from the string.
	s = removeSpaces(s)

	// Check if the string is empty.
	if s == "" {
		return 0, ErrInvalidYear
	}

	// Check if the string contains any non-alphanumeric characters.
	if nonAlphaNumericRegex.MatchString(s) {
		return 0, ErrInvalidYear
	}

	// Find the year in the string.
	year := yearRegex.FindString(s)
	yearLen := len(year)

	switch yearLen {
	case 0:
		// The year is not present in the string.
		// Return the current year.
		return time.Now().UTC().Year(), nil
	case 2:
		// The year is two-digits.
		// Convert the year to an integer.
		y := int(year[0]-'0')*10 + int(year[1]-'0')

		// Determine if the year is in the 21st or 20th century.
		if y < 50 {
			return 2000 + y, nil
		}

		return 1900 + y, nil
	case 4:
		// The year is four-digits.
		// Convert the year to an integer.
		y := int(year[0]-'0')*1000 + int(year[1]-'0')*100 +
			int(year[2]-'0')*10 + int(year[3]-'0')

		// Check if the year is out of range.
		if y < minYear || y > maxYear {
			return 0, ErrYearOutOfRange
		}

		return y, nil
	default:
		// The year is not in the expected format.
		return 0, ErrInvalidYear
	}
}
