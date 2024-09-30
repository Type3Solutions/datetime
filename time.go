package mildtg

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

const (
	minYear = 1941 // The U.S. entered World War II in 1941.
	maxYear = 9999 // The maximum year allowed.
)

const (
	dayMatch      = `^(0[1-9]|[12][0-9]|3[01])` // 01-31 is required
	hourMatch     = `(0[0-9]|1[0-9]|2[0-3])`    // 00-23 is required
	minuteMatch   = `(0[0-9]|[1-5][0-9])`       // 00-59 is required
	secondMatch   = `([0-5][0-9])?`             // 00-59 is optional
	timeZoneMatch = `([A-Z]{0,1})?`             // A-Z is optional for some date-time-groups
	monthsMatch   = `((JAN(\d{0,4}$)|JANUARY(\d{0,4}$))|` +
		`(FEB(\d{0,4}$)|FEBRUARY(\d{0,4}$))|` +
		`(MAR(\d{0,4}$)|MARCH(\d{0,4}$))|` +
		`(APR(\d{0,4}$)|APRIL(\d{0,4}$))|` +
		`(MAY(\d{0,4}$))|` +
		`(JUN(\d{0,4}$)|JUNE(\d{0,4}$))|` +
		`(JUL(\d{0,4}$)|JULY(\d{0,4}$))|` +
		`(AUG(\d{0,4}$)|AUGUST(\d{0,4}$))|` +
		`(SEP(\d{0,4}$)|SEPTEMBER(\d{0,4}$))|` +
		`(OCT(\d{0,4}$)|OCTOBER(\d{0,4}$))|` +
		`(NOV(\d{0,4}$)|NOVEMBER(\d{0,4}$))|` +
		`(DEC(\d{0,4}$)|DECEMBER(\d{0,4}$)))?` // Three-letter month or full month name is optional

	dtgMatch = `(?i)` + dayMatch + hourMatch + minuteMatch + secondMatch +
		timeZoneMatch + monthsMatch
)

var (
	// ErrNotEnoughChars is returned when an invalid date-time-group is provided.
	ErrNotEnoughChars = errors.New("date-time-group too short minimum is ddhhmm")

	// ErrInvalidDateTimeGroup is returned when an invalid date-time-group is provided.
	ErrInvalidDateTimeGroup = errors.New("invalid date-time-group")

	// ErrYearOutOfRange is returned when the year is out of range.
	ErrYearOutOfRange = errors.New("year out of range")
)

var (
	// dtgRegex is a regular expression to match a date-time-group string.
	dtgRegex = regexp.MustCompile(dtgMatch)

	// secondsRegex is a regular expression to match a date-time-group string with seconds.
	secondsRegex = regexp.MustCompile(`\d{6}([0-5][0-9])`)

	// timeZoneRegex is a regular expression to match a date-time-group string with a time zone.
	timeZoneRegex = regexp.MustCompile(`(?i)[A-Z]`)

	// monthRegex is a regular expression to match a three-letter month or
	// the full month name.
	monthRegex = regexp.MustCompile(`(?i)JAN(uary)?|FEB(ruary)?|MAR(ch)?|APR(il)?|` +
		`MAY|JUN(e)?|JUL(y)?|AUG(ust)?|SEP(tember)?|` +
		`OCT(ober)?|NOV(ember)?|DEC(ember)?`)

	// yearRegex is a regular expression to match a two or four digit year.
	yearRegex = regexp.MustCompile(`\d{2}$|\d{4}$`)
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

	// Check if the string is empty.
	if len(s) < 6 {
		return Time{}, ErrNotEnoughChars
	}

	// Check if the string is valid by comparing the length of the string
	// to the length of the date-time-group regular expression string.
	// There should be no additional or missing characters.
	if len(s) != len(dtgRegex.FindString(s)) {
		return Time{}, ErrInvalidDateTimeGroup
	}

	parts := dtgRegex.FindStringSubmatch(s)

	dayStr := parts[1]
	hourStr := parts[2]
	minuteStr := parts[3]
	day := int(dayStr[0]-'0')*10 + int(dayStr[1]-'0')
	hour := int(hourStr[0]-'0')*10 + int(hourStr[1]-'0')
	minute := int(minuteStr[0]-'0')*10 + int(minuteStr[1]-'0')

	// If there are no additional parts, we can return the time.
	if len(parts) == 4 {
		// Current year, month, and Zulu timezone are used if not provided.
		currentYear := time.Now().UTC().Year()
		currentMonth := time.Now().UTC().Month()

		return NewTime(time.Date(currentYear, currentMonth, day, hour,
			0, 0, 0, time.UTC)), nil
	}

	// If there are additional parts, we need to parse them.
	var seconds int                   // Default to 0.
	month := time.Now().UTC().Month() // Default to current month.
	year := time.Now().UTC().Year()   // Default to current year.
	tz := ZULU                        // Default to Zulu time zone.

	// Parse the additional parts.
	for i, part := range parts[4:] {
		switch {
		case secondsRegex.MatchString(part) && i == 0:
			// Seconds are present.
			seconds = int(part[0]-'0')*10 + int(part[1]-'0')
		case timeZoneRegex.MatchString(part) && len(part) == 1:
			// Time zone is present.
			tzStr := strings.ToUpper(part)
			tzOut, ok := timeZones[rune(tzStr[0])]
			if !ok {
				return Time{}, ErrInvalidTimeZone
			}
			tz = tzOut
		case monthRegex.MatchString(part):
			// Month is present but may contain the year.
			monthStr := monthRegex.FindString(part)
			monthOut, ok := months[strings.ToUpper(monthStr)]
			if !ok {
				return Time{}, ErrInvalidMonth
			}

			month = monthOut
		case yearRegex.MatchString(part) && len(part) > 0:
			switch len(part) {
			case 2:
				// Two-digit year.
				y := int(part[0]-'0')*10 + int(part[1]-'0')

				// Determine if the year is in the 21st or 20th century.
				if y < 50 {
					year = 2000 + y
				} else {
					year = 1900 + y
				}
			case 4:
				// Four-digit year.
				y := int(part[0]-'0')*1000 + int(part[1]-'0')*100 +
					int(part[2]-'0')*10 + int(part[3]-'0')

				// Check if the year is out of range.
				if y < minYear || y > maxYear {
					return Time{}, ErrYearOutOfRange
				}

				year = y
			}
		default:
			continue
		}
	}

	// Check if the day is valid for the month and year.
	if day > daysInMonth(month, year) {
		return Time{}, ErrInvalidDay
	}

	t := time.Date(year, month, day, hour, minute, seconds, 0, tz.Location())

	return NewTime(t), nil
}

