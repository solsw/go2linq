//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the second example from Enumerable.SkipWhile help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile

func main() {
	grades := go2linq.NewEnSlice(59, 82, 70, 56, 92, 98, 85)
	lowerGrades := go2linq.SkipWhileMust[int](
		go2linq.OrderByDescendingSelfMust(grades),
		func(grade int) bool { return grade >= 80 },
	)
	fmt.Println("All grades below 80:")
	enr := lowerGrades.GetEnumerator()
	for enr.MoveNext() {
		grade := enr.Current()
		fmt.Println(grade)
	}
}
