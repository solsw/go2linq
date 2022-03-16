//go:build go1.18

package go2linq

import (
	"golang.org/x/exp/constraints"
)

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending

// OrderBy sorts the elements of a sequence in ascending order.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby)
func OrderBy[Source constraints.Ordered](source Enumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OrderByKey(source, Identity[Source])
}

// OrderByMust is like OrderBy but panics in case of error.
func OrderByMust[Source constraints.Ordered](source Enumerable[Source]) *OrderedEnumerable[Source] {
	r, err := OrderBy(source)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByLs sorts the elements of a sequence in ascending order using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby)
func OrderByLs[Source any](source Enumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return OrderByKeyLs(source, Identity[Source], lesser)
}

// OrderByLsMust is like OrderByLs but panics in case of error.
func OrderByLsMust[Source any](source Enumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := OrderByLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByKey sorts the elements of a sequence in ascending order according to a key.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby)
func OrderByKey[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return OrderByKeyLs(source, keySelector, Lesser[Key](Order[Key]{}))
}

// OrderByKeyMust is like OrderByKey but panics in case of error.
func OrderByKeyMust[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	r, err := OrderByKey(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByKeyLs sorts the elements of a sequence in ascending order of keys using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby)
func OrderByKeyLs[Source, Key any](source Enumerable[Source],
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

// OrderByKeyLsMust is like OrderByKeyLs but panics in case of error.
func OrderByKeyLsMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := OrderByKeyLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDesc sorts the elements of a sequence in descending order.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending)
func OrderByDesc[Source constraints.Ordered](source Enumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OrderByKeyDesc(source, Identity[Source])
}

// OrderByDescMust is like OrderByDesc but panics in case of error.
func OrderByDescMust[Source constraints.Ordered](source Enumerable[Source]) *OrderedEnumerable[Source] {
	r, err := OrderByDesc(source)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByDescLs sorts the elements of a sequence in descending order using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending)
func OrderByDescLs[Source any](source Enumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return OrderByKeyDescLs(source, Identity[Source], lesser)
}

// OrderByDescLsMust is like OrderByDescLs but panics in case of error.
func OrderByDescLsMust[Source any](source Enumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := OrderByDescLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByKeyDesc sorts the elements of a sequence in descending order according to a key.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending)
func OrderByKeyDesc[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return OrderByKeyDescLs(source, keySelector, Lesser[Key](Order[Key]{}))
}

// OrderByKeyDescMust is like OrderByKeyDesc but panics in case of error.
func OrderByKeyDescMust[Source any, Key constraints.Ordered](source Enumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	r, err := OrderByKeyDesc(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByKeyDescLs sorts the elements of a sequence in descending order of keys using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending)
func OrderByKeyDescLs[Source, Key any](source Enumerable[Source],
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

// OrderByKeyDescLsMust is like OrderByKeyDescLs but panics in case of error.
func OrderByKeyDescLsMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := OrderByKeyDescLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}
