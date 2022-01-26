//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.Empty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.empty#examples

func main() {
	names1 := []string{"Hartono, Tommy"}
	names2 := []string{"Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	names3 := []string{"Solanki, Ajay", "Hoeing, Helge", "Andersen, Henriette Thaulow", "Potra, Cristina", "Iallo, Lucio"}
	namesList := go2linq.NewEnSlice(
		go2linq.NewEnSlice(names1...),
		go2linq.NewEnSlice(names2...),
		go2linq.NewEnSlice(names3...),
	)
	allNames := go2linq.AggregateSeedMust(namesList,
		go2linq.Empty[string](),
		func(current, next go2linq.Enumerable[string]) go2linq.Enumerable[string] {
			// Only include arrays that have four or more elements
			if go2linq.CountMust(next) > 3 {
				return go2linq.UnionMust(current, next)
			}
			return current
		},
	)
	enr := allNames.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Println(name)
	}
}
