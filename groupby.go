package go2linq

import (
	"iter"
	"slices"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// [GroupBy] groups the elements of a sequence according to a specified key selector function.
// The keys are compared using [generichelper.DeepEqual]. 'source' is enumerated immediately.
//
// [GroupBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBy[Source, Key any](source iter.Seq[Source], keySelector func(Source) Key) (iter.Seq[Grouping[Key, Source]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	r, err := GroupBySelEq(source, keySelector, Identity[Source], generichelper.DeepEqual[Key])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [GroupByEq] groups the elements of a sequence according to a specified key
// selector function and compares the keys using 'equal'. 'source' is enumerated immediately.
//
// [GroupByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupByEq[Source, Key any](source iter.Seq[Source], keySelector func(Source) Key,
	equal func(Key, Key) bool) (iter.Seq[Grouping[Key, Source]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	r, err := GroupBySelEq(source, keySelector, Identity[Source], equal)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [GroupBySel] groups the elements of a sequence according to a specified key selector function
// and projects the elements for each group using a specified function.
// The keys are compared using [generichelper.DeepEqual]. 'source' is enumerated immediately.
//
// [GroupBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySel[Source, Key, Element any](source iter.Seq[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element) (iter.Seq[Grouping[Key, Element]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	r, err := GroupBySelEq(source, keySelector, elementSelector, generichelper.DeepEqual[Key])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [GroupBySelEq] groups the elements of a sequence according to a key selector function.
// The keys are compared using 'equal' and each group's elements are projected using a specified function.
// 'source' is enumerated immediately.
//
// [GroupBySelEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySelEq[Source, Key, Element any](source iter.Seq[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, equal func(Key, Key) bool) (iter.Seq[Grouping[Key, Element]], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	lk, _ := ToLookupSelEq(source, keySelector, elementSelector, equal)
	return slices.Values(lk.groupings), nil
}

// [GroupByRes] groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using [generichelper.DeepEqual]. 'source' is enumerated immediately.
//
// [GroupByRes]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupByRes[Source, Key, Result any](source iter.Seq[Source], keySelector func(Source) Key,
	resultSelector func(Key, iter.Seq[Source]) Result) (iter.Seq[Result], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || resultSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	r, err := GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, generichelper.DeepEqual[Key])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [GroupByResEq] groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using 'equal'. 'source' is enumerated immediately.
//
// [GroupByResEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupByResEq[Source, Key, Result any](source iter.Seq[Source], keySelector func(Source) Key,
	resultSelector func(Key, iter.Seq[Source]) Result, equal func(Key, Key) bool) (iter.Seq[Result], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || resultSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	r, err := GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, equal)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [GroupBySelRes] groups the elements of a sequence according to a specified
// key selector function and creates a result value from each group and its key.
// The elements of each group are projected using a specified function.
// Key values are compared using [generichelper.DeepEqual]. 'source' is enumerated immediately.
//
// [GroupBySelRes]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySelRes[Source, Key, Element, Result any](source iter.Seq[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, iter.Seq[Element]) Result) (iter.Seq[Result], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	r, err := GroupBySelResEq(source, keySelector, elementSelector, resultSelector, generichelper.DeepEqual[Key])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// [GroupBySelResEq] groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// Key values are compared using 'equal' and the elements of each group are projected using a specified function.
// 'source' is enumerated immediately.
//
// [GroupBySelResEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySelResEq[Source, Key, Element, Result any](source iter.Seq[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, iter.Seq[Element]) Result,
	equal func(Key, Key) bool) (iter.Seq[Result], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if equal == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	gg, _ := GroupBySelEq(source, keySelector, elementSelector, equal)
	r, err := Select(gg, func(g Grouping[Key, Element]) Result {
		return resultSelector(g.key, g.Values())
	})
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
