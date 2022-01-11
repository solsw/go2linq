//go:build go1.18

package main

import (
	"fmt"
	"strconv"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Average help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.average

func main() {
	numbers := go2linq.NewOnSliceEn("10007", "37", "299846234235")
	average := go2linq.AverageMust(numbers, func(e string) int { r, _ := strconv.Atoi(e); return r })
	fmt.Printf("The average is %.f.\n", average)
}
