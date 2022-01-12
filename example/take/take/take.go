//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Take help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.take

func main() {
	grades := go2linq.NewOnSliceEn(59, 82, 70, 56, 92, 98, 85)
	topThreeGrades := go2linq.TakeMust(
		go2linq.OrderByDescendingLsMust(grades,
			go2linq.Identity[int],
			go2linq.Lesser[int](go2linq.Order[int]{}),
		).GetEnumerator(),
		3,
	)
	fmt.Println("The top three grades are:")
	for topThreeGrades.MoveNext() {
		grade := topThreeGrades.Current()
		fmt.Println(grade)
	}
}