// ParseDTGBytes parses a military date-time-group byte slice in the format
// DDHH[MM]|[MMSS]|(A-Z)[ MMM YY[YY] and returns a Time object.
func ParseDTGBytes(s string) (Time, error) {

	// The digitsBeforeChar slice is used to store the digits before any
	// characters in the date-time-group.
	// The slice has enough capacity to store the maximum number of digits
	// before a character in the date-time-group.
	// If there are no characters in the date-time-group, the slice will
	// store the maximum number of digits.
	digitsBeforeChar := make([]byte, 0, 2*4+4)
	maxDigitsBeforeChar := 2*4 + 4

	// The digitsAfterChar slice is used to store the digits after any
	// characters in the date-time-group.
	// This slice will have enough capacity to store a four-digit year.
	digitsAfterChar := make([]byte, 0, 4)
	maxDigitsAfterChar := 4

	// September is the longest month name plus one for the time zone designation.
	chars := make([]byte, 0, 9+1)
	maxChars := 9 + 1

	// The index is used to keep track of the current index in the digits slice.
	digitsBeforeIndex := 0
	digitsAfterIndex := 0
	charIndex := 0
	charsFound := false

	// Iterate over each byte in the byte slice.
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			// Digit.
			if !charsFound {
				// If the index is greater than or equal to the length of the slice,
				// return an error.
				if digitsBeforeIndex >= maxDigitsBeforeChar {
					return Time{}, ErrInvalidDateTimeGroup
				}

				digitsBeforeChar = append(digitsBeforeChar, s[i])
				digitsBeforeIndex++

			} else {
				// If the index is greater than or equal to the length of the slice,
				// return an error.
				// This could happen if the method receives a year with more than
				// four digits.
				if digitsAfterIndex >= maxDigitsAfterChar {
					return Time{}, ErrInvalidDateTimeGroup
				}

				digitsAfterChar = append(digitsAfterChar, s[i])
				digitsAfterIndex++
			}

		case s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z':
			// Character.
			if charIndex >= maxChars {
				return Time{}, ErrInvalidDateTimeGroup
			}

			var char byte
			if s[i] >= 'a' && s[i] <= 'z' {
				char = s[i] - 32
			} else {
				char = s[i]
			}

			chars = append(chars, char)
			charIndex++

			// Set the charsFound flag to true.
			charsFound = true

		case s[i] == ' ':
			// Do nothing.
			continue

		default:
			// Invalid character.
			return Time{}, ErrInvalidDateTimeGroup
		}
	}

	// The digitsBeforeChar slice must have at least six digits
	// and have a zero remainder if len(digitsBeforeChar) is divided by 2.
	if len(digitsBeforeChar) < 6 || len(digitsBeforeChar)%2 != 0 {
		return Time{}, ErrNotEnoughChars
	}

	// The day, hour, and minute are extracted from the digitsBeforeChar slice.
	day := int(digitsBeforeChar[0]-'0')*10 + int(digitsBeforeChar[1]-'0')
	hour := int(digitsBeforeChar[2]-'0')*10 + int(digitsBeforeChar[3]-'0')
	minute := int(digitsBeforeChar[4]-'0')*10 + int(digitsBeforeChar[5]-'0')
	seconds := 0
	year := time.Now().UTC().Year()
	month := time.Now().UTC().Month()
	tz := ZULU

	// Remove the day, hour, and minute from the slice.
	digitsBeforeChar = digitsBeforeChar[6:]

	switch {
	case len(digitsBeforeChar) == 0:
		// Do nothing.
	case len(digitsBeforeChar) == 2:
		// If the length of the remaining digits before the character is two, we
		// can assume these two digits represent the seconds.
		seconds = int(digitsBeforeChar[0]-'0')*10 + int(digitsBeforeChar[1]-'0')
		if seconds > 59 {
			return Time{}, ErrInvalidDateTimeGroup
		}
	case len(digitsBeforeChar) == 4:
		// If the length of the remaining digits before the character is four, we
		// can assume these four digits represent the four-digit year.
		// We do not attempt to parse a two-digit second with a two-digit year.
		year = int(digitsBeforeChar[0]-'0')*1000 + int(digitsBeforeChar[1]-'0')*100 +
			int(digitsBeforeChar[2]-'0')*10 + int(digitsBeforeChar[3]-'0')

	case len(digitsBeforeChar) == 6:
		// If the length of the remaining digits before the character is six, we
		// can assume we have a two-digit seconds and a four-digit year.
		seconds = int(digitsBeforeChar[0]-'0')*10 + int(digitsBeforeChar[1]-'0')
		if seconds > 59 {
			return Time{}, ErrInvalidDateTimeGroup
		}

		year = int(digitsBeforeChar[2]-'0')*1000 + int(digitsBeforeChar[3]-'0')*100 +
			int(digitsBeforeChar[4]-'0')*10 + int(digitsBeforeChar[5]-'0')
	}

	// Parse the month and time zone from the chars slice.
	switch {
	case len(chars) == 0:
		// Do nothing.
	case len(chars) == 1:
		// If the length of the chars slice is one, we can assume this character
		// represents the time zone.
		tzStr := strings.ToUpper(string(chars[0]))
		tzOut, ok := timeZones[rune(tzStr[0])]
		if !ok {
			// Get the local offset.
			offset := time.Now().UTC().Sub(time.Now()).Seconds() / 3600
			tzOut = timeZone{letter: rune(tzStr[0]), offset: int32(offset)}
		}

		tz = tzOut
	case len(chars) == 3:
		// If the length of the chars slice is three, we can assume this represents
		// the three-letter month abbreviation.
		monthStr := string(chars)
		monthOut, ok := months[strings.ToUpper(monthStr)]
		if !ok {
			return Time{}, ErrInvalidMonth
		}

		month = monthOut
	case len(chars) > 3:
		// If the length of the chars slice is greater than three, we either have a
		// time zone, a three-letter month abbreviation, or a full month name or a
		// combination of these.
		//
		// Check if the char string contains a month.
		m, ok := months[strings.ToUpper(string(chars))]
		if !ok {
			// Remove the first character from the chars slice.
			tzStr := strings.ToUpper(string(chars[0]))
			newMonthStr := string(chars[1:])

			// Check if the new month string is a valid month.
			m, ok = months[strings.ToUpper(newMonthStr)]
			if !ok {
				return Time{}, ErrInvalidMonth
			}

			tzOut, tzFound := timeZones[rune(tzStr[0])]
			if !tzFound {
				// Get the local offset.
				offset := time.Now().UTC().Sub(time.Now()).Seconds() / 3600
				tzOut = timeZone{letter: rune(tzStr[0]), offset: int32(offset)}
			}

			tz = tzOut
			month = m
		}

		month = m

	default:
		return Time{}, ErrInvalidDateTimeGroup
	}

	// Check the digitsAfterChar slice.
	if len(digitsAfterChar)%2 != 0 {
		return Time{}, ErrInvalidDateTimeGroup
	}

	// The maximum length of the digitsAfterChar slice is four,
	// and we should not see 1 or 3 digits.
	switch len(digitsAfterChar) {
	case 0:
		// Do nothing.
	case 2:
		// Two-digit year.
		y := int(digitsAfterChar[0]-'0')*10 + int(digitsAfterChar[1]-'0')

		// Determine if the year is in the 21st or 20th century.
		if y < 69 {
			year = 2000 + y
		} else {
			year = 1900 + y
		}
	case 4:
		// Four-digit year.
		y := int(digitsAfterChar[0]-'0')*1000 + int(digitsAfterChar[1]-'0')*100 +
			int(digitsAfterChar[2]-'0')*10 + int(digitsAfterChar[3]-'0')

		year = y
	}

	// Check if the day is valid for the month and year.
	if day > daysInMonth(month, year) || day < 1 {
		return Time{}, ErrInvalidDay
	}

	// Check hours and minutes.
	if hour > 23 || minute > 59 {
		return Time{}, ErrInvalidDateTimeGroup
	}

	t := time.Date(year, month, day, hour, minute, seconds, 0, tz.Location())

	return NewTime(t), nil
}

// removeSpaces removes all spaces from a string.
func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
