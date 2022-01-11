//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the example from Enumerable.Reverse help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.reverse#examples

func main() {
	apple := go2linq.NewOnSliceEn("a", "p", "p", "l", "e")
	reversed := go2linq.ReverseMust(apple)
	for reversed.MoveNext() {
		chr := reversed.Current()
		fmt.Printf("%v ", chr)
	}
	fmt.Println()
}
