package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Distinct returns distinct elements from a slice using go2linq.DeepEqualer to compare values.
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func Distinct[Source any](source []Source) ([]Source, error) {
	return DistinctEq(source, nil)
}

// DistinctMust is like Distinct but panics in case of error.
func DistinctMust[Source any](source []Source) []Source {
	r, err := Distinct(source)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctEq returns distinct elements from a slice using a specified equaler to compare values.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func DistinctEq[Source any](source []Source, equaler go2linq.Equaler[Source]) ([]Source, error) {
	return DistinctByEq(source, go2linq.Identity[Source], equaler)
}

// DistinctEqMust is like DistinctEq but panics in case of error.
func DistinctEqMust[Source any](source []Source, equaler go2linq.Equaler[Source]) []Source {
	r, err := DistinctEq(source, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// DistinctCmp returns distinct elements from a sequence using a specified comparer to compare values.
// (See go2linq.DistinctCmp function.)
// Order of elements in the result corresponds to the order of elements in 'source'.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func DistinctCmp[Source any](source []Source, comparer go2linq.Comparer[Source]) ([]Source, error) {
	return DistinctByCmp(source, go2linq.Identity[Source], comparer)
}

// DistinctCmpMust is like DistinctCmp but panics in case of error.
func DistinctCmpMust[Source any](source []Source, comparer go2linq.Comparer[Source]) []Source {
	r, err := DistinctCmp(source, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
