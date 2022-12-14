package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Aggregate applies an accumulator function over a slice.
// If 'source' is nil or empty, go2linq.ZeroValue[Source] is returned.
func Aggregate[Source any](source []Source, accumulator func(Source, Source) Source) (Source, error) {
	if len(source) == 0 {
		return go2linq.ZeroValue[Source](), nil
	}
	r, err := go2linq.Aggregate(go2linq.NewEnSlice(source...), accumulator)
	if err != nil {
		return go2linq.ZeroValue[Source](), err
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
		return go2linq.ZeroValue[Accumulate](), err
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
		return go2linq.ZeroValue[Result](), go2linq.ErrNilSelector
	}
	if len(source) == 0 {
		return resultSelector(seed), nil
	}
	r, err := go2linq.AggregateSeedSel(go2linq.NewEnSlice(source...), seed, accumulator, resultSelector)
	if err != nil {
		return go2linq.ZeroValue[Result](), err
	}
	return r, nil
}
