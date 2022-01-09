//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first two examples from Enumerable.LastOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

func main() {
	fruits := go2linq.NewOnSliceEn([]string{}...)
	last := go2linq.LastOrDefaultMust(fruits)
	if last == "" {
		fmt.Println("<string is empty>")
	} else {
		fmt.Println(last)
	}

	daysOfMonth := go2linq.NewOnSliceEn([]int{}...)
	// Setting the default value to 1 after the query.
	lastDay1 := go2linq.LastOrDefaultMust(daysOfMonth)
	if lastDay1 == 0 {
		lastDay1 = 1
	}
	fmt.Printf("The value of the lastDay1 variable is %v\n", lastDay1)

	daysOfMonth.Reset()
	// Setting the default value to 1 by using DefaultIfEmptyDef() in the query.
	lastDay2 := go2linq.LastMust(go2linq.DefaultIfEmptyDefMust(daysOfMonth, 1))
	fmt.Printf("The value of the lastDay2 variable is %d\n", lastDay2)
}
