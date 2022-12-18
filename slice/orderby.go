package slice

import (
	"sort"

	"github.com/solsw/go2linq/v2"
	"golang.org/x/exp/constraints"
)

// OrderByKey sorts the elements of a slice in ascending order of keys.
//
// To sort a slice by the values of the elements themselves, specify go2linq.Identity
// function for 'keySelector', also 'Source' must implement constraints.Ordered.
func OrderByKey[Source any, Key constraints.Ordered](source []Source, keySelector func(Source) Key) ([]Source, error) {
	return OrderByKeyLs(source, keySelector, go2linq.Lesser[Key](go2linq.Order[Key]{}))
}

// OrderByKeyLs sorts the elements of a slice in ascending order of keys using a specified lesser.
//
// To sort a slice by the values of the elements themselves, specify go2linq.Identity function for 'keySelector'.
func OrderByKeyLs[Source, Key any](source []Source, keySelector func(Source) Key, lesser go2linq.Lesser[Key]) ([]Source, error) {
	if lesser == nil {
		return nil, go2linq.ErrNilLesser
	}
	sort.SliceStable(source, func(i, j int) bool {
		return lesser.Less(keySelector(source[i]), keySelector(source[j]))
	})
	return source, nil
}

// OrderByDescKey sorts the elements of a slice in descending order of keys.
//
// To sort a slice by the values of the elements themselves, specify go2linq.Identity
// function for 'keySelector', also 'Source' must implement constraints.Ordered.
func OrderByDescKey[Source any, Key constraints.Ordered](source []Source, keySelector func(Source) Key) ([]Source, error) {
	return OrderByDescKeyLs(source, keySelector, go2linq.Lesser[Key](go2linq.Order[Key]{}))
}

// OrderByDescKeyLs sorts the elements of a slice in descending order of keys using a specified lesser.
//
// To sort a slice by the values of the elements themselves, specify go2linq.Identity function for 'keySelector'.
func OrderByDescKeyLs[Source, Key any](source []Source, keySelector func(Source) Key, lesser go2linq.Lesser[Key]) ([]Source, error) {
	if lesser == nil {
		return nil, go2linq.ErrNilLesser
	}
	sort.SliceStable(source, func(i, j int) bool {
		return lesser.Less(keySelector(source[j]), keySelector(source[i]))
	})
	return source, nil
}
