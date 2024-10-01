package mildtg

import (
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
	juliet   = 'J'
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
	secondsInHour int32 = 3600 // seconds in an hour
)

// Time zone designations skip J (Juliet).
// J is used to indicate the local time zone.

var (
	ZULU     = timeZone{zulu, 0 * secondsInHour}      // Zulu GMT +0
	ALPHA    = timeZone{alpha, 1 * secondsInHour}     // Alpha GMT +1
	BRAVO    = timeZone{bravo, 2 * secondsInHour}     // Bravo GMT +2
	CHARLIE  = timeZone{charlie, 3 * secondsInHour}   // Charlie GMT +3
	DELTA    = timeZone{delta, 4 * secondsInHour}     // Delta GMT +4
	ECHO     = timeZone{echo, 5 * secondsInHour}      // Echo GMT +5
	FOXTROT  = timeZone{foxtrot, 6 * secondsInHour}   // Foxtrot GMT +6
	GOLF     = timeZone{golf, 7 * secondsInHour}      // Golf GMT +7
	HOTEL    = timeZone{hotel, 8 * secondsInHour}     // Hotel GMT +8
	INDIA    = timeZone{india, 9 * secondsInHour}     // India GMT +9
	JULIET   = timeZone{juliet, 0}                    // Juliet local time zone
	KILO     = timeZone{kilo, 10 * secondsInHour}     // Kilo GMT +10
	LIMA     = timeZone{lima, 11 * secondsInHour}     // Lima GMT +11
	MIKE     = timeZone{mike, 12 * secondsInHour}     // Mike GMT +12
	NOVEMBER = timeZone{november, -1 * secondsInHour} // November GMT -1
	OSCAR    = timeZone{oscar, -2 * secondsInHour}    // Oscar GMT -2
	PAPA     = timeZone{papa, -3 * secondsInHour}     // Papa GMT -3
	QUEBEC   = timeZone{quebec, -4 * secondsInHour}   // Quebec GMT -4
	ROMEO    = timeZone{romeo, -5 * secondsInHour}    // Romeo GMT -5
	SIERRA   = timeZone{sierra, -6 * secondsInHour}   // Sierra GMT -6
	TANGO    = timeZone{tango, -7 * secondsInHour}    // Tango GMT -7
	UNIFORM  = timeZone{uniform, -8 * secondsInHour}  // Uniform GMT -8
	VICTOR   = timeZone{victor, -9 * secondsInHour}   // Victor GMT -9
	WHISKEY  = timeZone{whiskey, -10 * secondsInHour} // Whiskey GMT -10
	XRAY     = timeZone{xray, -11 * secondsInHour}    // X-ray GMT -11
	YANKEE   = timeZone{yankee, -12 * secondsInHour}  // Yankee GMT -12
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
