//go:build go1.18

package go2linq

import (
	"strings"
)

// Equaler defines a function to compare the objects of type T for equality.
type Equaler[T any] interface {
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.iequalitycomparer-1

	// Equal determines whether the specified objects are equal.
	Equal(T, T) bool
}

// EqualerFunc determines whether two objects are equal
// and implements Equaler interface.
//
// EqualerFunc is intended for use in functions that accept Equaler as parameter.
// E.g. Having equality function eqf = func(T, T) bool,
// DistinctEq may be called in the following way:
//
// var eq Equaler[T] = EqualerFunc[T](eqf)
// DistinctEq(source, eq)
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
// ComparerFunc implements Comparer, Equaler and Lesser interfaces.
//
// ComparerFunc is intended for use in functions that accept Equaler or Comparer as parameter.
// E.g. Having comparison function cmpf = func(T, T) int,
// DistinctCmp may be called in the following way:
//
// var cmp Comparer[T] = ComparerFunc[T](cmpf)
// DistinctCmp(source, cmp)
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
// and implements Lesser, Equaler and Comparer interfaces.
//
// LesserFunc is intended for use in functions that accept Equaler or Comparer as parameter.
// E.g. Having less function lsf = func(T, T) bool,
// DistinctCmp may be called in the following way:
//
// var cmp Comparer[T] = LesserFunc[T](lsf)
// DistinctCmp(source, cmp)
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

var (
	// IntEqualer is an Equaler for int.
	IntEqualer Equaler[int] =
		EqualerFunc[int](func(x, y int) bool { return x == y })

	// IntComparer is a Comparer for int.
	IntComparer Comparer[int] = ComparerFunc[int](func(x, y int) int {
		if x < y { return -1 }
		if x > y { return +1 }
		return 0
	})

	// IntLesser is a Lesser for int.
	IntLesser Lesser[int] = LesserFunc[int](func(x, y int) bool { return x < y })

	// Float64Equaler is an Equaler for float64.
	Float64Equaler Equaler[float64] =
		EqualerFunc[float64](func(x, y float64) bool { return x == y })

	// Float64Comparer is a Comparer for float64.
	Float64Comparer Comparer[float64] = ComparerFunc[float64](func(x, y float64) int {
		if x < y { return -1 }
		if x > y { return +1 }
		return 0
	})

	// Float64Lesser is a Lesser for float64.
	Float64Lesser Lesser[float64] = LesserFunc[float64](func(x, y float64) bool { return x < y })

	// StringEqualer is an Equaler for string.
	StringEqualer Equaler[string] =
		EqualerFunc[string](func(x, y string) bool { return x == y })

	// StringComparer is a Comparer for string.
	StringComparer Comparer[string] = ComparerFunc[string](func(x, y string) int {
		if x < y { return -1 }
		if x > y { return +1 }
		return 0
	})

	// StringLesser is a Lesser for string.
	StringLesser Lesser[string] = LesserFunc[string](func(x, y string) bool { return x < y })

	// CaseInsensitiveEqualer is a case insensitive Equaler for string.
	CaseInsensitiveEqualer Equaler[string] =
		EqualerFunc[string](func(x, y string) bool { return strings.ToLower(x) == strings.ToLower(y) })

	// CaseInsensitiveComparer is a case insensitive Comparer for string.
	CaseInsensitiveComparer Comparer[string] = ComparerFunc[string](func(x, y string) int {
		sx := strings.ToLower(x)
		sy := strings.ToLower(y)
		if sx < sy { return -1 }
		if sx > sy { return +1 }
		return 0
	})

	// CaseInsensitiveLesser is a case insensitive Lesser for string.
	CaseInsensitiveLesser Lesser[string] = LesserFunc[string](func(x, y string) bool {
		return strings.ToLower(x) < strings.ToLower(y)
	})
)
