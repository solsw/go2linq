//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	fruits1 := go2linq.NewEnSlice("orange")
	fruit1 := go2linq.SingleOrDefaultMust(fruits1)
	fmt.Println(fruit1)
}
