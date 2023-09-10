package go2linq

import (
	"fmt"
	"strings"
	"testing"

	"github.com/solsw/collate"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/JoinTest.cs

func TestJoinMust_string_rune(t *testing.T) {
	en := NewEnSlice("fs", "sf", "ff", "ss")
	type args struct {
		outer            Enumerable[string]
		inner            Enumerable[string]
		outerKeySelector func(string) rune
		innerKeySelector func(string) rune
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "SimpleJoin",
			args: args{
				outer:            NewEnSlice("first", "second", "third"),
				inner:            NewEnSlice("essence", "offer", "eating", "psalm"),
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewEnSlice("first:offer", "second:essence", "second:psalm"),
		},
		{name: "SameEnumerable",
			args: args{
				outer:            en,
				inner:            en,
				outerKeySelector: func(oel string) rune { return ([]rune(oel))[0] },
				innerKeySelector: func(iel string) rune { return ([]rune(iel))[1] },
				resultSelector:   func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewEnSlice("fs:sf", "fs:ff", "sf:fs", "sf:ss", "ff:sf", "ff:ff", "ss:fs", "ss:ss"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinMust(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("JoinMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestJoinMust_string(t *testing.T) {
	type args struct {
		outer            Enumerable[string]
		inner            Enumerable[string]
		outerKeySelector func(string) string
		innerKeySelector func(string) string
		resultSelector   func(string, string) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "CustomComparer",
			args: args{
				outer: NewEnSlice("ABCxxx", "abcyyy", "defzzz", "ghizzz"),
				inner: NewEnSlice("000abc", "111gHi", "222333"),
				outerKeySelector: func(oel string) string {
					return strings.ToLower(oel[:3])
				},
				innerKeySelector: func(iel string) string {
					return strings.ToLower(iel[3:])
				},
				resultSelector: func(oel, iel string) string { return oel + ":" + iel },
			},
			want: NewEnSlice("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JoinMust(tt.args.outer, tt.args.inner, tt.args.outerKeySelector, tt.args.innerKeySelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("JoinMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestJoinEqMust_CustomComparer(t *testing.T) {
	outer := []string{"ABCxxx", "abcyyy", "defzzz", "ghizzz"}
	inner := []string{"000abc", "111gHi", "222333"}
	got := JoinEqMust(
		NewEnSliceEn(outer...),
		NewEnSliceEn(inner...),
		func(oel string) string { return oel[:3] },
		func(iel string) string { return iel[3:] },
		func(oel, iel string) string { return oel + ":" + iel },
		collate.CaseInsensitiveOrder,
	)
	want := NewEnSliceEn("ABCxxx:000abc", "abcyyy:000abc", "ghizzz:111gHi")
	if !SequenceEqualMust(got, want) {
		t.Errorf("JoinEqMust_CustomComparer = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

func TestJoinMust_DifferentSourceTypes(t *testing.T) {
	outer := []int{5, 3, 7}
	inner := []string{"bee", "giraffe", "tiger", "badger", "ox", "cat", "dog"}
	got := JoinMust(
		NewEnSliceEn(outer...),
		NewEnSliceEn(inner...),
		Identity[int],
		func(iel string) int { return len(iel) },
		func(oel int, iel string) string { return fmt.Sprintf("%d:%s", oel, iel) },
	)
	want := NewEnSliceEn("5:tiger", "3:bee", "3:cat", "3:dog", "7:giraffe")
	if !SequenceEqualMust(got, want) {
		t.Errorf("JoinMust_DifferentSourceTypes = %v, want %v", ToStringDef(got), ToStringDef(want))
	}
}

// see JoinEx1 example from Enumerable.Join help
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.join
func ExampleJoinMust_ex1() {
	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	people := []Person{magnus, terry, charlotte}
	pets := []Pet{barley, boots, whiskers, daisy}

	// Create a list of Person-Pet pairs where each element is an OwnerNameAndPetName type that contains a
	// Pet's name and the name of the Person that owns the Pet.
	join := JoinMust(
		NewEnSliceEn(people...),
		NewEnSliceEn(pets...),
		Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pet Pet) OwnerNameAndPetName {
			return OwnerNameAndPetName{Owner: person.Name, Pet: pet.Name}
		},
	)
	enr := join.GetEnumerator()
	for enr.MoveNext() {
		obj := enr.Current()
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
func ExampleJoinMust_ex2() {
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
	join := JoinMust(
		NewEnSliceEn(products...),
		NewEnSliceEn(categories...),
		func(product Product) int { return product.CategoryId },
		func(category Category) int { return category.Id },
		func(product Product, category Category) string {
			return fmt.Sprintf("%s - %s", product.Name, category.CategoryName)
		},
	)
	enr := join.GetEnumerator()
	for enr.MoveNext() {
		item := enr.Current()
		fmt.Println(item)
	}
	// Output:
	// Cola - Beverage
	// Tea - Beverage
	// Apple - Fruit
	// Kiwi - Fruit
	// Carrot - Vegetable
}
