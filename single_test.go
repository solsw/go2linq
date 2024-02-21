package go2linq

import (
	"fmt"
	"iter"
	"reflect"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SingleTest.cs
// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SingleOrDefaultTest.cs

func TestSingle_int(t *testing.T) {
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
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
		{name: "EarlyOutWithoutPredicate",
			args: args{
				source: errorhelper.Must(Select(VarAll(1, 2, 0), func(x int) int { return 10 / x })),
			},
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Single(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Single() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("Single() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Single() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglePred_int(t *testing.T) {
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
				source:    VarAll(1, 3, 5, 4, 2),
				predicate: func(x int) bool { return x > 4 },
			},
			want: 5,
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			args: args{
				source:    VarAll(1, 2, 5, 10, 2, 1),
				predicate: func(x int) bool { return x > 3 },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
		{name: "EarlyOutWithPredicate",
			args: args{
				source:    errorhelper.Must(Select(VarAll(1, 2, 0), func(x int) int { return 10 / x })),
				predicate: func(int) bool { return true },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SinglePred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SinglePred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SinglePred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SinglePred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleOrDefault_int(t *testing.T) {
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
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
		{name: "EarlyOutWithoutPredicate",
			args: args{
				source: errorhelper.Must(Select(VarAll(1, 2, 0), func(x int) int { return 10 / x })),
			},
			wantErr:     true,
			expectedErr: ErrMultipleElements,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingleOrDefault(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("SingleOrDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SingleOrDefault() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleOrDefaultPred(t *testing.T) {
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
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
		{name: "EarlyOutWithPredicate",
			args: args{
				source:    errorhelper.Must(Select(VarAll(1, 2, 0), func(x int) int { return 10 / x })),
				predicate: func(int) bool { return true },
			},
			wantErr:     true,
			expectedErr: ErrMultipleMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SingleOrDefaultPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("SingleOrDefaultPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("SingleOrDefaultPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleOrDefaultPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func ExampleSingle_ex1() {
	fruits := []string{"orange"}
	fruit, _ := Single(SliceAll(fruits))
	fmt.Println(fruit)
	// Output:
	// orange
}

// third example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
func ExampleSingle_ex2() {
	pageNumbers := []int{}
	// Setting the default value to 1 by using DefaultIfEmpty() in the query.
	pageNumber, _ := Single(errorhelper.Must(DefaultIfEmptyDef(SliceAll(pageNumbers), 1)))
	fmt.Printf("The value of the pageNumber2 variable is %d\n", pageNumber)
	// Output:
	// The value of the pageNumber2 variable is 1
}

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func ExampleSingle() {
	fruits := []string{"orange", "apple"}
	fruit, err := Single(SliceAll(fruits))
	if err == ErrMultipleElements {
		fmt.Println("The collection does not contain exactly one element.")
	} else {
		fmt.Println(fruit)
	}
	// Output:
	// The collection does not contain exactly one element.
}

// third and fourth examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
func ExampleSinglePred() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}

	fruit1, _ := SinglePred(SliceAll(fruits), func(fr string) bool { return len(fr) > 10 })
	fmt.Println(fruit1)

	fruit2, err := SinglePred(SliceAll(fruits), func(fr string) bool { return len(fr) > 15 })
	if err == ErrNoMatch {
		fmt.Println("The collection does not contain exactly one element whose length is greater than 15.")
	} else {
		fmt.Println(fruit2)
	}

	fruit3, err := SinglePred(
		SliceAll(fruits),
		func(fr string) bool { return len(fr) > 5 },
	)
	if err == ErrMultipleMatch {
		fmt.Println("The collection does not contain exactly one element whose length is greater than 5.")
	} else {
		fmt.Println(fruit3)
	}
	// Output:
	// passionfruit
	// The collection does not contain exactly one element whose length is greater than 15.
	// The collection does not contain exactly one element whose length is greater than 5.
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
func ExampleSingleOrDefault_ex1() {
	fruits := []string{"orange"}
	fruit, _ := SingleOrDefault(SliceAll(fruits))
	fmt.Println(fruit)
	// Output:
	// orange
}

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
func ExampleSingleOrDefault_ex2() {
	fruits := []string{}
	fruit, _ := SingleOrDefault(SliceAll(fruits))
	var what string
	if fruit == "" {
		what = "No such string!"
	} else {
		what = fruit
	}
	fmt.Println(what)
	// Output:
	// No such string!
}

// third example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
func ExampleSingleOrDefault_ex3() {
	var pageNumbers []int = nil
	// Setting the default value to 1 after the query.
	pageNumber, _ := SingleOrDefault(SliceAll(pageNumbers))
	if pageNumber == 0 {
		pageNumber = 1
	}
	fmt.Printf("The value of the pageNumber1 variable is %d\n", pageNumber)
	// Output:
	// The value of the pageNumber1 variable is 1
}

// fourth and fifth examples from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
func ExampleSingleOrDefaultPred() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	fruit1, _ := SingleOrDefaultPred(SliceAll(fruits), func(fr string) bool { return len(fr) > 10 })
	fmt.Println(fruit1)

	fruit2, _ := SingleOrDefaultPred(SliceAll(fruits), func(fr string) bool { return len(fr) > 15 })
	var what string
	if fruit2 == "" {
		what = "No such string!"
	} else {
		what = fruit2
	}
	fmt.Println(what)
	// Output:
	// passionfruit
	// No such string!
}
