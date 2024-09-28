package mildtg

import (
	"errors"
	"testing"
)

func TestParseTimeZones(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  timeZone
		error error
	}{
		{name: "Zulu", input: "Z", want: ZULU, error: nil},
		{name: "Alpha", input: "A", want: ALPHA, error: nil},
		{name: "Bravo", input: "B", want: BRAVO, error: nil},
		{name: "Charlie", input: "C", want: CHARLIE, error: nil},
		{name: "Delta", input: "D", want: DELTA, error: nil},
		{name: "Echo", input: "E", want: ECHO, error: nil},
		{name: "Foxtrot", input: "F", want: FOXTROT, error: nil},
		{name: "Golf", input: "G", want: GOLF, error: nil},
		{name: "Hotel", input: "H", want: HOTEL, error: nil},
		{name: "India", input: "I", want: INDIA, error: nil},
		{name: "Kilo", input: "K", want: KILO, error: nil},
		{name: "Lima", input: "L", want: LIMA, error: nil},
		{name: "Mike", input: "M", want: MIKE, error: nil},
		{name: "November", input: "N", want: NOVEMBER, error: nil},
		{name: "Oscar", input: "O", want: OSCAR, error: nil},
		{name: "Papa", input: "P", want: PAPA, error: nil},
		{name: "Quebec", input: "Q", want: QUEBEC, error: nil},
		{name: "Romeo", input: "R", want: ROMEO, error: nil},
		{name: "Sierra", input: "S", want: SIERRA, error: nil},
		{name: "Tango", input: "T", want: TANGO, error: nil},
		{name: "Uniform", input: "U", want: UNIFORM, error: nil},
		{name: "Victor", input: "V", want: VICTOR, error: nil},
		{name: "Whiskey", input: "W", want: WHISKEY, error: nil},
		{name: "Xray", input: "X", want: XRAY, error: nil},
		{name: "Yankee", input: "Y", want: YANKEE, error: nil},
		{name: "Invalid", input: "invalid", want: timeZone{}, error: ErrInvalidTimeZone},
		{name: "Empty", input: "", want: timeZone{}, error: ErrInvalidTimeZone},
		{name: "Emoji", input: "🕒", want: timeZone{}, error: ErrInvalidTimeZone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimeZone(tt.input)
			if !errors.Is(err, tt.error) {
				t.Errorf("got %v, want %v", err, tt.error)
			}

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
