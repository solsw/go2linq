//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.Select help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

type indexstr struct {
	index int
	str   string
}

func main() {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	query := go2linq.SelectIdxMust(
		go2linq.NewEnSlice(fruits...),
		func(fruit string, index int) indexstr {
			return indexstr{index: index, str: fruit[:index]}
		},
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		obj := enr.Current()
		fmt.Printf("%+v\n", obj)
	}
}
