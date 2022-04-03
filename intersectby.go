//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersectby

// IntersectBy produces the set intersection of two sequences according to a specified key selector function
// and using DeepEqualer as key equaler. 'second' is enumerated on the first MoveNext call.
// Order of elements in the result corresponds to the order of elements in 'first'.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersectby)
func IntersectBy[Source, Key any](first Enumerable[Source], second Enumerable[Key], keySelector func(Source) Key) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return IntersectByEq(first, second, keySelector, nil)
}

// IntersectByMust is like IntersectBy but panics in case of an error.
func IntersectByMust[Source, Key any](first Enumerable[Source], second Enumerable[Key], keySelector func(Source) Key) Enumerable[Source] {
	r, err := IntersectBy(first, second, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

func enrIntersectByEq[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler Equaler[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, Equaler[Source](DeepEqualer[Source]{})).GetEnumerator()
		var once sync.Once
		var dsl2 []Key
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { dsl2 = ToSliceMust(DistinctEqMust(second, equaler)) })
				for enrD1.MoveNext() {
					c = enrD1.Current()
					k := keySelector(c)
					if elInElelEq(k, dsl2, equaler) {
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enrD1.Reset() },
		}
	}
}

// IntersectByEq produces the set intersection of two sequences according to a specified key selector function
// and using a specified key equaler.
// If 'equaler' is nil DeepEqualer is used. 'second' is enumerated on the first MoveNext call.
// Order of elements in the result corresponds to the order of elements in 'first'.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersectby)
func IntersectByEq[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler Equaler[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqualer[Key]{}
	}
	return OnFactory(enrIntersectByEq(first, second, keySelector, equaler)), nil
}

// IntersectByEqMust is like IntersectByEq but panics in case of an error.
func IntersectByEqMust[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler Equaler[Key]) Enumerable[Source] {
	r, err := IntersectByEq(first, second, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

func enrIntersectByCmp[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer Comparer[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, Equaler[Source](DeepEqualer[Source]{})).GetEnumerator()
		var once sync.Once
		var dsl2 []Key
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() {
					dsl2 = ToSliceMust(DistinctCmpMust(second, comparer))
					sort.Slice(dsl2, func(i, j int) bool { return comparer.Compare(dsl2[i], dsl2[j]) < 0 })
				})
				for enrD1.MoveNext() {
					c = enrD1.Current()
					k := keySelector(c)
					if elInElelCmp(k, dsl2, comparer) {
						return true
					}
				}
				return false
			},
			crrnt: func() Source { return c },
			rst:   func() { enrD1.Reset() },
		}
	}
}

// IntersectByCmp produces the set intersection of two sequences according to a specified key selector function
// and using a specified key comparer. (See DistinctCmp function.)
// 'second' is enumerated on the first MoveNext call.
// Order of elements in the result corresponds to the order of elements in 'first'.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersectby)
func IntersectByCmp[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer Comparer[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return OnFactory(enrIntersectByCmp(first, second, keySelector, comparer)), nil
}

// IntersectByCmpMust is like IntersectByCmp but panics in case of an error.
func IntersectByCmpMust[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer Comparer[Key]) Enumerable[Source] {
	r, err := IntersectByCmp(first, second, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
