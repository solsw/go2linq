package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Concat] concatenates two sequences.
//
// [Concat]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.concat
func Concat[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	return func(yield func(Source) bool) {
			for s1 := range first {
				if !yield(s1) {
					return
				}
			}
			for s2 := range second {
				if !yield(s2) {
					return
				}
			}
		},
		nil
}
