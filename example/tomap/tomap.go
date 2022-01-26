//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see ToDictionaryEx1 example from Enumerable.ToDictionary help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.todictionary

type Package struct {
	Company        string
	Weight         float64
	TrackingNumber int64
}

func main() {
	packages := go2linq.NewEnSlice(
		Package{Company: "Coho Vineyard", Weight: 25.2, TrackingNumber: 89453312},
		Package{Company: "Lucerne Publishing", Weight: 18.7, TrackingNumber: 89112755},
		Package{Company: "Wingtip Toys", Weight: 6.0, TrackingNumber: 299456122},
		Package{Company: "Adventure Works", Weight: 33.8, TrackingNumber: 4665518773},
	)
	// Create a Dictionary of Package objects, using TrackingNumber as the key.
	dictionary := go2linq.OnMap(
		go2linq.ToMapMust(packages,
			func(p Package) int64 { return p.TrackingNumber },
		),
	)
	enr := dictionary.GetEnumerator()
	for enr.MoveNext() {
		ke := enr.Current()
		p := ke.Element()
		fmt.Printf("Key %d: %s, %g pounds\n", ke.Key(), p.Company, p.Weight)
	}
}
