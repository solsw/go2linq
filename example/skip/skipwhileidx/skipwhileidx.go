//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

func main() {
	amounts := go2linq.NewEnSlice(5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500)
	query := go2linq.SkipWhileIdxMust(amounts, func(amount, index int) bool { return amount > index*1000 })
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		amount := enr.Current()
		fmt.Println(amount)
	}
}
