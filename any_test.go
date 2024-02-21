package go2linq

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AnyTest.cs

func TestAny_int(t *testing.T) {
	type args struct {
		source iter.Seq[int]
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
				source: VarAll(0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Any(tt.args.source)
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyPred_int(t *testing.T) {
	type args struct {
		source    iter.Seq[int]
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
				source:    VarAll(1, 3, 5),
				predicate: nil,
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
				source:    VarAll(1, 5, 20, 30),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "NonEmptySequenceWithPredicateNotMatchingElement",
			args: args{
				source:    VarAll(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 10 },
			},
			want: false,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstMatch",
			args: args{
				source: errorhelper.Must(Select(
					VarAll(10, 2, 0, 3),
					func(x int) int { return 10 / x },
				)),
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

func TestAnyPred_any(t *testing.T) {
	type args struct {
		source    iter.Seq[any]
		predicate func(any) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1",
			args: args{
				source:    VarAll[any](1, 2, 3, 4),
				predicate: func(e any) bool { return e.(int) == 4 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    VarAll[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) == 4 },
			},
			want: true,
		},
		{name: "3",
			args: args{
				source:    VarAll[any](1, 2, "three", "four"),
				predicate: func(e any) bool { _, ok := e.(int); return ok },
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := AnyPred(tt.args.source, tt.args.predicate)
			if got != tt.want {
				t.Errorf("AnyPred() = %v, want %v", got, tt.want)
			}
		})
	}
}

// first example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func ExampleAny_ex1() {
	numbers := []int{1, 2}
	hasElements, _ := Any(VarAll(numbers...))
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

// see AnyEx2 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func ExampleAny_ex2() {
	people := []Person{
		{
			LastName: "Haas",
			Pets: []Pet{
				{Name: "Barley", Age: 10},
				{Name: "Boots", Age: 14},
				{Name: "Whiskers", Age: 6},
			},
		},
		{
			LastName: "Fakhouri",
			Pets: []Pet{
				{Name: "Snowball", Age: 1},
			},
		},
		{
			LastName: "Antebi",
			Pets:     []Pet{},
		},
		{
			LastName: "Philips",
			Pets: []Pet{
				{Name: "Sweetie", Age: 2},
				{Name: "Rover", Age: 13},
			},
		},
	}
	// Determine which people have a non-empty Pet array.
	where, _ := Where(
		SliceAll(people),
		func(person Person) bool { return errorhelper.Must(Any(SliceAll(person.Pets))) },
	)
	names, _ := Select(
		where,
		func(person Person) string { return person.LastName },
	)
	for name := range names {
		fmt.Println(name)
	}
	// Output:
	// Haas
	// Fakhouri
	// Philips
}

// see AnyEx3 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func ExampleAnyPred_ex1() {
	pets := []Pet{
		{Name: "Barley", Age: 8, Vaccinated: true},
		{Name: "Boots", Age: 4, Vaccinated: false},
		{Name: "Whiskers", Age: 1, Vaccinated: false},
	}
	// Determine whether any pets over Age 1 are also unvaccinated.
	unvaccinated, _ := AnyPred(
		SliceAll(pets),
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

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#query-expression-syntax-examples
// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#any
func ExampleAnyPred_ex2() {
	markets := []Market{
		{Name: "Emily's", Items: []string{"kiwi", "cheery", "banana"}},
		{Name: "Kim's", Items: []string{"melon", "mango", "olive"}},
		{Name: "Adam's", Items: []string{"kiwi", "apple", "orange"}},
	}
	where, _ := Where(
		SliceAll(markets),
		func(m Market) bool {
			return errorhelper.Must(AnyPred(
				SliceAll(m.Items),
				func(item string) bool { return strings.HasPrefix(item, "o") },
			))
		},
	)
	names, _ := Select(where, func(m Market) string { return m.Name })
	for name := range names {
		fmt.Printf("%s market\n", name)
	}
	// Output:
	// Kim's market
	// Adam's market
}
