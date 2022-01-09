//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the last example from Enumerable.FirstOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault

func main() {
	names := go2linq.NewOnSliceEn("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	firstLongName := go2linq.FirstOrDefaultPredMust(names, func(name string) bool { return len(name) > 20 })
	fmt.Printf("The first long name is '%v'.\n", firstLongName)

	names.Reset()
	firstVeryLongName := go2linq.FirstOrDefaultPredMust(names, func(name string) bool { return len(name) > 30 })
	var what string
	if firstVeryLongName == "" {
		what = "not a"
	} else {
		what = "a"
	}
	fmt.Printf("There is %v name longer than 30 characters.\n", what)
}
