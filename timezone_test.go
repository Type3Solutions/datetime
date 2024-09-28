package mildtg

import (
	"errors"
	"reflect"
	"testing"
)

func Test_parseTimeZone(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    timeZone
		wantErr error
	}{
		{
			name:    "Zulu",
			args:    args{s: "270000Z JAN 20"},
			want:    ZULU,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimeZone(tt.args.s)

			if errors.Is(err, tt.wantErr) {
				t.Errorf("parseTimeZone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTimeZone() got = %v, want %v", got, tt.want)
			}
		})
	}
}
