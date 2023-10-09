package go2linq

import (
	"fmt"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SequenceEqualTest.cs

func TestSequenceEqualMust_int(t *testing.T) {
	r0 := RangeMust(0, 0)
	r1 := RangeMust(0, 1)
	r2 := RangeMust(0, 2)
	r3 := RepeatMust(1, 4)
	type args struct {
		first  Enumerable[int]
		second Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "EmptyEmpty",
			args: args{
				first:  Empty[int](),
				second: Empty[int](),
			},
			want: true,
		},
		{name: "EmptySecond",
			args: args{
				first:  NewEnSlice(1),
				second: Empty[int](),
			},
			want: false,
		},
		{name: "EmptyFirst",
			args: args{
				first:  Empty[int](),
				second: NewEnSlice(2),
			},
			want: false,
		},
		{name: "EqualSequences",
			args: args{
				first:  NewEnSlice(1),
				second: NewEnSlice(1),
			},
			want: true,
		},
		{name: "UnequalLengthsBothArrays",
			args: args{
				first:  NewEnSlice(1, 5, 3),
				second: NewEnSlice(1, 5, 3, 10),
			},
			want: false,
		},
		{name: "UnequalLengthsBothRangesFirstLonger",
			args: args{
				first:  RangeMust(0, 11),
				second: RangeMust(0, 10),
			},
			want: false,
		},
		{name: "UnequalLengthsBothRangesSecondLonger",
			args: args{
				first:  RangeMust(0, 10),
				second: RangeMust(0, 11),
			},
			want: false,
		},
		{name: "UnequalData",
			args: args{
				first:  NewEnSlice(1, 5, 3, 9),
				second: NewEnSlice(1, 5, 3, 10),
			},
			want: false,
		},
		{name: "EqualDataBothArrays",
			args: args{
				first:  NewEnSlice(1, 5, 3, 10),
				second: NewEnSlice(1, 5, 3, 10),
			},
			want: true,
		},
		{name: "EqualDataBothRanges",
			args: args{
				first:  RangeMust(0, 10),
				second: RangeMust(0, 10),
			},
			want: true,
		},
		{name: "OrderMatters",
			args: args{
				first:  NewEnSlice(1, 2),
				second: NewEnSlice(2, 1),
			},
			want: false,
		},
		{name: "ReturnAtFirstDifference",
			args: args{
				first: SelectMust(
					NewEnSlice(1, 5, 10, 2, 0),
					func(i int) int { return 10 / i },
				),
				second: SelectMust(
					NewEnSlice(1, 5, 10, 1, 0),
					func(i int) int { return 10 / i },
				),
			},
			want: false,
		},
		{name: "EqualQueries",
			args: args{
				first:  SkipMust(RangeMust(0, 8), 4),
				second: TakeMust(RangeMust(4, 8), 4),
			},
			want: true,
		},
		{name: "Same0",
			args: args{
				first:  r0,
				second: r0,
			},
			want: true,
		},
		{name: "Same1",
			args: args{
				first:  r1,
				second: r1,
			},
			want: true,
		},
		{name: "Same2",
			args: args{
				first:  r2,
				second: r2,
			},
			want: true,
		},
		{name: "Same3",
			args: args{
				first:  TakeMust(r3, 2),
				second: SkipMust(r3, 2),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SequenceEqualMust(tt.args.first, tt.args.second)
			if got != tt.want {
				t.Errorf("SequenceEqualMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequenceEqualMust_string(t *testing.T) {
	type args struct {
		first  Enumerable[string]
		second Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "2",
			args: args{
				first:  NewEnSlice("one", "two", "three", "four"),
				second: NewEnSlice("one", "two", "three", "four"),
			},
			want: true,
		},
		{name: "4",
			args: args{
				first:  NewEnSlice("a", "b"),
				second: NewEnSlice("a"),
			},
			want: false,
		},
		{name: "5",
			args: args{
				first:  NewEnSlice("a"),
				second: NewEnSlice("a", "b"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SequenceEqualMust(tt.args.first, tt.args.second)
			if got != tt.want {
				t.Errorf("SequenceEqualMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequenceEqualEqMust_string(t *testing.T) {
	type args struct {
		first   Enumerable[string]
		second  Enumerable[string]
		equaler collate.Equaler[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				first:   NewEnSlice("a", "b"),
				second:  NewEnSlice("a", "B"),
				equaler: collate.CaseInsensitiveOrder,
			},
			want: true,
		},
		{name: "CustomEqualityComparer",
			args: args{
				first:   NewEnSlice("foo", "BAR", "baz"),
				second:  NewEnSlice("FOO", "bar", "Baz"),
				equaler: collate.CaseInsensitiveOrder,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SequenceEqualEqMust(tt.args.first, tt.args.second, tt.args.equaler)
			if got != tt.want {
				t.Errorf("SequenceEqualEqMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see SequenceEqualEx1 example from Enumerable.SequenceEqual help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func ExampleSequenceEqualMust() {
	pet1 := Pet{Name: "Turbo", Age: 2}
	pet2 := Pet{Name: "Peanut", Age: 8}
	pets1 := []Pet{pet1, pet2}
	pets2 := []Pet{pet1, pet2}
	equal := SequenceEqualMust(
		NewEnSlice(pets1...),
		NewEnSlice(pets2...),
	)
	var what string
	if equal {
		what = "are"
	} else {
		what = "are not"
	}
	fmt.Printf("The lists %s equal.\n", what)
	// Output:
	// The lists are equal.
}

// see the last example from Enumerable.SequenceEqual help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func ExampleSequenceEqualEqMust() {
	storeA := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	storeB := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	equalEq := SequenceEqualEqMust(
		NewEnSlice(storeA...),
		NewEnSlice(storeB...),
		collate.Equaler[Product](
			collate.EqualerFunc[Product](
				func(p1, p2 Product) bool {
					return p1.Code == p2.Code && p1.Name == p2.Name
				},
			),
		),
	)
	fmt.Printf("Equal? %t\n", equalEq)
	// Output:
	// Equal? true
}
