//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq"
)

// see AllEx example from Enumerable.All help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all#examples

type Pet struct {
	Name string
	Age  int
}

func main() {
	pets := []Pet{
		Pet{Name: "Barley", Age: 10},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 6},
	}
	// Determine whether all Pet names in the array start with 'B'.
	allStartWithB := go2linq.AllMust(go2linq.NewOnSliceEn(pets...),
		func(pet Pet) bool { return strings.HasPrefix(pet.Name, "B") },
	)
	var what string
	if allStartWithB {
		what = "All"
	} else {
		what = "Not all"
	}
	fmt.Printf("%s pet names start with 'B'.\n", what)
}
