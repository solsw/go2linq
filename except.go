//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// Reimplementing LINQ to Objects: Part 17 – Except
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-17-except/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

// Except produces the set difference of two sequences by using reflect.DeepEqual to compare values.
// 'second' is enumerated immediately.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ExceptSelf instead.
// Except panics if 'first' or 'second' is nil.
func Except[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	return ExceptEq(first, second, nil)
}

// ExceptErr is like UnionEqSelf but returns an error instead of panicking.
func ExceptErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return Except(first, second), nil
}

// ExceptSelf produces the set difference of two sequences by using reflect.DeepEqual to compare values.
// 'second' is enumerated immediately.
// 'first' and 'second' may be based on the same Enumerator,
// 'first' must have real Reset method.
// ExceptSelf panics if 'first' or 'second' is nil.
func ExceptSelf[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return Except(first, NewOnSlice(sl2...))
}

// ExceptSelfErr is like ExceptSelf but returns an error instead of panicking.
func ExceptSelfErr[Source any](first, second Enumerator[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return ExceptSelf(first, second), nil
}

// ExceptEq produces the set difference of two sequences by using the specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ExceptEqSelf instead.
// ExceptEq panics if 'first' or 'second' is nil.
func ExceptEq[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
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
		MvNxt: func() bool {
			once.Do(func() { dsl2 = Slice(DistinctEq(second, eq)) })
			for d1.MoveNext() {
				c = d1.Current()
				if !elInElelEq(c, dsl2, eq) {
					return true
				}
			}
			return false
		},
		Crrnt: func() Source { return c },
		Rst:   func() { d1.Reset() },
	}
}

// ExceptEqErr is like ExceptEq but returns an error instead of panicking.
func ExceptEqErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return ExceptEq(first, second, eq), nil
}

// ExceptEqSelf produces the set difference of two sequences by using the specified Equaler to compare values.
// If 'eq' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
// ExceptEqSelf panics if 'first' or 'second' is nil.
func ExceptEqSelf[Source any](first, second Enumerator[Source], eq Equaler[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	sl2 := Slice(second)
	first.Reset()
	return ExceptEq(first, NewOnSlice(sl2...), eq)
}

// ExceptEqSelfErr is like ExceptEqSelf but returns an error instead of panicking.
func ExceptEqSelfErr[Source any](first, second Enumerator[Source], eq Equaler[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return ExceptEqSelf(first, second, eq), nil
}

// ExceptCmp produces the set difference of two sequences by using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ExceptCmpSelf instead.
// ExceptCmp panics if 'first' or 'second' or 'cmp' is nil.
func ExceptCmp[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
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
		MvNxt: func() bool {
			once.Do(func() {
				dsl2 = Slice(DistinctCmp(second, cmp))
				sort.Slice(dsl2, func(i, j int) bool { return cmp.Compare(dsl2[i], dsl2[j]) < 0 })
			})
			for d1.MoveNext() {
				c = d1.Current()
				if !elInElelCmp(c, dsl2, cmp) {
					return true
				}
			}
			return false
		},
		Crrnt: func() Source { return c },
		Rst:   func() { d1.Reset() },
	}
}

// ExceptCmpErr is like ExceptCmp but returns an error instead of panicking.
func ExceptCmpErr[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return ExceptCmp(first, second, cmp), nil
}

// ExceptCmpSelf produces the set difference of two sequences by using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
// ExceptCmpSelf panics if 'first' or 'second' or 'cmp' is nil.
func ExceptCmpSelf[Source any](first, second Enumerator[Source], cmp Comparer[Source]) Enumerator[Source] {
	if first == nil || second == nil {
		panic(ErrNilSource)
	}
	if cmp == nil {
		panic(ErrNilComparer)
	}
	sl2 := Slice(second)
	first.Reset()
	return ExceptCmp(first, NewOnSlice(sl2...), cmp)
}

// ExceptCmpSelfErr is like ExceptCmpSelf but returns an error instead of panicking.
func ExceptCmpSelfErr[Source any](first, second Enumerator[Source], cmp Comparer[Source]) (res Enumerator[Source], err error) {
	defer func() {
		catchPanic[Enumerator[Source]](recover(), &res, &err)
	}()
	return ExceptCmpSelf(first, second, cmp), nil
}
