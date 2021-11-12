//go:build go1.18

package go2linq

import (
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeWhileTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipWhileTest.cs

func Test_Take_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
		count int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "NegativeCount",
			args: args{
				source: Range(0, 5),
				count: -5,
			},
			want: Empty[int](),
		},
		{name: "ZeroCount",
			args: args{
				source: Range(0, 5),
				count: 0,
			},
			want: Empty[int](),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: Range(0, 5),
				count: 3,
			},
			want: NewOnSlice(0, 1, 2),
		},
		{name: "CountShorterThanSource2",
			args: args{
				source: NewOnSlice(1, 2, 3, 4),
				count: 3,
			},
			want: NewOnSlice(1, 2, 3),
		},
		{name: "CountEqualToSourceLength",
			args: args{
				source: Range(1, 5),
				count: 5,
			},
			want: NewOnSlice(1, 2, 3, 4, 5),
		},
		{name: "CountGreaterThanSourceLength",
			args: args{
				source: Range(2, 5),
				count: 100,
			},
			want: NewOnSlice(2, 3, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Take(tt.args.source, tt.args.count); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Take() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_Skip_int(t *testing.T) {
	type args struct {
		source Enumerator[int]
		count int
	}
	tests := []struct {
		name string
		args args
		want Enumerator[int]
	}{
		{name: "NegativeCount",
			args: args{
				source: Range(0, 5),
				count: -5,
			},
			want: NewOnSlice(0, 1, 2, 3, 4),
		},
		{name: "ZeroCount",
			args: args{
				source: Range(0, 5),
				count: 0,
			},
			want: NewOnSlice(0, 1, 2, 3, 4),
		},
		{name: "CountShorterThanSource",
			args: args{
				source: Range(0, 5),
				count: 3,
			},
			want: NewOnSlice(3, 4),
		},
		{name: "CountEqualToSourceLength",
			args: args{
				source: Range(0, 5),
				count: 5,
			},
			want: Empty[int](),
		},
		{name: "CountGreaterThanSourceLength",
			args: args{
				source: Range(0, 5),
				count: 100,
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Skip(tt.args.source, tt.args.count); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("Skip() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_TakeWhile_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "PredicateFailingFirstElement",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five", "six"),
				predicate: func(s string) bool { return len(s) > 4 },
			},
			want: Empty[string](),
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 5 },
			},
			want: NewOnSlice("zero", "one", "two"),
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 100 },
			},
			want: NewOnSlice("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TakeWhile(tt.args.source, tt.args.predicate); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("TakeWhile() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_TakeWhileIdx_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "PredicateWithIndexFailingFirstElement",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return i+len(s) > 4 },
			},
			want: Empty[string](),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) != i },
			},
			want: NewOnSlice("zero", "one", "two", "three"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) < 100 },
			},
			want: NewOnSlice("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TakeWhileIdx(tt.args.source, tt.args.predicate); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("TakeWhileIdx() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SkipWhile_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "PredicateFailingFirstElement",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) > 4 },
			},
			want: NewOnSlice("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 5 },
			},
			want: NewOnSlice("three", "four", "five"),
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return len(s) < 100 },
			},
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SkipWhile(tt.args.source, tt.args.predicate); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SkipWhile() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}

func Test_SkipWhileIdx_string(t *testing.T) {
	type args struct {
		source Enumerator[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerator[string]
	}{
		{name: "PredicateWithIndexFailingFirstElement",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return i+len(s) > 4 },
			},
			want: NewOnSlice("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) > i },
			},
			want: NewOnSlice("four", "five"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source: NewOnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) < 100 },
			},
			want: Empty[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SkipWhileIdx(tt.args.source, tt.args.predicate); !SequenceEqual(got, tt.want) {
				got.Reset()
				tt.want.Reset()
				t.Errorf("SkipWhileIdx() = '%v', want '%v'", String(got), String(tt.want))
			}
		})
	}
}
