//go:build go1.18

package go2linq

import (
	"sync"
)

// Reimplementing LINQ to Objects: Part 22 – GroupJoin
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-22-groupjoin/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin

// see example/groupjoin

func enrGroupJoinEq[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerable[Inner]) Result, equaler Equaler[Key]) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enr := outer.GetEnumerator()
		var once sync.Once
		var ilk *Lookup[Key, Inner]
		return enrFunc[Result]{
			mvNxt: func() bool {
				once.Do(func() { ilk = ToLookupEqMust(inner, innerKeySelector, equaler) })
				return enr.MoveNext()
			},
			crrnt: func() Result {
				c := enr.Current()
				return resultSelector(c, ilk.Item(outerKeySelector(c)))
			},
			rst: func() { enr.Reset() },
		}
	}
}

// GroupJoinEq correlates the elements of two sequences based on key equality and groups the results.
// A specified Equaler is used to compare keys.
// If 'equaler' is nil DeepEqual is used. 'inner' is enumerated immediately.
func GroupJoinEq[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerable[Inner]) Result, equaler Equaler[Key]) (Enumerable[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqual[Key]{}
	}
	return OnFactory(enrGroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, equaler)), nil
}

// GroupJoinEqMust is like GroupJoinEq but panics in case of error.
func GroupJoinEqMust[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerable[Inner]) Result, equaler Equaler[Key]) Enumerable[Result] {
	r, err := GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoin correlates the elements of two sequences based on equality of keys and groups the results.
// DeepEqual is used to compare keys. 'inner' is enumerated immediately.
func GroupJoin[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerable[Inner]) Result) (Enumerable[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, nil)
}

// GroupJoinMust is like GroupJoin but panics in case of error.
func GroupJoinMust[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerable[Inner]) Result) Enumerable[Result] {
	r, err := GroupJoin(outer, inner, outerKeySelector, innerKeySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
