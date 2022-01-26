//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/ReverseTest.cs

func Test_ReverseMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "ReversedRange",
			args: args{
				source: RangeMust(5, 5),
			},
			want: NewEnSlice(9, 8, 7, 6, 5),
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
			got := ReverseMust(tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ReverseMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_ReverseMust_string(t *testing.T) {
	type args struct {
		source Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "ReversedStrs",
			args: args{
				source: NewEnSlice("one", "two", "three", "four", "five"),
			},
			want: NewEnSlice("five", "four", "three", "two", "one"),
		},
		{name: "1",
			args: args{
				source: NewEnSlice("1"),
			},
			want: NewEnSlice("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseMust(tt.args.source)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("ReverseMust() = '%v', want '%v'", ToString(got), ToString(tt.want))
			}
		})
	}
}
