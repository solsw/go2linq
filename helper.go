package go2linq

import (
	"sort"
)

// elInElelEq determines (using 'equal') whether 'ee' contains 'el'
func elInElelEq[T any](el T, ee []T, equal func(T, T) bool) bool {
	for _, e := range ee {
		if equal(e, el) {
			return true
		}
	}
	return false
}

// elIdxInElelCmp searches (using 'compare') for 'el' in sorted 'ee'
// elIdxInElelCmp returns corresponding index (see https://pkg.go.dev/sort#Search)
func elIdxInElelCmp[T any](el T, ee []T, compare func(T, T) int) int {
	return sort.Search(len(ee), func(i int) bool {
		return compare(el, ee[i]) <= 0
	})
}

// elInElelCmp determines (using 'compare') whether sorted 'ee' contains 'el'
func elInElelCmp[T any](el T, ee []T, compare func(T, T) int) bool {
	i := elIdxInElelCmp(el, ee, compare)
	return i < len(ee) && compare(el, ee[i]) == 0
}

// elIntoElelAtIdx inserts 'el' into 'ee' at index 'i'
func elIntoElelAtIdx[T any](el T, ee *[]T, i int) {
	*ee = append(*ee, el)
	if i < len(*ee)-1 {
		copy((*ee)[i+1:], (*ee)[i:])
		(*ee)[i] = el
	}
}

// ReverseLess reverses the provided 'less'.
func ReverseLess[T any](less func(T, T) bool) func(T, T) bool {
	return func(x, y T) bool {
		return less(y, x)
	}
}

// ThenLess combines two less functions.
// At first 'less1' is applied to 'x' and 'y', if they appear different the result is returned.
// Otherwise, the result of 'less2' is returned.
func ThenLess[T any](less1, less2 func(T, T) bool) func(T, T) bool {
	return func(x, y T) bool {
		if less1(x, y) {
			return true
		}
		if less1(y, x) {
			return false
		}
		return less2(x, y)
	}
}
