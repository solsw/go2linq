//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Union help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

func main() {
	ints1 := go2linq.NewOnSliceEn(5, 3, 9, 7, 5, 9, 3, 7)
	ints2 := go2linq.NewOnSliceEn(8, 3, 6, 4, 4, 9, 1, 0)
	union := go2linq.UnionMust(ints1, ints2)
	for union.MoveNext() {
		num := union.Current()
		fmt.Printf("%d ", num)
	}
}
