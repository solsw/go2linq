//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the last example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

func main() {
	numbers := go2linq.DefaultIfEmptyMust(go2linq.NewOnSliceEn([]int{}...))
	for numbers.MoveNext() {
		number := numbers.Current()
		fmt.Println(number)
	}
}
