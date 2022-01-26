//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Union help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

func main() {
	ints1 := go2linq.NewEnSlice(5, 3, 9, 7, 5, 9, 3, 7)
	ints2 := go2linq.NewEnSlice(8, 3, 6, 4, 4, 9, 1, 0)
	union := go2linq.UnionMust(ints1, ints2)
	enr := union.GetEnumerator()
	for enr.MoveNext() {
		num := enr.Current()
		fmt.Printf("%d ", num)
	}
}
