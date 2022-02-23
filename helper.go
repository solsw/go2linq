//go:build go1.18

package go2linq

import (
	"fmt"
	"sort"
)

func asStringPrim[T any](t T, isStringer bool) string {
	if isStringer {
		return any(t).(fmt.Stringer).String()
	}
	return fmt.Sprint(t)
}

func typeIsStringer[T any]() bool {
	var i any = ZeroValue[T]()
	_, isStringer := i.(fmt.Stringer)
	return isStringer
}

// elInElelEq determines (using 'equaler') whether 'ee' contains 'el'
func elInElelEq[T any](el T, ee []T, equaler Equaler[T]) bool {
	for _, e := range ee {
		if equaler.Equal(e, el) {
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

func catchErr(panicArg interface{}, err *error) {
	if panicArg == nil {
		return
	}
	if err == nil {
		return
	}
	e, errCatched := panicArg.(error)
	if errCatched {
		*err = e
		return
	}
	panic(panicArg)
}
