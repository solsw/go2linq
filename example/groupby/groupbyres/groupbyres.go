//go:build go1.18

package main

import (
	"fmt"
	"math"

	"github.com/solsw/go2linq/v2"
)

// see GroupByEx3 example from Enumerable.GroupBy help
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
	petsList := go2linq.NewEnSlice(
		Pet{Name: "Barley", Age: 8.3},
		Pet{Name: "Boots", Age: 4.9},
		Pet{Name: "Whiskers", Age: 1.5},
		Pet{Name: "Daisy", Age: 4.3},
	)
	// Group Pet objects by the math.Floor of their Age.
	// Then project a Result type from each group that consists of the Key,
	// the Count of the group's elements, and the minimum and maximum Age in the group.
	query := go2linq.GroupByResMust(petsList,
		func(pet Pet) float64 { return math.Floor(pet.Age) },
		func(age float64, pets go2linq.Enumerable[Pet]) Result {
			count := go2linq.CountMust(pets)
			mn := go2linq.MinMust(pets,
				func(pet Pet) float64 { return pet.Age },
				go2linq.Lesser[float64](go2linq.Order[float64]{}),
			)
			mx := go2linq.MaxMust(pets,
				func(pet Pet) float64 { return pet.Age },
				go2linq.Lesser[float64](go2linq.Order[float64]{}),
			)
			return Result{Key: age, Count: count, Min: mn, Max: mx}
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
