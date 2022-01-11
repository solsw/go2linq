//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the example from Enumerable.Skip help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skip#examples

func main() {
	grades := go2linq.NewOnSliceEn(59, 82, 70, 56, 92, 98, 85)
	lowerGrades := go2linq.SkipMust(
		go2linq.OrderByDescendingLsMust(grades,
			go2linq.Identity[int],
			go2linq.Lesser[int](go2linq.Order[int]{}),
		).GetEnumerator(),
		3,
	)
	fmt.Println("All grades except the top three are:")
	for lowerGrades.MoveNext() {
		grade := lowerGrades.Current()
		fmt.Println(grade)
	}
}
