//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see MinEx4 example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min

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
	min := go2linq.MinMust(pets,
		func(pet Pet) int { return pet.Age },
		go2linq.Lesser[int](go2linq.Order[int]{}),
	)
	fmt.Printf("The youngest animal is age %d.\n", min)
}
