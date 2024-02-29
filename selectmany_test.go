package go2linq

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
)

// https://github.com/jskeet/edulinq/blob/master/src/Edulinq.Tests/SelectManyTest.cs

func TestSelectMany_int_rune(t *testing.T) {
	type args struct {
		source   iter.Seq[int]
		selector func(int) iter.Seq[rune]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[rune]
	}{
		{name: "SimpleFlatten",
			args: args{
				source: VarAll(3, 5, 20, 15),
				selector: func(x int) iter.Seq[rune] {
					return SliceAll([]rune(fmt.Sprint(x)))
				},
			},
			want: VarAll('3', '5', '2', '0', '1', '5'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectMany(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectMany() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectMany_int_int(t *testing.T) {
	type args struct {
		source   iter.Seq[int]
		selector func(int) iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "SimpleFlatten1",
			args: args{
				source: VarAll(1, 2, 3, 4),
				selector: func(i int) iter.Seq[int] {
					return VarAll(i, i*i)
				},
			},
			want: VarAll(1, 1, 2, 4, 3, 9, 4, 16),
		},
		{name: "SimpleFlatten2",
			args: args{
				source: VarAll(1, 2, 3, 4),
				selector: func(i int) iter.Seq[int] {
					if i%2 == 0 {
						return Empty[int]()
					}
					return VarAll(i, i*i)
				},
			},
			want: VarAll(1, 1, 3, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectMany(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectMany() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectMany_string_string(t *testing.T) {
	type args struct {
		source   iter.Seq[string]
		selector func(string) iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/projection-operations#selectmany
		{name: "SelectMany",
			args: args{
				source: VarAll("an apple a day", "the quick brown fox"),
				selector: func(s string) iter.Seq[string] {
					return SliceAll(strings.Fields(s))
				},
			},
			want: VarAll("an", "apple", "a", "day", "the", "quick", "brown", "fox"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectMany(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectMany() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectManyIdx_int_rune(t *testing.T) {
	type args struct {
		source   iter.Seq[int]
		selector func(int, int) iter.Seq[rune]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[rune]
	}{
		{name: "SimpleFlattenWithIndex",
			args: args{
				source: VarAll(3, 5, 20, 15),
				selector: func(x, idx int) iter.Seq[rune] {
					return SliceAll([]rune(fmt.Sprint(x + idx)))
				},
			},
			want: VarAll('3', '6', '2', '2', '1', '8'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectManyIdx(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectManyIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectManyIdx_int_int(t *testing.T) {
	type args struct {
		source   iter.Seq[int]
		selector func(int, int) iter.Seq[int]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{name: "SimpleFlatten",
			args: args{
				source: VarAll(1, 2, 3, 4),
				selector: func(i, idx int) iter.Seq[int] {
					if idx%2 == 0 {
						return Empty[int]()
					}
					return VarAll(i, i*i)
				},
			},
			want: VarAll(2, 4, 4, 16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectManyIdx(tt.args.source, tt.args.selector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectManyIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectManyColl_int_rune_string(t *testing.T) {
	type args struct {
		source             iter.Seq[int]
		collectionSelector func(int) iter.Seq[rune]
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "FlattenWithProjection",
			args: args{
				source: VarAll(3, 5, 20, 15),
				collectionSelector: func(x int) iter.Seq[rune] {
					return SliceAll([]rune(fmt.Sprint(x)))
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: VarAll("3: 3", "5: 5", "20: 2", "20: 0", "15: 1", "15: 5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectManyColl(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectManyColl() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

func TestSelectManyCollIdx_int_rune_string(t *testing.T) {
	type args struct {
		source             iter.Seq[int]
		collectionSelector func(int, int) iter.Seq[rune]
		resultSelector     func(int, rune) string
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[string]
	}{
		{name: "FlattenWithProjectionAndIndex",
			args: args{
				source: VarAll(3, 5, 20, 15),
				collectionSelector: func(x, idx int) iter.Seq[rune] {
					return SliceAll([]rune(fmt.Sprint(x + idx)))
				},
				resultSelector: func(x int, c rune) string {
					return fmt.Sprintf("%d: %s", x, string(c))
				},
			},
			want: VarAll("3: 3", "5: 6", "20: 2", "20: 2", "15: 1", "15: 8"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SelectManyCollIdx(tt.args.source, tt.args.collectionSelector, tt.args.resultSelector)
			equal, _ := SequenceEqual(got, tt.want)
			if !equal {
				t.Errorf("SelectManyCollIdx() = %v, want %v", StringDef(got), StringDef(tt.want))
			}
		})
	}
}

// second example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.concat#examples
func ExampleSelectMany_ex1() {
	cats := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
	}
	dogs := []Pet{
		{Name: "Bounder", Age: 3},
		{Name: "Snoopy", Age: 14},
		{Name: "Fido", Age: 9},
	}
	select1, _ := Select(SliceAll(cats), func(cat Pet) string { return cat.Name })
	select2, _ := Select(SliceAll(dogs), func(dog Pet) string { return dog.Name })
	selectMany, _ := SelectMany(VarAll(select1, select2), Identity[iter.Seq[string]])
	for name := range selectMany {
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

// SelectManyEx1 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func ExampleSelectMany_ex2() {
	petOwners := []PetOwner{
		{Name: "Higa, Sidney", Pets: []string{"Scruffy", "Sam"}},
		{Name: "Ashkenazi, Ronen", Pets: []string{"Walker", "Sugar"}},
		{Name: "Price, Vernette", Pets: []string{"Scratches", "Diesel"}},
	}

	// Query using SelectMany().
	selectMany, _ := SelectMany(
		SliceAll(petOwners),
		func(petOwner PetOwner) iter.Seq[string] { return SliceAll(petOwner.Pets) },
	)
	fmt.Println("Using SelectMany():")
	// Only one loop is required to iterate through the results since it is a one-dimensional collection.
	for pet := range selectMany {
		fmt.Println(pet)
	}

	// This code shows how to use Select() instead of SelectMany().
	petLists, _ := Select(
		SliceAll(petOwners),
		func(petOwner PetOwner) iter.Seq[string] {
			return SliceAll(petOwner.Pets)
		},
	)
	fmt.Println("\nUsing Select():")
	// Notice that two loops are required to iterate through the results
	// because the query returns a collection of sequences.
	for petList := range petLists {
		for pet := range petList {
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

// SelectManyEx2 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func ExampleSelectManyIdx() {
	petOwners := []PetOwner{
		{Name: "Higa, Sidney", Pets: []string{"Scruffy", "Sam"}},
		{Name: "Ashkenazi, Ronen", Pets: []string{"Walker", "Sugar"}},
		{Name: "Price, Vernette", Pets: []string{"Scratches", "Diesel"}},
		{Name: "Hines, Patrick", Pets: []string{"Dusty"}},
	}
	// Project the items in the array by appending the index of each PetOwner
	// to each pet's name in that petOwner's slice of pets.
	selectManyIdx, _ := SelectManyIdx(
		SliceAll(petOwners),
		func(petOwner PetOwner, index int) iter.Seq[string] {
			return errorhelper.Must(
				Select(
					SliceAll(petOwner.Pets),
					func(pet string) string { return strconv.Itoa(index) + pet },
				),
			)
		},
	)
	for pet := range selectManyIdx {
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

// SelectManyEx3 example from
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func ExampleSelectManyColl() {
	petOwners := []PetOwner{
		{Name: "Higa", Pets: []string{"Scruffy", "Sam"}},
		{Name: "Ashkenazi", Pets: []string{"Walker", "Sugar"}},
		{Name: "Price", Pets: []string{"Scratches", "Diesel"}},
		{Name: "Hines", Pets: []string{"Dusty"}},
	}
	// Project all pet's names together with the pet's owner.
	selectManyColl, _ := SelectManyColl(
		SliceAll(petOwners),
		func(petOwner PetOwner) iter.Seq[string] {
			return SliceAll(petOwner.Pets)
		},
		func(petOwner PetOwner, petName string) OwnerAndPet {
			return OwnerAndPet{petOwner: petOwner, petName: petName}
		},
	)
	// Filter only pet's names that start with S.
	where, _ := Where(selectManyColl,
		func(ownerAndPet OwnerAndPet) bool {
			return strings.HasPrefix(ownerAndPet.petName, "S")
		},
	)
	// Project the pet owner's name and the pet's name.
	ownerPetNames, _ := Select(where,
		func(ownerAndPet OwnerAndPet) OwnerNameAndPetName {
			return OwnerNameAndPetName{Owner: ownerAndPet.petOwner.Name, Pet: ownerAndPet.petName}
		},
	)
	for ownerPetName := range ownerPetNames {
		fmt.Printf("%+v\n", ownerPetName)
	}
	// Output:
	// {Owner:Higa Pet:Scruffy}
	// {Owner:Higa Pet:Sam}
	// {Owner:Ashkenazi Pet:Sugar}
	// {Owner:Price Pet:Scratches}
}
