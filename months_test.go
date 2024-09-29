package mildtg

import (
	"testing"
	"time"
)

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
