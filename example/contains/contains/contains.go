//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Contains help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

func main() {
	fruits := go2linq.NewOnSliceEn("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit := "mango"
	hasMango := go2linq.ContainsMust(fruits, fruit)
	var what string
	if hasMango {
		what = "does"
	} else {
		what = "does not"
	}
	fmt.Printf("The array %s contain '%s'.\n", what, fruit)
}
