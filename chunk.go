//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.chunk

func enrChunk[Source any](source Enumerable[Source], size int) func() Enumerator[[]Source] {
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
				if len(c) > 0 {
					return true
				}
				return false
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

// Chunk splits the elements of a sequence into chunks of size at most 'size'.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.chunk)
func Chunk[Source any](source Enumerable[Source], size int) (Enumerable[[]Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if size <= 0 {
		return nil, ErrSizeOutOfRange
	}
	return OnFactory(enrChunk(source, size)), nil
}

// ChunkMust is like Chunk but panics in case of an error.
func ChunkMust[Source any](source Enumerable[Source], size int) Enumerable[[]Source] {
	r, err := Chunk(source, size)
	if err != nil {
		panic(err)
	}
	return r
}
