//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Cast casts the elements of a slice to the specified type.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func Cast[Source, Result any](source []Source) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.Cast[Source, Result](go2linq.NewEnSlice(source...))
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// CastMust is like Cast but panics in case of error.
func CastMust[Source, Result any](source []Source) []Result {
	r, err := Cast[Source, Result](source)
	if err != nil {
		panic(err)
	}
	return r
}
