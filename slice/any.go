//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Any determines whether a slice contains any elements.
// If 'source' is nil or empty, false is returned.
func Any[Source any](source []Source) (bool, error) {
	return len(source) > 0, nil
}

// AnyMust is like Any but panics in case of error.
func AnyMust[Source any](source []Source) bool {
	return len(source) > 0
}

// AnyPred determines whether any element of a slice satisfies a condition.
// If 'source' is nil or empty, false is returned.
func AnyPred[Source any](source []Source, predicate func(Source) bool) (bool, error) {
	if len(source) == 0 {
		return false, nil
	}
	r, err := go2linq.AnyPred(go2linq.NewEnSlice(source...), predicate)
	if err != nil {
		return false, err
	}
	return r, nil
}

// AnyPredMust is like AnyPred but panics in case of error.
func AnyPredMust[Source any](source []Source, predicate func(Source) bool) bool {
	r, err := AnyPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
