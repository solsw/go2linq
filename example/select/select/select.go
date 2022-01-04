//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see example from Enumerable.Select help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

func main() {
	squares := go2linq.SelectMust(
		go2linq.RangeMust(1, 10),
		func(x int) int { return x * x },
	)
	for squares.MoveNext() {
		num := squares.Current()
		fmt.Println(num)
	}
}
