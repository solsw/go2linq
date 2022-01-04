//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Where help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func main() {
	fruits := []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry", "grape", "strawberry"}
	query := go2linq.WhereMust(
		go2linq.NewOnSliceEn(fruits...),
		func(fruit string) bool { return len(fruit) < 6 },
	)
	for query.MoveNext() {
		fruit := query.Current()
		fmt.Println(fruit)
	}
}
