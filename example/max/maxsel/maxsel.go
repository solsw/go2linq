//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see MaxEx4 example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

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
	max := go2linq.MaxMust(pets,
		func(pet Pet) int { return pet.Age + len(pet.Name) },
		go2linq.Lesser[int](go2linq.Orderer[int]{}),
	)
	fmt.Printf("The maximum pet age plus name length is %d.\n", max)
}
