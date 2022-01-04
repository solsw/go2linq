//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Where help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func main() {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}
	query := go2linq.WhereIdxMust(
		go2linq.NewOnSliceEn(numbers...),
		func(number, index int) bool { return number <= index*10 },
	)
	for query.MoveNext() {
		number := query.Current()
		fmt.Println(number)
	}
}
