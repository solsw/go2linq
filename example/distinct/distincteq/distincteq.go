//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq"
)

// see the last two examples from Enumerable.Distinct help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

type Product struct {
	Name string
	Code int
}

func main() {
	products := go2linq.NewOnSliceEn(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
		Product{Name: "Apple", Code: 9},
		Product{Name: "lemon", Code: 12},
	)
	eqf := go2linq.EqualerFunc[Product](func(p1, p2 Product) bool {
		return p1.Code == p2.Code && strings.ToUpper(p1.Name) == strings.ToUpper(p2.Name)
	})
	//Exclude duplicates.
	noduplicates := go2linq.DistinctEqMust(products, go2linq.Equaler[Product](eqf))
	for noduplicates.MoveNext() {
		product := noduplicates.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
}
