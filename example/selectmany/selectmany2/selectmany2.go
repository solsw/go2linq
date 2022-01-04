//go:build go1.18

package main

import (
	"fmt"
	"strconv"

	"github.com/solsw/go2linq"
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
	query := go2linq.SelectManyIdxMust(
		go2linq.NewOnSliceEn(petOwners...),
		func(petOwner PetOwner, index int) go2linq.Enumerator[string] {
			return go2linq.SelectMust(
				go2linq.NewOnSliceEn(petOwner.Pets...),
				func(pet string) string { return strconv.Itoa(index) + pet },
			)
		})
	for query.MoveNext() {
		pet := query.Current()
		fmt.Println(pet)
	}
}
