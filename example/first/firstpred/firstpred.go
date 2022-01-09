//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.First help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first

func main() {
	numbers := go2linq.NewOnSliceEn(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19)
	first := go2linq.FirstPredMust(numbers, func(number int) bool { return number > 80 })
	fmt.Println(first)
}
