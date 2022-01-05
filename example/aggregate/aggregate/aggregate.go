//go:build go1.18

package main

import (
	"fmt"
	"strings"

	"github.com/solsw/go2linq"
)

// see the last example from Enumerable.Aggregate help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

func main() {
	sentence := "the quick brown fox jumps over the lazy dog"
	// Split the string into individual words.
	words := go2linq.NewOnSliceEn(strings.Fields(sentence)...)
	// Prepend each word to the beginning of the new sentence to reverse the word order.
	reversed := go2linq.AggregateMust(words,
		func(workingSentence, next string) string {
			return next + " " + workingSentence
		},
	)
	fmt.Println(reversed)
}
