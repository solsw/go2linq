package go2linq

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastOrDefaultTest.cs

func TestLast_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithoutPredicate",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "EmptySequenceWithoutPredicate",
			args: args{
				source: Empty[int](),
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SingleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5, 10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Last(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Last() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullSourceWithPredicate",
			args: args{
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullPredicate",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceWithPredicate",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrEmptySource,
		},
		{name: "SingleElementSequenceWithMatchingPredicate",
			args: args{
				source:    NewEnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewEnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewEnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("LastPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastOrDefaultMust_int(t *testing.T) {
	type args struct {
		source Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "EmptySequenceWithoutPredicate",
			args: args{
				source: Empty[int](),
			},
			want: 0,
		},
		{name: "SingleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(5, 10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastOrDefaultMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefaultMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastOrDefaultPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{name: "NullPredicate",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceWithPredicate",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "SingleElementSequenceWithMatchingPredicate",
			args: args{
				source:    NewEnSlice(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    NewEnSlice(2),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    NewEnSlice(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    NewEnSlice(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LastOrDefaultPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastOrDefaultPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("LastOrDefaultPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefaultPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the first example from Enumerable.Last help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func ExampleLastMust() {
	numbers := NewEnSlice(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 67, 12, 19)
	last := LastMust(numbers)
	fmt.Println(last)
	// Output:
	// 19
}

// see the last example from Enumerable.Last help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func ExampleLastPredMust() {
	numbers := NewEnSlice(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 67, 12, 19)
	lastPred := LastPredMust(numbers, func(number int) bool { return number > 80 })
	fmt.Println(lastPred)
	// Output:
	// 87
}

// see the first two examples from Enumerable.LastOrDefault help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
func ExampleLastOrDefaultMust() {
	fruits := NewEnSlice([]string{}...)
	last := LastOrDefaultMust(fruits)
	if last == "" {
		fmt.Println("<string is empty>")
	} else {
		fmt.Println(last)
	}

	daysOfMonth := NewEnSlice([]int{}...)
	// Setting the default value to 1 after the query.
	lastDay1 := LastOrDefaultMust(daysOfMonth)
	if lastDay1 == 0 {
		lastDay1 = 1
	}
	fmt.Printf("The value of the lastDay1 variable is %v\n", lastDay1)

	// Setting the default value to 1 by using DefaultIfEmptyDef() in the query.
	lastDay2 := LastMust(DefaultIfEmptyDefMust(daysOfMonth, 1))
	fmt.Printf("The value of the lastDay2 variable is %d\n", lastDay2)
	// Output:
	// <string is empty>
	// The value of the lastDay1 variable is 1
	// The value of the lastDay2 variable is 1
}

// see the last example from Enumerable.LastOrDefault help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
func ExampleLastOrDefaultPredMust() {
	numbers := NewEnSlice(49.6, 52.3, 51.0, 49.4, 50.2, 48.3)
	last50 := LastOrDefaultPredMust(numbers, func(n float64) bool { return math.Round(n) == 50.0 })
	fmt.Printf("The last number that rounds to 50 is %v.\n", last50)

	last40 := LastOrDefaultPredMust(numbers, func(n float64) bool { return math.Round(n) == 40.0 })
	var what string
	if last40 == 0.0 {
		what = "<DOES NOT EXIST>"
	} else {
		what = fmt.Sprint(last40)
	}
	fmt.Printf("The last number that rounds to 40 is %v.\n", what)
	// Output:
	// The last number that rounds to 50 is 50.2.
	// The last number that rounds to 40 is <DOES NOT EXIST>.
}
