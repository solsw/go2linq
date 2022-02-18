//go:build go1.18

package go2linq

import (
	"golang.org/x/exp/constraints"
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
	return OrderByKey(source, Identity[Source])
}

// OrderBySelfMust is like OrderBySelf but panics in case of error.
func OrderBySelfMust[Source constraints.Ordered](source Enumerable[Source]) *OrderedEnumerable[Source] {
	r, err := OrderBySelf(source)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderBySelfLs sorts the elements of a sequence in ascending order using a specified lesser.
func OrderBySelfLs[Source any](source Enumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return OrderByKeyLs(source, Identity[Source], lesser)
}

// OrderBySelfLsMust is like OrderBySelfLs but panics in case of error.
func OrderBySelfLsMust[Source any](source Enumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := OrderBySelfLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByKey sorts the elements of a sequence in ascending order according to a key.
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

// OrderBySelfDesc sorts the elements of a sequence in descending order.
func OrderBySelfDesc[Source constraints.Ordered](source Enumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return OrderByKeyDesc(source, Identity[Source])
}

// OrderBySelfDescMust is like OrderBySelfDesc but panics in case of error.
func OrderBySelfDescMust[Source constraints.Ordered](source Enumerable[Source]) *OrderedEnumerable[Source] {
	r, err := OrderBySelfDesc(source)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderBySelfDescLs sorts the elements of a sequence in descending order using a specified lesser.
func OrderBySelfDescLs[Source any](source Enumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return OrderByKeyDescLs(source, Identity[Source], lesser)
}

// OrderBySelfDescLsMust is like OrderBySelfDescLs but panics in case of error.
func OrderBySelfDescLsMust[Source any](source Enumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := OrderBySelfDescLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// OrderByKeyDesc sorts the elements of a sequence in descending order according to a key.
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
