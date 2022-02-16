//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min

func main() {
	doubles := go2linq.NewEnSlice(1.5e+104, 9e+103, -2e+103)
	min := go2linq.MinMust(doubles)
	fmt.Printf("The smallest number is %G.\n", min)
}
