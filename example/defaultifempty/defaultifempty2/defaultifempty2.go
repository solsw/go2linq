//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see DefaultIfEmptyEx1 example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

type Pet struct {
	Name string
	Age  int
}

func main() {
	pets := go2linq.NewOnSliceEn(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	en := go2linq.DefaultIfEmptyMust(pets)
	for en.MoveNext() {
		pet := en.Current()
		fmt.Println(pet.Name)
	}
}
