package go2linq

import (
	"iter"
)

// [Any] determines whether a sequence contains any elements.
//
// [Any]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func Any[Source any](source iter.Seq[Source]) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	for _ = range source {
		return true, nil
	}
	return false, nil
}

// [AnyPred] determines whether any element of a sequence satisfies a condition.
//
// [AnyPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func AnyPred[Source any](source iter.Seq[Source], predicate func(Source) bool) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if predicate == nil {
		return false, ErrNilPredicate
	}
	for s := range source {
		if predicate(s) {
			return true, nil
		}
	}
	return false, nil
}
