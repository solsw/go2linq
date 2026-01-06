package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [CountBy] returns the count of elements in the source sequence grouped by key.
// Key values are compared using [reflect.DeepEqual]. 'source' is enumerated immediately.
//
// [CountBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.countby
func CountBy[Source, Key any](source iter.Seq[Source], keySelector func(Source) Key) (iter.Seq[generichelper.Tuple2[Key, int]], error) {
	return CountByEq(source, keySelector, generichelper.DeepEqual[Key])
}

// [CountByEq] returns the count of elements in the source sequence grouped by key.
// Key values are compared using 'keyEqual'. 'source' is enumerated immediately.
//
// [CountByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.countby
func CountByEq[Source, Key any](source iter.Seq[Source], keySelector func(Source) Key,
	keyEqual func(Key, Key) bool) (iter.Seq[generichelper.Tuple2[Key, int]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if keyEqual == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	gg, _ := GroupByEq(source, keySelector, keyEqual)
	r, _ := Select(gg, func(g Grouping[Key, Source]) generichelper.Tuple2[Key, int] {
		c, _ := Count(g.Values())
		return generichelper.Tuple2[Key, int]{Item1: g.key, Item2: c}
	})
	return r, nil
}
