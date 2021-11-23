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
// GroupJoinEq panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func GroupJoinEq[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	if eq == nil {
		eq = EqualerFunc[Key](DeepEqual[Key])
	}
	var once sync.Once
	var ilk *Lookup[Key, Inner]
	return OnFunc[Result]{
		mvNxt: func() bool {
			once.Do(func() { ilk = ToLookupEq(inner, innerKeySelector, eq) })
			return outer.MoveNext()
		},
		crrnt: func() Result {
			c := outer.Current()
			return resultSelector(c, ilk.Item(outerKeySelector(c)))
		},
		rst: func() { outer.Reset() },
	}
}

// GroupJoinEqErr is like GroupJoinEq but returns an error instead of panicking.
func GroupJoinEqErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq), nil
}

// GroupJoinEqSelf correlates the elements of two sequences based on key equality and groups the results.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
// GroupJoinEqSelf panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func GroupJoinEqSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	isl := Slice(inner)
	outer.Reset()
	return GroupJoinEq(outer, NewOnSlice(isl...), outerKeySelector, innerKeySelector, resultSelector, eq)
}

// GroupJoinEqSelfErr is like GroupJoinEqSelf but returns an error instead of panicking.
func GroupJoinEqSelfErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result, eq Equaler[Key]) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupJoinEqSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq), nil
}

// GroupJoin correlates the elements of two sequences based on equality of keys and groups the results.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' must not be based on the same Enumerator, otherwise use GroupJoinSelf instead.
// GroupJoin panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func GroupJoin[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	return GroupJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, nil)
}

// GroupJoinErr is like GroupJoin but returns an error instead of panicking.
func GroupJoinErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupJoin(outer, inner, outerKeySelector, innerKeySelector, resultSelector), nil
}

// GroupJoinSelf correlates the elements of two sequences based on equality of keys and groups the results.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
// GroupJoinSelf panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func GroupJoinSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	isl := Slice(inner)
	outer.Reset()
	return GroupJoin(outer, NewOnSlice(isl...), outerKeySelector, innerKeySelector, resultSelector)
}

// GroupJoinSelfErr is like GroupJoinSelf but returns an error instead of panicking.
func GroupJoinSelfErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Enumerator[Inner]) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchErrPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupJoinSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector), nil
}
