//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the second and third examples from Enumerable.Intersect help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect

type ProductA struct {
	Name string
	Code int
}

func main() {
	store1 := go2linq.NewOnSliceEn(ProductA{Name: "apple", Code: 9}, ProductA{Name: "orange", Code: 4})
	store2 := go2linq.NewOnSliceEn(ProductA{Name: "apple", Code: 9}, ProductA{Name: "lemon", Code: 12})
	// Get the products from the first array that have duplicates in the second array.
	var equaler go2linq.Equaler[ProductA] = go2linq.EqualerFunc[ProductA](
		func(p1, p2 ProductA) bool {
			return p1.Name == p2.Name && p1.Code == p2.Code
		},
	)
	duplicates := go2linq.IntersectEqMust(store1, store2, equaler)
	for duplicates.MoveNext() {
		product := duplicates.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
}
