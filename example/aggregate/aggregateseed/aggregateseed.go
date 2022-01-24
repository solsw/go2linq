//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the second example from Enumerable.Aggregate help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

func main() {
	ints := go2linq.NewEnSlice(4, 8, 8, 3, 9, 0, 7, 8, 2)
	// Count the even numbers in the array, using a seed value of 0.
	numEven := go2linq.AggregateSeedMust(ints,
		0,
		func(total, next int) int {
			if next%2 == 0 {
				return total + 1
			}
			return total
		},
	)
	fmt.Printf("The number of even integers is: %d\n", numEven)
}
