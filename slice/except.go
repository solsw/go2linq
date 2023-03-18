package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// Except produces the set difference of two slices using 'equaler' to compare values.
//
// If 'equaler' is nil [collate.DeepEqualer] is used.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func Except[Source any](first, second []Source, equaler collate.Equaler[Source]) ([]Source, error) {
	return ExceptBy(first, second, go2linq.Identity[Source], equaler)
}

// ExceptCmp produces the set difference of two slices using 'comparer' to compare values.
// (See [go2linq.DistinctCmp].)
//
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func ExceptCmp[Source any](first, second []Source, comparer collate.Comparer[Source]) ([]Source, error) {
	return ExceptByCmp(first, second, go2linq.Identity[Source], comparer)
}
