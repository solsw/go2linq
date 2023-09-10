package go2linq

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/GroupJoinTest.cs

func TestGroupJoinMust_SimpleGroupJoin(t *testing.T) {
	outer := []string{"first", "second", "third"}
	inner := []string{"essence", "offer", "eating", "psalm"}
	got := GroupJoinMust(
		NewEnSliceEn(outer...),
		NewEnSliceEn(inner...),
		func(oel string) rune { return []rune(oel)[0] },
		func(iel string) rune { return []rune(iel)[1] },
		func(oel string, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		},
	)
	want := NewEnSliceEn("first:offer", "second:essence;psalm", "third:")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupJoinMust_SimpleGroupJoin = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func TestGroupJoinMust_SameEnumerable(t *testing.T) {
	outer := []string{"fs", "sf", "ff", "ss"}
	inner := outer
	got := ToSliceMust(
		GroupJoinMust(
			NewEnSliceEn(outer...),
			NewEnSliceEn(inner...),
			func(oel string) rune { return []rune(oel)[0] },
			func(iel string) rune { return []rune(iel)[1] },
			func(oel string, iels Enumerable[string]) string {
				return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
			},
		),
	)
	want := []string{"fs:sf;ff", "sf:fs;ss", "ff:sf;ff", "ss:fs;ss"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupJoinMust_SameEnumerable = %v, want %v", got, want)
	}
}

func TestGroupJoinEqMust_CustomComparer(t *testing.T) {
	outer := []string{"ABCxxx", "abcyyy", "defzzz", "ghizzz"}
	inner := []string{"000abc", "111gHi", "222333", "333AbC"}
	got := GroupJoinEqMust(
		NewEnSliceEn(outer...),
		NewEnSliceEn(inner...),
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel string, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		},
		collate.CaseInsensitiveOrder,
	)
	want := NewEnSliceEn("ABCxxx:000abc;333AbC", "abcyyy:000abc;333AbC", "defzzz:", "ghizzz:111gHi")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupJoinEqMust_CustomComparer = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func TestGroupJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := []int{5, 3, 7, 4}
	inner := []string{"bee", "giraffe", "tiger", "badger", "ox", "cat", "dog"}
	got := GroupJoinMust(
		NewEnSliceEn(outer...),
		NewEnSliceEn(inner...),
		Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iels Enumerable[string]) string {
			return fmt.Sprintf("%v:%v", oel, strings.Join(ToStrings(iels), ";"))
		},
	)
	want := NewEnSliceEn("5:tiger", "3:bee;cat;dog", "7:giraffe", "4:")
	if !SequenceEqualMust(got, want) {
		t.Errorf("GroupJoinMust_DifferentSourceTypes = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

// see GroupJoinEx1 example from Enumerable.GroupJoin help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin
func ExampleGroupJoinMust_ex1() {
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

	query := GroupJoinMust(
		NewEnSliceEn(people...),
		NewEnSliceEn(pets...),
		Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pets Enumerable[Pet]) OwnerAndPets {
			return OwnerAndPets{
				OwnerName: person.Name,
				Pets:      SelectMust(pets, func(pet Pet) string { return pet.Name })}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		obj := enr.Current()
		// Output the owner's name.
		fmt.Printf("%s:\n", obj.OwnerName)
		// Output each of the owner's pet's names.
		enrPets := obj.Pets.GetEnumerator()
		for enrPets.MoveNext() {
			pet := enrPets.Current()
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
func ExampleGroupJoinMust_ex2() {
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
	productGroups := GroupJoinMust(
		NewEnSliceEn(categories...),
		NewEnSliceEn(products...),
		func(category Category) int { return category.Id },
		func(product Product) int { return product.CategoryId },
		func(category Category, products Enumerable[Product]) Enumerable[Product] {
			return products
		},
	)
	enrGroupJoin := productGroups.GetEnumerator()
	for enrGroupJoin.MoveNext() {
		fmt.Println("Group")
		productGroup := enrGroupJoin.Current()
		enrProductGroup := productGroup.GetEnumerator()
		for enrProductGroup.MoveNext() {
			product := enrProductGroup.Current()
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
