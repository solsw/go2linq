//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq/v2"
)

// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/quantifier-operations

type Market struct {
	Name  string
	Items []string
}

func printNames(where go2linq.Enumerable[Market]) {
	names := go2linq.SelectMust(where, func(m Market) string { return m.Name })
	enr := names.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Printf("%s market\n", name)
	}
}

func main() {
	markets := go2linq.NewEnSlice(
		Market{Name: "Emily's", Items: []string{"kiwi", "cheery", "banana"}},
		Market{Name: "Kim's", Items: []string{"melon", "mango", "olive"}},
		Market{Name: "Adam's", Items: []string{"kiwi", "apple", "orange"}},
	)

	whereAll := go2linq.WhereMust(markets, func(m Market) bool {
		items := go2linq.NewEnSlice(m.Items...)
		return go2linq.AllMust(items, func(item string) bool { return len(item) == 5 })
	})
	printNames(whereAll)

	fmt.Println()
	whereAny := go2linq.WhereMust(markets, func(m Market) bool {
		items := go2linq.NewEnSlice(m.Items...)
		return go2linq.AnyPredMust(items, func(item string) bool { return strings.HasPrefix(item, "o") })
	})
	printNames(whereAny)

	fmt.Println()
	whereContains := go2linq.WhereMust(markets, func(m Market) bool {
		items := go2linq.NewEnSlice(m.Items...)
		return go2linq.ContainsMust(items, "kiwi")
	})
	printNames(whereContains)
}
