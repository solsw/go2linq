//go:build go1.18

package go2linq

import (
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"
)

// Equaler defines a function to compare the objects of type T for equality.
type Equaler[T any] interface {
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.iequalitycomparer-1

	// Equal determines whether the specified objects are equal.
	Equal(T, T) bool
}

// EqualerFunc determines whether the specified objects are equal and implements the Equaler interface.
// EqualerFunc is intended for use in the functions that accept Equaler as a parameter.
//
// E.g. Having equality function:
//
//	var eqf func(T, T) bool
//
// DistinctEq may be called in the following way:
//
//	var equaler Equaler[T] = EqualerFunc[T](eqf)
//	DistinctEq(source, equaler)
type EqualerFunc[T any] func(T, T) bool

// Equal implements the Equaler interface.
func (eqf EqualerFunc[T]) Equal(x, y T) bool {
	return eqf(x, y)
}

// Lesser defines a function to compare the specified objects of type T.
type Lesser[T any] interface {

	// Less determines whether the first object is less than the second.
	Less(T, T) bool
}

// LesserFunc determines whether the first specified object is less than the second
// and implements the Equaler, Lesser and Comparer interfaces.
// LesserFunc is intended for use in the functions that accept Equaler, Lesser or Comparer as a parameter.
//
// E.g. Having less function:
//
//	var lsf = func(T, T) bool
//
// DistinctCmp may be called in the following way:
//
//	var cmp Comparer[T] = LesserFunc[T](lsf)
//	DistinctCmp(source, cmp)
type LesserFunc[T any] func(T, T) bool

// Equal implements the Equaler interface.
func (lsf LesserFunc[T]) Equal(x, y T) bool {
	return lsf.Compare(x, y) == 0
}

// Less implements the Lesser interface.
func (lsf LesserFunc[T]) Less(x, y T) bool {
	return lsf(x, y)
}

// Compare implements the Comparer interface.
func (lsf LesserFunc[T]) Compare(x, y T) int {
	if lsf.Less(x, y) {
		return -1
	}
	if lsf.Less(y, x) {
		return +1
	}
	return 0
}

// Comparer defines a function to compare the specified objects of type T.
type Comparer[T any] interface {
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.icomparer-1

	// Compare compares the specified objects and returns negative if the first one is less than the second,
	// zero if the first one is equal to the second and positive if the first one is greater than the second.
	Compare(T, T) int
}

// ComparerFunc compares the specified objects and returns negative if the first one is less than the second,
// zero if the first one is equal to the second and positive if the first one is greater than the second.
// ComparerFunc implements the Equaler, Lesser and Comparer interfaces.
// ComparerFunc is intended for use in the functions that accept Equaler, Lesser or Comparer as a parameter.
//
// E.g. Having comparison function:
//
//	var cmpf = func(T, T) int
//
// DistinctCmp may be called in the following way:
//
//	var cmp Comparer[T] = ComparerFunc[T](cmpf)
//	DistinctCmp(source, cmp)
type ComparerFunc[T any] func(T, T) int

// Equal implements the Equaler interface.
func (cmpf ComparerFunc[T]) Equal(x, y T) bool {
	return cmpf.Compare(x, y) == 0
}

// Less implements the Lesser interface.
func (cmpf ComparerFunc[T]) Less(x, y T) bool {
	return cmpf.Compare(x, y) < 0
}

// Compare implements the Comparer interface.
func (cmpf ComparerFunc[T]) Compare(x, y T) int {
	return cmpf(x, y)
}

// Order implements the Equaler, Comparer and Lesser interfaces for ordered types.
type Order[T constraints.Ordered] struct{}

// Equal implements the Equaler interface.
func (Order[T]) Equal(x, y T) bool {
	return x == y
}

// Less implements the Lesser interface.
func (Order[T]) Less(x, y T) bool {
	return x < y
}

// Compare implements the Comparer interface.
func (Order[T]) Compare(x, y T) int {
	if x < y {
		return -1
	}
	if x > y {
		return +1
	}
	return 0
}

// DeepEqualer is an Equaler implementation that is a generic wrapper around reflect.DeepEqual.
type DeepEqualer[T any] struct{}

func (DeepEqualer[T]) Equal(x, y T) bool {
	return reflect.DeepEqual(x, y)
}

var (
	// BoolEqualer is an Equaler for bool.
	BoolEqualer Equaler[bool] = EqualerFunc[bool](func(x, y bool) bool { return x == y })

	boolLesserFunc = LesserFunc[bool](func(x, y bool) bool { return !x && y })

	// BoolLesser is a Lesser for bool.
	BoolLesser Lesser[bool] = boolLesserFunc

	// BoolComparer is a Comparer for bool.
	BoolComparer Comparer[bool] = boolLesserFunc

	// CaseInsensitiveEqualer is a case insensitive Equaler for string.
	CaseInsensitiveEqualer Equaler[string] = EqualerFunc[string](func(x, y string) bool {
		// strings.EqualFold(x, y) not used to comply with CaseInsensitiveLesser and CaseInsensitiveComparer
		return strings.ToLower(x) == strings.ToLower(y)
	})

	// CaseInsensitiveLesser is a case insensitive Lesser for string.
	CaseInsensitiveLesser Lesser[string] = LesserFunc[string](func(x, y string) bool {
		return strings.ToLower(x) < strings.ToLower(y)
	})

	// CaseInsensitiveComparer is a case insensitive Comparer for string.
	CaseInsensitiveComparer Comparer[string] = ComparerFunc[string](func(x, y string) int {
		lx := strings.ToLower(x)
		ly := strings.ToLower(y)
		if lx < ly {
			return -1
		}
		if lx > ly {
			return +1
		}
		return 0
	})
)
