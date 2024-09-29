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
		dtg   string
		month string
		want  time.Month
		error error
	}{
		{name: "january", dtg: "270000Z JAN 2020", month: "JAN", want: time.January, error: nil},
		{name: "february", dtg: "270000Z FEB 2020", month: "FEB", want: time.February, error: nil},
		{name: "march", dtg: "270000Z MAR 2020", month: "MAR", want: time.March, error: nil},
		{name: "april", dtg: "270000Z APR 2020", month: "APR", want: time.April, error: nil},
		{name: "may", dtg: "270000Z MAY 2020", month: "MAY", want: time.May, error: nil},
		{name: "june", dtg: "270000Z JUN 2020", month: "JUN", want: time.June, error: nil},
		{name: "july", dtg: "270000Z JUL 2020", month: "JUL", want: time.July, error: nil},
		{name: "august", dtg: "270000Z AUG 2020", month: "AUG", want: time.August, error: nil},
		{name: "september", dtg: "270000Z SEP 2020", month: "SEP", want: time.September, error: nil},
		{name: "october", dtg: "270000Z OCT 2020", month: "OCT", want: time.October, error: nil},
		{name: "november", dtg: "270000Z NOV 2020", month: "NOV", want: time.November, error: nil},
		{name: "december", dtg: "270000Z DEC 2020", month: "DEC", want: time.December, error: nil},
		{name: "january mixed case", dtg: "270000Z jan 2020", month: "jan", want: time.January, error: nil},
		{name: "invalid month", dtg: "270000Z invalid 2020", month: "", want: 0, error: ErrInvalidMonth},
		{name: "empty", dtg: "", month: "", want: 0, error: ErrInvalidMonth},
		{name: "emoji", dtg: "🕒", month: "", want: 0, error: ErrInvalidMonth},
		{name: "invalid month length", dtg: "270000Z JANUARYY 2020", month: "", want: 0, error: ErrInvalidMonth},
		{name: "full january", dtg: "270000Z JANUARY 2020", month: "JANUARY", want: time.January, error: nil},
		{name: "full february", dtg: "270000Z FEBRUARY 2020", month: "FEBRUARY", want: time.February, error: nil},
		{name: "full march", dtg: "270000Z MARCH 2020", month: "MARCH", want: time.March, error: nil},
		{name: "full april", dtg: "270000Z APRIL 2020", month: "APRIL", want: time.April, error: nil},
		{name: "full may", dtg: "270000Z MAY 2020", month: "MAY", want: time.May, error: nil},
		{name: "full june", dtg: "270000Z JUNE 2020", month: "JUNE", want: time.June, error: nil},
		{name: "full july", dtg: "270000Z JULY 2020", month: "JULY", want: time.July, error: nil},
		{name: "full august", dtg: "270000Z AUGUST 2020", month: "AUGUST", want: time.August, error: nil},
		{name: "full september", dtg: "270000Z SEPTEMBER 2020", month: "SEPTEMBER", want: time.September, error: nil},
		{name: "full october", dtg: "270000Z OCTOBER 2020", month: "OCTOBER", want: time.October, error: nil},
		{name: "full november", dtg: "270000Z NOVEMBER 2020", month: "NOVEMBER", want: time.November, error: nil},
		{name: "full december", dtg: "270000Z DECEMBER 2020", month: "DECEMBER", want: time.December, error: nil},
		{name: "no month", dtg: "270000Z", month: "", want: time.Now().UTC().Month(), error: ErrInvalidMonth},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, month, err := parseMonth(tt.dtg)

			if !errors.Is(err, tt.error) {
				t.Errorf("got %v, want %v", err, tt.error)
			}

			if str != tt.month {
				t.Errorf("got %v, want %v", str, tt.month)
			}

			if month != tt.want {
				t.Errorf("got %v, want %v", month, tt.want)
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
