//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RepeatTest.cs

func Test_RepeatErr_string(t *testing.T) {
	type args struct {
		element string
		count   int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerator[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "SimpleRepeat",
			args: args{
				element: "foo",
				count:   3,
			},
			want: NewOnSlice("foo", "foo", "foo"),
		},
		{name: "EmptyRepeat",
			args: args{
				element: "foo",
				count:   0,
			},
			want: Empty[string](),
		},
		{name: "NegativeCount",
			args: args{
				element: "foo",
				count:   -1,
			},
			wantErr:     true,
			expectedErr: ErrNegativeCount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RepeatErr(tt.args.element, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepeatErr() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("RepeatErr() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("RepeatErr() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Repeat_int(t *testing.T) {
	type args struct {
		element int
		count   int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "1",
			args: args{
				element: 0,
				count:   0,
			},
			want: Empty[int](),
		},
		{name: "2",
			args: args{
				element: 2,
				count:   2,
			},
			want: NewOnSlice(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.element, tt.args.count); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Repeat() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
