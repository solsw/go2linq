//go:build go1.18

package go2linq

import (
	"fmt"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/AllTest.cs

func Test_All_int(t *testing.T) {
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
		{name: "EmptySequenceReturnsTrue",
			args: args{
				source:    Empty[int](),
				predicate: func(x int) bool { return x > 10 },
			},
			want: true,
		},
		{name: "PredicateMatchingNoElements",
			args: args{
				source:    NewEnSlice(1, 5, 20, 30),
				predicate: func(x int) bool { return x < 0 },
			},
			want: false,
		},
		{name: "PredicateMatchingSomeElements",
			args: args{
				source:    NewEnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 3 },
			},
			want: false,
		},
		{name: "PredicateMatchingAllElements",
			args: args{
				source:    NewEnSlice(1, 5, 8, 9),
				predicate: func(x int) bool { return x > 0 },
			},
			want: true,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstNonMatch",
			args: args{
				source:    SelectMust(NewEnSlice(10, 2, 0, 3), func(x int) int { return 10 / x }),
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

func Test_AllMust_any(t *testing.T) {
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
				source:    NewEnSlice[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) >= 3 },
			},
			want: true,
		},
		{name: "2",
			args: args{
				source:    NewEnSlice[any]("one", "two", "three", "four"),
				predicate: func(e any) bool { return len(e.(string)) > 3 },
			},
			want: false,
		},
		{name: "3",
			args: args{
				source:    NewEnSlice[any](1, 2, "three", "four"),
				predicate: func(e any) bool { _, ok := e.(int); return ok },
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllMust(tt.args.source, tt.args.predicate)
			if got != tt.want {
				t.Errorf("AllMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

// see AllEx example from Enumerable.All help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all#examples
func ExampleAllMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 10},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 6},
	)
	// Determine whether all Pet names in the array start with 'B'.
	allStartWithB := AllMust(pets,
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

// see AllEx2 example from Enumerable.All help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all#examples
func ExampleAllMust_2() {
	people := NewEnSlice(
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
			Pets: []Pet{
				Pet{Name: "Belle", Age: 8},
			},
		},
		Person{
			LastName: "Philips",
			Pets: []Pet{
				Pet{Name: "Sweetie", Age: 2},
				Pet{Name: "Rover", Age: 13},
			},
		},
	)
	// Determine which people have Pets that are all older than 5.
	where := WhereMust(people,
		func(person Person) bool {
			return AllMust(NewEnSlice(person.Pets...),
				func(pet Pet) bool { return pet.Age > 5 },
			)
		},
	)
	names := SelectMust(where,
		func(person Person) string { return person.LastName },
	)
	enr := names.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Println(name)
	}
	// Output:
	// Haas
	// Antebi
}

// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#query-expression-syntax-examples
// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations#all
func ExampleAllMust_3() {
	markets := NewEnSlice(
		Market{Name: "Emily's", Items: []string{"kiwi", "cheery", "banana"}},
		Market{Name: "Kim's", Items: []string{"melon", "mango", "olive"}},
		Market{Name: "Adam's", Items: []string{"kiwi", "apple", "orange"}},
	)
	where := WhereMust(markets,
		func(m Market) bool {
			items := NewEnSlice(m.Items...)
			return AllMust(items, func(item string) bool { return len(item) == 5 })
		},
	)
	names := SelectMust(where, func(m Market) string { return m.Name })
	enr := names.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Printf("%s market\n", name)
	}
	// Output:
	// Kim's market
}
