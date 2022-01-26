//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.TakeWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile

func main() {
	fruits := go2linq.NewEnSlice("apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry")
	query := go2linq.TakeWhileIdxMust(fruits,
		func(fruit string, index int) bool {
			return len(fruit) >= index
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
}
