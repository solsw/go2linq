//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see CountEx2 example from Enumerable.Count help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count

type Pet struct {
	Name       string
	Vaccinated bool
}

func main() {
	pets := go2linq.NewOnSliceEn(
		Pet{Name: "Barley", Vaccinated: true},
		Pet{Name: "Boots", Vaccinated: false},
		Pet{Name: "Whiskers", Vaccinated: false},
	)
	numberUnvaccinated := go2linq.CountPredMust(pets, func(p Pet) bool { return p.Vaccinated == false })
	fmt.Printf("There are %d unvaccinated animals.\n", numberUnvaccinated)
}
