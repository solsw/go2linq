//go:build go1.18

package go2linq

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupByTest.cs

func Test_GroupByMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupByMust(en, func(el string) int { return len(el) }))
	if len(grs) != 3 {
		t.Errorf("len(GroupByMust) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupByMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupByMust[0].Key = %v, want %v", gr0.key, 3)
	}
	got0 := NewEnSlice(gr0.values...)
	want0 := NewEnSlice("abc", "def")
	if !SequenceEqualMust(got0, want0) {
		t.Errorf("GroupByMust[0].values = %v, want %v", ToStringDef(got0), ToStringDef(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupByMust[1].Key = %v, want %v", gr1.key, 5)
	}
	got1 := NewEnSlice(gr1.values...)
	want1 := NewEnSlice("hello", "there")
	if !SequenceEqualMust(got1, want1) {
		t.Errorf("GroupByMust[1].values = %v, want %v", ToStringDef(got1), ToStringDef(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupByMust[2].Key = %v, want %v", gr2.key, 4)
	}
	got2 := NewEnSlice(gr2.values...)
	want2 := NewEnSlice("four")
	if !SequenceEqualMust(got2, want2) {
		t.Errorf("GroupByMust[2].values = %v, want %v", ToStringDef(got2), ToStringDef(want2))
	}
}

func Test_GroupBySelMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupBySelMust(en,
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] }),
	)
	if len(grs) != 3 {
		t.Errorf("len(GroupBySelMust) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupBySelMust[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupBySelMust[0].Key = %v, want %v", gr0.key, 3)
	}
	got0 := NewEnSlice(gr0.values...)
	want0 := NewEnSlice('a', 'd')
	if !SequenceEqualMust(got0, want0) {
		t.Errorf("GroupBySelMust[0].values = %v, want %v", ToStringDef(got0), ToStringDef(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupBySelMust[1].Key = %v, want %v", gr1, 3)
	}
	got1 := NewEnSlice(gr1.values...)
	want1 := NewEnSlice('h', 't')
	if !SequenceEqualMust(got1, want1) {
		t.Errorf("GroupBySelMust[1].values = %v, want %v", ToStringDef(got1), ToStringDef(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupBySelMust[2].Key = %v, want %v", gr2, 3)
	}
	got2 := NewEnSlice(gr2.values...)
	want2 := NewEnSlice('f')
	if !SequenceEqualMust(got2, want2) {
		t.Errorf("GroupBySelMust[2].values = %v, want %v", ToStringDef(got2), ToStringDef(want2))
	}
}

func Test_GroupByResMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupByResMust(en,
		func(el string) int { return len(el) },
		func(el int, en Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", el, strings.Join(ToStrings(en), ";"))
		}))
	got := NewEnSlice(grs...)
	want := NewEnSlice("3:abc;def", "5:hello;there", "4:four")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupByResMust = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func Test_GroupBySelResMust(t *testing.T) {
	en := NewEnSlice("abc", "hello", "def", "there", "four")
	grs := ToSliceMust(GroupBySelResMust(en,
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] },
		func(el int, en Enumerable[rune]) string {
			vv := func() []string {
				var r []string
				enr := en.GetEnumerator()
				for enr.MoveNext() {
					r = append(r, string(enr.Current()))
				}
				return r
			}()
			return fmt.Sprintf("%v:%v", el, strings.Join(vv, ";"))
		}))
	got := NewEnSlice(grs...)
	want := NewEnSlice("3:a;d", "5:h;t", "4:f")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupBySelResMust = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/grouping-data#query-expression-syntax-example
func ExampleGroupByMust() {
	numbers := NewEnSlice(35, 44, 200, 84, 3987, 4, 199, 329, 446, 208)
	query := GroupByMust(numbers, func(i int) int { return i % 2 })
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		group := enr.Current()
		if group.Key() == 0 {
			fmt.Println("\nEven numbers:")
		} else {
			fmt.Println("\nOdd numbers:")
		}
		enrGroup := group.GetEnumerator()
		for enrGroup.MoveNext() {
			i := enrGroup.Current()
			fmt.Println(i)
		}
	}
	// Output:
	// Odd numbers:
	// 35
	// 3987
	// 199
	// 329
	//
	// Even numbers:
	// 44
	// 200
	// 84
	// 4
	// 446
	// 208
}

