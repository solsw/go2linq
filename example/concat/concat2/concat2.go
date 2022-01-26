//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the second example from Enumerable.Concat help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.concat#examples

type Pet struct {
	Name string
	Age  int
}

func main() {
	cats := go2linq.NewEnSlice(
		Pet{Name: "Barley", Age: 8},
		Pet{Name: "Boots", Age: 4},
		Pet{Name: "Whiskers", Age: 1},
	)
	dogs := go2linq.NewEnSlice(
		Pet{Name: "Bounder", Age: 3},
		Pet{Name: "Snoopy", Age: 14},
		Pet{Name: "Fido", Age: 9},
	)
	query := go2linq.SelectManyMust(
		go2linq.NewEnSlice(
			go2linq.SelectMust(cats, func(cat Pet) string { return cat.Name }),
			go2linq.SelectMust(dogs, func(dog Pet) string { return dog.Name }),
		),
		go2linq.Identity[go2linq.Enumerable[string]],
	)
	enr := query.GetEnumerator()
	for enr.MoveNext() {
		name := enr.Current()
		fmt.Println(name)
	}
}
