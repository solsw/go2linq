//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.chunk

// Chunk splits the elements of a sequence into chunks of size at most 'size'.
func Chunk[Source any](source Enumerator[Source], size int) (Enumerator[[]Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if size <= 0 {
		return nil, ErrSizeOutOfRange
	}
	c := make([]Source, 0, size)
	return OnFunc[[]Source]{
			mvNxt: func() bool {
				for source.MoveNext() {
					c = append(c, source.Current())
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
				source.Reset()
			},
		},
		nil
}
