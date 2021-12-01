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
func Intersect[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return IntersectEq(first, second, nil)
}

// IntersectMust is like Intersect but panics in case of error.
func IntersectMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := Intersect(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectSelf produces the set intersection of two sequences by using reflect.DeepEqual as equality comparer.
// 'second' is enumerated immediately.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
func IntersectSelf[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return Intersect(first, NewOnSlice(sl2...))
}

// IntersectSelfMust is like IntersectSelf but panics in case of error.
func IntersectSelfMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := IntersectSelf(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectEq produces the set intersection of two sequences by using the specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use IntersectEqSelf instead.
func IntersectEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if eq == nil {
		eq = EqualerFunc[Source](DeepEqual[Source])
	}
	var once sync.Once
	var dsl2 []Source
	d1 := DistinctEqMust(first, eq)
	var c Source
	return OnFunc[Source]{
		mvNxt: func() bool {
			once.Do(func() { dsl2 = Slice(DistinctEqMust(second, eq)) })
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
	},
	nil
}

// IntersectEqMust is like IntersectEq but panics in case of error.
func IntersectEqMust[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	r, err := IntersectEq(first, second, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectEqSelf produces the set intersection of two sequences by using the specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
func IntersectEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return IntersectEq(first, NewOnSlice(sl2...), eq)
}

// IntersectEqSelfMust is like IntersectEqSelf but panics in case of error.
func IntersectEqSelfMust[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	r, err := IntersectEqSelf(first, second, eq)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectCmp produces the set intersection of two sequences by using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use IntersectCmpSelf instead.
func IntersectCmp[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if cmp == nil {
		return nil, ErrNilComparer
	}
	var once sync.Once
	var dsl2 []Source
	d1 := DistinctCmpMust(first, cmp)
	var c Source
	return OnFunc[Source]{
		mvNxt: func() bool {
			once.Do(func() {
				dsl2 = Slice(DistinctCmpMust(second, cmp))
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
	},
	nil
}

// IntersectCmpMust is like IntersectCmp but panics in case of error.
func IntersectCmpMust[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	r, err := IntersectCmp(first, second, cmp)
	if err != nil {
		panic(err)
	}
	return r
}

// IntersectCmpSelf produces the set intersection of two sequences by using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
func IntersectCmpSelf[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if cmp == nil {
		return nil, ErrNilComparer
	}
	sl2 := Slice(second)
	first.Reset()
	return IntersectCmp(first, NewOnSlice(sl2...), cmp)
}

// IntersectCmpSelfMust is like IntersectCmpSelf but panics in case of error.
func IntersectCmpSelfMust[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	r, err := IntersectCmpSelf(first, second, cmp)
	if err != nil {
		panic(err)
	}
	return r
}
