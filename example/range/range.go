//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the example from Enumerable.Range help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.range#examples

func main() {
	// Generate a sequence of integers from 1 to 10 and then select their squares.
	squares := go2linq.SelectMust(go2linq.RangeMust(1, 10), func(x int) int { return x * x })
	enr := squares.GetEnumerator()
	for enr.MoveNext() {
		num := enr.Current()
		fmt.Println(num)
	}
}
