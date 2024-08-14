package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [Except] produces the set difference of two sequences using [generichelper.DeepEqual] to compare values.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [Except]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func Except[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	r, err := ExceptEq(first, second, generichelper.DeepEqual[Source])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [ExceptEq] produces the set difference of two sequences using 'equal' to compare values.
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExceptEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	r, err := ExceptByEq(first, second, Identity[Source], equal)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [ExceptCmp] produces the set difference of two sequences using 'compare' to compare values. (See [DistinctCmp].)
// 'second' is enumerated on the first iteration over the result.
// Order of elements in the result corresponds to the order of elements in 'first'.
//
// [ExceptCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.except
func ExceptCmp[Source any](first, second iter.Seq[Source], compare func(Source, Source) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if compare == nil {
		return nil, errorhelper.CallerError(ErrNilCompare)
	}
	r, err := ExceptByCmp(first, second, Identity[Source], compare)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
