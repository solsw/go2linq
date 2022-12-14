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
