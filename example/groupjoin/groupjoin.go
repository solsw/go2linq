//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see GroupJoinEx1 example from Enumerable.GroupJoin help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin

type (
	Person struct {
		Name string
	}
	Pet struct {
		Name  string
		Owner Person
	}
	OwnerAndPets struct {
		OwnerName string
		Pets      go2linq.Enumerator[string]
	}
)

func main() {
	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	// Create a list where each element is an OwnerAndPets type that contains a person's name and
	// a collection of names of the pets they own.
	people := go2linq.NewOnSliceEn(magnus, terry, charlotte)
	pets := go2linq.NewOnSliceEn(barley, boots, whiskers, daisy)

	query := go2linq.GroupJoinMust(people, pets,
		go2linq.Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, petCollection go2linq.Enumerator[Pet]) OwnerAndPets {
			return OwnerAndPets{
				OwnerName: person.Name,
				Pets:      go2linq.SelectMust(petCollection, func(pet Pet) string { return pet.Name })}
		})
	for query.MoveNext() {
		obj := query.Current()
		// Output the owner's name.
		fmt.Printf("%s:\n", obj.OwnerName)
		// Output each of the owner's pet's names.
		for obj.Pets.MoveNext() {
			pet := obj.Pets.Current()
			fmt.Printf("  %s\n", pet)
		}
	}
}
