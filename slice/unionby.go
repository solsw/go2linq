package slice

import (
	"github.com/solsw/go2linq/v2"
)

// UnionBy produces the set union of two slices according to a specified key selector function
// and using a specified key equaler.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If both 'first' and 'second' are nil, nil is returned.
func UnionBy[Source, Key any](first, second []Source,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return DistinctBy(append(first, second...), keySelector, equaler)
}

// UnionByCmp produces the set union of two slices according to a specified key selector function
// and using a specified key comparer. (See go2linq.DistinctCmp function.)
// If both 'first' and 'second' are nil, nil is returned.
func UnionByCmp[Source, Key any](first, second []Source,
	keySelector func(Source) Key, comparer go2linq.Comparer[Key]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return DistinctByCmp(append(first, second...), keySelector, comparer)
}
