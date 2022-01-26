//go:build go1.18

package go2linq

import (
	"sort"
	"sync"
)

// Reimplementing LINQ to Objects: Part 16 – IntersectErr (and build fiddling)
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-16-intersect-and-build-fiddling/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.intersect

// Intersect produces the set intersection of two sequences using reflect.DeepEqual as an equaler.
// 'second' is enumerated immediately.
func Intersect[Source any](first, second Enumerable[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return IntersectEq(first, second, nil)
}

// IntersectMust is like Intersect but panics in case of error.
func IntersectMust[Source any](first, second Enumerable[Source]) Enumerable[Source] {
	r, err := Intersect(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

func enrIntersectEq[Source any](first, second Enumerable[Source], equaler Equaler[Source]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctEqMust(first, equaler).GetEnumerator()
		var once sync.Once
		var dsl2 []Source
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() { dsl2 = ToSliceMust(DistinctEqMust(second, equaler)) })
				for enrD1.MoveNext() {
					c = enrD1.Current()
					if elInElelEq(c, dsl2, equaler) {
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

// IntersectEq produces the set intersection of two sequences using the specified Equaler to compare values.
// If 'equaler' is nil DeepEqual is used.
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
func IntersectEq[Source any](first, second Enumerable[Source], equaler Equaler[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equaler == nil {
		equaler = DeepEqual[Source]{}
	}
	return OnFactory(enrIntersectEq(first, second, equaler)), nil
}

// IntersectEqMust is like IntersectEq but panics in case of error.
func IntersectEqMust[Source any](first, second Enumerable[Source], equaler Equaler[Source]) Enumerable[Source] {
	r, err := IntersectEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

func enrIntersectCmp[Source any](first, second Enumerable[Source], comparer Comparer[Source]) func() Enumerator[Source] {
	return func() Enumerator[Source] {
		enrD1 := DistinctCmpMust(first, comparer).GetEnumerator()
		var once sync.Once
		var dsl2 []Source
		var c Source
		return enrFunc[Source]{
			mvNxt: func() bool {
				once.Do(func() {
					dsl2 = ToSliceMust(DistinctCmpMust(second, comparer))
					sort.Slice(dsl2, func(i, j int) bool { return comparer.Compare(dsl2[i], dsl2[j]) < 0 })
				})
				for enrD1.MoveNext() {
					c = enrD1.Current()
					if elInElelCmp(c, dsl2, comparer) {
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

// IntersectCmp produces the set intersection of two sequences using the specified Comparer to compare values.
// (See DistinctCmp function.)
// 'second' is enumerated immediately.
// Order of elements in the result corresponds to the order of elements in 'first'.
func IntersectCmp[Source any](first, second Enumerable[Source], comparer Comparer[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return OnFactory(enrIntersectCmp(first, second, comparer)), nil
}

// IntersectCmpMust is like IntersectCmp but panics in case of error.
func IntersectCmpMust[Source any](first, second Enumerable[Source], comparer Comparer[Source]) Enumerable[Source] {
	r, err := IntersectCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
