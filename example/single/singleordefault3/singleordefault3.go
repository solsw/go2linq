//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the third example from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	pageNumbers := go2linq.NewEnSlice[int]()

	// Setting the default value to 1 after the query.
	pageNumber1 := go2linq.SingleOrDefaultMust(pageNumbers)
	if pageNumber1 == 0 {
		pageNumber1 = 1
	}
	fmt.Printf("The value of the pageNumber1 variable is %d\n", pageNumber1)

	// Setting the default value to 1 by using DefaultIfEmpty() in the query.
	pageNumber2 := go2linq.SingleMust(go2linq.DefaultIfEmptyDefMust(pageNumbers, 1))
	fmt.Printf("The value of the pageNumber2 variable is %d\n", pageNumber2)
}
