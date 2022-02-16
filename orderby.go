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

// OrderByDescendingSelf sorts the elements of a sequence in descending order.
func OrderByDescendingSelf[Source constraints.Ordered](source Enumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OrderByDescending(source, Identity[Source])
}

// OrderByDescendingSelfMust is like OrderByDescendingSelf but panics in case of error.
func OrderByDescendingSelfMust[Source constraints.Ordered](source Enumerable[Source]) *OrderedEnumerable[Source] {
	r, err := OrderByDescendingSelf(source)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDescending sorts the elements of a sequence in descending order according to a key.
func OrderByDescending[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return OrderByDescendingLs(source, keySelector, Lesser[Key](Order[Key]{}))
}

// OrderByDescendingMust is like OrderByDescending but panics in case of error.
func OrderByDescendingMust[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	r, err := OrderByDescending(source, keySelector)
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
