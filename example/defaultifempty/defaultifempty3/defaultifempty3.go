//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see DefaultIfEmptyEx2 example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

type Pet struct {
	Name string
	Age  int
}

func main() {
	defaultPet := Pet{Name: "Default Pet", Age: 0}
	pets1 := go2linq.NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	enr1 := go2linq.DefaultIfEmptyDefMust(pets1, defaultPet).GetEnumerator()
	for enr1.MoveNext() {
		pet := enr1.Current()
		fmt.Printf("Name: %s\n", pet.Name)
	}
	pets2 := go2linq.NewEnSlice([]Pet{}...)
	enr2 := go2linq.DefaultIfEmptyDefMust(pets2, defaultPet).GetEnumerator()
	for enr2.MoveNext() {
		pet := enr2.Current()
		fmt.Printf("\nName: %s\n", pet.Name)
	}
}
