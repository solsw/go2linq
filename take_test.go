package go2linq

import (
	"fmt"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/TakeWhileTest.cs

func TestTakeMust_int(t *testing.T) {
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
				t.Errorf("TakeMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestTakeWhileMust_string(t *testing.T) {
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
				t.Errorf("TakeWhileMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestTakeWhileIdxMust_string(t *testing.T) {
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
			want: Empty[string](),
		},
		{name: "PredicateWithIndexMatchingSomeElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) != idx },
			},
			want: NewEnSlice("zero", "one", "two", "three"),
		},
		{name: "PredicateWithIndexMatchingAllElements",
			args: args{
				source:    NewEnSlice("zero", "one", "two", "three", "four", "five"),
				predicate: func(s string, _ int) bool { return len(s) < 100 },
			},
			want: NewEnSlice("zero", "one", "two", "three", "four", "five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeWhileIdxMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("TakeWhileIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the example from Enumerable.Take help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.take
func ExampleTakeMust() {
	grades := []int{59, 82, 70, 56, 92, 98, 85}
	orderedGrades := OrderByDescMust(
		NewEnSliceEn(grades...),
	)
	topThreeGrades := TakeMust[int](orderedGrades, 3)
	fmt.Println("The top three grades are:")
	enr := topThreeGrades.GetEnumerator()
	for enr.MoveNext() {
		grade := enr.Current()
		fmt.Println(grade)
	}
	// Output:
	// The top three grades are:
	// 98
	// 92
	// 85
}

// see the second example from Enumerable.TakeWhile help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func ExampleTakeWhileMust() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	takeWhile := TakeWhileMust(
		NewEnSliceEn(fruits...),
		func(fruit string) bool {
			return collate.CaseInsensitiveOrder.Compare("orange", fruit) != 0
		},
	)
	enr := takeWhile.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// banana
	// mango
}

// see the first example from Enumerable.TakeWhile help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func ExampleTakeWhileIdxMust() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	takeWhileIdx := TakeWhileIdxMust(
		NewEnSliceEn(fruits...),
		func(fruit string, index int) bool {
			return len(fruit) >= index
		},
	)
	enr := takeWhileIdx.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// passionfruit
	// banana
	// mango
	// orange
	// blueberry
}
