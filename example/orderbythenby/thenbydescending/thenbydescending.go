//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see ThenByDescendingEx1 example from Enumerable.ThenByDescending help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

func main() {
	fruits := go2linq.NewEnSlice("apPLe", "baNanA", "apple", "APple", "orange", "BAnana", "ORANGE", "apPLE")
	// Sort the strings first ascending by their length and then descending using a custom case insensitive comparer.
	query := go2linq.ThenBySelfDescLsMust(
		go2linq.OrderByKeyMust(fruits, func(fruit string) int { return len(fruit) }),
		go2linq.CaseInsensitiveLesser,
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
}
