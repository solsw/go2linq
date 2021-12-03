//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ReverseTest.cs

func Test_ReverseMust_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "ReversedRange",
			args: args{
				source: RangeMust(5, 5),
			},
			want: NewOnSlice(9, 8, 7, 6, 5),
		},
		{name: "EmptyInput",
			args: args{
				source: Empty[int](),
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseMust(tt.args.source); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ReverseMust() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_ReverseMust_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "ReversedStrs",
			args: args{
				source: NewOnSlice("one", "two", "three", "four", "five"),
			},
			want: NewOnSlice("five", "four", "three", "two", "one"),
		},
		{name: "1",
			args: args{
				source: NewOnSlice("1"),
			},
			want: NewOnSlice("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseMust(tt.args.source); !SequenceEqualMust(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("ReverseMust() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
