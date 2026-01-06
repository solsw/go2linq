package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [Union] produces the set union of two sequences using [reflect.DeepEqual] to compare values.
//
// [Union]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func Union[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	r, err := UnionEq(first, second, generichelper.DeepEqual[Source])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [UnionEq] produces the set union of two sequences using 'equal' to compare values.
//
// [UnionEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func UnionEq[Source any](first, second iter.Seq[Source], equal func(Source, Source) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	concat, _ := Concat(first, second)
	r, err := DistinctEq(concat, equal)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [UnionCmp] produces the set union of two sequences using 'comparer' to compare values. (See [DistinctCmp].)
//
// [UnionCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.union
func UnionCmp[Source any](first, second iter.Seq[Source], compare func(Source, Source) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if compare == nil {
		return nil, errorhelper.CallerError(ErrNilCompare)
	}
	concat, _ := Concat(first, second)
	r, err := DistinctCmp(concat, compare)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
