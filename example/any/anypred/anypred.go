//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see AnyEx3 example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

type Pet struct {
	Name       string
	Age        int
	Vaccinated bool
}

func main() {
	pets := []Pet{
		Pet{Name: "Barley", Age: 8, Vaccinated: true},
		Pet{Name: "Boots", Age: 4, Vaccinated: false},
		Pet{Name: "Whiskers", Age: 1, Vaccinated: false},
	}
	// Determine whether any pets over Age 1 are also unvaccinated.
	unvaccinated := go2linq.AnyPredMust(
		go2linq.NewEnSlice(pets...),
		func(pet Pet) bool { return pet.Age > 1 && pet.Vaccinated == false },
	)
	var what string
	if unvaccinated {
		what = "are"
	} else {
		what = "are not any"
	}
	fmt.Printf("There %s unvaccinated animals over age one.\n", what)
}
