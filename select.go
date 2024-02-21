package go2linq

import (
	"iter"
)

// [Select] projects each element of a sequence into a new form.
//
// [Select]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.select
func Select[Source, Result any](source iter.Seq[Source], selector func(Source) Result) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			for s := range source {
				if !yield(selector(s)) {
					return
				}
			}
		},
		nil
}

// [SelectIdx] projects each element of a sequence into a new form by incorporating the element's index.
//
// [SelectIdx]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.select
func SelectIdx[Source, Result any](source iter.Seq[Source], selector func(Source, int) Result) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return func(yield func(Result) bool) {
			i := 0
			for s := range source {
				if !yield(selector(s, i)) {
					return
				}
				i++
			}
		},
		nil
}
