//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the example from Enumerable.ThenBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby

func main() {
	fruits := go2linq.NewOnSliceEn("grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry")
	// Sort the strings first by their length and then alphabetically.
	query := go2linq.ThenByLsMust(
		go2linq.OrderByLsMust(fruits,
			func(fruit string) int { return len(fruit) },
			go2linq.Lesser[int](go2linq.Order[int]{}),
		),
		go2linq.Identity[string],
		go2linq.Lesser[string](go2linq.Order[string]{}),
	).GetEnumerator()
	for query.MoveNext() {
		fruit := query.Current()
		fmt.Println(fruit)
	}
}
