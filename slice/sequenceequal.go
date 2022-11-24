package slice

import (
	"github.com/solsw/go2linq/v2"
	"golang.org/x/exp/slices"
)

// SequenceEqual determines whether two slices are equal by comparing the elements using go2linq.DeepEqualer.
func SequenceEqual[Source any](first, second []Source) bool {
	return SequenceEqualEq(first, second, nil)
}

// SequenceEqualEq determines whether two slices are equal by comparing their elements using a specified equaler.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func SequenceEqualEq[Source any](first, second []Source, equaler go2linq.Equaler[Source]) bool {
	if equaler == nil {
		equaler = go2linq.DeepEqualer[Source]{}
	}
	return slices.EqualFunc(first, second, equaler.Equal)
}
