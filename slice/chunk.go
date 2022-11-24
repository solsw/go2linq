package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Chunk splits the elements of a slice into chunks of size at most 'size'.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func Chunk[Source any](source []Source, size int) ([][]Source, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return [][]Source{}, nil
	}
	en, err := go2linq.Chunk(go2linq.NewEnSlice(source...), size)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// ChunkMust is like Chunk but panics in case of error.
func ChunkMust[Source any](source []Source, size int) [][]Source {
	r, err := Chunk(source, size)
	if err != nil {
		panic(err)
	}
	return r
}
