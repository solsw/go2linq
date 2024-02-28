package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Union] produces the set union of two sequences using [generichelper.DeepEqual] to compare values.
//
// [Union]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func Union[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return UnionEq(first, second, generichelper.DeepEqual[Source])
}

// [UnionEq] produces the set union of two sequences using 'equal' to compare values.
//
// [UnionEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func UnionEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	concat, _ := Concat(first, second)
	return DistinctEq(concat, equal)
}

// [UnionCmp] produces the set union of two sequences using 'comparer' to compare values. (See [DistinctCmp].)
//
// [UnionCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func UnionCmp[Source any](first, second iter.Seq[Source], compare func(Source, Source) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if compare == nil {
		return nil, ErrNilCompare
	}
	concat, _ := Concat(first, second)
	return DistinctCmp(concat, compare)
}
