package go2linq

import (
	"sort"
	"sync"

	"github.com/solsw/collate"
)

// Reimplementing LINQ to Objects: Part 16 – Intersect (and build fiddling)
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-16-intersect-and-build-fiddling/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect

// Intersect produces the set intersection of two sequences using collate.DeepEqualer to compare values.
// 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
// (https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect)
func Intersect[Source any](first, second Enumerable[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return IntersectEq(first, second, nil)
}

// IntersectMust is like [Intersect] but panics in case of error.
func IntersectMust[Source any](first, second Enumerable[Source]) Enumerable[Source] {
	r, err := Intersect(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

func factoryIntersectEq[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) func() Enumerator[Source] {
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

// IntersectEq produces the set intersection of two sequences using 'equaler' to compare values.
// If 'equaler' is nil collate.DeepEqualer is used. 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
// (https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect)
func IntersectEq[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Source]{}
	}
	return OnFactory(factoryIntersectEq(first, second, equaler)), nil
}

// IntersectEqMust is like [IntersectEq] but panics in case of error.
func IntersectEqMust[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) Enumerable[Source] {
	r, err := IntersectEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

func factoryIntersectCmp[Source any](first, second Enumerable[Source], comparer collate.Comparer[Source]) func() Enumerator[Source] {
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

// IntersectCmp produces the set intersection of two sequences using 'comparer' to compare values. (See [DistinctCmp].)
// 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
// (https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect)
func IntersectCmp[Source any](first, second Enumerable[Source], comparer collate.Comparer[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return OnFactory(factoryIntersectCmp(first, second, comparer)), nil
}

// IntersectCmpMust is like [IntersectCmp] but panics in case of error.
func IntersectCmpMust[Source any](first, second Enumerable[Source], comparer collate.Comparer[Source]) Enumerable[Source] {
	r, err := IntersectCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
