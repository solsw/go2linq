//go:build go1.18

package main

import (
	"fmt"
	"math"

	"github.com/solsw/go2linq/v2"
)

// see the last example from Enumerable.LastOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

func main() {
	numbers := go2linq.NewEnSlice(49.6, 52.3, 51.0, 49.4, 50.2, 48.3)
	last50 := go2linq.LastOrDefaultPredMust(numbers, func(n float64) bool { return math.Round(n) == 50.0 })
	fmt.Printf("The last number that rounds to 50 is %v.\n", last50)

	last40 := go2linq.LastOrDefaultPredMust(numbers, func(n float64) bool { return math.Round(n) == 40.0 })
	var what string
	if last40 == 0.0 {
		what = "<DOES NOT EXIST>"
	} else {
		what = fmt.Sprint(last40)
	}
	fmt.Printf("The last number that rounds to 40 is %v.\n", what)
}
