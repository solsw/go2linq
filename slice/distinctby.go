package slice

import (
	"github.com/solsw/go2linq/v2"
)

// DistinctBy returns distinct elements from a slice according to a specified key selector function
// and using a specified equaler to compare keys.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func DistinctBy[Source, Key any](source []Source, keySelector func(Source) Key, equaler go2linq.Equaler[Key]) ([]Source, error) {
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

// DistinctByCmp returns distinct elements from a slice according to a specified key selector function
// and using a specified comparer to compare keys. (See go2linq.DistinctCmp function.)
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
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
