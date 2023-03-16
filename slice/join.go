package slice

import (
	"github.com/solsw/collate"
	"github.com/solsw/go2linq/v2"
)

// Join correlates the elements of two slices based on matching keys.
// 'equaler' is used to compare keys.
// If 'equaler' is nil collate.DeepEqualer is used.
// If 'outer' or 'inner' is nil, nil is returned.
// If 'outer' or 'inner' is empty, new empty slice is returned.
func Join[Outer, Inner, Key, Result any](outer []Outer, inner []Inner, outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, equaler collate.Equaler[Key]) ([]Result, error) {
	if outer == nil || inner == nil {
		return nil, nil
	}
	if len(outer) == 0 || len(inner) == 0 {
		return []Result{}, nil
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	en, err := go2linq.JoinEq(go2linq.NewEnSlice(outer...), go2linq.NewEnSlice(inner...),
		outerKeySelector, innerKeySelector, resultSelector, equaler)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
