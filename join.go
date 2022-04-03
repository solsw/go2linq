//go:build go1.18

package go2linq

import (
	"sync"
)

// Reimplementing LINQ to Objects: Part 19 – Join
// https://codeblog.jonskeet.uk/2010/12/31/reimplementing-linq-to-objects-part-19-join/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join

// Join correlates the elements of two sequences based on matching keys.
// DeepEqualer is used to compare keys. 'inner' is enumerated on the first MoveNext call.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join)
func Join[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) (Enumerable[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, nil)
}

// JoinMust is like Join but panics in case of an error.
func JoinMust[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner],
	outerKeySelector func(Outer) Key, innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result) Enumerable[Result] {
	r, err := Join(outer, inner, outerKeySelector, innerKeySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

func enrJoinEq[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, equaler Equaler[Key]) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enrO := outer.GetEnumerator()
		enrT := Empty[Inner]().GetEnumerator()
		var once sync.Once
		var ilk *Lookup[Key, Inner]
		var oel Outer
		var iel Inner
		return enrFunc[Result]{
			mvNxt: func() bool {
				once.Do(func() { ilk = ToLookupEqMust(inner, innerKeySelector, equaler) })
				for {
					if enrT.MoveNext() {
						iel = enrT.Current()
						return true
					}
					if !enrO.MoveNext() {
						return false
					}
					oel = enrO.Current()
					enrT = ilk.Item(outerKeySelector(oel)).GetEnumerator()
				}
			},
			crrnt: func() Result { return resultSelector(oel, iel) },
			rst: func() {
				enrT = Empty[Inner]().GetEnumerator()
				enrO.Reset()
			},
		}
	}
}

// JoinEq correlates the elements of two sequences based on matching keys.
// A specified Equaler is used to compare keys.
// If 'equaler' is nil DeepEqualer is used. 'inner' is enumerated on the first MoveNext call.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.join)
//
// Similar to the keys equality functionality may be achieved using appropriate key selectors.
// See CustomComparer test for usage of case insensitive string keys.
func JoinEq[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, equaler Equaler[Key]) (Enumerable[Result], error) {
	if outer == nil || inner == nil {
		return nil, ErrNilSource
	}
	if outerKeySelector == nil || innerKeySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqualer[Key]{}
	}
	return OnFactory(enrJoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, equaler)), nil
}

// JoinEqMust is like JoinEq but panics in case of an error.
func JoinEqMust[Outer, Inner, Key, Result any](outer Enumerable[Outer], inner Enumerable[Inner], outerKeySelector func(Outer) Key,
	innerKeySelector func(Inner) Key, resultSelector func(Outer, Inner) Result, equaler Equaler[Key]) Enumerable[Result] {
	r, err := JoinEq(outer, inner, outerKeySelector, innerKeySelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
