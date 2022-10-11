//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// DistinctBy returns distinct elements from a slice according to a specified key selector function
// and using go2linq.DeepEqualer to compare keys.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func DistinctBy[Source, Key any](source []Source, keySelector func(Source) Key) ([]Source, error) {
	return DistinctByEq(source, keySelector, nil)
}

// DistinctByMust is like DistinctBy but panics in case of error.
func DistinctByMust[Source, Key any](source []Source, keySelector func(Source) Key) []Source {
	r, err := DistinctBy(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctByEq returns distinct elements from a slice according to a specified key selector function
// and using a specified equaler to compare keys.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func DistinctByEq[Source, Key any](source []Source, keySelector func(Source) Key, equaler go2linq.Equaler[Key]) ([]Source, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Source{}, nil
	}
	en, err := go2linq.DistinctByEq(go2linq.NewEnSlice(source...), keySelector, equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// DistinctByEqMust is like DistinctByEq but panics in case of error.
func DistinctByEqMust[Source, Key any](source []Source, keySelector func(Source) Key, equaler go2linq.Equaler[Key]) []Source {
	r, err := DistinctByEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctByCmp returns distinct elements from a slice according to a specified key selector function
// and using a specified comparer to compare keys. (See go2linq.DistinctCmp function.)
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func DistinctByCmp[Source, Key any](source []Source, keySelector func(Source) Key, comparer go2linq.Comparer[Key]) ([]Source, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Source{}, nil
	}
	en, err := go2linq.DistinctByCmp(go2linq.NewEnSlice(source...), keySelector, comparer)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// DistinctByCmpMust is like DistinctByCmp but panics in case of error.
func DistinctByCmpMust[Source, Key any](source []Source, keySelector func(Source) Key, comparer go2linq.Comparer[Key]) []Source {
	r, err := DistinctByCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
