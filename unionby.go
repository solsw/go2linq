//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.unionby

// UnionBy produces the set union of two sequences according to a specified key selector function
// and using reflect.DeepEqual as a key equaler.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionBySelf instead.
func UnionBy[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return UnionByEq(first, second, keySelector, nil)
}

// UnionByMust is like UnionBy but panics in case of error.
func UnionByMust[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key) Enumerator[Source] {
	r, err := UnionBy(first, second, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionBySelf produces the set union of two sequences according to a specified key selector function
// and using reflect.DeepEqual as a key equaler.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func UnionBySelf[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionBy(first, NewOnSliceEn(sl2...), keySelector)
}

// UnionBySelfMust is like UnionBySelf but panics in case of error.
func UnionBySelfMust[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key) Enumerator[Source] {
	r, err := UnionBySelf(first, second, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionByEq produces the set union of two sequences according to a specified key selector function
// and using a specified key equaler. If 'equaler' is nil reflect.DeepEqual is used.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionByEqSelf instead.
func UnionByEq[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, equaler Equaler[Key]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return DistinctByEq(ConcatMust(first, second), keySelector, equaler)
}

// UnionByEqMust is like UnionByEq but panics in case of error.
func UnionByEqMust[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, equaler Equaler[Key]) Enumerator[Source] {
	r, err := UnionByEq(first, second, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionByEqSelf produces the set union of two sequences according to a specified key selector function
// and using a specified key equaler. If 'equaler' is nil reflect.DeepEqual is used.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func UnionByEqSelf[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, equaler Equaler[Key]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionByEq(first, NewOnSliceEn(sl2...), keySelector, equaler)
}

// UnionByEqSelfMust is like UnionByEqSelf but panics in case of error.
func UnionByEqSelfMust[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, equaler Equaler[Key]) Enumerator[Source] {
	r, err := UnionByEqSelf(first, second, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionByCmp produces the set union of two sequences according to a specified key selector function
// and using a specified Comparer. (See DistinctCmp function.)
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionByCmpSelf instead.
func UnionByCmp[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, comparer Comparer[Key]) (Enumerator[Source], error) {
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

// UnionByCmpMust is like UnionByCmp but panics in case of error.
func UnionByCmpMust[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, comparer Comparer[Key]) Enumerator[Source] {
	r, err := UnionByCmp(first, second, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionByCmpSelf produces the set union of two sequences according to a specified key selector function
// and using a specified Comparer. (See DistinctCmp function.)
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func UnionByCmpSelf[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, comparer Comparer[Key]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionByCmp(first, NewOnSliceEn(sl2...), keySelector, comparer)
}

// UnionByCmpSelfMust is like UnionByCmpSelf but panics in case of error.
func UnionByCmpSelfMust[Source, Key any](first, second Enumerator[Source], keySelector func(Source) Key, comparer Comparer[Key]) Enumerator[Source] {
	r, err := UnionByCmpSelf(first, second, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
