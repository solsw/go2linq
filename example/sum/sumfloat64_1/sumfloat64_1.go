//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Sum help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

func main() {
	numbers := go2linq.NewOnSliceEn(43.68, 1.25, 583.7, 6.5)
	sum := go2linq.SumMust(numbers, func(f float64) float64 { return f })
	fmt.Printf("The sum of the numbers is %g.\n", sum)
}
