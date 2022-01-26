//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the fourth and fifth examples from Enumerable.SingleOrDefault help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

func main() {
	fruits := go2linq.NewEnSlice("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit1 := go2linq.SingleOrDefaultPredMust(fruits, func(fruit string) bool { return len(fruit) > 10 })
	fmt.Println(fruit1)

	fruit2 := go2linq.SingleOrDefaultPredMust(fruits, func(fruit string) bool { return len(fruit) > 15 })
	var what string
	if fruit2 == "" {
		what = "No such string!"
	} else {
		what = fruit2
	}
	fmt.Println(what)
}
