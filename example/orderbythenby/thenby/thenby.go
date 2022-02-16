//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the example from Enumerable.ThenBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby

func main() {
	fruits := go2linq.NewEnSlice("grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry")
	// Sort the strings first by their length and then alphabetically.
	query := go2linq.ThenBySelfMust(
		go2linq.OrderByKeyMust(fruits, func(fruit string) int { return len(fruit) }),
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
}
