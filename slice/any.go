package slice

import (
	"github.com/solsw/go2linq/v3"
)

// Any determines whether a slice contains any elements.
//
// If 'source' is nil or empty, false is returned.
func Any[Source any](source []Source) (bool, error) {
	return len(source) > 0, nil
}

// AnyPred determines whether any element of a slice satisfies a condition.
//
// If 'source' is nil or empty, false is returned.
func AnyPred[Source any](source []Source, predicate func(Source) bool) (bool, error) {
	if len(source) == 0 {
		return false, nil
	}
	r, err := go2linq.AnyPred(
		go2linq.NewEnSliceEn(source...),
		predicate,
	)
	if err != nil {
		return false, err
	}
	return r, nil
}
