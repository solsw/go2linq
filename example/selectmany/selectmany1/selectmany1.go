//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see SelectManyEx1 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

type PetOwner struct {
	Name string
	Pets []string
}

func main() {
	petOwners := []PetOwner{
		PetOwner{Name: "Higa, Sidney", Pets: []string{"Scruffy", "Sam"}},
		PetOwner{Name: "Ashkenazi, Ronen", Pets: []string{"Walker", "Sugar"}},
		PetOwner{Name: "Price, Vernette", Pets: []string{"Scratches", "Diesel"}},
	}

	// Query using SelectMany().
	query1 := go2linq.SelectManyMust(
		go2linq.NewEnSlice(petOwners...),
		func(petOwner PetOwner) go2linq.Enumerable[string] { return go2linq.NewEnSlice(petOwner.Pets...) },
	)
	fmt.Println("Using SelectMany():")
	// Only one for loop is required to iterate through the results since it is a one-dimensional collection.
	enr1 := query1.GetEnumerator()
	for enr1.MoveNext() {
		pet := enr1.Current()
		fmt.Println(pet)
	}

	// This code shows how to use Select() instead of SelectMany().
	query2 := go2linq.SelectMust(
		go2linq.NewEnSlice(petOwners...),
		func(petOwner PetOwner) go2linq.Enumerable[string] { 
			return go2linq.NewEnSlice(petOwner.Pets...)
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
}
