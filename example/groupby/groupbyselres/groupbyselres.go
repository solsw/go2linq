//go:build go1.18

package main

import (
	"fmt"
	"math"

	"github.com/solsw/go2linq/v2"
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
	pets := go2linq.NewEnSlice(
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
		func(baseAge float64, ages go2linq.Enumerable[float64]) Result {
			count := go2linq.CountMust(ages)
			mn := go2linq.MinMust(ages)
			mx := go2linq.MaxMust(ages)
			return Result{Key: baseAge, Count: count, Min: mn, Max: mx}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		result := enr.Current()
		fmt.Printf("\nAge group: %g\n", result.Key)
		fmt.Printf("Number of pets in this Age group: %d\n", result.Count)
		fmt.Printf("Minimum Age: %g\n", result.Min)
		fmt.Printf("Maximum Age: %g\n", result.Max)
	}
}
