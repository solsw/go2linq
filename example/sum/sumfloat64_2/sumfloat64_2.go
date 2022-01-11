//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see SumEx1 example from Enumerable.Sum help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sum

type Package struct {
	Company string
	Weight  float64
}

func main() {
	packages := go2linq.NewOnSliceEn(
		Package{Company: "Coho Vineyard", Weight: 25.2},
		Package{Company: "Lucerne Publishing", Weight: 18.7},
		Package{Company: "Wingtip Toys", Weight: 6.0},
		Package{Company: "Adventure Works", Weight: 33.8},
	)
	totalWeight := go2linq.SumMust(packages, func(pkg Package) float64 { return pkg.Weight })
	fmt.Printf("The total weight of the packages is: %.1f\n", totalWeight)
}
