package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// IntersectBy produces the set intersection of two slices according to
// a specified key selector function and using a specified key equaler.
// If 'equaler' is nil go2linq.collate.DeepEqualer is used.
// collate.Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func IntersectBy[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, equaler collate.Equaler[Key]) ([]Source, error) {
	if first == nil || second == nil {
		return nil, nil
	}
	if len(first) == 0 || len(second) == 0 {
		return []Source{}, nil
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	en, err := go2linq.IntersectByEq(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), keySelector, equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// IntersectByCmp produces the set intersection of two slices according to a specified
// key selector function and using a specified key comparer. (See go2linq.DistinctCmp function.)
// collate.Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func IntersectByCmp[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, comparer collate.Comparer[Key]) ([]Source, error) {
	if first == nil || second == nil {
		return nil, nil
	}
	if len(first) == 0 || len(second) == 0 {
		return []Source{}, nil
	}
	en, err := go2linq.IntersectByCmp(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), keySelector, comparer)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
