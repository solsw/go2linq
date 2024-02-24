package go2linq

import (
	"fmt"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SequenceEqualTest.cs

func TestSequenceEqual_int(t *testing.T) {
	r0, _ := Range(0, 0)
	r1, _ := Range(0, 1)
	r2, _ := Range(0, 2)
	r3, _ := Repeat(1, 4)
	type args struct {
		first  iter.Seq[int]
		second iter.Seq[int]
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
		{name: "EmptyFirst",
			args: args{
				first:  Empty[int](),
				second: VarAll(2),
			},
			want: false,
		},
		{name: "EmptySecond",
			args: args{
				first:  VarAll(1),
				second: Empty[int](),
			},
			want: false,
		},
		{name: "EqualSequences",
			args: args{
				first:  VarAll(1),
				second: VarAll(1),
			},
			want: true,
		},
		{name: "UnequalLengthsBothArrays",
			args: args{
				first:  VarAll(1, 5, 3),
				second: VarAll(1, 5, 3, 10),
			},
			want: false,
		},
		{name: "UnequalLengthsBothRangesFirstLonger",
			args: args{
				first:  errorhelper.Must(Range(0, 11)),
				second: errorhelper.Must(Range(0, 10)),
			},
			want: false,
		},
		{name: "UnequalLengthsBothRangesSecondLonger",
			args: args{
				first:  errorhelper.Must(Range(0, 10)),
				second: errorhelper.Must(Range(0, 11)),
			},
			want: false,
		},
		{name: "UnequalData",
			args: args{
				first:  VarAll(1, 5, 3, 9),
				second: VarAll(1, 5, 3, 10),
			},
			want: false,
		},
		{name: "EqualDataBothArrays",
			args: args{
				first:  VarAll(1, 5, 3, 10),
				second: VarAll(1, 5, 3, 10),
			},
			want: true,
		},
		{name: "EqualDataBothRanges",
			args: args{
				first:  errorhelper.Must(Range(0, 10)),
				second: errorhelper.Must(Range(0, 10)),
			},
			want: true,
		},
		{name: "OrderMatters",
			args: args{
				first:  VarAll(1, 2),
				second: VarAll(2, 1),
			},
			want: false,
		},
		{name: "ReturnAtFirstDifference",
			args: args{
				first: errorhelper.Must(Select(
					VarAll(1, 5, 10, 2, 0),
					func(i int) int { return 10 / i },
				)),
				second: errorhelper.Must(Select(
					VarAll(1, 5, 10, 1, 0),
					func(i int) int { return 10 / i },
				)),
			},
			want: false,
		},
		{name: "EqualQueries",
			args: args{
				first:  errorhelper.Must(Skip(errorhelper.Must(Range(0, 8)), 4)),
				second: errorhelper.Must(Take(errorhelper.Must(Range(4, 8)), 4)),
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
				first:  errorhelper.Must(Take(r3, 2)),
				second: errorhelper.Must(Skip(r3, 2)),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SequenceEqual(tt.args.first, tt.args.second)
			if got != tt.want {
				t.Errorf("SequenceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequenceEqual_string(t *testing.T) {
	type args struct {
		first  iter.Seq[string]
		second iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "2",
			args: args{
				first:  VarAll("one", "two", "three", "four"),
				second: VarAll("one", "two", "three", "four"),
			},
			want: true,
		},
		{name: "4",
			args: args{
				first:  VarAll("a", "b"),
				second: VarAll("a"),
			},
			want: false,
		},
		{name: "5",
			args: args{
				first:  VarAll("a"),
				second: VarAll("a", "b"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SequenceEqual(tt.args.first, tt.args.second)
			if got != tt.want {
				t.Errorf("SequenceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequenceEqualEq_string(t *testing.T) {
	type args struct {
		first  iter.Seq[string]
		second iter.Seq[string]
		equal  func(string, string) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				first:  VarAll("a", "b"),
				second: VarAll("a", "B"),
				equal:  caseInsensitiveEqual,
			},
			want: true,
		},
		{name: "CustomEqualityComparer",
			args: args{
				first:  VarAll("foo", "BAR", "baz"),
				second: VarAll("FOO", "bar", "Baz"),
				equal:  caseInsensitiveEqual,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SequenceEqualEq(tt.args.first, tt.args.second, tt.args.equal)
			if got != tt.want {
				t.Errorf("SequenceEqualEq() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see SequenceEqualEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func ExampleSequenceEqual() {
	pet1 := Pet{Name: "Turbo", Age: 2}
	pet2 := Pet{Name: "Peanut", Age: 8}
	pets1 := []Pet{pet1, pet2}
	pets2 := []Pet{pet1, pet2}
	sequenceEqual, _ := SequenceEqual(SliceAll(pets1), SliceAll(pets2))
	var what string
	if sequenceEqual {
		what = "are"
	} else {
		what = "are not"
	}
	fmt.Printf("The lists %s equal.\n", what)
	// Output:
	// The lists are equal.
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal
func ExampleSequenceEqualEq() {
	storeA := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	storeB := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
	}
	equalEq, _ := SequenceEqualEq(
		SliceAll(storeA),
		SliceAll(storeB),
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && p1.Name == p2.Name
		},
	)
	fmt.Printf("Equal? %t\n", equalEq)
	// Output:
	// Equal? true
}

func TestSequenceEqual2_int_string(t *testing.T) {
	type args struct {
		first  iter.Seq2[int, string]
		second iter.Seq2[int, string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "EmptyEmpty",
			args: args{
				first:  Empty2[int, string](),
				second: Empty2[int, string](),
			},
			want: true,
		},
		{name: "EmptyFirst",
			args: args{
				first:  Empty2[int, string](),
				second: Sec2_int_string(1),
			},
			want: false,
		},
		{name: "EmptySecond",
			args: args{
				first:  Sec2_int_string(1),
				second: Empty2[int, string](),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SequenceEqual2(tt.args.first, tt.args.second)
			if got != tt.want {
				t.Errorf("SequenceEqual2() = %v, want %v", got, tt.want)
			}
		})
	}
}
