//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/CastTest.cs

func Test_CastErr_interface_int(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSource",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "SimpleCasting",
			args: args{
				source: NewOnSlice[interface{}](1, 2, 3, 4),
			},
			want: NewOnSlice[int](1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CastErr[interface{}, int](tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("CastErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("CastErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("CastErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_OfType_interface_int(t *testing.T) {
	type args struct {
		source Enumerator[interface{}]
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[int]
	}{
		{name: "OfType",
			args: args{
				source: NewOnSlice[interface{}](1, 2, "two", 3, 3.14, 4, nil),
			},
			want: NewOnSlice[int](1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OfType[interface{}, int](tt.args.source)
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("OfType() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
