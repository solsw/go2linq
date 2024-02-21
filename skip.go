package go2linq

import (
	"iter"
)

// [Skip] bypasses a specified number of elements in a sequence and then returns the remaining elements.
//
// [Skip]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skip
func Skip[Source any](source iter.Seq[Source], count int) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return source, nil
	}
	return func(yield func(Source) bool) {
			i := 0
			for s := range source {
				if i < count {
					i++
					continue
				}
				if !yield(s) {
					return
				}
			}
		},
		nil
}

// [SkipLast] returns a new sequence that contains the elements from 'source'
// with the last 'count' elements of the source collection omitted.
//
// [SkipLast]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skiplast
func SkipLast[Source any](source iter.Seq[Source], count int) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if count <= 0 {
		return source, nil
	}
	ss, _ := ToSlice(source)
	return SliceAll(ss[:len(ss)-count]), nil
}

// [SkipWhile] bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
//
// [SkipWhile]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile
func SkipWhile[Source any](source iter.Seq[Source], predicate func(Source) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return func(yield func(Source) bool) {
			rest := false
			for s := range source {
				if !rest {
					if predicate(s) {
						continue
					} else {
						rest = true
					}
				}
				if !yield(s) {
					return
				}
			}
		},
		nil
}

// [SkipWhileIdx] bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
// The element's index is used in the logic of the predicate function.
//
// [SkipWhileIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.skipwhile
func SkipWhileIdx[Source any](source iter.Seq[Source], predicate func(Source, int) bool) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if predicate == nil {
		return nil, ErrNilPredicate
	}
	return func(yield func(Source) bool) {
			rest := false
			i := 0
			for s := range source {
				if !rest {
					if predicate(s, i) {
						i++
						continue
					} else {
						rest = true
					}
				}
				if !yield(s) {
					return
				}
			}
		},
		nil
}
