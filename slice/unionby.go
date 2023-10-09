package slice

import (
	"github.com/solsw/collate"
)

// UnionBy produces the set union of two slices according to
// a specified key selector function and using a specified key equaler.
//
// If 'equaler' is nil, [collate.DeepEqualer] is used.
// If both 'first' and 'second' are nil, nil is returned.
func UnionBy[Source, Key any](first, second []Source,
	keySelector func(Source) Key, equaler collate.Equaler[Key]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return DistinctBy(append(first, second...), keySelector, equaler)
}

// UnionByCmp produces the set union of two slices according to a specified key selector function
// and using a specified key comparer. (See [go2linq.DistinctCmp].)
//
// If both 'first' and 'second' are nil, nil is returned.
func UnionByCmp[Source, Key any](first, second []Source,
	keySelector func(Source) Key, comparer collate.Comparer[Key]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return DistinctByCmp(append(first, second...), keySelector, comparer)
}
