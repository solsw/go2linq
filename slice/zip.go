package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Zip applies a specified function to the corresponding elements
// of two slices, producing a slice of the results.
//
// If 'first' or 'second' is nil, nil is returned.
// If 'first' or 'second' is empty, new empty slice is returned.
func Zip[First, Second, Result any](first []First, second []Second,
	resultSelector func(First, Second) Result) ([]Result, error) {
	if first == nil || second == nil {
		return nil, nil
	}
	if len(first) == 0 || len(second) == 0 {
		return []Result{}, nil
	}
	en, err := go2linq.Zip(
		go2linq.NewEnSliceEn(first...),
		go2linq.NewEnSliceEn(second...),
		resultSelector,
	)
	if err != nil {
		return nil, err
	}
	return go2linq.ToSlice(en)
}
