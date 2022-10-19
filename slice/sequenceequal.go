//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
	"golang.org/x/exp/slices"
)

// SequenceEqual determines whether two slices are equal by comparing the elements using go2linq.DeepEqualer.
func SequenceEqual[Source any](first, second []Source) (bool, error) {
	return SequenceEqualEq(first, second, nil)
}

// SequenceEqualMust is like SequenceEqual but panics in case of error.
func SequenceEqualMust[Source any](first, second []Source) bool {
	return SequenceEqualEqMust(first, second, nil)
}

// SequenceEqualEq determines whether two slices are equal by comparing their elements using a specified equaler.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func SequenceEqualEq[Source any](first, second []Source, equaler go2linq.Equaler[Source]) (res bool, err error) {
	return SequenceEqualEqMust(first, second, equaler), nil
}

// SequenceEqualEqMust is like SequenceEqualEq but panics in case of error.
func SequenceEqualEqMust[Source any](first, second []Source, equaler go2linq.Equaler[Source]) bool {
	if equaler == nil {
		equaler = go2linq.DeepEqualer[Source]{}
	}
	return slices.EqualFunc(first, second, equaler.Equal)
}
