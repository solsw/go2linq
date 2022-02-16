//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/WhereTest.cs

func Test_Where_int(t *testing.T) {
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
				predicate: func(i int) bool { return i > 5 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullPredicateThrowsNullArgumentException",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "SimpleFiltering",
			args: args{
				source:    NewEnSlice(1, 3, 4, 2, 8, 1),
				predicate: func(i int) bool { return i < 4 },
			},
			want: NewEnSlice(1, 3, 2, 1),
		},
		{name: "EmptySource",
			args: args{
				source:    Empty[int](),
				predicate: func(i int) bool { return i > 5 },
			},
			want: Empty[int](),
		},
		{name: "1",
			args: args{
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return false },
			},
			want: Empty[int](),
		},
		{name: "2",
			args: args{
				source:    NewEnSlice(1, 2, 3, 4),
				predicate: func(int) bool { return true },
			},
			want: NewEnSlice(1, 2, 3, 4),
		},
		{name: "3",
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

func Test_WhereMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "4",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four", "five"),
				predicate: func(string) bool { return true },
			},
			want: NewEnSlice("one", "two", "three", "four", "five"),
		},
		{name: "5",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four", "five"),
				predicate: func(s string) bool { return strings.HasPrefix(s, "t") },
			},
			want: NewEnSlice("two", "three"),
		},
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/filtering-data#query-expression-syntax-example
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

func Test_WhereIdx_int(t *testing.T) {
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
				predicate: func(x, index int) bool { return x > 5 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "WithIndexNullPredicateThrowsNullArgumentException",
			args: args{
				source: NewEnSlice(1, 3, 7, 9, 10),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "WithIndexSimpleFiltering",
			args: args{
				source:    NewEnSlice(1, 3, 4, 2, 8, 1),
				predicate: func(x, index int) bool { return x < index },
			},
			want: NewEnSlice(2, 1),
		},
		{name: "WithIndexEmptySource",
			args: args{
				source:    Empty[int](),
				predicate: func(x, index int) bool { return x < 4 },
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

func Test_WhereIdxMust_string(t *testing.T) {
	type args struct {
		source    Enumerable[string]
		predicate func(string, int) bool
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "1",
			args: args{
				source:    NewEnSlice("one", "two", "three", "four", "five"),
				predicate: func(s string, i int) bool { return len(s) == i },
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

func Example_WhereMust() {
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

func Example_WhereIdxMust() {
	fmt.Println(ToStringDef(
		WhereIdxMust(
			NewEnSlice("one", "two", "three", "four", "five"),
			func(s string, i int) bool { return len(s) == i },
		),
	))
	fmt.Println(ToStringDef(
		WhereIdxMust(
			ReverseMust(NewEnSlice("one", "two", "three", "four", "five")),
			func(s string, i int) bool { return len(s) == i },
		),
	))
	fmt.Println(ToStringDef(
		WhereIdxMust[string](
			OrderBySelfMust(NewEnSlice("one", "two", "three", "four", "five")),
			func(s string, i int) bool { return len(s) > i },
		),
	))
	// Output:
	// [five]
	// [two]
	// [five four one three]
}
