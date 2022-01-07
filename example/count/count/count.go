//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Count help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count

func main() {
	fruits := go2linq.NewOnSliceEn("apple", "banana", "mango", "orange", "passionfruit", "grape")
	numberOfFruits := go2linq.CountMust(fruits)
	fmt.Printf("There are %d fruits in the collection.\n", numberOfFruits)
}
