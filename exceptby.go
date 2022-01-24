//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.exceptby

// ExceptBy produces the set difference of two sequences according to a specified key selector function
// and using reflect.DeepEqual as key equaler.
func ExceptBy[Source, Key any](first Enumerable[Source], second Enumerable[Key], keySelector func(Source) Key) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ExceptByEq(first, second, keySelector, nil)
}

// ExceptByMust is like ExceptBy but panics in case of error.
func ExceptByMust[Source, Key any](first Enumerable[Source], second Enumerable[Key], keySelector func(Source) Key) Enumerable[Source] {
	r, err := ExceptBy(first, second, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

func enrExceptByEq[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler Equaler[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, Equaler[Source](DeepEqual[Source]{})).GetEnumerator()
		var once sync.Once
		var dsl2 []Key
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { dsl2 = EnToSlice(DistinctEqMust(second, equaler)) })
				for enrD1.MoveNext() {
					c = enrD1.Current()
					k := keySelector(c)
					if !elInElelEq(k, dsl2, equaler) {
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

// ExceptByEq produces the set difference of two sequences according to a specified key selector function
// and using a specified key equaler. If 'equaler' is nil DeepEqual is used. 'second' is enumerated immediately.
func ExceptByEq[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler Equaler[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqual[Key]{}
	}
	return EnOnFactory(enrExceptByEq(first, second, keySelector, equaler)), nil
}

// ExceptByEqMust is like ExceptByEq but panics in case of error.
func ExceptByEqMust[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler Equaler[Key]) Enumerable[Source] {
	r, err := ExceptByEq(first, second, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

func enrExceptByCmp[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer Comparer[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, Equaler[Source](DeepEqual[Source]{})).GetEnumerator()
		var once sync.Once
		var dsl2 []Key
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() {
					dsl2 = EnToSlice(DistinctCmpMust(second, comparer))
					sort.Slice(dsl2, func(i, j int) bool { return comparer.Compare(dsl2[i], dsl2[j]) < 0 })
				})
				for enrD1.MoveNext() {
					c = enrD1.Current()
					k := keySelector(c)
					if !elInElelCmp(k, dsl2, comparer) {
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

// ExceptByCmp produces the set difference of two sequences according to a specified key selector function
// and using a specified Comparer to compare values.
// If 'equaler' is nil DeepEqual is used. 'second' is enumerated immediately.
func ExceptByCmp[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer Comparer[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return EnOnFactory(enrExceptByCmp(first, second, keySelector, comparer)), nil
}

// ExceptByCmpMust is like ExceptByCmp but panics in case of error.
func ExceptByCmpMust[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer Comparer[Key]) Enumerable[Source] {
	r, err := ExceptByCmp(first, second, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
