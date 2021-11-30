//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 13 - Aggregate
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-13-aggregate/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

// Aggregate applies an accumulator function over a sequence.
func Aggregate[Source any](source Enumerator[Source], accumulator func(Source, Source) Source) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if accumulator == nil {
		var s0 Source
		return s0, ErrNilAccumulator
	}
	if !source.MoveNext() {
		var s0 Source
		return s0, ErrEmptySource
	}
	r := source.Current()
	for source.MoveNext() {
		r = accumulator(r, source.Current())
	}
	return r, nil
}

// AggregateMust is like Aggregate but panics in case of error.
func AggregateMust[Source any](source Enumerator[Source], accumulator func(Source, Source) Source) Source {
	r, err := Aggregate(source, accumulator)
	if err != nil {
		panic(err)
	}
	return r
}

// AggregateSeed applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value.
func AggregateSeed[Source, Accumulate any](source Enumerator[Source],
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) (Accumulate, error) {
	if source == nil {
		var a0 Accumulate
		return a0, ErrNilSource
	}
	if accumulator == nil {
		var a0 Accumulate
		return a0, ErrNilAccumulator
	}
	r := seed
	for source.MoveNext() {
		r = accumulator(r, source.Current())
	}
	return r, nil
}

// AggregateSeedMust is like AggregateSeed but panics in case of error.
func AggregateSeedMust[Source, Accumulate any](source Enumerator[Source],
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) Accumulate {
	r, err := AggregateSeed(source, seed, accumulator)
	if err != nil {
		panic(err)
	}
	return r
}

// AggregateSeedSel applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value,
// and the specified function is used to select the result value.
func AggregateSeedSel[Source, Accumulate, Result any](source Enumerator[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) (Result, error) {
	if source == nil {
		var r0 Result
		return r0, ErrNilSource
	}
	if accumulator == nil {
		var r0 Result
		return r0, ErrNilAccumulator
	}
	if resultSelector == nil {
		var r0 Result
		return r0, ErrNilSelector
	}
	r := seed
	for source.MoveNext() {
		r = accumulator(r, source.Current())
	}
	return resultSelector(r), nil
}

// AggregateSeedSelMust is like AggregateSeedSel but panics in case of error.
func AggregateSeedSelMust[Source, Accumulate, Result any](source Enumerator[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) Result {
	r, err := AggregateSeedSel(source, seed, accumulator, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
