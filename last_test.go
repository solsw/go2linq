package go2linq

import (
	"fmt"
	"iter"
	"math"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/LastOrDefaultTest.cs

func TestLast_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
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
				source: VarAll(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: VarAll(5, 10),
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
		source    iter.Seq[int]
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
				source: VarAll(1, 2, 3, 4),
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
				source:    VarAll(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    VarAll(2),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    VarAll(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrNoMatch,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    VarAll(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    VarAll(1, 2, 5, 10, 2, 1),
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

func TestLastOrDefault_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
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
				source: VarAll(5),
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithoutPredicate",
			args: args{
				source: VarAll(5, 10),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := LastOrDefault(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastOrDefaultPred_int(t *testing.T) {
	type args struct {
		source    iter.Seq[int]
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
				source: VarAll(1, 2, 3, 4),
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
				source:    VarAll(5),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			args: args{
				source:    VarAll(2),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			args: args{
				source:    VarAll(1, 2, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 0,
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			args: args{
				source:    VarAll(1, 2, 5, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    VarAll(1, 2, 5, 10, 2, 1),
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

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func ExampleLast() {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 67, 12, 19}
	last, _ := Last(SliceAll(numbers))
	fmt.Println(last)
	// Output:
	// 19
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func ExampleLastPred() {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 67, 12, 19}
	lastPred, _ := LastPred(
		SliceAll(numbers),
		func(number int) bool { return number > 80 },
	)
	fmt.Println(lastPred)
	// Output:
	// 87
}

// first two examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
func ExampleLastOrDefault() {
	fruits := []string{}
	last, _ := LastOrDefault(SliceAll(fruits))
	if last == "" {
		fmt.Println("<string is empty>")
	} else {
		fmt.Println(last)
	}

	daysOfMonth := []int{}
	// Setting the default value to 1 after the query.
	lastDay1, _ := LastOrDefault(SliceAll(daysOfMonth))
	if lastDay1 == 0 {
		lastDay1 = 1
	}
	fmt.Printf("The value of the lastDay1 variable is %v\n", lastDay1)

	// Setting the default value to 1 by using DefaultIfEmptyDef() in the query.
	defaultIfEmptyDef, _ := DefaultIfEmptyDef(SliceAll(daysOfMonth), 1)
	lastDay2, _ := Last(defaultIfEmptyDef)
	fmt.Printf("The value of the lastDay2 variable is %d\n", lastDay2)
	// Output:
	// <string is empty>
	// The value of the lastDay1 variable is 1
	// The value of the lastDay2 variable is 1
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
func ExampleLastOrDefaultPred() {
	numbers := []float64{49.6, 52.3, 51.0, 49.4, 50.2, 48.3}
	last50, _ := LastOrDefaultPred(
		SliceAll(numbers),
		func(n float64) bool { return math.Round(n) == 50.0 },
	)
	fmt.Printf("The last number that rounds to 50 is %v.\n", last50)

	last40, _ := LastOrDefaultPred(
		SliceAll(numbers),
		func(n float64) bool { return math.Round(n) == 40.0 },
	)
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
