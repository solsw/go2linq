//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq"
)

// see AllEx2 example from Enumerable.All help
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all#examples

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
			Pets: []Pet{
				Pet{Name: "Belle", Age: 8},
			},
		},
		Person{
			LastName: "Philips",
			Pets: []Pet{
				Pet{Name: "Sweetie", Age: 2},
				Pet{Name: "Rover", Age: 13},
			},
		},
	}
	// Determine which people have Pets that are all older than 5.
	whereQuery := go2linq.WhereMust(go2linq.NewOnSliceEn(people...),
		func(person Person) bool {
			return go2linq.AllMust(go2linq.NewOnSliceEn(person.Pets...), func(pet Pet) bool { return pet.Age > 5 })
		},
	)
	names := go2linq.SelectMust(whereQuery,
		func(person Person) string { return person.LastName },
	)
	for names.MoveNext() {
		name := names.Current()
		fmt.Println(name)
	}
}
