//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq/v2"
)

// see the last two examples from Enumerable.Except help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

type ProductA struct {
	Name string
	Code int
}

func main() {
	fruits1 := go2linq.NewEnSlice(
		ProductA{Name: "apple", Code: 9},
		ProductA{Name: "orange", Code: 4},
		ProductA{Name: "lemon", Code: 12},
	)
	fruits2 := go2linq.NewEnSlice(
		ProductA{Name: "APPLE", Code: 9},
	)
	var equaler go2linq.Equaler[ProductA] = go2linq.EqualerFunc[ProductA](
		func(p1, p2 ProductA) bool {
			return p1.Code == p2.Code && strings.ToUpper(p1.Name) == strings.ToUpper(p2.Name)
		},
	)
	//Get all the elements from the first array except for the elements from the second array.
	except := go2linq.ExceptEqMust(fruits1, fruits2, equaler)
	enr := except.GetEnumerator()
	for enr.MoveNext() {
		product := enr.Current()
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}
}
