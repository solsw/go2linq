package slice

import (
	"github.com/solsw/go2linq/v2"
	"golang.org/x/exp/slices"
)

// Contains determines whether a sequence contains a specified element using go2linq.DeepEqualer.
// If 'source' is nil or empty, false is returned.
func Contains[Source any](source []Source, value Source) bool {
	return ContainsEq(source, value, nil)
}

// ContainsEq determines whether a sequence contains a specified element using a specified equaler.
// If 'source' is nil or empty, false is returned.
// If 'equaler' is nil go2linq.DeepEqualer is used.
func ContainsEq[Source any](source []Source, value Source, equaler go2linq.Equaler[Source]) bool {
	if len(source) == 0 {
		return false
	}
	if equaler == nil {
		equaler = go2linq.DeepEqualer[Source]{}
	}
	return slices.IndexFunc(source, func(v Source) bool { return equaler.Equal(v, value) }) >= 0
}
