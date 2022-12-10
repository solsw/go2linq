package slice

import (
	"github.com/solsw/go2linq/v2"
)

// ExceptBy produces the set difference of two slices according to
// a specified key selector function and using a specified key equaler.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func ExceptBy[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) ([]Source, error) {
	if first == nil {
		return nil, nil
	}
	if len(first) == 0 {
		return []Source{}, nil
	}
	if len(second) == 0 {
		return first, nil
	}
	en, err := go2linq.ExceptByEq(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), keySelector, equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// ExceptByMust is like ExceptBy but panics in case of error.
func ExceptByMust[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, equaler go2linq.Equaler[Key]) []Source {
	r, err := ExceptBy(first, second, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptByCmp produces the set difference of two slices according to a specified key selector function
// and using a specified key comparer. (See go2linq.DistinctCmp function.)
// Order of elements in the result corresponds to the order of elements in 'first'.
// If 'first' is nil, nil is returned.
// If 'first' is empty, new empty slice is returned.
// If 'second' is nil or empty, 'first' is returned.
func ExceptByCmp[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, comparer go2linq.Comparer[Key]) ([]Source, error) {
	if first == nil {
		return nil, nil
	}
	if len(first) == 0 {
		return []Source{}, nil
	}
	if len(second) == 0 {
		return first, nil
	}
	en, err := go2linq.ExceptByCmp(go2linq.NewEnSlice(first...), go2linq.NewEnSlice(second...), keySelector, comparer)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// ExceptByCmpMust is like ExceptByCmp but panics in case of error.
func ExceptByCmpMust[Source, Key any](first []Source, second []Key,
	keySelector func(Source) Key, comparer go2linq.Comparer[Key]) []Source {
	r, err := ExceptByCmp(first, second, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
