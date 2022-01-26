//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Except help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

func main() {
	numbers1 := go2linq.NewEnSlice(2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5)
	numbers2 := go2linq.NewEnSlice(2.2)
	onlyInFirstSet := go2linq.ExceptMust(numbers1, numbers2)
	enr := onlyInFirstSet.GetEnumerator()
	for enr.MoveNext() {
		number := enr.Current()
		fmt.Println(number)
	}
}
