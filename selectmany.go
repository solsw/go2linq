package go2linq

import (
	"iter"
)

// [SelectMany] projects each element of a sequence to another sequence
// and flattens the resulting sequences into one sequence.
//
// [SelectMany]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectMany[Source, Result any](source iter.Seq[Source], selector func(Source) iter.Seq[Result]) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			for s := range source {
				for r := range selector(s) {
					if !yield(r) {
						return
					}
				}
			}
		},
		nil
}

// [SelectManyIdx] projects each element of a sequence and its index to another sequence
// and flattens the resulting sequences into one sequence.
//
// [SelectManyIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectManyIdx[Source, Result any](source iter.Seq[Source], selector func(Source, int) iter.Seq[Result]) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			i := 0
			for s := range source {
				for r := range selector(s, i) {
					if !yield(r) {
						return
					}
				}
				i++
			}
		},
		nil
}

// [SelectManyColl] projects each element of a sequence to another sequence,
// flattens the resulting sequences into one sequence
// and invokes a result selector function on each element therein.
//
// [SelectManyColl]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectManyColl[Source, Collection, Result any](source iter.Seq[Source],
	collectionSelector func(Source) iter.Seq[Collection], resultSelector func(Source, Collection) Result) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if collectionSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			for s := range source {
				for c := range collectionSelector(s) {
					if !yield(resultSelector(s, c)) {
						return
					}
				}
			}
		},
		nil
}

// [SelectManyCollIdx] projects each element of a sequence and its index to another sequence,
// flattens the resulting sequences into one sequence
// and invokes a result selector function on each element therein.
//
// [SelectManyCollIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.selectmany
func SelectManyCollIdx[Source, Collection, Result any](source iter.Seq[Source],
	collectionSelector func(Source, int) iter.Seq[Collection], resultSelector func(Source, Collection) Result) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if collectionSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			i := 0
			for s := range source {
				for c := range collectionSelector(s, i) {
					if !yield(resultSelector(s, c)) {
						return
					}
				}
				i++
			}
		},
		nil
}
