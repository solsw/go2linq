//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the example from Enumerable.Reverse help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.reverse#examples

func main() {
	apple := go2linq.NewEnSlice("a", "p", "p", "l", "e")
	reversed := go2linq.ReverseMust(apple)
	enr := reversed.GetEnumerator()
	for enr.MoveNext() {
		chr := enr.Current()
		fmt.Printf("%v ", chr)
	}
	fmt.Println()
}
