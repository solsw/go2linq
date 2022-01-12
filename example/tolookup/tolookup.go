//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see LookupExample example from Lookup Class help
// https://docs.microsoft.com/dotnet/api/system.linq.Lookup-2#examples

type Package struct {
	Company        string
	Weight         float64
	TrackingNumber int64
}

func main() {
	// Create a list of Packages to put into a Lookup data structure.
	packages := go2linq.NewOnSliceEn(
		Package{Company: "Coho Vineyard", Weight: 25.2, TrackingNumber: 89453312},
		Package{Company: "Lucerne Publishing", Weight: 18.7, TrackingNumber: 89112755},
		Package{Company: "Wingtip Toys", Weight: 6.0, TrackingNumber: 299456122},
		Package{Company: "Contoso Pharmaceuticals", Weight: 9.3, TrackingNumber: 670053128},
		Package{Company: "Wide World Importers", Weight: 33.8, TrackingNumber: 4665518773},
	)
	// Create a Lookup to organize the packages.
	// Use the first character of Company as the key value.
	// Select Company appended to TrackingNumber for each element value in the Lookup.
	lookup := go2linq.ToLookupSelMust(packages,
		func(p Package) rune {
			return []rune(p.Company)[0]
		},
		func(p Package) string {
			return fmt.Sprintf("%s %d", p.Company, p.TrackingNumber)
		},
	)

	// Iterate through each Grouping in the Lookup and output the contents.
	enLk := lookup.GetEnumerator()
	for enLk.MoveNext() {
		packageGroup := enLk.Current()
		// Print the key value of the Grouping.
		fmt.Println(string(packageGroup.Key()))
		// Iterate through each value in the Grouping and print its value.
		enGr := packageGroup.GetEnumerator()
		for enGr.MoveNext() {
			str := enGr.Current()
			fmt.Printf("    %s\n", str)
		}
	}

	// Get the number of key-collection pairs in the Lookup.
	count := lookup.Count()
	fmt.Printf("\n%d\n", count)

	// Select a collection of Packages by indexing directly into the Lookup.
	cgroup := lookup.Item('C')
	// Output the results.
	fmt.Println("\nPackages that have a key of 'C'")
	for cgroup.MoveNext() {
		str := cgroup.Current()
		fmt.Println(str)
	}

	// Determine if there is a key with the value 'G' in the Lookup.
	hasG := lookup.Contains('G')
	fmt.Printf("\n%t\n", hasG)
}
