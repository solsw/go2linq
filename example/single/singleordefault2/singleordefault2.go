//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the second example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	fruits2 := go2linq.NewOnSliceEn[string]()
	fruit2 := go2linq.SingleOrDefaultMust(fruits2)
	var what string
	if fruit2 == "" {
		what = "No such string!"
	} else {
		what = fruit2
	}
	fmt.Println(what)
}
