package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Intersect] produces the set intersection of two sequences using [generichelper.DeepEqual] to compare values.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [Intersect]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func Intersect[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return IntersectEq(first, second, generichelper.DeepEqual[Source])
}

// [IntersectEq] produces the set intersection of two sequences using 'equal' to compare values.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [IntersectEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func IntersectEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	return seqIntersectByEq(first, second, Identity[Source], equal, equal, nil),
		nil
}

// [IntersectCmp] produces the set intersection of two sequences using 'comparer' to compare values. (See [DistinctCmp].)
// 'second' is enumerated on the first [Enumerator.MoveNext] call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [IntersectCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.intersect
func IntersectCmp[Source any](first, second iter.Seq[Source], compare func(Source, Source) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if compare == nil {
		return nil, ErrNilCompare
	}
	return seqIntersectByEq(first, second, Identity[Source],
			func(a, b Source) bool { return compare(a, b) == 0 }, nil, compare),
		nil
}
