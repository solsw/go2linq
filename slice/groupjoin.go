package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// GroupJoin correlates the elements of two slices based on key equality and groups the results.
// 'equaler' is used to compare keys.
// If 'equaler' is nil collate.DeepEqualer is used.
// If 'outer' is nil, nil is returned.
// If 'outer' is empty, new empty slice is returned.
func GroupJoin[Outer, Inner, Key, Result any](outer []Outer, inner []Inner, outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, []Inner) Result, equaler collate.Equaler[Key]) ([]Result, error) {
	if outer == nil {
		return nil, nil
	}
	if len(outer) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.GroupJoinEq(go2linq.NewEnSlice(outer...), go2linq.NewEnSlice(inner...),
		outerKeySelector, innerKeySelector,
		func(o Outer, en go2linq.Enumerable[Inner]) Result { return resultSelector(o, go2linq.ToSliceMust(en)) },
		equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
