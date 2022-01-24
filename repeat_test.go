//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/RepeatTest.cs

func Test_Repeat_string(t *testing.T) {
	type args struct {
		element string
		count   int
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[string]
		wantErr     bool
		expectedErr error
	}{
		{name: "SimpleRepeat",
			args: args{
				element: "foo",
				count:   3,
			},
			want: NewEnSlice("foo", "foo", "foo"),
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
			got, err := Repeat(tt.args.element, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repeat() error = '%v', wantErr '%v'", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Repeat() error = '%v', expectedErr '%v'", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Repeat() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}

func Test_RepeatMust_int(t *testing.T) {
	type args struct {
		element int
		count   int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
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
			want: NewEnSlice(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RepeatMust(tt.args.element, tt.args.count)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("RepeatMust() = '%v', want '%v'", EnToString(got), EnToString(tt.want))
			}
		})
	}
}
