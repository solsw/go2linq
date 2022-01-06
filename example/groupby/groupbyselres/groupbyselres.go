//go:build go1.18

package main

import (
	"fmt"
	"math"

	"github.com/solsw/go2linq"
)

// see GroupByEx4 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

type Pet struct {
	Name string
	Age  float64
}

type Result struct {
	Key      float64
	Count    int
	Min, Max float64
}

func main() {
	pets := go2linq.NewOnSliceEn(
		Pet{Name: "Barley", Age: 8.3},
		Pet{Name: "Boots", Age: 4.9},
		Pet{Name: "Whiskers", Age: 1.5},
		Pet{Name: "Daisy", Age: 4.3},
	)
	// Group Pet.Age values by the math.Floor of the age.
	// Then project a Result type from each group that consists of the Key,
	// the Count of the group's elements, and the minimum and maximum Age in the group.
	query := go2linq.GroupBySelResMust(pets,
		func(pet Pet) float64 { return math.Floor(pet.Age) },
		func(pet Pet) float64 { return pet.Age },
		func(baseAge float64, ages go2linq.Enumerator[float64]) Result {
			c := go2linq.CountMust(ages)
			ages.Reset()
			mn := go2linq.MinMust(ages,
				go2linq.Identity[float64],
				go2linq.Lesser[float64](go2linq.Orderer[float64]{}),
			)
			ages.Reset()
			mx := go2linq.MaxMust(ages,
				go2linq.Identity[float64],
				go2linq.Lesser[float64](go2linq.Orderer[float64]{}),
			)
			return Result{Key: baseAge, Count: c, Min: mn, Max: mx}
		},
	)
	for query.MoveNext() {
		result := query.Current()
		fmt.Printf("\nAge group: %g\n", result.Key)
		fmt.Printf("Number of pets in this Age group: %d\n", result.Count)
		fmt.Printf("Minimum Age: %g\n", result.Min)
		fmt.Printf("Maximum Age: %g\n", result.Max)
	}
}
