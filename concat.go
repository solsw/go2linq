package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Concat] concatenates two [sequences].
//
// [Concat]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.concat
// [sequences]: https://pkg.go.dev/iter#Seq
func Concat[Source any](first, second iter.Seq[Source]) (iter.Seq[Source], error) {
	return ConcatMany(first, second)
}

// ConcatMany concatenates [sequences].
//
// [sequences]: https://pkg.go.dev/iter#Seq
func ConcatMany[V any](seqs ...iter.Seq[V]) (iter.Seq[V], error) {
	if len(seqs) == 0 {
		return nil, errorhelper.CallerError(ErrEmptySource)
	}
	for _, seq := range seqs {
		if seq == nil {
			return nil, errorhelper.CallerError(ErrNilSource)
		}
	}
	return func(yield func(V) bool) {
			for _, seq := range seqs {
				for v := range seq {
					if !yield(v) {
						return
					}
				}
			}
		},
		nil
}
