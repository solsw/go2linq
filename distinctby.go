package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [DistinctBy] returns distinct elements from a sequence according to
// a specified key selector function and using [generichelper.DeepEqual] to compare keys.
//
// [DistinctBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
func DistinctBy[Source, Key any](source iter.Seq[Source], keySelector func(Source) Key) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	return DistinctByEq(source, keySelector, generichelper.DeepEqual[Key])
}

// [DistinctByEq] returns distinct elements from a sequence according to
// a specified key selector function and using a specified 'equal' to compare keys.
//
// [DistinctByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
func DistinctByEq[Source, Key any](source iter.Seq[Source],
	keySelector func(Source) Key, equal func(Key, Key) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	return func(yield func(Source) bool) {
			var seen []Key
			for s := range source {
				k := keySelector(s)
				if !elInElelEq(k, seen, equal) {
					seen = append(seen, k)
					if !yield(s) {
						return
					}
				}
			}
		},
		nil
}

// [DistinctByCmp] returns distinct elements from a sequence according to a specified key selector function
// and using a specified 'compare' to compare keys. (See [DistinctCmp].)
//
// [DistinctByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
func DistinctByCmp[Source, Key any](source iter.Seq[Source],
	keySelector func(Source) Key, compare func(Key, Key) int) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if compare == nil {
		return nil, errorhelper.CallerError(ErrNilCompare)
	}
	return func(yield func(Source) bool) {
			seen := make([]Key, 0)
			for s := range source {
				k := keySelector(s)
				i := elIdxInElelCmp(k, seen, compare)
				if i == len(seen) || compare(k, seen[i]) != 0 {
					elIntoElelAtIdx(k, &seen, i)
					if !yield(s) {
						return
					}
				}
			}
		},
		nil
}
