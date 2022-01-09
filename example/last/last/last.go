//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Last help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last

func main() {
	numbers := go2linq.NewOnSliceEn(9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 67, 12, 19)
	last := go2linq.LastMust(numbers)
	fmt.Println(last)
}
