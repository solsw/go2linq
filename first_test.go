package go2linq

import (
	"fmt"
	"reflect"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstOrDefaultTest.cs

func TestFirst_int(t *testing.T) {
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := First(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("First() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstPred_int(t *testing.T) {
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source:    SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 5 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOrDefaultMust_int(t *testing.T) {
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstOrDefaultMust(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefaultMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOrDefaultPred_int(t *testing.T) {
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
		{name: "NullSourceWithoutPredicate",
			wantErr:     true,
			expectedErr: ErrNilSource,
		},
		{name: "NullSourceWithPredicate",
			args:        args{predicate: func(x int) bool { return x > 3 }},
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source:    SelectMust(NewEnSlice(15, 1, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 5 },
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstOrDefaultPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrDefaultPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("FirstOrDefaultPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefaultPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the second example from Enumerable.First help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first
func ExampleFirstMust() {
	numbers := NewEnSlice(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19)
	first := FirstMust(numbers)
	fmt.Println(first)
	// Output:
	// 9
}

// see the first two examples from Enumerable.FirstOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
func ExampleFirstOrDefaultMust() {
	numbers := NewEnSlice([]int{}...)
	firstOrDefault := FirstOrDefaultMust(numbers)
	fmt.Println(firstOrDefault)

	months := NewEnSlice([]int{}...)
	// Setting the default value to 1 after the query.
	firstOrDefault1 := FirstOrDefaultMust(months)
	if firstOrDefault1 == 0 {
		firstOrDefault1 = 1
	}
	fmt.Printf("The value of the firstMonth1 variable is %v\n", firstOrDefault1)

	// Setting the default value to 1 by using DefaultIfEmptyDef() in the query.
	firstOrDefault2 := FirstMust(DefaultIfEmptyDefMust(months, 1))
	fmt.Printf("The value of the firstMonth2 variable is %v\n", firstOrDefault2)
	// Output:
	// 0
	// The value of the firstMonth1 variable is 1
	// The value of the firstMonth2 variable is 1
}

// see the last example from Enumerable.FirstOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
func ExampleFirstOrDefaultPredMust() {
	names := NewEnSlice("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	firstLongName := FirstOrDefaultPredMust(names, func(name string) bool { return len(name) > 20 })
	fmt.Printf("The first long name is '%v'.\n", firstLongName)

	firstVeryLongName := FirstOrDefaultPredMust(names, func(name string) bool { return len(name) > 30 })
	var what string
	if firstVeryLongName == "" {
		what = "not a"
	} else {
		what = "a"
	}
	fmt.Printf("There is %v name longer than 30 characters.\n", what)
	// Output:
	// The first long name is 'Andersen, Henriette Thaulow'.
	// There is not a name longer than 30 characters.
}

// see the first example from Enumerable.First help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first
func ExampleFirstPredMust() {
	numbers := NewEnSlice(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19)
	firstPred := FirstPredMust(numbers, func(number int) bool { return number > 80 })
	fmt.Println(firstPred)
	// Output:
	// 92
}
