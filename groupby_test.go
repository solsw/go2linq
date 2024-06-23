package go2linq

import (
	"fmt"
	"iter"
	"math"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupByTest.cs

func TestGroupBy(t *testing.T) {
	groupBy, _ := GroupBy(
		VarToSeq("abc", "hello", "def", "there", "four"),
		func(el string) int { return len(el) },
	)
	grs, _ := ToSlice(groupBy)
	if len(grs) != 3 {
		t.Errorf("len(GroupBy) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupBy[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupBy[0].Key = %v, want %v", gr0.key, 3)
	}
	got0 := SliceToSeq(gr0.values)
	want0 := VarToSeq("abc", "def")
	equal, _ := SequenceEqual(got0, want0)
	if !equal {
		t.Errorf("GroupBy[0].values = %v, want %v", StringDef(got0), StringDef(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupBy[1].Key = %v, want %v", gr1.key, 5)
	}
	got1 := SliceToSeq(gr1.values)
	want1 := VarToSeq("hello", "there")
	equal, _ = SequenceEqual(got1, want1)
	if !equal {
		t.Errorf("GroupBy[1].values = %v, want %v", StringDef(got1), StringDef(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupBy[2].Key = %v, want %v", gr2.key, 4)
	}
	got2 := SliceToSeq(gr2.values)
	want2 := VarToSeq("four")
	equal, _ = SequenceEqual(got2, want2)
	if !equal {
		t.Errorf("GroupBy[2].values = %v, want %v", StringDef(got2), StringDef(want2))
	}
}

func TestGroupBySel(t *testing.T) {
	groupBySel, _ := GroupBySel(
		VarToSeq("abc", "hello", "def", "there", "four"),
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] },
	)
	grs, _ := ToSlice(groupBySel)
	if len(grs) != 3 {
		t.Errorf("len(GroupBySel) = %v, want %v", len(grs), 3)
	}
	lg0 := len(grs[0].values)
	if lg0 != 2 {
		t.Errorf("len(GroupBySel[0].values) = %v, want %v", lg0, 2)
	}

	gr0 := grs[0]
	if gr0.key != 3 {
		t.Errorf("GroupBySel[0].Key = %v, want %v", gr0.key, 3)
	}
	got0 := SliceToSeq(gr0.values)
	want0 := VarToSeq('a', 'd')
	equal, _ := SequenceEqual(got0, want0)
	if !equal {
		t.Errorf("GroupBySel[0].values = %v, want %v", StringDef(got0), StringDef(want0))
	}

	gr1 := grs[1]
	if gr1.key != 5 {
		t.Errorf("GroupBySel[1].Key = %v, want %v", gr1, 3)
	}
	got1 := SliceToSeq(gr1.values)
	want1 := VarToSeq('h', 't')
	equal, _ = SequenceEqual(got1, want1)
	if !equal {
		t.Errorf("GroupBySel[1].values = %v, want %v", StringDef(got1), StringDef(want1))
	}

	gr2 := grs[2]
	if gr2.key != 4 {
		t.Errorf("GroupBySel[2].Key = %v, want %v", gr2, 3)
	}
	got2 := SliceToSeq(gr2.values)
	want2 := VarToSeq('f')
	equal, _ = SequenceEqual(got2, want2)
	if !equal {
		t.Errorf("GroupBySel[2].values = %v, want %v", StringDef(got2), StringDef(want2))
	}
}

func TestGroupByRes(t *testing.T) {
	groupByRes, _ := GroupByRes(
		VarToSeq("abc", "hello", "def", "there", "four"),
		func(el string) int { return len(el) },
		func(el int, seq iter.Seq[string]) string {
			ss, _ := Strings(seq)
			return fmt.Sprintf("%v:%v", el, strings.Join(ss, ";"))
		})
	grs, _ := ToSlice(groupByRes)
	got := SliceToSeq(grs)
	want := VarToSeq("3:abc;def", "5:hello;there", "4:four")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("GroupByRes = %v, want %v", StringDef(got), StringDef(want))
	}
}

func TestGroupBySelRes(t *testing.T) {
	groupBySelRes, _ := GroupBySelRes(
		VarToSeq("abc", "hello", "def", "there", "four"),
		func(el string) int { return len(el) },
		func(el string) rune { return []rune(el)[0] },
		func(el int, seq iter.Seq[rune]) string {
			vv := func() []string {
				var ss []string
				for s := range seq {
					ss = append(ss, string(s))
				}
				return ss
			}()
			return fmt.Sprintf("%v:%v", el, strings.Join(vv, ";"))
		})
	grs, _ := ToSlice(groupBySelRes)
	got := SliceToSeq(grs)
	want := VarToSeq("3:a;d", "5:h;t", "4:f")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("GroupBySelRes = %v, want %v", StringDef(got), StringDef(want))
	}
}

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/grouping-data#query-expression-syntax-example
func ExampleGroupBy() {
	groupBy, _ := GroupBy(
		VarToSeq(35, 44, 200, 84, 3987, 4, 199, 329, 446, 208),
		func(i int) int { return i % 2 },
	)
	for group := range groupBy {
		if group.Key() == 0 {
			fmt.Println("\nEven numbers:")
		} else {
			fmt.Println("\nOdd numbers:")
		}
		for i := range group.Values() {
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

// GroupByEx3 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func ExampleGroupByRes() {
	pets := []PetF{
		{Name: "Barley", Age: 8.3},
		{Name: "Boots", Age: 4.9},
		{Name: "Whiskers", Age: 1.5},
		{Name: "Daisy", Age: 4.3},
	}
	// Group PetF objects by the math.Floor of their Age.
	// Then project a Result type from each group that consists of the Key,
	// Count of the group's elements, and the minimum and maximum Age in the group.
	groupByRes, _ := GroupByRes(
		SliceToSeq(pets),
		func(pet PetF) float64 { return math.Floor(pet.Age) },
		func(age float64, pets iter.Seq[PetF]) Result {
			count, _ := Count(pets)
			min, _ := MinSel(pets, func(pet PetF) float64 { return pet.Age })
			max, _ := MaxSel(pets, func(pet PetF) float64 { return pet.Age })
			return Result{Key: age, Count: count, Min: min, Max: max}
		},
	)
	for result := range groupByRes {
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

// GroupByEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func ExampleGroupBySel() {
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
		{Name: "Daisy", Age: 4},
	}
	// Group the pets using Age as the key value and selecting only the Pet's Name for each value.
	groupBySel, _ := GroupBySel(
		SliceToSeq(pets),
		func(pet Pet) int { return pet.Age },
		func(pet Pet) string { return pet.Name },
	)
	// Iterate over each Grouping in the collection.
	for petGroup := range groupBySel {
		// Print the key value of the Grouping.
		fmt.Println(petGroup.Key())
		// Iterate over each value in the Grouping and print the value.
		for name := range petGroup.Values() {
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

// GroupByEx4 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func ExampleGroupBySelRes() {
	pets := []PetF{
		{Name: "Barley", Age: 8.3},
		{Name: "Boots", Age: 4.9},
		{Name: "Whiskers", Age: 1.5},
		{Name: "Daisy", Age: 4.3},
	}
	// Group PetF.Age values by the math.Floor of the age.
	// Then project a Result type from each group that consists of the Key,
	// Count of the group's elements, and the minimum and maximum Age in the group.
	groupBySelRes, _ := GroupBySelRes(
		SliceToSeq(pets),
		func(pet PetF) float64 { return math.Floor(pet.Age) },
		func(pet PetF) float64 { return pet.Age },
		func(baseAge float64, ages iter.Seq[float64]) Result {
			count, _ := Count(ages)
			min, _ := Min(ages)
			max, _ := Max(ages)
			return Result{Key: baseAge, Count: count, Min: min, Max: max}
		},
	)
	for result := range groupBySelRes {
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
