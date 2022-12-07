package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Intersect produces the set intersection of two slices using go2linq.DeepEqualer to compare values.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func Intersect[Source any](first, second []Source) ([]Source, error) {
	return IntersectEq(first, second, nil)
}

// IntersectMust is like Intersect but panics in case of error.
func IntersectMust[Source any](first, second []Source) []Source {
	r, err := Intersect(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectEq produces the set intersection of two slices using 'equaler' to compare values.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func IntersectEq[Source any](first, second []Source, equaler go2linq.Equaler[Source]) ([]Source, error) {
	if first == nil || second == nil {
		return nil, nil
	}
	if len(first) == 0 || len(second) == 0 {
		return []Source{}, nil
	}
	if equaler == nil {
		equaler = go2linq.DeepEqualer[Source]{}
	}
	en, err := go2linq.IntersectEq(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// IntersectEqMust is like IntersectEq but panics in case of error.
func IntersectEqMust[Source any](first, second []Source, equaler go2linq.Equaler[Source]) []Source {
	r, err := IntersectEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectCmp produces the set intersection of two slices using a 'comparer' to compare values.
// (See go2linq.DistinctCmp function.)
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func IntersectCmp[Source any](first, second []Source, comparer go2linq.Comparer[Source]) ([]Source, error) {
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

// IntersectCmpMust is like IntersectCmp but panics in case of error.
func IntersectCmpMust[Source any](first, second []Source, comparer go2linq.Comparer[Source]) []Source {
	r, err := IntersectCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
