package mildtg

import (
	"errors"
	"testing"
	"time"
)

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
		{
			name:  "invalid timezone",
			input: "010100🇺🇸JAN21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "20th century year",
			input: "010100ZJAN99",
			want:  NewTime(time.Date(1999, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "21st century year",
			input: "010100ZJAN00",
			want:  NewTime(time.Date(2000, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "invalid year",
			input: "010100ZJAN1940",
			want:  Time{},
			error: ErrYearOutOfRange,
		},
		{
			name:  "invalid day for month",
			input: "310100ZFEB21",
			want:  Time{},
			error: ErrInvalidDay,
		},
		{
			name:  "valid dtg with 4-digit year",
			input: "010100ZJAN2021",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "invalid dtg",
			input: "🇺🇸",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "invalid dtg with valid month",
			input: "JAN",
			want:  Time{},
			error: ErrNotEnoughChars,
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

func TestParseDTGBytes(t *testing.T) {
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
			error: ErrInvalidMonth,
		},
		{
			name:  "another misspelled month",
			input: "010100ZJJANUARY21",
			want:  Time{},
			error: ErrInvalidMonth,
		},
		{
			name:  "yet another misspelled month",
			input: "010100ZJANAURY21",
			want:  Time{},
			error: ErrInvalidMonth,
		},
		{
			name:  "invalid day",
			input: "000100ZJAN21",
			want:  Time{},
			error: ErrInvalidDay,
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
		{
			name:  "invalid timezone",
			input: "010100🇺🇸JAN21",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "20th century year",
			input: "010100ZJAN99",
			want:  NewTime(time.Date(1999, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "21st century year",
			input: "010100ZJAN00",
			want:  NewTime(time.Date(2000, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "invalid year",
			input: "010100ZJAN19401",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "invalid day for month",
			input: "310100ZFEB21",
			want:  Time{},
			error: ErrInvalidDay,
		},
		{
			name:  "valid dtg with 4-digit year",
			input: "010100ZJAN2021",
			want:  NewTime(time.Date(2021, 1, 1, 1, 0, 0, 0, ZULU.Location())),
			error: nil,
		},
		{
			name:  "invalid dtg",
			input: "🇺🇸",
			want:  Time{},
			error: ErrInvalidDateTimeGroup,
		},
		{
			name:  "invalid dtg with valid month",
			input: "JAN",
			want:  Time{},
			error: ErrNotEnoughChars,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDTGBytes(tt.input)
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

func BenchmarkParseDTG(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t, err := ParseDTG("010100ZJAN21")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}

		_ = t
	}
}

func BenchmarkParseDTGBytes(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		t, err := ParseDTGBytes("010100ZJAN21")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}

		_ = t
	}
}
