//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Empty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.empty#examples

func main() {
	names1 := []string{"Hartono, Tommy"}
	names2 := []string{"Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	names3 := []string{"Solanki, Ajay", "Hoeing, Helge", "Andersen, Henriette Thaulow", "Potra, Cristina", "Iallo, Lucio"}
	namesList := go2linq.NewOnSliceEn(
		go2linq.NewOnSliceEn(names1...),
		go2linq.NewOnSliceEn(names2...),
		go2linq.NewOnSliceEn(names3...),
	)
	allNames := go2linq.AggregateSeedMust(namesList,
		go2linq.Empty[string](),
		func(current, next go2linq.Enumerator[string]) go2linq.Enumerator[string] {
			// Only include arrays that have four or more elements
			if go2linq.CountMust(next) > 3 {
				return go2linq.UnionMust(current, next)
			}
			return current
		},
	)
	allNames.Reset()
	for allNames.MoveNext() {
		name := allNames.Current()
		fmt.Println(name)
	}
}
