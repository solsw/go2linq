package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Select projects each element of a slice into a new form.
//
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func Select[Source, Result any](source []Source, selector func(Source) Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.Select(
		go2linq.NewEnSliceEn(source...),
		selector,
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}

// SelectIdx projects each element of a slice into a new form by incorporating the element's index.
//
// If 'source' is nil, nil is returned.
// If 'source' is empty, new empty slice is returned.
func SelectIdx[Source, Result any](source []Source, selector func(Source, int) Result) ([]Result, error) {
	if source == nil {
		return nil, nil
	}
	if len(source) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.SelectIdx(
		go2linq.NewEnSliceEn(source...),
		selector,
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
