package go2linq

import (
	"github.com/solsw/collate"
)

// Reimplementing LINQ to Objects: Part 15 – Union
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-15-union/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union

// Union produces the set union of two sequences using collate.DeepEqualer to compare values.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union)
func Union[Source any](first, second Enumerable[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return UnionEq(first, second, nil)
}

// UnionMust is like Union but panics in case of error.
func UnionMust[Source any](first, second Enumerable[Source]) Enumerable[Source] {
	r, err := Union(first, second)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionEq produces the set union of two sequences using a specified collate.Equaler to compare values.
// If 'equaler' is nil collate.DeepEqualer is used.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union)
func UnionEq[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return DistinctEq(ConcatMust(first, second), equaler)
}

// UnionEqMust is like UnionEq but panics in case of error.
func UnionEqMust[Source any](first, second Enumerable[Source], equaler collate.Equaler[Source]) Enumerable[Source] {
	r, err := UnionEq(first, second, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// UnionCmp produces the set union of two sequences using a specified collate.Comparer to compare values.
// (See DistinctCmp function.)
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.union)
func UnionCmp[Source any](first, second Enumerable[Source], comparer collate.Comparer[Source]) (Enumerable[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	return DistinctCmp(ConcatMust(first, second), comparer)
}

// UnionCmpMust is like UnionCmp but panics in case of error.
func UnionCmpMust[Source any](first, second Enumerable[Source], comparer collate.Comparer[Source]) Enumerable[Source] {
	r, err := UnionCmp(first, second, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
