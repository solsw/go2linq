//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 13 - Aggregate
// https://codeblog.jonskeet.uk/2010/12/30/reimplementing-linq-to-objects-part-13-aggregate/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.aggregate

// Aggregate applies an accumulator function over a sequence.
// Aggregate panics if 'source' or 'accumulator' is nil or 'source' contains no elements.
func Aggregate[Source any](source Enumerator[Source], accumulator func(Source, Source) Source) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if accumulator == nil {
		panic(ErrNilAccumulator)
	}
	if !source.MoveNext() {
		panic(ErrEmptySource)
	}
	r := source.Current()
	for source.MoveNext() {
		r = accumulator(r, source.Current())
	}
	return r
}

// AggregateErr is like Aggregate but returns an error instead of panicking.
func AggregateErr[Source any](source Enumerator[Source], accumulator func(Source, Source) Source) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return Aggregate(source, accumulator), nil
}

// AggregateSeed applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value.
// AggregateSeed panics if 'source' or 'accumulator' is nil.
func AggregateSeed[Source, Accumulate any](source Enumerator[Source],
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) Accumulate {
	if source == nil {
		panic(ErrNilSource)
	}
	if accumulator == nil {
		panic(ErrNilAccumulator)
	}
	r := seed
	for source.MoveNext() {
		r = accumulator(r, source.Current())
	}
	return r
}

// AggregateSeedErr is like AggregateSeed but returns an error instead of panicking.
func AggregateSeedErr[Source, Accumulate any](source Enumerator[Source],
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) (res Accumulate, err error) {
	defer func() {
		catchErrPanic[Accumulate](recover(), &res, &err)
	}()
	return AggregateSeed(source, seed, accumulator), nil
}

// AggregateSeedSel applies an accumulator function over a sequence.
// The specified seed value is used as the initial accumulator value,
// and the specified function is used to select the result value.
// AggregateSeedSel panics if 'source' or 'accumulator' or 'resultSelector' is nil.
func AggregateSeedSel[Source, Accumulate, Result any](source Enumerator[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) Result {
	if source == nil {
		panic(ErrNilSource)
	}
	if accumulator == nil {
		panic(ErrNilAccumulator)
	}
	if resultSelector == nil {
		panic(ErrNilSelector)
	}
	r := seed
	for source.MoveNext() {
		r = accumulator(r, source.Current())
	}
	return resultSelector(r)
}

// AggregateSeedSelErr is like AggregateSeedSel but returns an error instead of panicking.
func AggregateSeedSelErr[Source, Accumulate, Result any](source Enumerator[Source], seed Accumulate,
	accumulator func(Accumulate, Source) Accumulate, resultSelector func(Accumulate) Result) (res Result, err error) {
	defer func() {
		catchErrPanic[Result](recover(), &res, &err)
	}()
	return AggregateSeedSel(source, seed, accumulator, resultSelector), nil
}
