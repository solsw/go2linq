//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Average help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func main() {
	fruits := go2linq.NewOnSliceEn("apple", "banana", "mango", "orange", "passionfruit", "grape")
	average := go2linq.AverageMust(fruits, func(e string) int { return len(e) })
	fmt.Printf("The average string length is %g.\n", average)
}
