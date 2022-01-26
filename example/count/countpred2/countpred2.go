//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see LongCountEx2 example from Enumerable.LongCount help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount

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
	const Age = 3
	count := go2linq.CountPredMust(pets, func(pet Pet) bool { return pet.Age > Age })
	fmt.Printf("There are %d animals over age %d.\n", count, Age)
}
