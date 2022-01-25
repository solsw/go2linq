//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see JoinEx1 example from Enumerable.Join help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join

type (
	Person struct {
		Name string
	}
	Pet struct {
		Name  string
		Owner Person
	}
	OwnerNamePet struct {
		OwnerName string
		Pet       string
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

	people := go2linq.NewEnSlice(magnus, terry, charlotte)
	pets := go2linq.NewEnSlice(barley, boots, whiskers, daisy)

	// Create a list of Person-Pet pairs where each element is an OwnerNamePet type that contains a
	// Pet's name and the name of the Person that owns the Pet.
	query := go2linq.JoinMust(people, pets,
		go2linq.Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pet Pet) OwnerNamePet {
			return OwnerNamePet{OwnerName: person.Name, Pet: pet.Name}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		obj := enr.Current()
		fmt.Printf("%s - %s\n", obj.OwnerName, obj.Pet)
	}
}
