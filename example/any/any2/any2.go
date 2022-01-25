//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// see AnyEx2 example from Enumerable.Any help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

type (
	Pet struct {
		Name string
		Age  int
	}
	Person struct {
		LastName string
		Pets     []Pet
	}
)

func main() {
	people := []Person{
		Person{
			LastName: "Haas",
			Pets: []Pet{
				Pet{Name: "Barley", Age: 10},
				Pet{Name: "Boots", Age: 14},
				Pet{Name: "Whiskers", Age: 6},
			},
		},
		Person{
			LastName: "Fakhouri",
			Pets: []Pet{
				Pet{Name: "Snowball", Age: 1},
			},
		},
		Person{
			LastName: "Antebi",
			Pets:     []Pet{},
		},
		Person{
			LastName: "Philips",
			Pets: []Pet{
				Pet{Name: "Sweetie", Age: 2},
				Pet{Name: "Rover", Age: 13},
			},
		},
	}
	// Determine which people have a non-empty Pet array.
	names := go2linq.SelectMust(
		go2linq.WhereMust(
			go2linq.NewEnSlice(people...),
			func(person Person) bool { return go2linq.AnyMust(go2linq.NewEnSlice(person.Pets...)) },
		),
		func(person Person) string { return person.LastName },
	)
	enrNames := names.GetEnumerator()
	for enrNames.MoveNext() {
		name := enrNames.Current()
		fmt.Println(name)
	}
}
