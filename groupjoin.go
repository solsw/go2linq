package go2linq

import (
	"iter"
	"sync"

	"github.com/solsw/generichelper"
)

// [GroupJoin] correlates the elements of two sequences based on equality of keys and groups the results.
// [generichelper.DeepEqual] is used to compare keys.
// 'inner' is enumerated on the first iteration over the result.
//
// [GroupJoin]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin
func GroupJoin[Outer, Inner, Key, Result any](outer iter.Seq[Outer], inner iter.Seq[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, iter.Seq[Inner]) Result) (iter.Seq[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, generichelper.DeepEqual[Key])
}

// [GroupJoinEq] correlates the elements of two sequences based on key equality and groups the results.
// 'equal' is used to compare keys. 'inner' is enumerated on the first iteration over the result.
//
// [GroupJoinEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin
func GroupJoinEq[Outer, Inner, Key, Result any](outer iter.Seq[Outer], inner iter.Seq[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key,
	resultSelector func(Outer, iter.Seq[Inner]) Result, equal func(Key, Key) bool) (iter.Seq[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	return func(yield func(Result) bool) {
			var once sync.Once
			var ilk *Lookup[Key, Inner]
			for o := range outer {
				once.Do(func() { ilk, _ = ToLookupEq(inner, innerKeySelector, equal) })
				if !yield(resultSelector(o, ilk.Item(outerKeySelector(o)))) {
					return
				}
			}
		},
		nil
}
