//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the example from Enumerable.Repeat help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.repeat#examples

func main() {
	strings := go2linq.RepeatMust("I like programming.", 15)
	enr := strings.GetEnumerator()
	for enr.MoveNext() {
		str := enr.Current()
		fmt.Println(str)
	}
}
