package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Aggregate] applies an accumulator function over a sequence.
//
// [Aggregate]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregate
func Aggregate[Source any](source iter.Seq[Source], accumulator func(Source, Source) Source) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if accumulator == nil {
		return generichelper.ZeroValue[Source](), ErrNilAccumulator
	}
	var res Source
	empty := true
	first := true
	for s := range source {
		empty = false
		if first {
			first = false
			res = s
			continue
		}
		res = accumulator(res, s)
	}
	if empty {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return res, nil
}

// [AggregateSeed] applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value.
//
// [AggregateSeed]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregate
func AggregateSeed[Source, Accumulate any](source iter.Seq[Source],
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) (Accumulate, error) {
	if source == nil {
		return generichelper.ZeroValue[Accumulate](), ErrNilSource
	}
	if accumulator == nil {
		return generichelper.ZeroValue[Accumulate](), ErrNilAccumulator
	}
	res := seed
	for s := range source {
		res = accumulator(res, s)
	}
	return res, nil
}

// [AggregateSeedSel] applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value,
// and the specified function is used to select the result value.
//
// [AggregateSeedSel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregate
func AggregateSeedSel[Source, Accumulate, Result any](source iter.Seq[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) (Result, error) {
	if source == nil {
		return generichelper.ZeroValue[Result](), ErrNilSource
	}
	if accumulator == nil {
		return generichelper.ZeroValue[Result](), ErrNilAccumulator
	}
	if resultSelector == nil {
		return generichelper.ZeroValue[Result](), ErrNilSelector
	}
	res := seed
	for s := range source {
		res = accumulator(res, s)
	}
	return resultSelector(res), nil
}
