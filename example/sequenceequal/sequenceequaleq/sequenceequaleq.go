//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the last example from Enumerable.SequenceEqual help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.sequenceequal

type Product struct {
	Name string
	Code int
}

func main() {
	storeA := go2linq.NewEnSlice(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
	)
	storeB := go2linq.NewEnSlice(
		Product{Name: "apple", Code: 9},
		Product{Name: "orange", Code: 4},
	)
	equalAB := go2linq.SequenceEqualEqMust(storeA, storeB,
		go2linq.Equaler[Product](go2linq.EqualerFunc[Product](
			func(p1, p2 Product) bool {
				return p1.Code == p2.Code && p1.Name == p2.Name
			},
		)),
	)
	fmt.Printf("Equal? %t\n", equalAB)
}
