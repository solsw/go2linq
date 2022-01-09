//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see the first example from Enumerable.Intersect help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect

func main() {
	id1 := go2linq.NewOnSliceEn(44, 26, 92, 30, 71, 38)
	id2 := go2linq.NewOnSliceEn(39, 59, 83, 47, 26, 4, 30)
	both := go2linq.IntersectMust(id1, id2)
	for both.MoveNext() {
		id := both.Current()
		fmt.Println(id)
	}
}
