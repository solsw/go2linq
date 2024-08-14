package go2linq

import (
	"fmt"
	"iter"
	"slices"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/JoinTest.cs

func TestJoin_string_rune(t *testing.T) {
	seq := VarToSeq("fs", "sf", "ff", "ss")
	type args struct {
		outer            iter.Seq[string]
		inner            iter.Seq[string]
		outerKeySelector func(string) rune
		innerKeySelector func(string) rune
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "SimpleJoin",
			args: args{
				outer:            VarToSeq("first", "second", "third"),
				inner:            VarToSeq("essence", "offer", "eating", "psalm"),
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: VarToSeq("first:offer", "second:essence", "second:psalm"),
		},
		{name: "SameEnumerable",
			args: args{
				outer:            seq,
				inner:            seq,
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: VarToSeq("fs:sf", "fs:ff", "sf:fs", "sf:ss", "ff:sf", "ff:ff", "ss:fs", "ss:ss"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Join(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Join() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestJoin_string(t *testing.T) {
	type args struct {
		outer            iter.Seq[string]
		inner            iter.Seq[string]
		outerKeySelector func(string) string
		innerKeySelector func(string) string
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "CustomComparer",
			args: args{
				outer: VarToSeq("ABCxxx", "abcyyy", "defzzz", "ghizzz"),
				inner: VarToSeq("000abc", "111gHi", "222333"),
				outerKeySelector: func(oel string) string {
					return strings.ToLower(oel[:3])
				},
				innerKeySelector: func(iel string) string {
					return strings.ToLower(iel[3:])
				},
				resultSelector: func(oel, iel string) string { return oel + ":" + iel },
			},
			want: VarToSeq("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Join(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("Join() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestJoinEq_CustomComparer(t *testing.T) {
	got, _ := JoinEq(
		VarToSeq("ABCxxx", "abcyyy", "defzzz", "ghizzz"),
		VarToSeq("000abc", "111gHi", "222333"),
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel, iel string) string { return oel + ":" + iel },
		caseInsensitiveEqual,
	)
	want := VarToSeq("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("JoinEq = %v, want %v", StringDef(got), StringDef(want))
	}
}

func TestJoin_DifferentSourceTypes(t *testing.T) {
	got, _ := Join(
		VarToSeq(5, 3, 7),
		VarToSeq("bee", "giraffe", "tiger", "badger", "ox", "cat", "dog"),
		Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iel string) string { return fmt.Sprintf("%d:%s", oel, iel) },
	)
	want := VarToSeq("5:tiger", "3:bee", "3:cat", "3:dog", "7:giraffe")
	equal, _ := SequenceEqual(got, want)
	if !equal {
		t.Errorf("Join = %v, want %v", StringDef(got), StringDef(want))
	}
}

// JoinEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.join
func ExampleJoin_ex1() {
	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	// Create a list of Person-Pet pairs where each element is an OwnerNameAndPetName type that contains a
	// Pet's name and the name of the Person that owns the Pet.
	join, _ := Join(
		VarToSeq(magnus, terry, charlotte),
		VarToSeq(barley, boots, whiskers, daisy),
		Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pet Pet) OwnerNameAndPetName {
			return OwnerNameAndPetName{Owner: person.Name, Pet: pet.Name}
		},
	)
	for obj := range join {
		fmt.Printf("%s - %s\n", obj.Owner, obj.Pet)
	}
	// Output:
	// Hedlund, Magnus - Daisy
	// Adams, Terry - Barley
	// Adams, Terry - Boots
	// Weiss, Charlotte - Whiskers
}

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#query-expression-syntax-examples
// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/join-operations#join
func ExampleJoin_ex2() {
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
	// Join products and categories based on CategoryId
	join, _ := Join(
		slices.Values(products),
		slices.Values(categories),
		func(product Product) int { return product.CategoryId },
		func(category Category) int { return category.Id },
		func(product Product, category Category) string {
			return fmt.Sprintf("%s - %s", product.Name, category.CategoryName)
		},
	)
	for item := range join {
		fmt.Println(item)
	}
	// Output:
	// Cola - Beverage
	// Tea - Beverage
	// Apple - Fruit
	// Kiwi - Fruit
	// Carrot - Vegetable
}
