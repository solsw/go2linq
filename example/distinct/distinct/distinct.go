//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the first example from Enumerable.Distinct help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.distinct

func main() {
	ages := go2linq.NewEnSlice(21, 46, 46, 55, 17, 21, 55, 55)
	distinctAges := go2linq.DistinctMust(ages)
	fmt.Println("Distinct ages:")
	enr := distinctAges.GetEnumerator()
	for enr.MoveNext() {
		age := enr.Current()
		fmt.Println(age)
	}
}
