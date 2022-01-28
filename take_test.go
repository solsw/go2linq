//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeWhileTest.cs

func Test_TakeMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
		count  int
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "NegativeCount",
			args: args{
				source: RangeMust(0, 5),
				count:  -5,
			},
			want: Empty[int](),
		},
		{name: "ZeroCount",
			args: args{
				source: RangeMust(0, 5),
				count:  0,
			},
			want: Empty[int](),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: RangeMust(0, 5),
				count:  3,
			},
			want: NewEnSlice(0, 1, 2),
		},
		{name: "CountShorterThanSource2",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				count:  3,
			},
			want: NewEnSlice(1, 2, 3),
		},
		{name: "CountEqualToSourceLength",
			args: args{
				source: RangeMust(1, 5),
				count:  5,
			},
			want: NewEnSlice(1, 2, 3, 4, 5),
		},
		{name: "CountGreaterThanSourceLength",
			args: args{
				source: RangeMust(2, 5),
				count:  100,
			},
			want: NewEnSlice(2, 3, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeMust(tt.args.source, tt.args.count)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("TakeMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_TakeWhileMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "PredicateFailingFirstElement",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five", "six"),
				predicate: func(s string) bool { return len(s) > 4 },
			},
			want: Empty[string](),
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 5 },
			},
			want: NewEnSlice("zero", "one", "two"),
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 100 },
			},
			want: NewEnSlice("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeWhileMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("TakeWhileMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}

func Test_TakeWhileIdxMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "PredicateWithIndexFailingFirstElement",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return i+len(s) > 4 },
			},
			want: Empty[string](),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) != i },
			},
			want: NewEnSlice("zero", "one", "two", "three"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) < 100 },
			},
			want: NewEnSlice("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeWhileIdxMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("TakeWhileIdxMust() = %v, want %v", ToString(got), ToString(tt.want))
			}
		})
	}
}
