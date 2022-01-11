//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the second example from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits2 := go2linq.NewOnSliceEn("orange", "apple")
	fruit2, err := go2linq.Single(fruits2)
	if err == go2linq.ErrMultipleElements {
		fmt.Println("The collection does not contain exactly one element.")
	} else {
		fmt.Println(fruit2)
	}
}
