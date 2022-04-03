//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 13 - Aggregate
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-13-aggregate/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

// Aggregate applies an accumulator function over a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate)
func Aggregate[Source any](source Enumerable[Source], accumulator func(Source, Source) Source) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if accumulator == nil {
		return ZeroValue[Source](), ErrNilAccumulator
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return ZeroValue[Source](), ErrEmptySource
	}
	r := enr.Current()
	for enr.MoveNext() {
		r = accumulator(r, enr.Current())
	}
	return r, nil
}

// AggregateMust is like Aggregate but panics in case of an error.
func AggregateMust[Source any](source Enumerable[Source], accumulator func(Source, Source) Source) Source {
	r, err := Aggregate(source, accumulator)
	if err != nil {
		panic(err)
	}
	return r
}

// AggregateSeed applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate)
func AggregateSeed[Source, Accumulate any](source Enumerable[Source],
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) (Accumulate, error) {
	if source == nil {
		return ZeroValue[Accumulate](), ErrNilSource
	}
	if accumulator == nil {
		return ZeroValue[Accumulate](), ErrNilAccumulator
	}
	enr := source.GetEnumerator()
	r := seed
	for enr.MoveNext() {
		r = accumulator(r, enr.Current())
	}
	return r, nil
}

// AggregateSeedMust is like AggregateSeed but panics in case of an error.
func AggregateSeedMust[Source, Accumulate any](source Enumerable[Source],
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
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate)
func AggregateSeedSel[Source, Accumulate, Result any](source Enumerable[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) (Result, error) {
	if source == nil {
		return ZeroValue[Result](), ErrNilSource
	}
	if accumulator == nil {
		return ZeroValue[Result](), ErrNilAccumulator
	}
	if resultSelector == nil {
		return ZeroValue[Result](), ErrNilSelector
	}
	enr := source.GetEnumerator()
	r := seed
	for enr.MoveNext() {
		r = accumulator(r, enr.Current())
	}
	return resultSelector(r), nil
}

// AggregateSeedSelMust is like AggregateSeedSel but panics in case of an error.
func AggregateSeedSelMust[Source, Accumulate, Result any](source Enumerable[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) Result {
	r, err := AggregateSeedSel(source, seed, accumulator, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
