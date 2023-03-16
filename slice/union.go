package slice

import (
	"github.com/solsw/collate"
)

// Union produces the set union of two slices using 'equaler' to compare values.
// If 'equaler' is nil collate.DeepEqualer is used.
// If both 'first' and 'second' are nil, nil is returned.
func Union[Source any](first, second []Source, equaler collate.Equaler[Source]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return Distinct(append(first, second...), equaler)
}

// UnionCmp produces the set union of two slices using 'comparer' to compare values.
// (See go2linq.DistinctCmp function.)
// If both 'first' and 'second' are nil, nil is returned.
func UnionCmp[Source any](first, second []Source, comparer collate.Comparer[Source]) ([]Source, error) {
	if first == nil && second == nil {
		return nil, nil
	}
	return DistinctCmp(append(first, second...), comparer)
}
