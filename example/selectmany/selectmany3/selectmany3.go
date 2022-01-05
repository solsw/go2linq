//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq"
)

// see SelectManyEx3 example from Enumerable.SelectMany help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

type (
	PetOwner struct {
		Name string
		Pets []string
	}
	OwnerAndPet struct {
		petOwner PetOwner
		petName  string
	}
	OwnerNameAndPetName struct {
		Owner string
		Pet   string
	}
)

func main() {
	petOwners := []PetOwner{
		PetOwner{Name: "Higa", Pets: []string{"Scruffy", "Sam"}},
		PetOwner{Name: "Ashkenazi", Pets: []string{"Walker", "Sugar"}},
		PetOwner{Name: "Price", Pets: []string{"Scratches", "Diesel"}},
		PetOwner{Name: "Hines", Pets: []string{"Dusty"}},
	}
	// Project all pet's names together with the pet's owner.
	selectManyQuery := go2linq.SelectManyCollMust(
		go2linq.NewOnSliceEn(petOwners...),
		func(petOwner PetOwner) go2linq.Enumerator[string] {
			return go2linq.NewOnSliceEn(petOwner.Pets...)
		},
		func(petOwner PetOwner, petName string) OwnerAndPet {
			return OwnerAndPet{petOwner: petOwner, petName: petName}
		},
	)
	// Filter only pet's names that start with S.
	whereQuery := go2linq.WhereMust(
		selectManyQuery,
		func(ownerAndPet OwnerAndPet) bool {
			return strings.HasPrefix(ownerAndPet.petName, "S")
		},
	)
	// Project the pet owner's name and the pet's name.
	selectQuery := go2linq.SelectMust(
		whereQuery,
		func(ownerAndPet OwnerAndPet) OwnerNameAndPetName {
			return OwnerNameAndPetName{Owner: ownerAndPet.petOwner.Name, Pet: ownerAndPet.petName}
		},
	)
	for selectQuery.MoveNext() {
		obj := selectQuery.Current()
		fmt.Printf("%+v\n", obj)
	}
}
