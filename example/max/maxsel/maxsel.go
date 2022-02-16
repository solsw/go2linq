//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see MaxEx4 example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

type Pet struct {
	Name string
	Age  int
}

func main() {
	pets := go2linq.NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	max := go2linq.MaxSelMust(pets, func(pet Pet) int { return pet.Age + len(pet.Name) })
	fmt.Printf("The maximum pet age plus name length is %d.\n", max)
}
