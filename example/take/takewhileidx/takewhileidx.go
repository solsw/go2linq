//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.TakeWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile

func main() {
	fruits := go2linq.NewOnSliceEn("apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry")
	query := go2linq.TakeWhileIdxMust(fruits,
		func(fruit string, index int) bool {
			return len(fruit) >= index
		},
	)
	for query.MoveNext() {
		fruit := query.Current()
		fmt.Println(fruit)
	}
}
