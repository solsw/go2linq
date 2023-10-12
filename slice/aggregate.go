package slice

import (
	"github.com/solsw/generichelper"
	"github.com/solsw/go2linq/v3"
)

// Aggregate applies an accumulator function over a slice.
// If 'source' is nil or empty, Source's [zero value] is returned.
//
// [zero value]: https://go.dev/ref/spec#The_zero_value
func Aggregate[Source any](source []Source, accumulator func(Source, Source) Source) (Source, error) {
	if len(source) == 0 {
		return generichelper.ZeroValue[Source](), nil
	}
	r, err := go2linq.Aggregate(go2linq.NewEnSlice(source...), accumulator)
	if err != nil {
		return generichelper.ZeroValue[Source](), err
	}
	return r, nil
}

// AggregateSeed applies an accumulator function over a slice.
// The specified seed value is used as the initial accumulator value.
// If 'source' is nil or empty, 'seed' is returned.
func AggregateSeed[Source, Accumulate any](source []Source,
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) (Accumulate, error) {
	if len(source) == 0 {
		return seed, nil
	}
	r, err := go2linq.AggregateSeed(go2linq.NewEnSlice(source...), seed, accumulator)
	if err != nil {
		return generichelper.ZeroValue[Accumulate](), err
	}
	return r, nil
}

// AggregateSeedSel applies an accumulator function over a slice.
// The specified seed value is used as the initial accumulator value,
// and the specified function is used to select the result value.
// If 'source' is nil or empty, 'resultSelector(seed)' is returned.
func AggregateSeedSel[Source, Accumulate, Result any](source []Source, seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) (Result, error) {
	if resultSelector == nil {
		return generichelper.ZeroValue[Result](), go2linq.ErrNilSelector
	}
	if len(source) == 0 {
		return resultSelector(seed), nil
	}
	r, err := go2linq.AggregateSeedSel(
		go2linq.NewEnSlice(source...),
		seed,
		accumulator,
		resultSelector,
	)
	if err != nil {
		return generichelper.ZeroValue[Result](), err
	}
	return r, nil
}
