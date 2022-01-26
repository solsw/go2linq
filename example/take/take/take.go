//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Take help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.take

func main() {
	grades := go2linq.NewEnSlice(59, 82, 70, 56, 92, 98, 85)
	topThreeGrades := go2linq.TakeMust[int](
		go2linq.OrderByDescendingLsMust(grades,
			go2linq.Identity[int],
			go2linq.Lesser[int](go2linq.Order[int]{}),
		),
		3,
	)
	fmt.Println("The top three grades are:")
	enr := topThreeGrades.GetEnumerator()
	for enr.MoveNext() {
		grade := enr.Current()
		fmt.Println(grade)
	}
}
