//go:build go1.18

package go2linq

import (
	"constraints"
)

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending

// OrderBySelf sorts the elements of a sequence in ascending order.
func OrderBySelf[Source constraints.Ordered](source Enumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OrderBy(source, Identity[Source])
}

// OrderBySelfMust is like OrderBySelf but panics in case of error.
func OrderBySelfMust[Source constraints.Ordered](source Enumerable[Source]) *OrderedEnumerable[Source] {
	r, err := OrderBySelf(source)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderBy sorts the elements of a sequence in ascending order according to a key.
func OrderBy[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return OrderByLs(source, keySelector, Lesser[Key](Order[Key]{}))
}

// OrderByMust is like OrderBy but panics in case of error.
func OrderByMust[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	r, err := OrderBy(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByLs sorts the elements of a sequence in ascending order using a specified lesser.
func OrderByLs[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return &OrderedEnumerable[Source]{
			source,
			projectionLesser(lesser, keySelector),
		},
		nil
}

// OrderByLsMust is like OrderByLs but panics in case of error.
func OrderByLsMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := OrderByLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByCmp sorts the elements of a sequence in ascending order using a specified comparer.
func OrderByCmp[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	var ls Lesser[Key] = ComparerFunc[Key](comparer.Compare)
	return OrderByLs(source, keySelector, ls)
}

// OrderByCmpMust is like OrderByCmp but panics in case of error.
func OrderByCmpMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	r, err := OrderByCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDescendingLs sorts the elements of a sequence in descending order using a specified lesser.
func OrderByDescendingLs[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return &OrderedEnumerable[Source]{
			source,
			reverseLesser(projectionLesser(lesser, keySelector)),
		},
		nil
}

// OrderByDescendingLsMust is like OrderByDescendingLs but panics in case of error.
func OrderByDescendingLsMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := OrderByDescendingLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDescendingCmp sorts the elements of a sequence in descending order using a specified comparer.
func OrderByDescendingCmp[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if comparer == nil {
		return nil, ErrNilComparer
	}
	var ls Lesser[Key] = ComparerFunc[Key](comparer.Compare)
	return OrderByDescendingLs(source, keySelector, ls)
}

// OrderByDescendingCmpMust is like OrderByDescendingCmp but panics in case of error.
func OrderByDescendingCmpMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	r, err := OrderByDescendingCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
