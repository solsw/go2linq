//go:build go1.18

package go2linq

import (
	"sync"
)

// Reimplementing LINQ to Objects: Part 22 – GroupJoin
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-22-groupjoin/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupjoin

// see example/groupjoin

// GroupJoinEq correlates the elements of two sequences based on key equality and groups the results.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' must not be based on the same Enumerator, otherwise use GroupJoinEqSelf instead.
func GroupJoinEq[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	if eq == nil {
		eq = EqualerFunc[Key](DeepEqual[Key])
	}
	var once sync.Once
	var ilk *Lookup[Key, Inner]
	return OnFunc[Result]{
			mvNxt: func() bool {
				once.Do(func() { ilk = ToLookupEqMust(inner, innerKeySelector, eq) })
				return outer.MoveNext()
			},
			crrnt: func() Result {
				c := outer.Current()
				return resultSelector(c, ilk.Item(outerKeySelector(c)))
			},
			rst: func() { outer.Reset() },
		},
		nil
}

// GroupJoinEqMust is like GroupJoinEq but panics in case of error.
func GroupJoinEqMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) Enumerator[Result] {
	r, err := GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoinEqSelf correlates the elements of two sequences based on key equality and groups the results.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
func GroupJoinEqSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	isl := Slice(inner)
	outer.Reset()
	return GroupJoinEq(outer, NewOnSliceEn(isl...), outerKeySelector, innerKeySelector, resultSelector, eq)
}

// GroupJoinEqSelfMust is like GroupJoinEqSelf but panics in case of error.
func GroupJoinEqSelfMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) Enumerator[Result] {
	r, err := GroupJoinEqSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoin correlates the elements of two sequences based on equality of keys and groups the results.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' must not be based on the same Enumerator, otherwise use GroupJoinSelf instead.
func GroupJoin[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, nil)
}

// GroupJoinMust is like GroupJoin but panics in case of error.
func GroupJoinMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) Enumerator[Result] {
	r, err := GroupJoin(outer, inner, outerKeySelector, innerKeySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupJoinSelf correlates the elements of two sequences based on equality of keys and groups the results.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
func GroupJoinSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	isl := Slice(inner)
	outer.Reset()
	return GroupJoin(outer, NewOnSliceEn(isl...), outerKeySelector, innerKeySelector, resultSelector)
}

// GroupJoinSelfMust is like GroupJoinSelf but panics in case of error.
func GroupJoinSelfMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) Enumerator[Result] {
	r, err := GroupJoinSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
