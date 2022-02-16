//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipWhileTest.cs

func Test_SkipMust_int(t *testing.T) {
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
			want: NewEnSlice(0, 1, 2, 3, 4),
		},
		{name: "ZeroCount",
			args: args{
				source: RangeMust(0, 5),
				count:  0,
			},
			want: NewEnSlice(0, 1, 2, 3, 4),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: RangeMust(0, 5),
				count:  3,
			},
			want: NewEnSlice(3, 4),
		},
		{name: "CountEqualToSourceLength",
			args: args{
				source: RangeMust(0, 5),
				count:  5,
			},
			want: Empty[int](),
		},
		{name: "CountGreaterThanSourceLength",
			args: args{
				source: RangeMust(0, 5),
				count:  100,
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SkipMust(tt.args.source, tt.args.count)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SkipMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SkipWhileMust_string(t *testing.T) {
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
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) > 4 },
			},
			want: NewEnSlice("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 5 },
			},
			want: NewEnSlice("three", "four", "five"),
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 100 },
			},
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SkipWhileMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SkipWhileMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func Test_SkipWhileIdxMust_string(t *testing.T) {
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
			want: NewEnSlice("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) > i },
			},
			want: NewEnSlice("four", "five"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) < 100 },
			},
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SkipWhileIdxMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SkipWhileIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}
