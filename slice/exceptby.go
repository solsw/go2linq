package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// ExceptBy produces the set difference of two slices according to
// a specified key selector function and using a specified key equaler.
// If 'equaler' is nil collate.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func ExceptBy[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, equaler collate.Equaler[Key]) ([]Source, error) {
	if first == nil {
		return nil, nil
	}
	if len(first) == 0 {
		return []Source{}, nil
	}
	if len(second) == 0 {
		return first, nil
	}
	en, err := go2linq.ExceptByEq(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), keySelector, equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// ExceptByCmp produces the set difference of two slices according to a specified key selector function
// and using a specified key comparer. (See [go2linq.DistinctCmp].)
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func ExceptByCmp[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, comparer collate.Comparer[Key]) ([]Source, error) {
	if first == nil {
		return nil, nil
	}
	if len(first) == 0 {
		return []Source{}, nil
	}
	if len(second) == 0 {
		return first, nil
	}
	en, err := go2linq.ExceptByCmp(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), keySelector, comparer)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
