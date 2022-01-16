//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the second example from Enumerable.Contains help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.contains

type Product struct {
	Name string
	Code int
}

func main() {
	fruits := go2linq.NewOnSliceEn(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
		Product{Name: "lemon", Code: 12},
	)
	apple := Product{Name: "apple", Code: 9}
	kiwi := Product{Name: "kiwi", Code: 8}
	var equaler go2linq.Equaler[Product] = go2linq.EqualerFunc[Product](
		func(p1, p2 Product) bool {
			return p1.Code == p2.Code && p1.Name == p2.Name
		},
	)
	hasApple := go2linq.ContainsEqMust(fruits, apple, equaler)
	hasKiwi := go2linq.ContainsEqMust(fruits, kiwi, equaler)
	fmt.Printf("Apple? %t\n", hasApple)
	fmt.Printf("Kiwi? %t\n", hasKiwi)
}
