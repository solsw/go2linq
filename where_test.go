package go2linq

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/WhereTest.cs

func TestWhere_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
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
				source:    NewEnSlice(1, 2, 3, 4),
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
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return false },
			},
			want: Empty[int](),
		},
		{name: "AlwaysTruePredicate",
			args: args{
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return true },
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "SimpleFiltering1",
			args: args{
				source:    NewEnSlice(1, 3, 4, 2, 8, 1),
				predicate: func(i int) bool { return i < 4 },
			},
			want: NewEnSlice(1, 3, 2, 1),
		},
		{name: "SimpleFiltering2",
			args: args{
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(i int) bool { return i%2 == 1 },
			},
			want: NewEnSlice(1, 3),
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
				if err != tt.expectedErr {
					t.Errorf("Where() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("Where() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestWhereMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "AlwaysTruePredicate",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four", "five"),
				predicate: func(string) bool { return true },
			},
			want: NewEnSlice("one", "two", "three", "four", "five"),
		},
		{name: "SimpleFiltering",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return strings.HasPrefix(s, "t") },
			},
			want: NewEnSlice("two", "three"),
		},
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/filtering-data#query-expression-syntax-example
		{name: "Where",
			args: args{
				source:    NewEnSlice("the", "quick", "brown", "fox", "jumps"),
				predicate: func(s string) bool { return len(s) == 3 },
			},
			want: NewEnSlice("the", "fox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WhereMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("WhereMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestWhereIdx_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int, int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        Enumerable[int]
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
				source:    NewEnSlice(1, 3, 7, 9, 10),
				predicate: nil,
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "WithIndexSimpleFiltering",
			args: args{
				source:    NewEnSlice(1, 3, 4, 2, 8, 1),
				predicate: func(x, idx int) bool { return x < idx },
			},
			want: NewEnSlice(2, 1),
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
				if err != tt.expectedErr {
					t.Errorf("WhereIdx() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("WhereIdx() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestWhereIdxMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SimpleFiltering",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four", "five"),
				predicate: func(s string, idx int) bool { return len(s) == idx },
			},
			want: NewEnSlice("five"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WhereIdxMust(tt.args.source, tt.args.predicate)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("WhereIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the first example from Enumerable.Where help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.where
func ExampleWhereMust_ex1() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	where := WhereMust(
		NewEnSlice(fruits...),
		func(fruit string) bool { return len(fruit) < 6 },
	)
	enr := where.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
	// Output:
	// apple
	// mango
	// grape
}

func ExampleWhereMust_ex2() {
	fmt.Println(ToStringDef(
		WhereMust(
			RangeMust(1, 10),
			func(i int) bool { return i%2 == 0 },
		),
	))
	fmt.Println(ToStringDef(
		WhereMust(
			NewEnSlice("one", "two", "three", "four", "five"),
			func(s string) bool { return strings.HasSuffix(s, "e") },
		),
	))
	// Output:
	// [2 4 6 8 10]
	// [one three five]
}

// see the last example from Enumerable.Where help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.where
func ExampleWhereIdxMust_ex1() {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}
	query := WhereIdxMust(
		NewEnSlice(numbers...),
		func(number, index int) bool { return number <= index*10 },
	)
	_ = ForEach(context.Background(), query,
		func(number int) error {
			fmt.Println(number)
			return nil
		},
	)
	// Output:
	// 0
	// 20
	// 15
	// 40
}

func ExampleWhereIdxMust_ex2() {
	fmt.Println(ToStringDef(
		WhereIdxMust(
			NewEnSlice("one", "two", "three", "four", "five"),
			func(s string, i int) bool { return len(s) == i },
		),
	))
	fmt.Println(ToStringDef(
		WhereIdxMust(
			ReverseMust(
				NewEnSlice("one", "two", "three", "four", "five"),
			),
			func(s string, i int) bool { return len(s) == i },
		),
	))
	fmt.Println(ToStringDef(
		WhereIdxMust[string](
			OrderByMust(
				NewEnSlice("one", "two", "three", "four", "five"),
			),
			func(s string, i int) bool { return len(s) > i },
		),
	))
	// Output:
	// [five]
	// [two]
	// [five four one three]
}
