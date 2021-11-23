//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// Reimplementing LINQ to Objects: Part 16 – IntersectErr (and build fiddling)
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-16-intersect-and-build-fiddling/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect

// Intersect produces the set intersection of two sequences by using reflect.DeepEqual as equality comparer.
// 'second' is enumerated immediately.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use IntersectSelf instead.
// Intersect panics if 'first' or 'second' is nil.
func Intersect[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	return IntersectEq(first, second, nil)
}

// IntersectErr is like Intersect but returns an error instead of panicking.
func IntersectErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Intersect(first, second), nil
}

// IntersectSelf produces the set intersection of two sequences by using reflect.DeepEqual as equality comparer.
// 'second' is enumerated immediately.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
// IntersectSelf panics if 'first' or 'second' is nil.
func IntersectSelf[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return Intersect(first, NewOnSlice(sl2...))
}

// IntersectSelfErr is like IntersectSelf but returns an error instead of panicking.
func IntersectSelfErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return IntersectSelf(first, second), nil
}

// IntersectEq produces the set intersection of two sequences by using the specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use IntersectEqSelf instead.
// IntersectEq panics if 'first' or 'second' is nil.
func IntersectEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	var once sync.Once
	var dsl2 []Source
	d1 := DistinctEq(first, eq)
	var c Source
	return OnFunc[Source]{
		mvNxt: func() bool {
			once.Do(func() { dsl2 = Slice(DistinctEq(second, eq)) })
			for d1.MoveNext() {
				c = d1.Current()
				if elInElelEq(c, dsl2, eq) {
					return true
				}
			}
			return false
		},
		crrnt: func() Source { return c },
		rst:   func() { d1.Reset() },
	}
}

// IntersectEqErr is like IntersectEq but returns an error instead of panicking.
func IntersectEqErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return IntersectEq(first, second, eq), nil
}

// IntersectEqSelf produces the set intersection of two sequences by using the specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
// IntersectEqSelf panics if 'first' or 'second' is nil.
func IntersectEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return IntersectEq(first, NewOnSlice(sl2...), eq)
}

// IntersectEqSelfErr is like IntersectEqSelf but returns an error instead of panicking.
func IntersectEqSelfErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return IntersectEqSelf(first, second, eq), nil
}

// IntersectCmp produces the set intersection of two sequences by using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use IntersectCmpSelf instead.
// IntersectCmp panics if 'first' or 'second' or 'cmp' is nil.
func IntersectCmp[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if cmp == nil {
		panic(ErrNilComparer)
	}
	var once sync.Once
	var dsl2 []Source
	d1 := DistinctCmp(first, cmp)
	var c Source
	return OnFunc[Source]{
		mvNxt: func() bool {
			once.Do(func() {
				dsl2 = Slice(DistinctCmp(second, cmp))
				sort.Slice(dsl2, func(i, j int) bool { return cmp.Compare(dsl2[i], dsl2[j]) < 0 })
			})
			for d1.MoveNext() {
				c = d1.Current()
				if elInElelCmp(c, dsl2, cmp) {
					return true
				}
			}
			return false
		},
		crrnt: func() Source { return c },
		rst:   func() { d1.Reset() },
	}
}

// IntersectCmpErr is like IntersectCmp but returns an error instead of panicking.
func IntersectCmpErr[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return IntersectCmp(first, second, cmp), nil
}

// IntersectCmpSelf produces the set intersection of two sequences by using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
// IntersectCmpSelf panics if 'first' or 'second' or 'cmp' is nil.
func IntersectCmpSelf[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if cmp == nil {
		panic(ErrNilComparer)
	}
	sl2 := Slice(second)
	first.Reset()
	return IntersectCmp(first, NewOnSlice(sl2...), cmp)
}

// IntersectCmpSelfErr is like IntersectCmpSelf but returns an error instead of panicking.
func IntersectCmpSelfErr[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchErrPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return IntersectCmpSelf(first, second, cmp), nil
}
