//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// All determines whether all elements of a slice satisfy a condition.
// If 'source' is nil or empty, true is returned.
func All[Source any](source []Source, predicate func(Source) bool) (bool, error) {
	if len(source) == 0 {
		return true, nil
	}
	r, err := go2linq.All(go2linq.NewEnSlice(source...), predicate)
	if err != nil {
		return false, err
	}
	return r, nil
}

// AllMust is like All but panics in case of error.
func AllMust[Source any](source []Source, predicate func(Source) bool) bool {
	r, err := All(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}