package go2linq

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AllTest.cs

func TestAll_int(t *testing.T) {
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
		{name: "EmptySequenceReturnsTrue",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "PredicateMatchingNoElements",
			args: args{
				source:    VarAll(1, 5, 20, 30),
				predicate: func(x int) bool { return x < 0 },
			},
			want: false,
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    VarAll(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 3 },
			},
			want: false,
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    VarAll(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 0 },
			},
			want: true,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstNonMatch",
			args: args{
				source: errorhelper.Must(Select(
					VarAll(4, 6, 0, 3),
					func(x int) int { return 12 / x },
				)),
				predicate: func(y int) bool { return y > 2 },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := All(tt.args.source, tt.args.predicate)
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("All() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll_any(t *testing.T) {
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
				source:    VarAll[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) >= 3 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    VarAll[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) > 3 },
			},
			want: false,
		},
		{name: "3",
			args: args{
				source:    VarAll[any](1, 2, "three", "four"),
				predicate: func(e any) bool { _, ok := e.(int); return ok },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := All(tt.args.source, tt.args.predicate)
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see AllEx example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.all#examples
func ExampleAll_ex1() {
	pets := []Pet{
		{Name: "Barley", Age: 10},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 6},
	}
	// Determine whether all Pet names in the array start with 'B'.
	allStartWithB, _ := All(
		SliceAll(pets),
		func(pet Pet) bool { return strings.HasPrefix(pet.Name, "B") },
	)
	var what string
	if allStartWithB {
		what = "All"
	} else {
		what = "Not all"
	}
	fmt.Printf("%s pet names start with 'B'.\n", what)
	// Output:
	// Not all pet names start with 'B'.
}

// see AllEx2 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.all#examples
func ExampleAll_ex2() {
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
			Pets: []Pet{
				{Name: "Belle", Age: 8},
			},
		},
		{
			LastName: "Philips",
			Pets: []Pet{
				{Name: "Sweetie", Age: 2},
				{Name: "Rover", Age: 13},
			},
		},
	}
	// Determine which people have Pets that are all older than 5.
	where, _ := Where(
		SliceAll(people),
		func(person Person) bool {
			return errorhelper.Must(All(
				SliceAll(person.Pets),
				func(pet Pet) bool { return pet.Age > 5 },
			))
		},
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
	// Antebi
}

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#query-expression-syntax-examples
// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#all
func ExampleAll_ex3() {
	markets := []Market{
		{Name: "Emily's", Items: []string{"kiwi", "cheery", "banana"}},
		{Name: "Kim's", Items: []string{"melon", "mango", "olive"}},
		{Name: "Adam's", Items: []string{"kiwi", "apple", "orange"}},
	}
	where, _ := Where(
		SliceAll(markets),
		func(m Market) bool {
			return errorhelper.Must(All(
				SliceAll(m.Items),
				func(item string) bool { return len(item) == 5 },
			))
		},
	)
	names, _ := Select(where, func(m Market) string { return m.Name })
	for name := range names {
		fmt.Printf("%s market\n", name)
	}
	// Output:
	// Kim's market
}
