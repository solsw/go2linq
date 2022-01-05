//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

func main() {
	numbers := go2linq.NewOnSliceEn(1, 2)
	hasElements := go2linq.AnyMust(numbers)
	var what string
	if hasElements {
		what = "is not"
	} else {
		what = "is"
	}
	fmt.Printf("The list %s empty.\n", what)
}
