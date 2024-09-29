package mildtg

import (
	"errors"
	"testing"
	"time"
)

func TestParseYear(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  int
		error error
	}{
		{name: "valid four digit year", input: "270000Z JAN 2020", want: 2020, error: nil},
		{name: "valid two digit year", input: "270000Z JAN 20", want: 2020, error: nil},
		{name: "no year", input: "270000Z", want: time.Now().UTC().Year(), error: nil},
		{name: "invalid year", input: "270000Z JAN invalid", want: 0, error: ErrInvalidYear},
		{name: "empty", input: "", want: 0, error: ErrInvalidYear},
		{name: "emoji", input: "🕒", want: 0, error: ErrInvalidYear},
		{name: "invalid year length", input: "270000Z JAN 20200", want: 0, error: ErrInvalidYear},
		{name: "invalid year length", input: "270000Z JAN 202000", want: 0, error: ErrInvalidYear},
		{name: "year before min", input: "270000Z JAN 1940", want: 0, error: ErrYearOutOfRange},
		{name: "year out of range", input: "270000Z JAN 10000", want: 0, error: ErrInvalidYear},
		{name: "invalid year length", input: "270000Z JAN 100", want: 0, error: ErrInvalidYear},
		{name: "invalid year length", input: "270000Z JAN 1", want: 0, error: ErrInvalidYear},
		{name: "invalid year length", input: "270000Z JAN 0", want: 0, error: ErrInvalidYear},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseYear(tt.input)

			if !errors.Is(err, tt.error) {
				t.Errorf("got %v, want %v", err, tt.error)
			}

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
