//go:build go1.18

package slice

import (
	"github.com/solsw/go2linq/v2"
)

// Aggregate applies an accumulator function over a slice.
// If 'source' is nil or empty, go2linq.ZeroValue is returned.
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

// AggregateMust is like Aggregate but panics in case of error.
func AggregateMust[Source any](source []Source, accumulator func(Source, Source) Source) Source {
	r, err := Aggregate(source, accumulator)
	if err != nil {
		panic(err)
	}
	return r
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

// AggregateSeedMust is like AggregateSeed but panics in case of error.
func AggregateSeedMust[Source, Accumulate any](source []Source,
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) Accumulate {
	r, err := AggregateSeed(source, seed, accumulator)
	if err != nil {
		panic(err)
	}
	return r
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

// AggregateSeedSelMust is like AggregateSeedSel but panics in case of error.
func AggregateSeedSelMust[Source, Accumulate, Result any](source []Source, seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) Result {
	r, err := AggregateSeedSel(source, seed, accumulator, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
