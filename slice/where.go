//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Where filters a slice of T based on a predicate.
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func Where[T any](source []T, predicate func(T) bool) ([]T, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []T{}, nil
	}
	en, err := go2linq.Where(
		go2linq.NewEnSlice(source...),
		predicate,
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
