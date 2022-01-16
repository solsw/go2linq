//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the last two examples from Enumerable.Union help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

type Product struct {
	Name string
	Code int
}

func main() {
	store1 := go2linq.NewOnSliceEn(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
	)
	store2 := go2linq.NewOnSliceEn(
		Product{Name: "apple", Code: 9},
		Product{Name: "lemon", Code: 12},
	)
	//Get the products from the both arrays excluding duplicates.
	var equaler go2linq.Equaler[Product] = go2linq.EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && p1.Name == p2.Name
		},
	)
	union := go2linq.UnionEqMust(store1, store2, equaler)
	for union.MoveNext() {
		product := union.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
}
