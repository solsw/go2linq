package go2linq

import (
	"errors"
)

var ErrTestError = errors.New("test error")

type (
	elel[T any] struct {
		e1, e2 T
	}
	elelel[T any] struct {
		e1, e2, e3 T
	}
	elelelel[T any] struct {
		e1, e2, e3, e4 T
	}
)

type (
	Category struct {
		Id           int
		CategoryName string
	}
	Market struct {
		Name  string
		Items []string
	}
	OwnerAndPet struct {
		petOwner PetOwner
		petName  string
	}
	OwnerAndPets struct {
		OwnerName string
		Pets      Enumerable[string]
	}
	OwnerNameAndPetName struct {
		Owner string
		Pet   string
	}
	Package struct {
		Company        string
		Weight         float64
		TrackingNumber int64
	}
	Person struct {
		Name     string
		LastName string
		Pets     []Pet
	}
	Pet struct {
		Name       string
		Age        int
		Vaccinated bool
		Owner      Person
	}
	PetF struct {
		Name string
		Age  float64
	}
	PetOwner struct {
		Name string
		Pets []string
	}
	Product struct {
		Name       string
		Code       int
		CategoryId int
	}
	Result struct {
		Key      float64
		Count    int
		Min, Max float64
	}
)

type PlanetType int

// https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/set-operations#examples

const (
	Rock PlanetType = iota
	Ice
	Gas
	Liquid
)

type Planet struct {
	Name         string
	Type         PlanetType
	OrderFromSun int
}

var (
	Mercury Planet = Planet{Name: "Mercury", Type: Rock, OrderFromSun: 1}
	Venus   Planet = Planet{Name: "Venus", Type: Rock, OrderFromSun: 2}
	Earth   Planet = Planet{Name: "Earth", Type: Rock, OrderFromSun: 3}
	Mars    Planet = Planet{Name: "Mars", Type: Rock, OrderFromSun: 4}
	Jupiter Planet = Planet{Name: "Jupiter", Type: Gas, OrderFromSun: 5}
	Saturn  Planet = Planet{Name: "Saturn", Type: Gas, OrderFromSun: 6}
	Uranus  Planet = Planet{Name: "Uranus", Type: Liquid, OrderFromSun: 7}
	Neptune Planet = Planet{Name: "Neptune", Type: Liquid, OrderFromSun: 8}
	Pluto   Planet = Planet{Name: "Pluto", Type: Ice, OrderFromSun: 9}
)
