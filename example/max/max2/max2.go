//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see MaxEx3 example from Enumerable.Max help
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
		go2linq.Identity[Pet],
		// Compares Pets by summing each Pet's age and name length.
		go2linq.Lesser[Pet](go2linq.LesserFunc[Pet](
			func(p1, p2 Pet) bool { return p1.Age+len(p1.Name) < p2.Age+len(p2.Name) },
		)),
	)
	fmt.Printf("The 'maximum' animal is %s.\n", max.Name)
}
