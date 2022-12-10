package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Except produces the set difference of two slices using 'equaler' to compare values.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func Except[Source any](first, second []Source, equaler go2linq.Equaler[Source]) ([]Source, error) {
	return ExceptBy(first, second, go2linq.Identity[Source], equaler)
}

// ExceptMust is like Except but panics in case of error.
func ExceptMust[Source any](first, second []Source, equaler go2linq.Equaler[Source]) []Source {
	r, err := Except(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptCmp produces the set difference of two slices using 'comparer' to compare values.
// (See go2linq.DistinctCmp function.)
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func ExceptCmp[Source any](first, second []Source, comparer go2linq.Comparer[Source]) ([]Source, error) {
	return ExceptByCmp(first, second, go2linq.Identity[Source], comparer)
}

// ExceptCmpMust is like ExceptCmp but panics in case of error.
func ExceptCmpMust[Source any](first, second []Source, comparer go2linq.Comparer[Source]) []Source {
	r, err := ExceptCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
