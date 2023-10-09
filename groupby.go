package go2linq

import (
	"github.com/solsw/collate"
	"github.com/solsw/errorhelper"
)

// Reimplementing LINQ to Objects: Part 21 - GroupByErr
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-21-groupby/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby

// [GroupBy] groups the elements of a sequence according to a specified key selector function.
// The keys are compared using [collate.DeepEqualer]. 'source' is enumerated immediately.
//
// [GroupBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBy[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) (Enumerable[Grouping[Key, Source]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelEq(source, keySelector, Identity[Source], nil)
}

// GroupByMust is like [GroupBy] but panics in case of error.
func GroupByMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) Enumerable[Grouping[Key, Source]] {
	return errorhelper.Must(GroupBy(source, keySelector))
}

// [GroupByEq] groups the elements of a sequence according to
// a specified key selector function and compares the keys using 'equaler'.
// If 'equaler' is nil, [collate.DeepEqualer] is used. 'source' is enumerated immediately.
//
// [GroupByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupByEq[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) (Enumerable[Grouping[Key, Source]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelEq(source, keySelector, Identity[Source], equaler)
}

// GroupByEqMust is like [GroupByEq] but panics in case of error.
func GroupByEqMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, equaler collate.Equaler[Key]) Enumerable[Grouping[Key, Source]] {
	return errorhelper.Must(GroupByEq(source, keySelector, equaler))
}

// [GroupBySel] groups the elements of a sequence according to a specified key selector function
// and projects the elements for each group using a specified function.
// The keys are compared using [collate.DeepEqualer]. 'source' is enumerated immediately.
//
// [GroupBySel]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySel[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (Enumerable[Grouping[Key, Element]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelEq(source, keySelector, elementSelector, nil)
}

// GroupBySelMust is like [GroupBySel] but panics in case of error.
func GroupBySelMust[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) Enumerable[Grouping[Key, Element]] {
	return errorhelper.Must(GroupBySel(source, keySelector, elementSelector))
}

// [GroupBySelEq] groups the elements of a sequence according to a key selector function.
// The keys are compared using 'equaler' and each group's elements are projected using a specified function.
// If 'equaler' is nil, [collate.DeepEqualer] is used. 'source' is enumerated immediately.
//
// [GroupBySelEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySelEq[Source, Key, Element any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, equaler collate.Equaler[Key]) (Enumerable[Grouping[Key, Element]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Key]{}
	}
	lk := ToLookupSelEqMust(source, keySelector, elementSelector, equaler)
	return lk, nil
}

// GroupBySelEqMust is like [GroupBySelEq] but panics in case of error.
func GroupBySelEqMust[Source, Key, Element any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, equaler collate.Equaler[Key]) Enumerable[Grouping[Key, Element]] {
	return errorhelper.Must(GroupBySelEq(source, keySelector, elementSelector, equaler))
}

// [GroupByRes] groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using [collate.DeepEqualer]. 'source' is enumerated immediately.
//
// [GroupByRes]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupByRes[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, nil)
}

// GroupByResMust is like [GroupByRes] but panics in case of error.
func GroupByResMust[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result) Enumerable[Result] {
	return errorhelper.Must(GroupByRes(source, keySelector, resultSelector))
}

// [GroupByResEq] groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using 'equaler'.
// If 'equaler' is nil, [collate.DeepEqualer] is used. 'source' is enumerated immediately.
//
// [GroupByResEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupByResEq[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result, equaler collate.Equaler[Key]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, equaler)
}

// GroupByResEqMust is like [GroupByResEq] but panics in case of error.
func GroupByResEqMust[Source, Key, Result any](source Enumerable[Source], keySelector func(Source) Key,
	resultSelector func(Key, Enumerable[Source]) Result, equaler collate.Equaler[Key]) Enumerable[Result] {
	return errorhelper.Must(GroupByResEq(source, keySelector, resultSelector, equaler))
}

// [GroupBySelRes] groups the elements of a sequence according to a specified
// key selector function and creates a result value from each group and its key.
// The elements of each group are projected using a specified function.
// Key values are compared using [collate.DeepEqualer]. 'source' is enumerated immediately.
//
// [GroupBySelRes]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySelRes[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelResEq(source, keySelector, elementSelector, resultSelector, nil)
}

// GroupBySelResMust is like [GroupBySelRes] but panics in case of error.
func GroupBySelResMust[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result) Enumerable[Result] {
	return errorhelper.Must(GroupBySelRes(source, keySelector, elementSelector, resultSelector))
}

// [GroupBySelResEq] groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// Key values are compared using 'equaler' and the elements of each group are projected using a specified function.
// If 'equaler' is nil, [collate.DeepEqualer] is used. 'source' is enumerated immediately.
//
// [GroupBySelResEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.groupby
func GroupBySelResEq[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result, equaler collate.Equaler[Key]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	gg := GroupBySelEqMust(source, keySelector, elementSelector, equaler)
	return Select(gg, func(g Grouping[Key, Element]) Result {
		return resultSelector(g.key, &g)
	})
}

// GroupBySelResEqMust is like [GroupBySelResEq] but panics in case of error.
func GroupBySelResEqMust[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result, equaler collate.Equaler[Key]) Enumerable[Result] {
	return errorhelper.Must(GroupBySelResEq(source, keySelector, elementSelector, resultSelector, equaler))
}
