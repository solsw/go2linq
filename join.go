//go:build go1.18

package go2linq

import (
	"sync"
)

// Reimplementing LINQ to Objects: Part 19 – Join
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-19-join/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join

// Join correlates the elements of two sequences based on matching keys.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' must not be based on the same Enumerator, otherwise use JoinSelf instead.
func Join[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, nil)
}

// JoinMust is like Join but panics in case of error.
func JoinMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) Enumerator[Result] {
	r, err := Join(outer, inner, outerKeySelector, innerKeySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// JoinSelf correlates the elements of two sequences based on matching keys.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
func JoinSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	isl := Slice(inner)
	outer.Reset()
	return Join(outer, NewOnSlice(isl...), outerKeySelector, innerKeySelector, resultSelector)
}

// JoinSelfMust is like JoinSelf but panics in case of error.
func JoinSelfMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) Enumerator[Result] {
	r, err := JoinSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// JoinEq correlates the elements of two sequences based on matching keys.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' must not be based on the same Enumerator, otherwise use JoinEqSelf instead.
//
// (The similar to keys equality comparison functionality may be achieved using appropriate key selectors.
// See CustomComparer test for usage of case insensitive string keys.)
func JoinEq[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) (Enumerator[Result], error) {
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
	t := Empty[Inner]()
	var oel Outer
	var iel Inner
	return OnFunc[Result]{
		mvNxt: func() bool {
			once.Do(func() { ilk = ToLookupEqMust(inner, innerKeySelector, eq) })
			for {
				if t.MoveNext() {
					iel = t.Current()
					return true
				}
				if !outer.MoveNext() {
					return false
				}
				oel = outer.Current()
				t = ilk.Item(outerKeySelector(oel))
			}
		},
		crrnt: func() Result { return resultSelector(oel, iel) },
		rst:   func() { t = Empty[Inner](); outer.Reset() },
	},
	nil
}

// JoinEqMust is like JoinEq but panics in case of error.
func JoinEqMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) Enumerator[Result] {
	r, err := JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// JoinEqSelf correlates the elements of two sequences based on matching keys.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
func JoinEqSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) (Enumerator[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	isl := Slice(inner)
	outer.Reset()
	return JoinEq(outer, NewOnSlice(isl...), outerKeySelector, innerKeySelector, resultSelector, eq)
}

// JoinEqSelfMust is like JoinEqSelf but panics in case of error.
func JoinEqSelfMust[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) Enumerator[Result] {
	r, err := JoinEqSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq)
	if err != nil {
		panic(err)
	}
	return r
}
