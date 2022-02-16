//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the example from Enumerable.Skip help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip#examples

func main() {
	grades := go2linq.NewEnSlice(59, 82, 70, 56, 92, 98, 85)
	lowerGrades := go2linq.SkipMust[int](go2linq.OrderBySelfDescMust(grades), 3)
	fmt.Println("All grades except the top three are:")
	enr := lowerGrades.GetEnumerator()
	for enr.MoveNext() {
		grade := enr.Current()
		fmt.Println(grade)
	}
}
