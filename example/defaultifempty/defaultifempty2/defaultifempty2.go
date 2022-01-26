//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see DefaultIfEmptyEx1 example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

type Pet struct {
	Name string
	Age  int
}

func main() {
	pets := go2linq.NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	enr := go2linq.DefaultIfEmptyMust(pets).GetEnumerator()
	for enr.MoveNext() {
		pet := enr.Current()
		fmt.Println(pet.Name)
	}
}
