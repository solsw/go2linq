//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first two examples from Enumerable.FirstOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault

func main() {
	numbers := go2linq.NewOnSliceEn([]int{}...)
	first := go2linq.FirstOrDefaultMust(numbers)
	fmt.Println(first)

	months := go2linq.NewOnSliceEn([]int{}...)
	// Setting the default value to 1 after the query.
	firstMonth1 := go2linq.FirstOrDefaultMust(months)
	if firstMonth1 == 0 {
		firstMonth1 = 1
	}
	fmt.Printf("The value of the firstMonth1 variable is %v\n", firstMonth1)

	months.Reset()
	// Setting the default value to 1 by using DefaultIfEmptyDef() in the query.
	firstMonth2 := go2linq.FirstMust(go2linq.DefaultIfEmptyDefMust(months, 1))
	fmt.Printf("The value of the firstMonth2 variable is %v\n", firstMonth2)
}
