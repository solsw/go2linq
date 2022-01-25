//go:build go1.18

package main

import (
	"fmt"
	"strconv"

	"github.com/solsw/go2linq/v2"
)

// see SelectManyEx2 example from Enumerable.SelectMany help
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
		PetOwner{Name: "Hines, Patrick", Pets: []string{"Dusty"}},
	}
	// Project the items in the array by appending the index of each PetOwner
	// to each pet's name in that petOwner's array of pets.
	query := go2linq.SelectManyIdxMust(
		go2linq.NewEnSlice(petOwners...),
		func(petOwner PetOwner, index int) go2linq.Enumerable[string] {
			return go2linq.SelectMust(
				go2linq.NewEnSlice(petOwner.Pets...),
				func(pet string) string { return strconv.Itoa(index) + pet },
			)
		})
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		pet := enr.Current()
		fmt.Println(pet)
	}
}
