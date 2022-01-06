//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Distinct help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

func main() {
	ages := go2linq.NewOnSliceEn(21, 46, 46, 55, 17, 21, 55, 55)
	distinctAges := go2linq.DistinctMust(ages)
	fmt.Println("Distinct ages:")
	for distinctAges.MoveNext() {
		age := distinctAges.Current()
		fmt.Println(age)
	}
}
