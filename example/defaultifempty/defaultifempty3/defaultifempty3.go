//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see DefaultIfEmptyEx2 example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

type Pet struct {
	Name string
	Age  int
}

func main() {
	defaultPet := Pet{Name: "Default Pet", Age: 0}
	pets1 := go2linq.NewOnSliceEn(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	en1 := go2linq.DefaultIfEmptyDefMust(pets1, defaultPet)
	for en1.MoveNext() {
		pet := en1.Current()
		fmt.Printf("Name: %s\n", pet.Name)
	}
	pets2 := go2linq.NewOnSliceEn([]Pet{}...)
	en2 := go2linq.DefaultIfEmptyDefMust(pets2, defaultPet)
	for en2.MoveNext() {
		pet := en2.Current()
		fmt.Printf("\nName: %s\n", pet.Name)
	}
}
