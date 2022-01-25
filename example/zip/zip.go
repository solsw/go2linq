//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the example from Enumerable.Zip help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.zip

func main() {
	numbers := go2linq.NewEnSlice(1, 2, 3, 4)
	words := go2linq.NewEnSlice("one", "two", "three")
	numbersAndWords := go2linq.ZipMust(
		numbers,
		words,
		func(first int, second string) string { return fmt.Sprintf("%d %s", first, second) },
	)
	enr := numbersAndWords.GetEnumerator()
	for enr.MoveNext() {
		item := enr.Current()
		fmt.Println(item)
	}
}
