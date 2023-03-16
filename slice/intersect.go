package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// Intersect produces the set intersection of two slices using 'equaler' to compare values.
// If 'equaler' is nil [collate.DeepEqualer] is used.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func Intersect[Source any](first, second []Source, equaler collate.Equaler[Source]) ([]Source, error) {
	if first == nil || second == nil {
		return nil, nil
	}
	if len(first) == 0 || len(second) == 0 {
		return []Source{}, nil
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Source]{}
	}
	en, err := go2linq.IntersectEq(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// IntersectCmp produces the set intersection of two slices using a 'comparer' to compare values.
// (See [go2linq.DistinctCmp].)
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func IntersectCmp[Source any](first, second []Source, comparer collate.Comparer[Source]) ([]Source, error) {
	if first == nil || second == nil {
		return nil, nil
	}
	if len(first) == 0 || len(second) == 0 {
		return []Source{}, nil
	}
	en, err := go2linq.IntersectCmp(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), comparer)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
