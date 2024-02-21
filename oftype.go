package go2linq

import (
	"iter"
)

// [OfType] filters the elements of a sequence based on a specified type.
//
// [OfType]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.oftype
func OfType[Source, Result any](source iter.Seq[Source]) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return func(yield func(Result) bool) {
			for s := range source {
				var a any = s
				r, ok := a.(Result)
				if !ok {
					continue
				}
				if !yield(r) {
					return
				}
			}
		},
		nil
}
