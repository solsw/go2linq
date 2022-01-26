//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Max help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.max

func main() {
	longs := go2linq.NewEnSlice(4294967296, 466855135, 81125)
	max := go2linq.MaxMust(longs, go2linq.Identity[int], go2linq.Lesser[int](go2linq.Order[int]{}))
	fmt.Printf("The largest number is %d.\n", max)
}
