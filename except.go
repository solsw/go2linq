//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// Reimplementing LINQ to Objects: Part 17 – Except
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-17-except/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.except

// Except produces the set difference of two sequences using reflect.DeepEqual to compare values.
// 'second' is enumerated immediately.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ExceptSelf instead.
func Except[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return ExceptEq(first, second, nil)
}

// ExceptMust is like Except but panics in case of error.
func ExceptMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := Except(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptSelf produces the set difference of two sequences using reflect.DeepEqual to compare values.
// 'second' is enumerated immediately.
// 'first' and 'second' may be based on the same Enumerator,
// 'first' must have real Reset method.
func ExceptSelf[Source any](first, second Enumerator[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return Except(first, NewOnSliceEn(sl2...))
}

// ExceptSelfMust is like ExceptSelf but panics in case of error.
func ExceptSelfMust[Source any](first, second Enumerator[Source]) Enumerator[Source] {
	r, err := ExceptSelf(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptEq produces the set difference of two sequences using the specified Equaler to compare values.
// If 'equaler' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ExceptEqSelf instead.
func ExceptEq[Source any](first, second Enumerator[Source], equaler Equaler[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equaler == nil {
		equaler = EqualerFunc[Source](DeepEqual[Source])
	}
	var once sync.Once
	var dsl2 []Source
	d1 := DistinctEqMust(first, equaler)
	var c Source
	return OnFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { dsl2 = Slice(DistinctEqMust(second, equaler)) })
				for d1.MoveNext() {
					c = d1.Current()
					if !elInElelEq(c, dsl2, equaler) {
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

// ExceptEqMust is like ExceptEq but panics in case of error.
func ExceptEqMust[Source any](first, second Enumerator[Source], equaler Equaler[Source]) Enumerator[Source] {
	r, err := ExceptEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptEqSelf produces the set difference of two sequences using the specified Equaler to compare values.
// If 'equaler' is nil reflect.DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
func ExceptEqSelf[Source any](first, second Enumerator[Source], equaler Equaler[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	sl2 := Slice(second)
	first.Reset()
	return ExceptEq(first, NewOnSliceEn(sl2...), equaler)
}

// ExceptEqSelfMust is like ExceptEqSelf but panics in case of error.
func ExceptEqSelfMust[Source any](first, second Enumerator[Source], equaler Equaler[Source]) Enumerator[Source] {
	r, err := ExceptEqSelf(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptCmp produces the set difference of two sequences using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' must not be based on the same Enumerator, otherwise use ExceptCmpSelf instead.
func ExceptCmp[Source any](first, second Enumerator[Source], comparer Comparer[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	var once sync.Once
	var dsl2 []Source
	d1 := DistinctCmpMust(first, comparer)
	var c Source
	return OnFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() {
					dsl2 = Slice(DistinctCmpMust(second, comparer))
					sort.Slice(dsl2, func(i, j int) bool { return comparer.Compare(dsl2[i], dsl2[j]) < 0 })
				})
				for d1.MoveNext() {
					c = d1.Current()
					if !elInElelCmp(c, dsl2, comparer) {
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

// ExceptCmpMust is like ExceptCmp but panics in case of error.
func ExceptCmpMust[Source any](first, second Enumerator[Source], comparer Comparer[Source]) Enumerator[Source] {
	r, err := ExceptCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}

// ExceptCmpSelf produces the set difference of two sequences using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
// 'first' and 'second' may be based on the same Enumerator.
// 'first' must have real Reset method.
func ExceptCmpSelf[Source any](first, second Enumerator[Source], comparer Comparer[Source]) (Enumerator[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	sl2 := Slice(second)
	first.Reset()
	return ExceptCmp(first, NewOnSliceEn(sl2...), comparer)
}

// ExceptCmpSelfMust is like ExceptCmpSelf but panics in case of error.
func ExceptCmpSelfMust[Source any](first, second Enumerator[Source], comparer Comparer[Source]) Enumerator[Source] {
	r, err := ExceptCmpSelf(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
