//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// SelectMany projects each element of a slice to a slice and flattens the resulting slices into one slice.
// If 'source' is nil, nil is returned. If 'source' is empty, new empty slice is returned.
func SelectMany[Source, Result any](source []Source, selector func(Source) []Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.SelectMany(go2linq.NewEnSlice(source...),
		func(s Source) go2linq.Enumerable[Result] { return go2linq.NewEnSlice(selector(s)...) },
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// SelectManyMust is like SelectMany but panics in case of error.
func SelectManyMust[Source, Result any](source []Source, selector func(Source) []Result) []Result {
	r, err := SelectMany(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}
