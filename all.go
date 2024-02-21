package go2linq

import (
	"iter"
)

// [All] determines whether all elements of a sequence satisfy a condition.
//
// [All]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.all
func All[Source any](source iter.Seq[Source], predicate func(Source) bool) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if predicate == nil {
		return false, ErrNilPredicate
	}
	for s := range source {
		if !predicate(s) {
			return false, nil
		}
	}
	return true, nil
}
