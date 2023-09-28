package go2linq

import (
	"sort"
	"sync"

	"github.com/solsw/collate"
	"github.com/solsw/errorhelper"
)

// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby

// [ExceptBy] produces the set difference of two sequences according to
// a specified key selector function and using [collate.DeepEqualer] as key equaler.
// 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby
func ExceptBy[Source, Key any](first Enumerable[Source], second Enumerable[Key], keySelector func(Source) Key) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ExceptByEq(first, second, keySelector, nil)
}

// ExceptByMust is like [ExceptBy] but panics in case of error.
func ExceptByMust[Source, Key any](first Enumerable[Source], second Enumerable[Key], keySelector func(Source) Key) Enumerable[Source] {
	return errorhelper.Must(ExceptBy(first, second, keySelector))
}

func factoryExceptByEq[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, collate.Equaler[Source](collate.DeepEqualer[Source]{})).GetEnumerator()
		var once sync.Once
		var dsl2 []Key
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { dsl2 = ToSliceMust(DistinctEqMust(second, equaler)) })
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

// [ExceptByEq] produces the set difference of two sequences according to
// a specified key selector function and using a specified key equaler.
// If 'equaler' is nil [collate.DeepEqualer] is used.
// 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby
func ExceptByEq[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	return OnFactory(factoryExceptByEq(first, second, keySelector, equaler)), nil
}

// ExceptByEqMust is like [ExceptByEq] but panics in case of error.
func ExceptByEqMust[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) Enumerable[Source] {
	return errorhelper.Must(ExceptByEq(first, second, keySelector, equaler))
}

func factoryExceptByCmp[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer collate.Comparer[Key]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, collate.Equaler[Source](collate.DeepEqualer[Source]{})).GetEnumerator()
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

// [ExceptByCmp] produces the set difference of two sequences according to
// a specified key selector function and using a specified key comparer. (See [DistinctCmp].)
// 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.exceptby
func ExceptByCmp[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer collate.Comparer[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return OnFactory(factoryExceptByCmp(first, second, keySelector, comparer)), nil
}

// ExceptByCmpMust is like [ExceptByCmp] but panics in case of error.
func ExceptByCmpMust[Source, Key any](first Enumerable[Source], second Enumerable[Key],
	keySelector func(Source) Key, comparer collate.Comparer[Key]) Enumerable[Source] {
	return errorhelper.Must(ExceptByCmp(first, second, keySelector, comparer))
}
