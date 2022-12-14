package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Union produces the set union of two slices using 'equaler' to compare values.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If both 'first' and 'second' are nil, nil is returned.
func Union[Source any](first, second []Source, equaler go2linq.Equaler[Source]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return Distinct(append(first, second...), equaler)
}

// UnionCmp produces the set union of two slices using 'comparer' to compare values.
// (See go2linq.DistinctCmp function.)
// If both 'first' and 'second' are nil, nil is returned.
func UnionCmp[Source any](first, second []Source, comparer go2linq.Comparer[Source]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return DistinctCmp(append(first, second...), comparer)
}
