//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see the third and fourth examples from Enumerable.Single help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single

func main() {
	fruits := go2linq.NewEnSlice("apple", "banana", "mango", "orange", "passionfruit", "grape")
	fruit1 := go2linq.SinglePredMust(fruits, func(fruit string) bool { return len(fruit) > 10 })
	fmt.Println(fruit1)

	fruit2, err := go2linq.SinglePred(fruits, func(fruit string) bool { return len(fruit) > 15 })
	if err == go2linq.ErrNoMatch {
		fmt.Println("The collection does not contain exactly one element whose length is greater than 15.")
	} else {
		fmt.Println(fruit2)
	}

	fruit3, err := go2linq.SinglePred(fruits, func(fruit string) bool { return len(fruit) > 5 })
	if err == go2linq.ErrMultipleMatch {
		fmt.Println("The collection does not contain exactly one element whose length is greater than 5.")
	} else {
		fmt.Println(fruit3)
	}
}
