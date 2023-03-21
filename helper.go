package go2linq

import (
	"fmt"
	"sort"

	"github.com/solsw/collate"
)

func asStringPrim[T any](t T, isStringer bool) string {
	if isStringer {
		return any(t).(fmt.Stringer).String()
	}
	return fmt.Sprint(t)
}

// elInElelEq determines (using 'equaler') whether 'ee' contains 'el'
func elInElelEq[T any](el T, ee []T, equaler collate.Equaler[T]) bool {
	for _, e := range ee {
		if equaler.Equal(e, el) {
			return true
		}
	}
	return false
}

// elIdxInElelCmp searches (using 'comparer') for 'el' in sorted 'ee'
// elIdxInElelCmp returns corresponding index (see https://pkg.go.dev/sort#Search)
func elIdxInElelCmp[T any](el T, ee []T, comparer collate.Comparer[T]) int {
	return sort.Search(len(ee), func(i int) bool {
		return comparer.Compare(el, ee[i]) <= 0
	})
}

// elInElelCmp determines (using 'comparer') whether sorted 'ee' contains 'el'
func elInElelCmp[T any](el T, ee []T, comparer collate.Comparer[T]) bool {
	i := elIdxInElelCmp(el, ee, comparer)
	return i < len(ee) && comparer.Compare(el, ee[i]) == 0
}

// elIntoElelAtIdx inserts 'el' into 'ee' at index 'i'
func elIntoElelAtIdx[T any](el T, ee *[]T, i int) {
	*ee = append(*ee, el)
	if i < len(*ee)-1 {
		copy((*ee)[i+1:], (*ee)[i:])
		(*ee)[i] = el
	}
}

// projectionLesser converts [collate.Lesser[Key]] into [collate.Lesser[Source]] using 'sel'
func projectionLesser[Source, Key any](ls collate.Lesser[Key], sel func(Source) Key) collate.Lesser[Source] {
	return collate.LesserFunc[Source](
		func(x, y Source) bool {
			return ls.Less(sel(x), sel(y))
		},
	)
}

// reverseLesser reverses the provided [collate.Lesser]
func reverseLesser[T any](ls collate.Lesser[T]) collate.Lesser[T] {
	return collate.LesserFunc[T](
		func(x, y T) bool {
			return ls.Less(y, x)
		},
	)
}

// compoundLesser combines two [collate.Lesser]s
func compoundLesser[T any](ls1, ls2 collate.Lesser[T]) collate.Lesser[T] {
	return collate.LesserFunc[T](
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
