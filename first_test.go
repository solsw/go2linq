package go2linq

import (
	"fmt"
	"iter"
	"reflect"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/FirstOrDefaultTest.cs

func TestFirst_int(t *testing.T) {
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: errorhelper.Must(Select(
					VarAll(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				)),
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source: errorhelper.Must(Select(
					VarAll(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				)),
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

func TestFirstOrDefault_int(t *testing.T) {
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithoutPredicate",
			args: args{
				source: errorhelper.Must(Select(
					VarAll(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				)),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := FirstOrDefault(tt.args.source)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOrDefaultPred_int(t *testing.T) {
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
			want: 5,
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			args: args{
				source: errorhelper.Must(Select(
					VarAll(15, 1, 0, 3),
					func(x int) int { return 10 / x },
				)),
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

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
func ExampleFirst() {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19}
	first, _ := First(SliceAll(numbers))
	fmt.Println(first)
	// Output:
	// 9
}

// first two examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
func ExampleFirstOrDefault() {
	numbers := []int{}
	firstOrDefault1, _ := FirstOrDefault(SliceAll(numbers))
	fmt.Println(firstOrDefault1)

	months := []int{}
	// Setting the default value to 1 after the query.
	firstOrDefault2, _ := FirstOrDefault(VarAll(months...))
	if firstOrDefault2 == 0 {
		firstOrDefault2 = 1
	}
	fmt.Printf("The value of the firstMonth1 variable is %v\n", firstOrDefault2)

	// Setting the default value to 1 by using DefaultIfEmptyDef() in the query.
	defaultIfEmptyDef, _ := DefaultIfEmptyDef(SliceAll(months), 1)
	first, _ := First(defaultIfEmptyDef)
	fmt.Printf("The value of the firstMonth2 variable is %v\n", first)
	// Output:
	// 0
	// The value of the firstMonth1 variable is 1
	// The value of the firstMonth2 variable is 1
}

// last example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
func ExampleFirstOrDefaultPred() {
	names := []string{"Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	firstLongName, _ := FirstOrDefaultPred(
		SliceAll(names),
		func(name string) bool { return len(name) > 20 },
	)
	fmt.Printf("The first long name is '%v'.\n", firstLongName)

	firstVeryLongName, _ := FirstOrDefaultPred(
		SliceAll(names),
		func(name string) bool { return len(name) > 30 },
	)
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

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
func ExampleFirstPred() {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19}
	firstPred, _ := FirstPred(
		SliceAll(numbers),
		func(number int) bool { return number > 80 },
	)
	fmt.Println(firstPred)
	// Output:
	// 92
}
