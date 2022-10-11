//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SkipWhileTest.cs

func TestSkipMust_int(t *testing.T) {
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

func TestSkipWhileMust_string(t *testing.T) {
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

func TestSkipWhileIdxMust_string(t *testing.T) {
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
				predicate: func(s string, idx int) bool { return idx+len(s) > 4 },
			},
			want: NewEnSlice("zero", "one", "two", "three", "four", "five"),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) > idx },
			},
			want: NewEnSlice("four", "five"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, _ int) bool { return len(s) < 100 },
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

// see the example from Enumerable.Skip help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip#examples
func ExampleSkipMust() {
	grades := NewEnSlice(59, 82, 70, 56, 92, 98, 85)
	orderedGrades := OrderByDescMust(grades)
	lowerGrades := SkipMust[int](orderedGrades, 3)
	fmt.Println("All grades except the top three are:")
	enr := lowerGrades.GetEnumerator()
	for enr.MoveNext() {
		grade := enr.Current()
		fmt.Println(grade)
	}
	// Output:
	// All grades except the top three are:
	// 82
	// 70
	// 59
	// 56
}

// see the second example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile
func ExampleSkipWhileMust() {
	grades := NewEnSlice(59, 82, 70, 56, 92, 98, 85)
	orderedGrades := OrderByDescMust(grades)
	lowerGrades := SkipWhileMust[int](orderedGrades,
		func(grade int) bool { return grade >= 80 },
	)
	fmt.Println("All grades below 80:")
	enr := lowerGrades.GetEnumerator()
	for enr.MoveNext() {
		grade := enr.Current()
		fmt.Println(grade)
	}
	// Output:
	// All grades below 80:
	// 70
	// 59
	// 56
}

// see the first example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile
func ExampleSkipWhileIdxMust() {
	amounts := NewEnSlice(5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500)
	skipWhileIdx := SkipWhileIdxMust(amounts,
		func(amount, index int) bool { return amount > index*1000 },
	)
	enr := skipWhileIdx.GetEnumerator()
	for enr.MoveNext() {
		amount := enr.Current()
		fmt.Println(amount)
	}
	// Output:
	// 4000
	// 1500
	// 5500
}
