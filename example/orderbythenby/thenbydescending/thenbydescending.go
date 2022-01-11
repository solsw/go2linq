//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see ThenByDescendingEx1 example from Enumerable.ThenByDescending help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

func main() {
	fruits := go2linq.NewOnSliceEn("apPLe", "baNanA", "apple", "APple", "orange", "BAnana", "ORANGE", "apPLE")
	// Sort the strings first ascending by their length and then descending using a custom case insensitive comparer.
	query := go2linq.ThenByDescendingLsMust(
		go2linq.OrderByLsMust(fruits,
			func(fruit string) int { return len(fruit) },
			go2linq.Lesser[int](go2linq.Order[int]{}),
		),
		go2linq.Identity[string],
		go2linq.CaseInsensitiveLesser,
	).GetEnumerator()
	for query.MoveNext() {
		fruit := query.Current()
		fmt.Println(fruit)
	}
}
