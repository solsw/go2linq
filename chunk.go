package go2linq

import (
	"github.com/solsw/generichelper"
)

// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.chunk

func factoryChunk[Source any](source Enumerable[Source], size int) func() Enumerator[[]Source] {
	return func() Enumerator[[]Source] {
		enr := source.GetEnumerator()
		c := make([]Source, 0, size)
		return enrFunc[[]Source]{
			mvNxt: func() bool {
				for enr.MoveNext() {
					c = append(c, enr.Current())
					if len(c) == size {
						return true
					}
				}
				return len(c) > 0
			},
			crrnt: func() []Source {
				lc := c
				c = make([]Source, 0, size)
				return lc
			},
			rst: func() {
				c = make([]Source, 0, size)
				enr.Reset()
			},
		}
	}
}

// [Chunk] splits the elements of a sequence into chunks of size at most 'size'.
//
// [Chunk]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.chunk
func Chunk[Source any](source Enumerable[Source], size int) (Enumerable[[]Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if size <= 0 {
		return nil, ErrSizeOutOfRange
	}
	return OnFactory(factoryChunk(source, size)), nil
}

// ChunkMust is like [Chunk] but panics in case of error.
func ChunkMust[Source any](source Enumerable[Source], size int) Enumerable[[]Source] {
	return generichelper.Must(Chunk(source, size))
}
