package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [AggregateBy] applies an accumulator function over a sequence, grouping results by key.
// 'seed' is the initial accumulator value for each key.
// Key values are compared using [reflect.DeepEqual]. 'source' is enumerated immediately.
//
// [AggregateBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregateby
func AggregateBy[Source, Key, Accumulate any](source iter.Seq[Source], keySelector func(Source) Key,
	seed Accumulate, accumulator func(Accumulate, Source) Accumulate) (iter.Seq[generichelper.Tuple2[Key, Accumulate]], error) {
	return AggregateByEq(source, keySelector, seed, accumulator, generichelper.DeepEqual[Key])
}

// [AggregateByEq] applies an accumulator function over a sequence, grouping results by key.
// 'seed' is the initial accumulator value for each key.
// Key values are compared using 'keyEqual'. 'source' is enumerated immediately.
//
// [AggregateByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregateby
func AggregateByEq[Source, Key, Accumulate any](source iter.Seq[Source],
	keySelector func(Source) Key, seed Accumulate, accumulator func(Accumulate, Source) Accumulate,
	keyEqual func(Key, Key) bool) (iter.Seq[generichelper.Tuple2[Key, Accumulate]], error) {
	return AggregateBySelEq(source, keySelector, func(Key) Accumulate { return seed }, accumulator, keyEqual)
}

// [AggregateBySel] applies an accumulator function over a sequence, grouping results by key.
// 'seedSelector' is a factory for the initial accumulator value for each particular key.
// Key values are compared using [reflect.DeepEqual]. 'source' is enumerated immediately.
//
// [AggregateBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregateby
func AggregateBySel[Source, Key, Accumulate any](source iter.Seq[Source], keySelector func(Source) Key,
	seedSelector func(Key) Accumulate, accumulator func(Accumulate, Source) Accumulate) (iter.Seq[generichelper.Tuple2[Key, Accumulate]], error) {
	return AggregateBySelEq(source, keySelector, seedSelector, accumulator, generichelper.DeepEqual[Key])
}

// [AggregateBySelEq] applies an accumulator function over a sequence, grouping results by key.
// 'seedSelector' is a factory for the initial accumulator value for each particular key.
// Key values are compared using 'keyEqual'. 'source' is enumerated immediately.
//
// [AggregateBySelEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.aggregateby
func AggregateBySelEq[Source, Key, Accumulate any](source iter.Seq[Source], keySelector func(Source) Key,
	seedSelector func(Key) Accumulate, accumulator func(Accumulate, Source) Accumulate,
	keyEqual func(Key, Key) bool) (iter.Seq[generichelper.Tuple2[Key, Accumulate]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || seedSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if accumulator == nil {
		return nil, errorhelper.CallerError(ErrNilAccumulator)
	}
	if keyEqual == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	gg, _ := GroupByEq(source, keySelector, keyEqual)
	r, _ := Select(gg, func(g Grouping[Key, Source]) generichelper.Tuple2[Key, Accumulate] {
		ag, _ := AggregateSeed(g.Values(), seedSelector(g.key), accumulator)
		return generichelper.Tuple2[Key, Accumulate]{Item1: g.key, Item2: ag}
	})
	return r, nil
}
