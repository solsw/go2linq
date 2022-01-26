//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.Average help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func main() {
	grades := go2linq.NewEnSlice(78, 92, 100, 37, 81)
	average := go2linq.AverageMust(grades, go2linq.Identity[int])
	fmt.Printf("The average grade is %g.\n", average)
}
