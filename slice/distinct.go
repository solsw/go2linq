package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Distinct returns distinct elements from a slice using a specified equaler to compare values.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func Distinct[Source any](source []Source, equaler go2linq.Equaler[Source]) ([]Source, error) {
	return DistinctBy(source, go2linq.Identity[Source], equaler)
}

// DistinctCmp returns distinct elements from a sequence using a specified comparer to compare values.
// (See go2linq.DistinctCmp function.)
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func DistinctCmp[Source any](source []Source, comparer go2linq.Comparer[Source]) ([]Source, error) {
	return DistinctByCmp(source, go2linq.Identity[Source], comparer)
}
