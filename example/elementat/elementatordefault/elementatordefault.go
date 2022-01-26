//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.ElementAtOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault#examples

func main() {
	names := go2linq.NewEnSlice("Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu")
	index := 20
	name := go2linq.ElementAtOrDefaultMust(names, index)
	var what string
	if name == "" {
		what = "<no name at this index>"
	} else {
		what = name
	}
	fmt.Printf("The name chosen at index %d is '%s'.\n", index, what)
}
