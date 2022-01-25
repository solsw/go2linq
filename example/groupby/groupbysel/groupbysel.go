//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see GroupByEx1 example from Enumerable.GroupBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

type Pet struct {
	Name string
	Age  int
}

func main() {
	pets := go2linq.NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
		Pet{Name: "Daisy", Age: 4},
	)
	// Group the pets using Age as the key value and selecting only the Pet's Name for each value.
	query := go2linq.GroupBySelMust(pets,
		func(pet Pet) int { return pet.Age },
		func(pet Pet) string { return pet.Name },
	)
	// Iterate over each Grouping in the collection.
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		petGroup := enr.Current()
		// Print the key value of the Grouping.
		fmt.Println(petGroup.Key())
		names := petGroup.GetEnumerator()
		// Iterate over each value in the Grouping and print the value.
		for names.MoveNext() {
			name := names.Current()
			fmt.Printf("  %s\n", name)
		}
	}
}
