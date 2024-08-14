package go2linq

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/WhereTest.cs

func TestWhere_int(t *testing.T) {
	type args struct {
		source    iter.Seq[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceThrowsNullArgumentException",
			args: args{
				source:    nil,
				predicate: func(i int) bool { return i > 5 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullPredicateThrowsNullArgumentException",
			args: args{
				source:    VarToSeq(1, 2, 3, 4),
				predicate: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySource",
			args: args{
				source:    Empty[int](),
				predicate: func(i int) bool { return i > 5 },
			},
			want: Empty[int](),
		},
		{name: "AlwaysFalsePredicate",
			args: args{
				source:    VarToSeq(1, 2, 3, 4),
				predicate: func(int) bool { return false },
			},
			want: Empty[int](),
		},
		{name: "AlwaysTruePredicate",
			args: args{
				source:    VarToSeq(1, 2, 3, 4),
				predicate: func(int) bool { return true },
			},
			want: VarToSeq(1, 2, 3, 4),
		},
		{name: "SimpleFiltering1",
			args: args{
				source:    VarToSeq(1, 3, 4, 2, 8, 1),
				predicate: func(i int) bool { return i < 4 },
			},
			want: VarToSeq(1, 3, 2, 1),
		},
		{name: "SimpleFiltering2",
			args: args{
				source:    VarToSeq(1, 2, 3, 4),
				predicate: func(i int) bool { return i%2 == 1 },
			},
			want: VarToSeq(1, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Where(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Where() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Where() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Where() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestWhere_string(t *testing.T) {
	type args struct {
		source    iter.Seq[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "AlwaysTruePredicate",
			args: args{
				source:    VarToSeq("one", "two", "three", "four", "five"),
				predicate: func(string) bool { return true },
			},
			want: VarToSeq("one", "two", "three", "four", "five"),
		},
		{name: "SimpleFiltering",
			args: args{
				source:    VarToSeq("one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return strings.HasPrefix(s, "t") },
			},
			want: VarToSeq("two", "three"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/filtering-data#query-expression-syntax-example
		{name: "Where",
			args: args{
				source:    VarToSeq("the", "quick", "brown", "fox", "jumps"),
				predicate: func(s string) bool { return len(s) == 3 },
			},
			want: VarToSeq("the", "fox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Where(tt.args.source, tt.args.predicate)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Where() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestWhereIdx_int(t *testing.T) {
	type args struct {
		source    iter.Seq[int]
		predicate func(int, int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        iter.Seq[int]
		wantErr     bool
		expectedErr error
	}{
		{name: "WithIndexNullSourceThrowsNullArgumentException",
			args: args{
				source:    nil,
				predicate: func(x, _ int) bool { return x > 5 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "WithIndexNullPredicateThrowsNullArgumentException",
			args: args{
				source:    VarToSeq(1, 3, 7, 9, 10),
				predicate: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "WithIndexSimpleFiltering",
			args: args{
				source:    VarToSeq(1, 3, 4, 2, 8, 1),
				predicate: func(x, idx int) bool { return x < idx },
			},
			want: VarToSeq(2, 1),
		},
		{name: "WithIndexEmptySource",
			args: args{
				source:    Empty[int](),
				predicate: func(x, _ int) bool { return x < 4 },
			},
			want: Empty[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WhereIdx(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("WhereIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("WhereIdx() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("WhereIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestWhereIdx_string(t *testing.T) {
	type args struct {
		source    iter.Seq[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "SimpleFiltering",
			args: args{
				source:    VarToSeq("one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) == idx },
			},
			want: VarToSeq("five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := WhereIdx(tt.args.source, tt.args.predicate)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("WhereIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.where
func ExampleWhere_ex1() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	where, _ := Where(
		slices.Values(fruits),
		func(fruit string) bool { return len(fruit) < 6 },
	)
	for fruit := range where {
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// mango
	// grape
}

func ExampleWhere_ex2() {
	range1, _ := Range(1, 10)
	where1, _ := Where(
		range1,
		func(i int) bool { return i%2 == 0 },
	)
	fmt.Println(StringDef(where1))
	where2, _ := Where(
		VarToSeq("one", "two", "three", "four", "five"),
		func(s string) bool { return strings.HasSuffix(s, "e") },
	)
	fmt.Println(StringDef(where2))
	// Output:
	// [2 4 6 8 10]
	// [one three five]
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.where
func ExampleWhereIdx_ex1() {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}
	whereIdx, _ := WhereIdx(
		slices.Values(numbers),
		func(number, index int) bool { return number <= index*10 },
	)
	for number := range whereIdx {
		fmt.Println(number)
	}
	// Output:
	// 0
	// 20
	// 15
	// 40
}

func ExampleWhereIdx_ex2() {
	whereIdx1, _ := WhereIdx(
		VarToSeq("one", "two", "three", "four", "five"),
		func(s string, i int) bool { return len(s) == i },
	)
	fmt.Println(StringDef(whereIdx1))
	reverse2, _ := Reverse(VarToSeq("one", "two", "three", "four", "five"))
	whereIdx2, _ := WhereIdx(
		reverse2,
		func(s string, i int) bool { return len(s) == i },
	)
	fmt.Println(StringDef(whereIdx2))
	orderBy, _ := OrderBy(VarToSeq("one", "two", "three", "four", "five"))
	whereIdx3, _ := WhereIdx(
		orderBy,
		func(s string, i int) bool { return len(s) > i },
	)
	fmt.Println(StringDef(whereIdx3))
	// Output:
	// [five]
	// [two]
	// [five four one three]
}
