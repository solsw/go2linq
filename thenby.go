//go:build go1.18

package go2linq

import (
	"golang.org/x/exp/constraints"
)

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

// ThenBy performs a subsequent ordering of the elements in a sequence in ascending order.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby)
func ThenBy[Source constraints.Ordered](source *OrderedEnumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return ThenByKey(source, Identity[Source])
}

// ThenByMust is like ThenBy but panics in case of error.
func ThenByMust[Source constraints.Ordered](source *OrderedEnumerable[Source]) *OrderedEnumerable[Source] {
	r, err := ThenBy(source)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByLs performs a subsequent ordering of the elements in a sequence in ascending order using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby)
func ThenByLs[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return ThenByKeyLs(source, Identity[Source], lesser)
}

// ThenByLsMust is like ThenByLs but panics in case of error.
func ThenByLsMust[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := ThenByLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByKey performs a subsequent ordering of the elements in a sequence in ascending order according to a key.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby)
func ThenByKey[Source any, Key constraints.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ThenByKeyLs(source, keySelector, Lesser[Key](Order[Key]{}))
}

// ThenByKeyMust is like ThenByKey but panics in case of error.
func ThenByKeyMust[Source any, Key constraints.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	r, err := ThenByKey(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByKeyLs performs a subsequent ordering of the elements in a sequence in ascending order using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby)
func ThenByKeyLs[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if source.ls == nil || lesser == nil {
		return nil, ErrNilLesser
	}
	return &OrderedEnumerable[Source]{
			source.en,
			compoundLesser(source.ls, projectionLesser(lesser, keySelector)),
		},
		nil
}

// ThenByKeyLsMust is like ThenByKeyLs but panics in case of error.
func ThenByKeyLsMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := ThenByKeyLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByDesc performs a subsequent ordering of the elements in a sequence in descending order.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending)
func ThenByDesc[Source constraints.Ordered](source *OrderedEnumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return ThenByKeyDesc(source, Identity[Source])
}

// ThenByDescMust is like ThenByDesc but panics in case of error.
func ThenByDescMust[Source constraints.Ordered](source *OrderedEnumerable[Source]) *OrderedEnumerable[Source] {
	r, err := ThenByDesc(source)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByDescLs sorts the elements of a sequence in descending order using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending)
func ThenByDescLs[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return ThenByKeyDescLs(source, Identity[Source], lesser)
}

// ThenByDescLsMust is like ThenByDescLs but panics in case of error.
func ThenByDescLsMust[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := ThenByDescLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByKeyDesc performs a subsequent ordering of the elements in a sequence in descending order according to a key.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending)
func ThenByKeyDesc[Source any, Key constraints.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ThenByKeyDescLs(source, keySelector, Lesser[Key](Order[Key]{}))
}

// ThenByKeyDescMust is like ThenByKeyDesc but panics in case of error.
func ThenByKeyDescMust[Source any, Key constraints.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	r, err := ThenByKeyDesc(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByKeyDescLs performs a subsequent ordering of the elements in a sequence in descending order using a specified lesser.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending)
func ThenByKeyDescLs[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	if source.ls == nil || lesser == nil {
		return nil, ErrNilLesser
	}
	return &OrderedEnumerable[Source]{
			source.en,
			compoundLesser(source.ls, reverseLesser(projectionLesser(lesser, keySelector))),
		},
		nil
}

// ThenByKeyDescLsMust is like ThenByKeyDescLs but panics in case of error.
func ThenByKeyDescLsMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := ThenByKeyDescLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}
