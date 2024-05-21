package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Chunk] splits the elements of a sequence into chunks of size at most 'size'.
//
// [Chunk]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.chunk
func Chunk[Source any](source iter.Seq[Source], size int) (iter.Seq[[]Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if size <= 0 {
		return nil, errorhelper.CallerError(ErrSizeOutOfRange)
	}
	return func(yield func([]Source) bool) {
			next, stop := iter.Pull(source)
			defer stop()
			for {
				ss := make([]Source, 0, size)
				for range size {
					s, ok := next()
					if !ok {
						break
					}
					ss = append(ss, s)
				}
				if len(ss) == 0 {
					return
				}
				if !yield(ss) {
					return
				}
			}
		},
		nil
}
