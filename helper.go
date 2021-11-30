//go:build go1.18

package go2linq

import (
	"errors"
	"sort"
)

func catchErrPanic[T any](panicArg interface{}, res *T, err *error) {
	if panicArg == nil {
		return
	}
	e, panicCatched := panicArg.(error)
	if !panicCatched {
		s, panicCatched := panicArg.(string)
		if panicCatched {
			e = errors.New(s)
		}
	}
	if panicCatched {
		var t0 T
		*res = t0
		*err = e
		return
	}
	panic(panicArg)
}

// elInElelEq determines (using 'eq') whether 'ee' contains 'el'
func elInElelEq[T any](el T, ee []T, eq Equaler[T]) bool {
	for _, e := range ee {
		if eq.Equal(e, el) {
			return true
		}
	}
	return false
}

// elIdxInElelCmp searches (using 'cmp') for 'el' in sorted 'ee'
// elIdxInElelCmp returns corresponding index (see https://pkg.go.dev/sort#Search)
func elIdxInElelCmp[T any](el T, ee []T, cmp Comparer[T]) int {
	return sort.Search(len(ee), func(i int) bool {
		return cmp.Compare(el, ee[i]) <= 0
	})
}

// elInElelCmp determines (using 'cmp') whether sorted 'ee' contains 'el'
func elInElelCmp[T any](el T, ee []T, cmp Comparer[T]) bool {
	i := elIdxInElelCmp(el, ee, cmp)
	return i < len(ee) && cmp.Compare(el, ee[i]) == 0
}

// elIntoElelAtIdx inserts 'el' into 'ee' at index 'i'
func elIntoElelAtIdx[T any](el T, ee *[]T, i int) {
	*ee = append(*ee, el)
	if i < len(*ee)-1 {
		copy((*ee)[i+1:], (*ee)[i:])
		(*ee)[i] = el
	}
}

// projectionLesser converts Lesser[TKey] into Lesser[TSource] using 'sel'
func projectionLesser[TSource, TKey any](ls Lesser[TKey], sel func(TSource) TKey) Lesser[TSource] {
	return LesserFunc[TSource](
		func(x, y TSource) bool {
			return ls.Less(sel(x), sel(y))
		},
	)
}

// reverseLesser reverses the provided Lesser
func reverseLesser[T any](ls Lesser[T]) Lesser[T] {
	return LesserFunc[T](
		func(x, y T) bool {
			return ls.Less(y, x)
		},
	)
}

// compoundLesser combines two Lessers
func compoundLesser[T any](ls1, ls2 Lesser[T]) Lesser[T] {
	return LesserFunc[T](
		func(x, y T) bool {
			if ls1.Less(x, y) {
				return true
			}
			if ls1.Less(y, x) {
				return false
			}
			return ls2.Less(x, y)
		},
	)
}
