package go2linq

import (
	"iter"
)

// [Take] returns a specified number of contiguous elements from the start of a sequence.
//
// [Take]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.take
func Take[Source any](source iter.Seq[Source], count int) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return Empty[Source](), nil
	}
	return func(yield func(Source) bool) {
			i := 0
			for s := range source {
				if !yield(s) {
					return
				}
				i++
				if i >= count {
					return
				}
			}
		},
		nil
}

// [TakeLast] returns a new [iter.Seq] that contains the last 'count' elements from 'source'.
//
// [TakeLast]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takelast
func TakeLast[Source any](source iter.Seq[Source], count int) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return Empty[Source](), nil
	}
	sl, _ := ToSlice(source)
	return SliceAll(sl[len(sl)-count:]), nil
}

// [TakeWhile] returns elements from a sequence as long as a specified condition is true.
//
// [TakeWhile]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func TakeWhile[Source any](source iter.Seq[Source], predicate func(Source) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return func(yield func(Source) bool) {
			for s := range source {
				if !predicate(s) {
					return
				}
				if !yield(s) {
					return
				}
			}
		},
		nil
}

// [TakeWhileIdx] returns elements from a sequence as long as a specified condition is true.
// The element's index is used in the logic of the predicate function.
//
// [TakeWhileIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.takewhile
func TakeWhileIdx[Source any](source iter.Seq[Source], predicate func(Source, int) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return func(yield func(Source) bool) {
			i := 0
			for s := range source {
				if !predicate(s, i) {
					return
				}
				if !yield(s) {
					return
				}
				i++
			}
		},
		nil
}
