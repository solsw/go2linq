//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see OrderByEx1 example from Enumerable.OrderBy help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby

type Pet struct {
	Name string
	Age  int
}

func main() {
	pets := go2linq.NewOnSliceEn(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	var ls go2linq.Lesser[Pet] = go2linq.LesserFunc[Pet](func(p1, p2 Pet) bool { return p1.Age < p2.Age })
	query := go2linq.OrderByLsMust(pets, go2linq.Identity[Pet], ls).GetEnumerator()
	for query.MoveNext() {
		pet := query.Current()
		fmt.Printf("%s - %d\n", pet.Name, pet.Age)
	}
}
