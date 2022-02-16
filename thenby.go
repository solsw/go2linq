//go:build go1.18

package go2linq

import (
	"constraints"
)

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

// ThenBySelf performs a subsequent ordering of the elements in a sequence in ascending order.
func ThenBySelf[Source constraints.Ordered](source *OrderedEnumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return ThenByKey(source, Identity[Source])
}

// ThenBySelfMust is like ThenBySelf but panics in case of error.
func ThenBySelfMust[Source constraints.Ordered](source *OrderedEnumerable[Source]) *OrderedEnumerable[Source] {
	r, err := ThenBySelf(source)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenBySelfLs performs a subsequent ordering of the elements in a sequence in ascending order using a specified lesser.
func ThenBySelfLs[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return ThenByKeyLs(source, Identity[Source], lesser)
}

// ThenBySelfLsMust is like ThenBySelfLs but panics in case of error.
func ThenBySelfLsMust[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := ThenBySelfLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByKey performs a subsequent ordering of the elements in a sequence in ascending order according to a key.
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

// ThenBySelfDesc performs a subsequent ordering of the elements in a sequence in descending order.
func ThenBySelfDesc[Source constraints.Ordered](source *OrderedEnumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return ThenByKeyDesc(source, Identity[Source])
}

// ThenBySelfDescMust is like ThenBySelfDesc but panics in case of error.
func ThenBySelfDescMust[Source constraints.Ordered](source *OrderedEnumerable[Source]) *OrderedEnumerable[Source] {
	r, err := ThenBySelfDesc(source)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenBySelfDescLs sorts the elements of a sequence in descending order using a specified lesser.
func ThenBySelfDescLs[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return ThenByKeyDescLs(source, Identity[Source], lesser)
}

// ThenBySelfDescLsMust is like ThenBySelfDescLs but panics in case of error.
func ThenBySelfDescLsMust[Source any](source *OrderedEnumerable[Source], lesser Lesser[Source]) *OrderedEnumerable[Source] {
	r, err := ThenBySelfDescLs(source, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByKeyDesc performs a subsequent ordering of the elements in a sequence in descending order according to a key.
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
