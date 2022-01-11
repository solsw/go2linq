//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see SequenceEqualEx1 example from Enumerable.SequenceEqual help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

type Pet struct {
	Name string
	Age  int
}

func main() {
	pet1 := Pet{Name: "Turbo", Age: 2}
	pet2 := Pet{Name: "Peanut", Age: 8}
	pets1 := go2linq.NewOnSliceEn(pet1, pet2)
	pets2 := go2linq.NewOnSliceEn(pet1, pet2)
	equal := go2linq.SequenceEqualMust(pets1, pets2)
	var what string
	if equal {
		what = "are"
	} else {
		what = "are not"
	}
	fmt.Printf("The lists %s equal.\n", what)
}
