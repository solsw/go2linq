//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
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

	query1 := go2linq.SelectManyMust(
		go2linq.NewOnSliceEn(petOwners...),
		func(petOwner PetOwner) go2linq.Enumerator[string] { return go2linq.NewOnSlice(petOwner.Pets...) },
	)
	fmt.Println("Using SelectMany():")
	for query1.MoveNext() {
		pet := query1.Current()
		fmt.Println(pet)
	}

	query2 := go2linq.SelectMust(
		go2linq.NewOnSliceEn(petOwners...),
		func(petOwner PetOwner) go2linq.Enumerator[string] { 
			return go2linq.NewOnSliceEn(petOwner.Pets...)
		},
	)
	fmt.Println("\nUsing Select():")
	for query2.MoveNext() {
		petList := query2.Current()
		for petList.MoveNext() {
			pet := petList.Current()
			fmt.Println(pet)
		}
		fmt.Println()
	}
}
