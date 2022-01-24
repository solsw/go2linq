//go:build go1.18

package main

import (
	"context"
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.Where help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.where

func main() {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}
	query := go2linq.WhereIdxMust(
		go2linq.NewEnSlice(numbers...),
		func(number, index int) bool { return number <= index*10 },
	)
	go2linq.ForEachEn(context.Background(), query,
		func(_ context.Context, number int) error {
			fmt.Println(number)
			return nil
		},
	)
}
