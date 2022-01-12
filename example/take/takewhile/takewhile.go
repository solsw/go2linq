//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the second example from Enumerable.TakeWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.takewhile

func main() {
	fruits := go2linq.NewOnSliceEn("apple", "banana", "mango", "orange", "passionfruit", "grape")
	query := go2linq.TakeWhileMust(fruits,
		func(fruit string) bool {
			return go2linq.CaseInsensitiveComparer.Compare("orange", fruit) != 0
		},
	)
	for query.MoveNext() {
		fruit := query.Current()
		fmt.Println(fruit)
	}
}
