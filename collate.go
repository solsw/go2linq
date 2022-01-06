//go:build go1.18

package go2linq

import (
	"constraints"
	"strings"
)

// Equaler defines a function to compare the objects of type T for equality.
type Equaler[T any] interface {
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.iequalitycomparer-1

	// Equal determines whether the specified objects are equal.
	Equal(T, T) bool
}

// EqualerFunc determines whether two objects are equal and implements the Equaler interface.
//
// EqualerFunc is intended for use in functions that accept Equaler as parameter.
// E.g. Having equality function eqf = func(T, T) bool,
// DistinctEqErr may be called in the following way:
//
// var eq Equaler[T] = EqualerFunc[T](eqf)
// DistinctEqErr(source, eq)
type EqualerFunc[T any] func(T, T) bool

// Equal implements the Equaler interface.
func (eqf EqualerFunc[T]) Equal(x, y T) bool {
	return eqf(x, y)
}

// Comparer defines a function to compare two objects of type T.
type Comparer[T any] interface {
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.icomparer-1

	// Compare compares two objects and returns negative if the first one is less than the second,
	// zero if the first one is equal to the second and positive if the first one is greater than the second.
	Compare(T, T) int
}

// ComparerFunc compares two objects and returns negative if the first one is less than the second,
// zero if the first one is equal to the second and positive if the first one is greater than the second.
// ComparerFunc implements the Comparer, Equaler and Lesser interfaces.
//
// ComparerFunc is intended for use in functions that accept Equaler or Comparer as parameter.
// E.g. Having comparison function cmpf = func(T, T) int,
// DistinctCmpErr may be called in the following way:
//
// var cmp Comparer[T] = ComparerFunc[T](cmpf)
// DistinctCmpErr(source, cmp)
type ComparerFunc[T any] func(T, T) int

// Compare implements the Comparer interface.
func (cmpf ComparerFunc[T]) Compare(x, y T) int {
	return cmpf(x, y)
}

// Equal implements the Equaler interface.
func (cmpf ComparerFunc[T]) Equal(x, y T) bool {
	return cmpf(x, y) == 0
}

// Less implements the Lesser interface.
func (cmpf ComparerFunc[T]) Less(x, y T) bool {
	return cmpf(x, y) < 0
}

// Lesser defines a function to compare the objects of type T for equality.
type Lesser[T any] interface {
	// Less determines whether the first object is less than the second.
	Less(T, T) bool
}

// LesserFunc determines whether the first object is less than the second
// and implements the Lesser, Equaler and Comparer interfaces.
//
// LesserFunc is intended for use in functions that accept Equaler or Comparer as parameter.
// E.g. Having less function lsf = func(T, T) bool,
// DistinctCmpErr may be called in the following way:
//
// var cmp Comparer[T] = LesserFunc[T](lsf)
// DistinctCmpErr(source, cmp)
type LesserFunc[T any] func(T, T) bool

// Less implements the Lesser interface.
func (lsf LesserFunc[T]) Less(x, y T) bool {
	return lsf(x, y)
}

// Equal implements the Equaler interface.
func (lsf LesserFunc[T]) Equal(x, y T) bool {
	if lsf(x, y) || lsf(y, x) {
		return false
	}
	return true
}

// Compare implements the Comparer interface.
func (lsf LesserFunc[T]) Compare(x, y T) int {
	if lsf(x, y) {
		return -1
	}
	if lsf(y, x) {
		return +1
	}
	return 0
}

// Orderer implements the Equaler, Comparer and Lesser interfaces for ordered types.
type Orderer[T constraints.Ordered] struct{}

// Equal implements the Equaler interface.
func (Orderer[T]) Equal(x, y T) bool {
	return x == y
}

// Compare implements the Comparer interface.
func (Orderer[T]) Compare(x, y T) int {
	if x < y {
		return -1
	}
	if x > y {
		return +1
	}
	return 0
}

// Less implements the Lesser interface.
func (Orderer[T]) Less(x, y T) bool {
	return x < y
}

var (
	// CaseInsensitiveEqualer is a case insensitive Equaler for string.
	CaseInsensitiveEqualer Equaler[string] = EqualerFunc[string](func(x, y string) bool { return strings.ToLower(x) == strings.ToLower(y) })

	// CaseInsensitiveComparer is a case insensitive Comparer for string.
	CaseInsensitiveComparer Comparer[string] = ComparerFunc[string](func(x, y string) int {
		sx := strings.ToLower(x)
		sy := strings.ToLower(y)
		if sx < sy {
			return -1
		}
		if sx > sy {
			return +1
		}
		return 0
	})

	// CaseInsensitiveLesser is a case insensitive Lesser for string.
	CaseInsensitiveLesser Lesser[string] = LesserFunc[string](func(x, y string) bool {
		return strings.ToLower(x) < strings.ToLower(y)
	})
)
