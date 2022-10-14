//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Contains determines whether a sequence contains a specified element using go2linq.DeepEqualer.
// If 'source' is nil or empty, false is returned.
func Contains[Source any](source []Source, value Source) (bool, error) {
	return ContainsEq(source, value, nil)
}

// ContainsMust is like Contains but panics in case of error.
func ContainsMust[Source any](source []Source, value Source) bool {
	r, err := Contains(source, value)
	if err != nil {
		panic(err)
	}
	return r
}

// ContainsEq determines whether a sequence contains a specified element using a specified equaler.
// If 'equaler' is nil go2linq.DeepEqualer is used.
// If 'source' is nil or empty, false is returned.
func ContainsEq[Source any](source []Source, value Source, equaler go2linq.Equaler[Source]) (bool, error) {
	if len(source) == 0 {
		return false, nil
	}
	r, err := go2linq.ContainsEq(go2linq.NewEnSlice(source...), value, equaler)
	if err != nil {
		return false, err
	}
	return r, nil
}

// ContainsEqMust is like ContainsEq but panics in case of error.
func ContainsEqMust[Source any](source []Source, value Source, equaler go2linq.Equaler[Source]) bool {
	r, err := ContainsEq(source, value, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
