//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 15 – Union
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-15-union/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

// Union produces the set union of two sequences by using reflect.DeepEqual as equality comparer.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionSelf instead.
// Union panics if 'first' or 'second' is nil.
func Union[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	return UnionEq(first, second, nil)
}

// UnionErr is like Union but returns an error instead of panicking.
func UnionErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Union(first, second), nil
}

// UnionSelf produces the set union of two sequences by using reflect.DeepEqual as equality comparer.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// UnionSelf panics if 'first' or 'second' is nil.
func UnionSelf[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return Union(first, NewOnSlice(sl2...))
}

// UnionSelfErr is like UnionSelf but returns an error instead of panicking.
func UnionSelfErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return UnionSelf(first, second), nil
}

// UnionEq produces the set union of two sequences by using a specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionEqSelf instead.
// UnionEq panics if 'first' or 'second' is nil.
func UnionEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	return DistinctEq(Concat(first, second), eq)
}

// UnionEqErr is like UnionEq but returns an error instead of panicking.
func UnionEqErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return UnionEq(first, second, eq), nil
}

// UnionEqSelf produces the set union of two sequences by using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// UnionEqSelf panics if 'first' or 'second' is nil.
func UnionEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionEq(first, NewOnSlice(sl2...), eq)
}

// UnionEqSelfErr is like UnionEqSelf but returns an error instead of panicking.
func UnionEqSelfErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return UnionEqSelf(first, second, eq), nil
}

// UnionCmp produces the set union of two sequences by using a specified Comparer.
// (See DistinctCmp function.)
// 'first' and 'second' must not be based on the same Enumerator, otherwise use UnionCmpSelf instead.
// UnionCmp panics if 'first' or 'second' or 'cmp' is nil.
func UnionCmp[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if cmp == nil {
		panic(ErrNilComparer)
	}
	return DistinctCmp(Concat(first, second), cmp)
}

// UnionCmpErr is like UnionCmp but returns an error instead of panicking.
func UnionCmpErr[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return UnionCmp(first, second, cmp), nil
}

// UnionCmpSelf produces the set union of two sequences by using a specified Comparer.
// (See DistinctCmp function.)
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method. 'second' is enumerated immediately.
// UnionCmp panics if 'first' or 'second' or 'cmp' is nil.
func UnionCmpSelf[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if cmp == nil {
		panic(ErrNilComparer)
	}
	sl2 := Slice(second)
	first.Reset()
	return UnionCmp(first, NewOnSlice(sl2...), cmp)
}

// UnionCmpSelfErr is like UnionCmpSelf but returns an error instead of panicking.
func UnionCmpSelfErr[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return UnionCmpSelf(first, second, cmp), nil
}