// see GroupByEx3 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func ExampleGroupByResMust() {
	petsList := NewEnSlice(
		PetF{Name: "Barley", Age: 8.3},
		PetF{Name: "Boots", Age: 4.9},
		PetF{Name: "Whiskers", Age: 1.5},
		PetF{Name: "Daisy", Age: 4.3},
	)
	// Group PetF objects by the math.Floor of their Age.
	// Then project a Result type from each group that consists of the Key,
	// the Count of the group's elements, and the minimum and maximum Age in the group.
	query := GroupByResMust(petsList,
		func(pet PetF) float64 { return math.Floor(pet.Age) },
		func(age float64, pets Enumerable[PetF]) Result {
			count := CountMust(pets)
			mn := MinSelMust(pets, func(pet PetF) float64 { return pet.Age })
			mx := MaxSelMust(pets, func(pet PetF) float64 { return pet.Age })
			return Result{Key: age, Count: count, Min: mn, Max: mx}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		result := enr.Current()
		fmt.Printf("\nAge group: %g\n", result.Key)
		fmt.Printf("Number of pets in this age group: %d\n", result.Count)
		fmt.Printf("Minimum age: %g\n", result.Min)
		fmt.Printf("Maximum age: %g\n", result.Max)
	}
	// Output:
	// Age group: 8
	// Number of pets in this age group: 1
	// Minimum age: 8.3
	// Maximum age: 8.3
	//
	// Age group: 4
	// Number of pets in this age group: 2
	// Minimum age: 4.3
	// Maximum age: 4.9
	//
	// Age group: 1
	// Number of pets in this age group: 1
	// Minimum age: 1.5
	// Maximum age: 1.5
}

// see GroupByEx1 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func ExampleGroupBySelMust() {
	pets := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
		Pet{Name: "Daisy", Age: 4},
	)
	// Group the pets using Age as the key value and selecting only the Pet's Name for each value.
	query := GroupBySelMust(pets,
		func(pet Pet) int { return pet.Age },
		func(pet Pet) string { return pet.Name },
	)
	// Iterate over each Grouping in the collection.
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		petGroup := enr.Current()
		// Print the key value of the Grouping.
		fmt.Println(petGroup.Key())
		names := petGroup.GetEnumerator()
		// Iterate over each value in the Grouping and print the value.
		for names.MoveNext() {
			name := names.Current()
			fmt.Printf("  %s\n", name)
		}
	}
	// Output:
	// 8
	//   Barley
	// 4
	//   Boots
	//   Daisy
	// 1
	//   Whiskers
}

// see GroupByEx4 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func ExampleGroupBySelResMust() {
	pets := NewEnSlice(
		PetF{Name: "Barley", Age: 8.3},
		PetF{Name: "Boots", Age: 4.9},
		PetF{Name: "Whiskers", Age: 1.5},
		PetF{Name: "Daisy", Age: 4.3},
	)
	// Group PetF.Age values by the math.Floor of the age.
	// Then project a Result type from each group that consists of the Key,
	// the Count of the group's elements, and the minimum and maximum Age in the group.
	query := GroupBySelResMust(pets,
		func(pet PetF) float64 { return math.Floor(pet.Age) },
		func(pet PetF) float64 { return pet.Age },
		func(baseAge float64, ages Enumerable[float64]) Result {
			count := CountMust(ages)
			mn := MinMust(ages)
			mx := MaxMust(ages)
			return Result{Key: baseAge, Count: count, Min: mn, Max: mx}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		result := enr.Current()
		fmt.Printf("\nAge group: %g\n", result.Key)
		fmt.Printf("Number of pets in this age group: %d\n", result.Count)
		fmt.Printf("Minimum age: %g\n", result.Min)
		fmt.Printf("Maximum age: %g\n", result.Max)
	}
	// Output:
	// Age group: 8
	// Number of pets in this age group: 1
	// Minimum age: 8.3
	// Maximum age: 8.3
	//
	// Age group: 4
	// Number of pets in this age group: 2
	// Minimum age: 4.3
	// Maximum age: 4.9
	//
	// Age group: 1
	// Number of pets in this age group: 1
	// Minimum age: 1.5
	// Maximum age: 1.5
}
