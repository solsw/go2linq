package go2linq

import (
	"iter"
)

// [Count] returns the number of elements in a sequence.
//
// [Count]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
func Count[Source any](source iter.Seq[Source]) (int, error) {
	if source == nil {
		return -1, ErrNilSource
	}
	res := 0
	for _ = range source {
		res++
	}
	return res, nil
}

// [CountPred] returns a number that represents how many elements in a specified sequence satisfy a condition.
//
// [CountPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
func CountPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (int, error) {
	if source == nil {
		return -1, ErrNilSource
	}
	if predicate == nil {
		return -1, ErrNilPredicate
	}
	res := 0
	for s := range source {
		if predicate(s) {
			res++
		}
	}
	return res, nil
}
