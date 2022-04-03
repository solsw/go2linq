//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.unionby

// UnionBy produces the set union of two sequences according to a specified key selector function
// and using DeepEqualer as key equaler.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.unionby)
func UnionBy[Source, Key any](first, second Enumerable[Source], keySelector func(Source) Key) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return UnionByEq(first, second, keySelector, nil)
}

// UnionByMust is like UnionBy but panics in case of an error.
func UnionByMust[Source, Key any](first, second Enumerable[Source], keySelector func(Source) Key) Enumerable[Source] {
	r, err := UnionBy(first, second, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionByEq produces the set union of two sequences according to a specified key selector function
// and using a specified key equaler. If 'equaler' is nil DeepEqualer is used.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.unionby)
func UnionByEq[Source, Key any](first, second Enumerable[Source],
	keySelector func(Source) Key, equaler Equaler[Key]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return DistinctByEq(ConcatMust(first, second), keySelector, equaler)
}

// UnionByEqMust is like UnionByEq but panics in case of an error.
func UnionByEqMust[Source, Key any](first, second Enumerable[Source],
	keySelector func(Source) Key, equaler Equaler[Key]) Enumerable[Source] {
	r, err := UnionByEq(first, second, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionByCmp produces the set union of two sequences according to a specified key selector function
// and using a specified key comparer. (See DistinctCmp function.)
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.unionby)
func UnionByCmp[Source, Key any](first, second Enumerable[Source],
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
	return DistinctByCmp(ConcatMust(first, second), keySelector, comparer)
}

// UnionByCmpMust is like UnionByCmp but panics in case of an error.
func UnionByCmpMust[Source, Key any](first, second Enumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) Enumerable[Source] {
	r, err := UnionByCmp(first, second, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
