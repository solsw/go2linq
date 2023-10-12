package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v3"
)

// Distinct returns distinct elements from a slice using a specified equaler to compare values.
// If 'equaler' is nil, [collate.DeepEqualer] is used.
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func Distinct[Source any](source []Source, equaler collate.Equaler[Source]) ([]Source, error) {
	return DistinctBy(source, go2linq.Identity[Source], equaler)
}

// DistinctCmp returns distinct elements from a sequence using a specified comparer to compare values.
// (See [go2linq.DistinctCmp].)
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func DistinctCmp[Source any](source []Source, comparer collate.Comparer[Source]) ([]Source, error) {
	return DistinctByCmp(source, go2linq.Identity[Source], comparer)
}
