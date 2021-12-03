//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 15 – Union
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-15-union/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

// Union produces the set union of two sequences by using reflect.DeepEqual as equality comparer.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionSelf instead.
func Union[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return UnionEq(first, second, nil)
}

// UnionMust is like Union but panics in case of error.
func UnionMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := Union(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionSelf produces the set union of two sequences by using reflect.DeepEqual as equality comparer.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func UnionSelf[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return Union(first, NewOnSlice(sl2...))
}

// UnionSelfMust is like UnionSelf but panics in case of error.
func UnionSelfMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := UnionSelf(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionEq produces the set union of two sequences by using a specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionEqSelf instead.
func UnionEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return DistinctEq(ConcatMust(first, second), eq)
}

// UnionEqMust is like UnionEq but panics in case of error.
func UnionEqMust[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	r, err := UnionEq(first, second, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionEqSelf produces the set union of two sequences by using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func UnionEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionEq(first, NewOnSlice(sl2...), eq)
}

// UnionEqSelfMust is like UnionEqSelf but panics in case of error.
func UnionEqSelfMust[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	r, err := UnionEqSelf(first, second, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionCmp produces the set union of two sequences by using a specified Comparer.
// (See DistinctCmp function.)
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionCmpSelf instead.
func UnionCmp[Source any](first, second Enumerator[Source], comparer Comparer[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return DistinctCmp(ConcatMust(first, second), comparer)
}

// UnionCmpMust is like UnionCmp but panics in case of error.
func UnionCmpMust[Source any](first, second Enumerator[Source], comparer Comparer[Source]) Enumerator[Source] {
	r, err := UnionCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionCmpSelf produces the set union of two sequences by using a specified Comparer.
// (See DistinctCmp function.)
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
func UnionCmpSelf[Source any](first, second Enumerator[Source], comparer Comparer[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionCmp(first, NewOnSlice(sl2...), comparer)
}

// UnionCmpSelfMust is like UnionCmpSelf but panics in case of error.
func UnionCmpSelfMust[Source any](first, second Enumerator[Source], comparer Comparer[Source]) Enumerator[Source] {
	r, err := UnionCmpSelf(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
