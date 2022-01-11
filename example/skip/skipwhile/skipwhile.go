//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the second example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

func main() {
	grades := go2linq.NewOnSliceEn(59, 82, 70, 56, 92, 98, 85)
	lowerGrades := go2linq.SkipWhileMust(
		go2linq.OrderByDescendingLsMust(grades,
			go2linq.Identity[int],
			go2linq.Lesser[int](go2linq.Order[int]{}),
		).GetEnumerator(),
		func(grade int) bool { return grade >= 80 },
	)
	fmt.Println("All grades below 80:")
	for lowerGrades.MoveNext() {
		grade := lowerGrades.Current()
		fmt.Println(grade)
	}
}
