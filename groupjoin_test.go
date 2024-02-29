package go2linq

import (
	"fmt"
	"iter"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupJoinTest.cs

func TestGroupJoin_SimpleGroupJoin(t *testing.T) {
	got, _ := GroupJoin(
		VarAll("first", "second", "third"),
		VarAll("essence", "offer", "eating", "psalm"),
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels iter.Seq[string]) string {
			ss, _ := Strings(iels)
			return fmt.Sprintf("%v:%v", oel, strings.Join(ss, ";"))
		},
	)
	want := VarAll("first:offer", "second:essence;psalm", "third:")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("GroupJoin_SimpleGroupJoin = %v, want %v", StringDef(got), StringDef(want))
	}
}

func TestGroupJoin_SameEnumerable(t *testing.T) {
	outer := []string{"fs", "sf", "ff", "ss"}
	inner := outer
	groupJoin, _ := GroupJoin(
		SliceAll(outer),
		SliceAll(inner),
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels iter.Seq[string]) string {
			ss, _ := Strings(iels)
			return fmt.Sprintf("%v:%v", oel, strings.Join(ss, ";"))
		},
	)
	got, _ := ToSlice(groupJoin)
	want := []string{"fs:sf;ff", "sf:fs;ss", "ff:sf;ff", "ss:fs;ss"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoin_SameEnumerable = %v, want %v", got, want)
	}
}

func TestGroupJoinEq_CustomComparer(t *testing.T) {
	outer := []string{"ABCxxx", "abcyyy", "defzzz", "ghizzz"}
	inner := []string{"000abc", "111gHi", "222333", "333AbC"}
	got, _ := GroupJoinEq(
		SliceAll(outer),
		SliceAll(inner),
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel string, iels iter.Seq[string]) string {
			ss, _ := Strings(iels)
			return fmt.Sprintf("%v:%v", oel, strings.Join(ss, ";"))
		},
		caseInsensitiveEqual,
	)
	want := VarAll("ABCxxx:000abc;333AbC", "abcyyy:000abc;333AbC", "defzzz:", "ghizzz:111gHi")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("GroupJoinEq_CustomComparer = %v, want %v", StringDef(got), StringDef(want))
	}
}

func TestGroupJoin_DifferentSourceTypes(t *testing.T) {
	outer := []int{5, 3, 7, 4}
	inner := []string{"bee", "giraffe", "tiger", "badger", "ox", "cat", "dog"}
	got, _ := GroupJoin(
		SliceAll(outer),
		SliceAll(inner),
		Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iels iter.Seq[string]) string {
			ss, _ := Strings(iels)
			return fmt.Sprintf("%v:%v", oel, strings.Join(ss, ";"))
		},
	)
	want := VarAll("5:tiger", "3:bee;cat;dog", "7:giraffe", "4:")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("GroupJoin_DifferentSourceTypes = %v, want %v", StringDef(got), StringDef(want))
	}
}

// GroupJoinEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin
func ExampleGroupJoin_ex1() {
	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	// Create a list where each element is an OwnerAndPets type that contains a person's name and
	// a collection of names of the pets they own.
	people := []Person{magnus, terry, charlotte}
	pets := []Pet{barley, boots, whiskers, daisy}

	groupJoin, _ := GroupJoin(
		SliceAll(people),
		SliceAll(pets),
		Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pets iter.Seq[Pet]) OwnerAndPets {
			return OwnerAndPets{
				OwnerName: person.Name,
				Pets:      errorhelper.Must(Select(pets, func(pet Pet) string { return pet.Name }))}
		},
	)
	for obj := range groupJoin {
		// Output the owner's name.
		fmt.Printf("%s:\n", obj.OwnerName)
		// Output each of the owner's pet's names.
		for pet := range obj.Pets {
			fmt.Printf("  %s\n", pet)
		}
	}
	// Output:
	// Hedlund, Magnus:
	//   Daisy
	// Adams, Terry:
	//   Barley
	//   Boots
	// Weiss, Charlotte:
	//   Whiskers
}

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#query-expression-syntax-examples
// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#groupjoin
func ExampleGroupJoin_ex2() {
	products := []Product{
		{Name: "Cola", CategoryId: 0},
		{Name: "Tea", CategoryId: 0},
		{Name: "Apple", CategoryId: 1},
		{Name: "Kiwi", CategoryId: 1},
		{Name: "Carrot", CategoryId: 2},
	}
	categories := []Category{
		{Id: 0, CategoryName: "Beverage"},
		{Id: 1, CategoryName: "Fruit"},
		{Id: 2, CategoryName: "Vegetable"},
	}
	// Join categories and product based on CategoryId and grouping result
	productGroups, _ := GroupJoin(
		SliceAll(categories),
		SliceAll(products),
		func(category Category) int { return category.Id },
		func(product Product) int { return product.CategoryId },
		func(category Category, products iter.Seq[Product]) iter.Seq[Product] {
			return products
		},
	)
	for productGroup := range productGroups {
		fmt.Println("Group")
		for product := range productGroup {
			fmt.Printf("%8s\n", product.Name)
		}
	}
	// Output:
	// Group
	//     Cola
	//      Tea
	// Group
	//    Apple
	//     Kiwi
	// Group
	//   Carrot
}
