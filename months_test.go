package mildtg

import (
	"errors"
	"testing"
	"time"
)

func TestParseMonth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  time.Month
		error error
	}{
		{name: "january", input: "JAN", want: time.January, error: nil},
		{name: "february", input: "FEB", want: time.February, error: nil},
		{name: "march", input: "MAR", want: time.March, error: nil},
		{name: "april", input: "APR", want: time.April, error: nil},
		{name: "may", input: "MAY", want: time.May, error: nil},
		{name: "june", input: "JUN", want: time.June, error: nil},
		{name: "july", input: "JUL", want: time.July, error: nil},
		{name: "august", input: "AUG", want: time.August, error: nil},
		{name: "september", input: "SEP", want: time.September, error: nil},
		{name: "october", input: "OCT", want: time.October, error: nil},
		{name: "november", input: "NOV", want: time.November, error: nil},
		{name: "december", input: "DEC", want: time.December, error: nil},
		{name: "invalid", input: "invalid", want: time.Month(0), error: ErrInvalidMonth},
		{name: "empty", input: "", want: time.Month(0), error: ErrInvalidMonth},
		{name: "emoji", input: "🕒", want: time.Month(0), error: ErrInvalidMonth},
		{name: "lowercase", input: "jan", want: time.January, error: nil},
		{name: "mix case", input: "Jan", want: time.January, error: nil},
		{name: "full dtg", input: "270000Z JAN 20", want: time.January, error: nil},
		{name: "full dtg lowercase", input: "270000Z jan 20", want: time.January, error: nil},
		{name: "full dtg mix case", input: "270000Z Jan 20", want: time.January, error: nil},
		{name: "full dtg invalid", input: "270000Z invalid 20", want: time.Month(0), error: ErrInvalidMonth},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMonth(tt.input)

			if !errors.Is(err, tt.error) {
				t.Errorf("got %v, want %v", err, tt.error)
			}

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input time.Month
		year  int
		want  int
	}{
		{name: "january", input: time.January, year: 2020, want: 31},
		{name: "february", input: time.February, year: 2020, want: 29},
		{name: "march", input: time.March, year: 2020, want: 31},
		{name: "april", input: time.April, year: 2020, want: 30},
		{name: "may", input: time.May, year: 2020, want: 31},
		{name: "june", input: time.June, year: 2020, want: 30},
		{name: "july", input: time.July, year: 2020, want: 31},
		{name: "august", input: time.August, year: 2020, want: 31},
		{name: "september", input: time.September, year: 2020, want: 30},
		{name: "october", input: time.October, year: 2020, want: 31},
		{name: "november", input: time.November, year: 2020, want: 30},
		{name: "december", input: time.December, year: 2020, want: 31},
		{name: "february no leap year", input: time.February, year: 2021, want: 28},
		{name: "invalid", input: time.Month(0), year: 2020, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := daysInMonth(tt.input, tt.year)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
