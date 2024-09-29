package mildtg

import (
	"errors"
	"testing"
	"time"
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

func TestParseDTG(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  Time
		error error
	}{
		{
			name:  "valid dtg with no spaces",
			input: "010100ZJAN21",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "valid dtg with spaces",
			input: "01 01 00Z JAN 21",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "valid dtg with lowercase month",
			input: "010100Zjan21",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "valid dtg with lowercase timezone",
			input: "010100zJAN21",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "valid dtg with lowercase month and timezone",
			input: "010100zjan21",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "day, hour, and minute only",
			input: "010100",
			want: NewTime(time.Date(time.Now().UTC().Year(),
				time.Now().UTC().Month(), 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "misspelled month",
			input: "010100ZJANUARYY21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "another misspelled month",
			input: "010100ZJJANUARY21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "yet another misspelled month",
			input: "010100ZJANAURY21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "invalid day",
			input: "000100ZJAN21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "invalid minutes",
			input: "010160ZJAN21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "leap year days on non-leap year",
			input: "290200ZFEB21",
			want:  Time{},
			error: ErrInvalidDay,
		},
		{
			name:  "leap year days on leap year",
			input: "290200ZFEB20",
			want:  NewTime(time.Date(2020, 2, 29, 2, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDTG(tt.input)
			if !errors.Is(err, tt.error) {
				t.Errorf("got %v, want %v", err, tt.error)
			}

			t.Run("year", func(t *testing.T) {
				if got.Year() != tt.want.Year() {
					t.Errorf("got %v, want %v", got.Year(), tt.want.Year())
				}
			})

			t.Run("month", func(t *testing.T) {
				if got.Month() != tt.want.Month() {
					t.Errorf("got %v, want %v", got.Month(), tt.want.Month())
				}
			})

			t.Run("day", func(t *testing.T) {
				if got.Day() != tt.want.Day() {
					t.Errorf("got %v, want %v", got.Day(), tt.want.Day())
				}
			})

			t.Run("hour", func(t *testing.T) {
				if got.Hour() != tt.want.Hour() {
					t.Errorf("got %v, want %v", got.Hour(), tt.want.Hour())
				}
			})

			t.Run("minute", func(t *testing.T) {
				if got.Minute() != tt.want.Minute() {
					t.Errorf("got %v, want %v", got.Minute(), tt.want.Minute())
				}
			})

			t.Run("second", func(t *testing.T) {
				if got.Second() != tt.want.Second() {
					t.Errorf("got %v, want %v", got.Second(), tt.want.Second())
				}
			})

			t.Run("location", func(t *testing.T) {
				gotLocation := got.Location()
				wantLocation := tt.want.Location()

				t.Run("name", func(t *testing.T) {
					if gotLocation.String() != wantLocation.String() {
						t.Errorf("got %v, want %v", gotLocation.String(), wantLocation.String())
					}
				})
			})

			t.Run("string", func(t *testing.T) {
				if got.String() != tt.want.String() {
					t.Errorf("got %v, want %v", got.String(), tt.want.String())
				}
			})
		})
	}
}
