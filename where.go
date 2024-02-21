package go2linq

import (
	"iter"
)

// [Where] filters a sequence of values based on a predicate.
//
// [Where]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.where
func Where[Source any](source iter.Seq[Source], predicate func(Source) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return func(yield func(Source) bool) {
			for s := range source {
				if !predicate(s) {
					continue
				}
				if !yield(s) {
					return
				}
			}
		},
		nil
}

// [WhereIdx] filters a sequence of values based on a predicate.
// Each element's index is used in the logic of the predicate function.
//
// [WhereIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.where
func WhereIdx[Source any](source iter.Seq[Source], predicate func(Source, int) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return func(yield func(Source) bool) {
			i := -1
			for s := range source {
				i++
				if !predicate(s, i) {
					continue
				}
				if !yield(s) {
					return
				}
			}
		},
		nil
}
