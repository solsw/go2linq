//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.Where help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func main() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	query := go2linq.WhereMust(
		go2linq.NewEnSlice(fruits...),
		func(fruit string) bool { return len(fruit) < 6 },
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		fruit := enr.Current()
		fmt.Println(fruit)
	}
}
