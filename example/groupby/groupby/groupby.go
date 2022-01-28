//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/grouping-data#query-expression-syntax-example

func main() {
	numbers := go2linq.NewEnSlice(35, 44, 200, 84, 3987, 4, 199, 329, 446, 208)
	query := go2linq.GroupByMust(numbers, func(i int) int { return i % 2 })
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		group := enr.Current()
		if group.Key() == 0 {
			fmt.Println("\nEven numbers:")
		} else {
			fmt.Println("\nOdd numbers:")
		}
		enrGroup := group.GetEnumerator()
		for enrGroup.MoveNext() {
			i := enrGroup.Current()
			fmt.Println(i)
		}
	}
}
