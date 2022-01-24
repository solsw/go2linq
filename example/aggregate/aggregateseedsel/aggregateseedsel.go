//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Aggregate help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

func main() {
	fruits := go2linq.NewEnSlice("apple", "mango", "orange", "passionfruit", "grape")
	// Determine whether any string in the array is longer than "banana".
	longestName := go2linq.AggregateSeedSelMust(fruits,
		"banana",
		func(longest, next string) string {
			if len(next) > len(longest) {
				return next
			}
			return longest
		},
		// Return the final result as an upper case string.
		func(fruit string) string { return strings.ToUpper(fruit) },
	)
	fmt.Printf("The fruit with the longest name is %s.\n", longestName)
}
