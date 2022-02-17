//go:build go1.18

package go2linq

import (
	"fmt"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AnyTest.cs

func Test_AnyMust_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "EmptySequenceWithoutPredicate",
			args: args{
				source: Empty[int](),
			},
			want: false,
		},
		{name: "NonEmptySequenceWithoutPredicate",
			args: args{
				source: NewEnSlice(0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyMust(tt.args.source)
			if got != tt.want {
				t.Errorf("AnyMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyPred_int(t *testing.T) {
	type args struct {
		source    Enumerable[int]
		predicate func(int) bool
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantErr     bool
		expectedErr error
	}{
		{name: "NullPredicate",
			args: args{
				source: NewEnSlice(1, 3, 5),
			},
			wantErr:     true,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequenceWithPredicate",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "NonEmptySequenceWithPredicateMatchingElement",
			args: args{
				source:    NewEnSlice(1, 5, 20, 30),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "NonEmptySequenceWithPredicateNotMatchingElement",
			args: args{
				source:    NewEnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstMatch",
			args: args{
				source:    SelectMust(NewEnSlice(10, 2, 0, 3), func(x int) int { return 10 / x }),
				predicate: func(y int) bool { return y > 2 },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AnyPred(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnyPred() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("AnyPred() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("AnyPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AnyPredMust_interface(t *testing.T) {
	type args struct {
		source    Enumerable[any]
		predicate func(any) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				source:    NewEnSlice[any](1, 2, 3, 4),
				predicate: func(e any) bool { return e.(int) == 4 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    NewEnSlice[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) == 4 },
			},
			want: true,
		},
		{name: "3",
			args: args{
				source:    NewEnSlice[any](1, 2, "three", "four"),
				predicate: func(e any) bool { _, ok := e.(int); return ok },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyPredMust(tt.args.source, tt.args.predicate)
			if got != tt.want {
				t.Errorf("AnyPredMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see the first example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any
func ExampleAnyMust() {
	numbers := NewEnSlice(1, 2)
	hasElements := AnyMust(numbers)
	var what string
	if hasElements {
		what = "is not"
	} else {
		what = "is"
	}
	fmt.Printf("The list %s empty.\n", what)
	// Output:
	// The list is not empty.
}

// see AnyEx2 example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any
func ExampleAnyMust_2() {
	people := []Person{
		Person{
			LastName: "Haas",
			Pets: []Pet{
				Pet{Name: "Barley", Age: 10},
				Pet{Name: "Boots", Age: 14},
				Pet{Name: "Whiskers", Age: 6},
			},
		},
		Person{
			LastName: "Fakhouri",
			Pets: []Pet{
				Pet{Name: "Snowball", Age: 1},
			},
		},
		Person{
			LastName: "Antebi",
			Pets:     []Pet{},
		},
		Person{
			LastName: "Philips",
			Pets: []Pet{
				Pet{Name: "Sweetie", Age: 2},
				Pet{Name: "Rover", Age: 13},
			},
		},
	}
	// Determine which people have a non-empty Pet array.
	names := SelectMust(
		WhereMust(
			NewEnSlice(people...),
			func(person Person) bool { return AnyMust(NewEnSlice(person.Pets...)) },
		),
		func(person Person) string { return person.LastName },
	)
	enrNames := names.GetEnumerator()
	for enrNames.MoveNext() {
		name := enrNames.Current()
		fmt.Println(name)
	}
	// Output:
	// Haas
	// Fakhouri
	// Philips
}

// see AnyEx3 example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any
func ExampleAnyPredMust() {
	pets := []Pet{
		Pet{Name: "Barley", Age: 8, Vaccinated: true},
		Pet{Name: "Boots", Age: 4, Vaccinated: false},
		Pet{Name: "Whiskers", Age: 1, Vaccinated: false},
	}
	// Determine whether any pets over Age 1 are also unvaccinated.
	unvaccinated := AnyPredMust(
		NewEnSlice(pets...),
		func(pet Pet) bool { return pet.Age > 1 && pet.Vaccinated == false },
	)
	var what string
	if unvaccinated {
		what = "are"
	} else {
		what = "are not any"
	}
	fmt.Printf("There %s unvaccinated animals over age one.\n", what)
	// Output:
	// There are unvaccinated animals over age one.
}
