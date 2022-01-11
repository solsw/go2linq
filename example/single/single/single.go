//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits1 := go2linq.NewOnSliceEn("orange")
	fruit1 := go2linq.SingleMust(fruits1)
	fmt.Println(fruit1)
}
