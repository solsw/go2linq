package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Except] produces the set difference of two sequences using [generichelper.DeepEqual] to compare values.
// 'second' is enumerated on the first 'next' call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [Except]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func Except[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	return ExceptEq(first, second, generichelper.DeepEqual[Source])
}

// [ExceptEq] produces the set difference of two sequences using 'equal' to compare values.
// 'second' is enumerated on the first 'next' call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExceptEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if equal == nil {
		return nil, ErrNilEqual
	}
	return ExceptByEq(first, second, Identity[Source], equal)
}

// [ExceptCmp] produces the set difference of two sequences using 'compare' to compare values. (See [DistinctCmp].)
// 'second' is enumerated on the first 'next' call.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExceptCmp[Source any](first, second iter.Seq[Source], compare func(Source, Source) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, ErrNilSource
	}
	if compare == nil {
		return nil, ErrNilCompare
	}
	return ExceptByCmp(first, second, Identity[Source], compare)
}
