//go:build go1.18

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/solsw/go2linq/v2"
)

// see example from Enumerable.ElementAt help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat#examples

func main() {
	names := []string{"Hartono, Tommy", "Adams, Terry", "Andersen, Henriette Thaulow", "Hedlund, Magnus", "Ito, Shu"}
	namesEn := go2linq.NewEnSlice(names...)
	rand.Seed(time.Now().UnixNano())
	name := go2linq.ElementAtMust(namesEn, rand.Intn(len(names)))
	fmt.Printf("The name chosen at random is '%s'.\n", name)
}
