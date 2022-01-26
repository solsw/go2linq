//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the last example from Enumerable.DefaultIfEmpty help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty

func main() {
	numbers := go2linq.DefaultIfEmptyMust(go2linq.NewEnSlice([]int{}...))
	enr := numbers.GetEnumerator()
	for enr.MoveNext() {
		number := enr.Current()
		fmt.Println(number)
	}
}
