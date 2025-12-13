package go2linq

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"strings"
	"testing"

	"github.com/solsw/iterhelper"
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
				source:    iterhelper.Var(1, 2, 3, 4),
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
				source:    iterhelper.Var(1, 2, 3, 4),
				predicate: func(int) bool { return false },
			},
			want: Empty[int](),
		},
		{name: "AlwaysTruePredicate",
			args: args{
				source:    iterhelper.Var(1, 2, 3, 4),
				predicate: func(int) bool { return true },
			},
			want: iterhelper.Var(1, 2, 3, 4),
		},
		{name: "SimpleFiltering1",
			args: args{
				source:    iterhelper.Var(1, 3, 4, 2, 8, 1),
				predicate: func(i int) bool { return i < 4 },
			},
			want: iterhelper.Var(1, 3, 2, 1),
		},
		{name: "SimpleFiltering2",
			args: args{
				source:    iterhelper.Var(1, 2, 3, 4),
				predicate: func(i int) bool { return i%2 == 1 },
			},
			want: iterhelper.Var(1, 3),
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
				t.Errorf("Where() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
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
				source:    iterhelper.Var("one", "two", "three", "four", "five"),
				predicate: func(string) bool { return true },
			},
			want: iterhelper.Var("one", "two", "three", "four", "five"),
		},
		{name: "SimpleFiltering",
			args: args{
				source:    iterhelper.Var("one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return strings.HasPrefix(s, "t") },
			},
			want: iterhelper.Var("two", "three"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/filtering-data#query-expression-syntax-example
		{name: "Where",
			args: args{
				source:    iterhelper.Var("the", "quick", "brown", "fox", "jumps"),
				predicate: func(s string) bool { return len(s) == 3 },
			},
			want: iterhelper.Var("the", "fox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Where(tt.args.source, tt.args.predicate)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Where() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
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
				source:    iterhelper.Var(1, 3, 7, 9, 10),
				predicate: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "WithIndexSimpleFiltering",
			args: args{
				source:    iterhelper.Var(1, 3, 4, 2, 8, 1),
				predicate: func(x, idx int) bool { return x < idx },
			},
			want: iterhelper.Var(2, 1),
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
				t.Errorf("WhereIdx() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
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
				source:    iterhelper.Var("one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) == idx },
			},
			want: iterhelper.Var("five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := WhereIdx(tt.args.source, tt.args.predicate)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("WhereIdx() = %v, want %v", iterhelper.StringDef(got), iterhelper.StringDef(tt.want))
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
	fmt.Println(iterhelper.StringDef(where1))
	where2, _ := Where(
		iterhelper.Var("one", "two", "three", "four", "five"),
		func(s string) bool { return strings.HasSuffix(s, "e") },
	)
	fmt.Println(iterhelper.StringDef(where2))
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
		iterhelper.Var("one", "two", "three", "four", "five"),
		func(s string, i int) bool { return len(s) == i },
	)
	fmt.Println(iterhelper.StringDef(whereIdx1))
	reverse2, _ := Reverse(iterhelper.Var("one", "two", "three", "four", "five"))
	whereIdx2, _ := WhereIdx(
		reverse2,
		func(s string, i int) bool { return len(s) == i },
	)
	fmt.Println(iterhelper.StringDef(whereIdx2))
	orderBy, _ := OrderBy(iterhelper.Var("one", "two", "three", "four", "five"))
	whereIdx3, _ := WhereIdx(
		orderBy,
		func(s string, i int) bool { return len(s) > i },
	)
	fmt.Println(iterhelper.StringDef(whereIdx3))
	// Output:
	// [five]
	// [two]
	// [five four one three]
}
