//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see MinEx3 example from Enumerable.Min help
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
		go2linq.Identity[Pet],
		// Compares Pet's ages.
		go2linq.Lesser[Pet](go2linq.LesserFunc[Pet](func(p1, p2 Pet) bool { return p1.Age < p2.Age })),
	)
	fmt.Printf("The 'minimum' animal is %s.\n", min.Name)
}
