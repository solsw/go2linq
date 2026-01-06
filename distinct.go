package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [Distinct] returns distinct elements from a sequence using [reflect.DeepEqual] to compare values.
// Order of elements in the result corresponds to the order of elements in 'source'.
//
// [Distinct]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func Distinct[Source any](source iter.Seq[Source]) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	r, err := DistinctEq(source, generichelper.DeepEqual[Source])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [DistinctEq] returns distinct elements from a sequence using a specified 'equal' to compare values.
// Order of elements in the result corresponds to the order of elements in 'source'.
//
// [DistinctEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func DistinctEq[Source any](source iter.Seq[Source], equal func(Source, Source) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	r, err := DistinctByEq(source, Identity[Source], equal)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [DistinctCmp] returns distinct elements from a sequence using a specified 'compare' to compare values.
// Order of elements in the result corresponds to the order of elements in 'source'.
//
// Sorted slice of already seen elements is internally built.
// Sorted slice allows to use binary search to determine whether the element was seen or not.
// This may give performance gain when processing large sequences
// (though this is a subject for benchmarking, see [BenchmarkDistinctEqMust] and [BenchmarkDistinctCmpMust]).
//
// [DistinctCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinct
func DistinctCmp[Source any](source iter.Seq[Source], compare func(Source, Source) int) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if compare == nil {
		return nil, errorhelper.CallerError(ErrNilCompare)
	}
	r, err := DistinctByCmp(source, Identity[Source], compare)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
