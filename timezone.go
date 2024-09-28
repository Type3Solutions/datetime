package mildtg

import (
	"errors"
	"strings"
	"time"
)

const (
	alpha    = 'A'
	bravo    = 'B'
	charlie  = 'C'
	delta    = 'D'
	echo     = 'E'
	foxtrot  = 'F'
	golf     = 'G'
	hotel    = 'H'
	india    = 'I'
	kilo     = 'K'
	lima     = 'L'
	mike     = 'M'
	november = 'N'
	oscar    = 'O'
	papa     = 'P'
	quebec   = 'Q'
	romeo    = 'R'
	sierra   = 'S'
	tango    = 'T'
	uniform  = 'U'
	victor   = 'V'
	whiskey  = 'W'
	xray     = 'X'
	yankee   = 'Y'
	zulu     = 'Z'
)

// TimeZone represents a time zone designation.
type timeZone struct {
	letter rune  // time zone designation letter (Zulu, Alpha, Bravo, etc.)
	offset int32 // number of seconds east of UTC/GMT (positive) or west of UTC/GMT (negative)
}

// String returns the time zone designation letter.
func (tz timeZone) String() string {
	return strings.ToUpper(string(tz.letter))
}

// Offset returns seconds east of UTC/GMT (positive) or west of UTC/GMT (negative).
func (tz timeZone) Offset() int {
	return int(tz.offset)
}

// Location returns the time.Location for the time zone.
func (tz timeZone) Location() *time.Location {
	return time.FixedZone(tz.String(), int(tz.offset))
}

const (
	hour int32 = 3600 // seconds in an hour
)

var (
	ErrInvalidTimeZone = errors.New("invalid time zone")
)

// Time zone designations skip J (Juliet).
// J is used to indicate the local time zone.

var (
	ZULU     = timeZone{zulu, 0 * hour}      // Zulu GMT +0
	ALPHA    = timeZone{alpha, 1 * hour}     // Alpha GMT +1
	BRAVO    = timeZone{bravo, 2 * hour}     // Bravo GMT +2
	CHARLIE  = timeZone{charlie, 3 * hour}   // Charlie GMT +3
	DELTA    = timeZone{delta, 4 * hour}     // Delta GMT +4
	ECHO     = timeZone{echo, 5 * hour}      // Echo GMT +5
	FOXTROT  = timeZone{foxtrot, 6 * hour}   // Foxtrot GMT +6
	GOLF     = timeZone{golf, 7 * hour}      // Golf GMT +7
	HOTEL    = timeZone{hotel, 8 * hour}     // Hotel GMT +8
	INDIA    = timeZone{india, 9 * hour}     // India GMT +9
	KILO     = timeZone{kilo, 10 * hour}     // Kilo GMT +10
	LIMA     = timeZone{lima, 11 * hour}     // Lima GMT +11
	MIKE     = timeZone{mike, 12 * hour}     // Mike GMT +12
	NOVEMBER = timeZone{november, -1 * hour} // November GMT -1
	OSCAR    = timeZone{oscar, -2 * hour}    // Oscar GMT -2
	PAPA     = timeZone{papa, -3 * hour}     // Papa GMT -3
	QUEBEC   = timeZone{quebec, -4 * hour}   // Quebec GMT -4
	ROMEO    = timeZone{romeo, -5 * hour}    // Romeo GMT -5
	SIERRA   = timeZone{sierra, -6 * hour}   // Sierra GMT -6
	TANGO    = timeZone{tango, -7 * hour}    // Tango GMT -7
	UNIFORM  = timeZone{uniform, -8 * hour}  // Uniform GMT -8
	VICTOR   = timeZone{victor, -9 * hour}   // Victor GMT -9
	WHISKEY  = timeZone{whiskey, -10 * hour} // Whiskey GMT -10
	XRAY     = timeZone{xray, -11 * hour}    // X-ray GMT -11
	YANKEE   = timeZone{yankee, -12 * hour}  // Yankee GMT -12
)

var (
	timeZones = map[rune]timeZone{
		zulu:     ZULU,
		alpha:    ALPHA,
		bravo:    BRAVO,
		charlie:  CHARLIE,
		delta:    DELTA,
		echo:     ECHO,
		foxtrot:  FOXTROT,
		golf:     GOLF,
		hotel:    HOTEL,
		india:    INDIA,
		kilo:     KILO,
		lima:     LIMA,
		mike:     MIKE,
		november: NOVEMBER,
		oscar:    OSCAR,
		papa:     PAPA,
		quebec:   QUEBEC,
		romeo:    ROMEO,
		sierra:   SIERRA,
		tango:    TANGO,
		uniform:  UNIFORM,
		victor:   VICTOR,
		whiskey:  WHISKEY,
		xray:     XRAY,
		yankee:   YANKEE,
	}
)

// parseTimeZone parses a time zone designation string.
func parseTimeZone(s string) (timeZone, error) {
	tz := timeZone{}

	if len(s) < 1 || len(s) > 1 {
		return tz, ErrInvalidTimeZone
	}

	tz.letter = rune(s[0])

	if tzOut, ok := timeZones[tz.letter]; ok {
		return tzOut, nil
	} else {
		return tz, ErrInvalidTimeZone
	}
}
