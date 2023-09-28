package go2linq

import (
	"github.com/solsw/collate"
	"github.com/solsw/errorhelper"
)

// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby

// [UnionBy] produces the set union of two sequences according to
// a specified key selector function and using [collate.DeepEqualer] as key equaler.
//
// [UnionBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionBy[Source, Key any](first, second Enumerable[Source], keySelector func(Source) Key) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return UnionByEq(first, second, keySelector, nil)
}

// UnionByMust is like [UnionBy] but panics in case of error.
func UnionByMust[Source, Key any](first, second Enumerable[Source], keySelector func(Source) Key) Enumerable[Source] {
	return errorhelper.Must(UnionBy(first, second, keySelector))
}

// [UnionByEq] produces the set union of two sequences according to
// a specified key selector function and using a specified key equaler.
// If 'equaler' is nil [collate.DeepEqualer] is used.
//
// [UnionByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionByEq[Source, Key any](first, second Enumerable[Source],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return DistinctByEq(ConcatMust(first, second), keySelector, equaler)
}

// UnionByEqMust is like [UnionByEq] but panics in case of error.
func UnionByEqMust[Source, Key any](first, second Enumerable[Source],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) Enumerable[Source] {
	return errorhelper.Must(UnionByEq(first, second, keySelector, equaler))
}

// [UnionByCmp] produces the set union of two sequences according to a specified
// key selector function and using a specified key comparer. (See [DistinctCmp].)
//
// [UnionByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionByCmp[Source, Key any](first, second Enumerable[Source],
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
	return DistinctByCmp(ConcatMust(first, second), keySelector, comparer)
}

// UnionByCmpMust is like [UnionByCmp] but panics in case of error.
func UnionByCmpMust[Source, Key any](first, second Enumerable[Source],
	keySelector func(Source) Key, comparer collate.Comparer[Key]) Enumerable[Source] {
	return errorhelper.Must(UnionByCmp(first, second, keySelector, comparer))
}
