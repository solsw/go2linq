//go:build go1.18

package go2linq

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectManyTest.cs

func TestSelectManyMust_int_rune(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int) Enumerable[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[rune]
	}{
		{name: "SimpleFlatten",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				selector: func(x int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x))...)
				},
			},
			want: NewEnSlice('3', '5', '2', '0', '1', '5'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectManyMust_int_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int) Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "SimpleFlatten1",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				selector: func(i int) Enumerable[int] {
					return NewEnSlice(i, i*i)
				},
			},
			want: NewEnSlice(1, 1, 2, 4, 3, 9, 4, 16),
		},
		{name: "SimpleFlatten2",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				selector: func(i int) Enumerable[int] {
					if i%2 == 0 {
						return Empty[int]()
					}
					return NewEnSlice(i, i*i)
				},
			},
			want: NewEnSlice(1, 1, 3, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectManyMust_string_string(t *testing.T) {
	type args struct {
		source   Enumerable[string]
		selector func(string) Enumerable[string]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#selectmany
		{name: "SelectMany",
			args: args{
				source: NewEnSlice("an apple a day", "the quick brown fox"),
				selector: func(s string) Enumerable[string] {
					return NewEnSlice(strings.Fields(s)...)
				},
			},
			want: NewEnSlice("an", "apple", "a", "day", "the", "quick", "brown", "fox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectManyIdxMust_int_rune(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int, int) Enumerable[rune]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[rune]
	}{
		{name: "SimpleFlattenWithIndex",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				selector: func(x, idx int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x + idx))...)
				},
			},
			want: NewEnSlice('3', '6', '2', '2', '1', '8'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyIdxMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectManyIdxMust_int_int(t *testing.T) {
	type args struct {
		source   Enumerable[int]
		selector func(int, int) Enumerable[int]
	}
	tests := []struct {
		name string
		args args
		want Enumerable[int]
	}{
		{name: "1",
			args: args{
				source: NewEnSlice(1, 2, 3, 4),
				selector: func(i, idx int) Enumerable[int] {
					if idx%2 == 0 {
						return Empty[int]()
					}
					return NewEnSlice(i, i*i)
				},
			},
			want: NewEnSlice(2, 4, 4, 16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyIdxMust(tt.args.source, tt.args.selector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectManyCollMust_int_rune_string(t *testing.T) {
	type args struct {
		source             Enumerable[int]
		collectionSelector func(int) Enumerable[rune]
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "FlattenWithProjection",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				collectionSelector: func(x int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x))...)
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: NewEnSlice("3: 3", "5: 5", "20: 2", "20: 0", "15: 1", "15: 5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyCollMust(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyCollMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

func TestSelectManyCollIdxMust_int_rune_string(t *testing.T) {
	type args struct {
		source             Enumerable[int]
		collectionSelector func(int, int) Enumerable[rune]
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want Enumerable[string]
	}{
		{name: "FlattenWithProjectionAndIndex",
			args: args{
				source: NewEnSlice(3, 5, 20, 15),
				collectionSelector: func(x, idx int) Enumerable[rune] {
					return NewEnSlice([]rune(fmt.Sprint(x + idx))...)
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: NewEnSlice("3: 3", "5: 6", "20: 2", "20: 2", "15: 1", "15: 8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectManyCollIdxMust(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			if !SequenceEqualMust(got, tt.want) {
				t.Errorf("SelectManyCollIdxMust() = %v, want %v", ToStringDef(got), ToStringDef(tt.want))
			}
		})
	}
}

// see the second example from Enumerable.Concat help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat#examples
func ExampleSelectManyMust_ex1() {
	cats := NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	dogs := NewEnSlice(
		Pet{Name: "Bounder", Age: 3},
		Pet{Name: "Snoopy", Age: 14},
		Pet{Name: "Fido", Age: 9},
	)
	query := SelectManyMust(
		NewEnSlice(
			SelectMust(cats, func(cat Pet) string { return cat.Name }),
			SelectMust(dogs, func(dog Pet) string { return dog.Name }),
		),
		Identity[Enumerable[string]],
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Println(name)
	}
	// Output:
	// Barley
	// Boots
	// Whiskers
	// Bounder
	// Snoopy
	// Fido
}

// see SelectManyEx1 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func ExampleSelectManyMust_ex2() {
	petOwners := NewEnSlice(
		PetOwner{Name: "Higa, Sidney", Pets: []string{"Scruffy", "Sam"}},
		PetOwner{Name: "Ashkenazi, Ronen", Pets: []string{"Walker", "Sugar"}},
		PetOwner{Name: "Price, Vernette", Pets: []string{"Scratches", "Diesel"}},
	)

	// Query using SelectMany().
	query1 := SelectManyMust(petOwners,
		func(petOwner PetOwner) Enumerable[string] { return NewEnSlice(petOwner.Pets...) },
	)
	fmt.Println("Using SelectMany():")
	// Only one for loop is required to iterate through the results since it is a one-dimensional collection.
	enr1 := query1.GetEnumerator()
	for enr1.MoveNext() {
		pet := enr1.Current()
		fmt.Println(pet)
	}

	// This code shows how to use Select() instead of SelectMany().
	query2 := SelectMust(petOwners,
		func(petOwner PetOwner) Enumerable[string] {
			return NewEnSlice(petOwner.Pets...)
		},
	)
	fmt.Println("\nUsing Select():")
	// Notice that two foreach loops are required to iterate through the results
	// because the query returns a collection of arrays.
	enr2 := query2.GetEnumerator()
	for enr2.MoveNext() {
		petList := enr2.Current()
		enrPetList := petList.GetEnumerator()
		for enrPetList.MoveNext() {
			pet := enrPetList.Current()
			fmt.Println(pet)
		}
		fmt.Println()
	}
	// Output:
	// Using SelectMany():
	// Scruffy
	// Sam
	// Walker
	// Sugar
	// Scratches
	// Diesel
	//
	// Using Select():
	// Scruffy
	// Sam
	//
	// Walker
	// Sugar
	//
	// Scratches
	// Diesel
}

// see SelectManyEx2 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func ExampleSelectManyIdxMust() {
	petOwners := NewEnSlice(
		PetOwner{Name: "Higa, Sidney", Pets: []string{"Scruffy", "Sam"}},
		PetOwner{Name: "Ashkenazi, Ronen", Pets: []string{"Walker", "Sugar"}},
		PetOwner{Name: "Price, Vernette", Pets: []string{"Scratches", "Diesel"}},
		PetOwner{Name: "Hines, Patrick", Pets: []string{"Dusty"}},
	)
	// Project the items in the array by appending the index of each PetOwner
	// to each pet's name in that petOwner's array of pets.
	query := SelectManyIdxMust(petOwners,
		func(petOwner PetOwner, index int) Enumerable[string] {
			return SelectMust(
				NewEnSlice(petOwner.Pets...),
				func(pet string) string { return strconv.Itoa(index) + pet },
			)
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		pet := enr.Current()
		fmt.Println(pet)
	}
	// Output:
	// 0Scruffy
	// 0Sam
	// 1Walker
	// 1Sugar
	// 2Scratches
	// 2Diesel
	// 3Dusty
}

// see SelectManyEx3 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func ExampleSelectManyCollMust() {
	petOwners := NewEnSlice(
		PetOwner{Name: "Higa", Pets: []string{"Scruffy", "Sam"}},
		PetOwner{Name: "Ashkenazi", Pets: []string{"Walker", "Sugar"}},
		PetOwner{Name: "Price", Pets: []string{"Scratches", "Diesel"}},
		PetOwner{Name: "Hines", Pets: []string{"Dusty"}},
	)
	// Project all pet's names together with the pet's owner.
	selectMany := SelectManyCollMust(petOwners,
		func(petOwner PetOwner) Enumerable[string] {
			return NewEnSlice(petOwner.Pets...)
		},
		func(petOwner PetOwner, petName string) OwnerAndPet {
			return OwnerAndPet{petOwner: petOwner, petName: petName}
		},
	)
	// Filter only pet's names that start with S.
	where := WhereMust(selectMany,
		func(ownerAndPet OwnerAndPet) bool {
			return strings.HasPrefix(ownerAndPet.petName, "S")
		},
	)
	// Project the pet owner's name and the pet's name.
	selectEn := SelectMust(where,
		func(ownerAndPet OwnerAndPet) OwnerNameAndPetName {
			return OwnerNameAndPetName{Owner: ownerAndPet.petOwner.Name, Pet: ownerAndPet.petName}
		},
	)
	enr := selectEn.GetEnumerator()
	for enr.MoveNext() {
		obj := enr.Current()
		fmt.Printf("%+v\n", obj)
	}
	// Output:
	// {Owner:Higa Pet:Scruffy}
	// {Owner:Higa Pet:Sam}
	// {Owner:Ashkenazi Pet:Sugar}
	// {Owner:Price Pet:Scratches}
}
