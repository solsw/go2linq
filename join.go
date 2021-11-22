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
// Join panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func Join[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	return JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, nil)
}

// JoinErr is like Join but returns an error instead of panicking.
func JoinErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return Join(outer, inner, outerKeySelector, innerKeySelector, resultSelector), nil
}

// JoinSelf correlates the elements of two sequences based on matching keys.
// reflect.DeepEqual is used to compare keys. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
// JoinSelf panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func JoinSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	isl := Slice(inner)
	outer.Reset()
	return Join(outer, NewOnSlice(isl...), outerKeySelector, innerKeySelector, resultSelector)
}

// JoinSelfErr is like JoinSelf but returns an error instead of panicking.
func JoinSelfErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return JoinSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector), nil
}

// JoinEq correlates the elements of two sequences based on matching keys.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' must not be based on the same Enumerator, otherwise use JoinEqSelf instead.
// JoinEq panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
//
// (The similar to keys equality comparison functionality may be achieved using appropriate key selectors.
// See CustomComparer test for usage of case insensitive string keys.)
func JoinEq[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
  innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) Enumerator[Result] {
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
	t := Empty[Inner]()
	var oel Outer
	var iel Inner
	return OnFunc[Result]{
		MvNxt: func() bool {
			once.Do(func() { ilk = ToLookupEq(inner, innerKeySelector, eq) })
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
		Crrnt: func() Result { return resultSelector(oel, iel) },
		Rst:   func() { t = Empty[Inner](); outer.Reset() },
	}
}

// JoinEqErr is like JoinEq but returns an error instead of panicking.
func JoinEqErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
  innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq), nil
}

// JoinEqSelf correlates the elements of two sequences based on matching keys.
// A specified Equaler is used to compare keys.
// If 'eq' is nil reflect.DeepEqual is used. 'inner' is enumerated immediately.
// 'outer' and 'inner' may be based on the same Enumerator.
// 'outer' must have real Reset method.
// JoinEqSelf panics if 'outer' or 'inner' or 'outerKeySelector' or 'innerKeySelector' or 'resultSelector' is nil.
func JoinEqSelf[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
  innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) Enumerator[Result] {
	if outer == nil || inner == nil {
		panic(ErrNilSource)
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	isl := Slice(inner)
	outer.Reset()
	return JoinEq(outer, NewOnSlice(isl...), outerKeySelector, innerKeySelector, resultSelector, eq)
}

// JoinEqSelfErr is like JoinEqSelf but returns an error instead of panicking.
func JoinEqSelfErr[Outer, Inner, Key, Result any](outer Enumerator[Outer], inner Enumerator[Inner], outerKeySelector func(Outer) Key,
  innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, eq Equaler[Key]) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return JoinEqSelf(outer, inner, outerKeySelector, innerKeySelector, resultSelector, eq), nil
}
