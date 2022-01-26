//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the second example from Enumerable.First help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first

func main() {
	numbers := go2linq.NewEnSlice(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19)
	first := go2linq.FirstMust(numbers)
	fmt.Println(first)
}
