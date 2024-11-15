package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [Index] returns a sequence that incorporates the element's index into a tuple.
//
// [Index]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.index
func Index[Source any](source iter.Seq[Source]) (iter.Seq[generichelper.Tuple2[int, Source]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	return func(yield func(generichelper.Tuple2[int, Source]) bool) {
			i := 0
			for s := range source {
				if !yield(generichelper.Tuple2[int, Source]{Item1: i, Item2: s}) {
					return
				}
				i++
			}
		},
		nil
}
