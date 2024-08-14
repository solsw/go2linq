package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [UnionBy] produces the set union of two sequences according to
// a specified key selector function and using [generichelper.DeepEqual] as key equaler.
//
// [UnionBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionBy[Source, Key any](first, second iter.Seq[Source], keySelector func(Source) Key) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	r, err := UnionByEq(first, second, keySelector, generichelper.DeepEqual[Key])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [UnionByEq] produces the set union of two sequences according to
// a specified key selector function and using a specified key equaler.
//
// [UnionByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionByEq[Source, Key any](first, second iter.Seq[Source],
	keySelector func(Source) Key, equal func(Key, Key) bool) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	concat, _ := Concat(first, second)
	r, err := DistinctByEq(concat, keySelector, equal)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [UnionByCmp] produces the set union of two sequences according to a specified
// key selector function and using a specified key comparer. (See [DistinctCmp].)
//
// [UnionByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.unionby
func UnionByCmp[Source, Key any](first, second iter.Seq[Source],
	keySelector func(Source) Key, compare func(Key, Key) int) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if compare == nil {
		return nil, errorhelper.CallerError(ErrNilCompare)
	}
	concat, _ := Concat(first, second)
	r, err := DistinctByCmp(concat, keySelector, compare)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
