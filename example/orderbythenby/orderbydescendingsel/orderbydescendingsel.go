//go:build go1.18

package main

import (
	"fmt"
	"math"

	"github.com/solsw/go2linq/v2"
)

// see OrderByDescendingEx1 example from Enumerable.OrderByDescending help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending

func main() {
	decimals := go2linq.NewEnSlice(6.2, 8.3, 0.5, 1.3, 6.3, 9.7)
	var ls go2linq.Lesser[float64] = go2linq.LesserFunc[float64](
		func(f1, f2 float64) bool {
			_, fr1 := math.Modf(f1)
			_, fr2 := math.Modf(f2)
			if math.Abs(fr1-fr2) < 0.001 {
				return f1 < f2
			}
			return fr1 < fr2
		},
	)
	query := go2linq.OrderBySelfDescLsMust(decimals, ls)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		num := enr.Current()
		fmt.Println(num)
	}
}
