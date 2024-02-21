package go2linq

import (
	"iter"
)

// [Cast] casts the elements of a sequence to a specified type.
//
// [Cast]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.cast
func Cast[Source, Result any](source iter.Seq[Source]) (iter.Seq[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return func(yield func(Result) bool) {
			for s := range source {
				var a any = s
				if !yield(a.(Result)) {
					return
				}
			}
		},
		nil
}
