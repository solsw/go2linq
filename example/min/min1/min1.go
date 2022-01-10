//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Min help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.min

func main() {
	doubles := go2linq.NewOnSliceEn(1.5e+104, 9e+103, -2e+103)
	min := go2linq.MinMust(doubles, go2linq.Identity[float64], go2linq.Lesser[float64](go2linq.Orderer[float64]{}))
	fmt.Printf("The smallest number is %G.\n", min)
}
