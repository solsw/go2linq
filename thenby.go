package go2linq

import (
	"cmp"

	"github.com/solsw/collate"
	"github.com/solsw/errorhelper"
)

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

// [ThenBy] performs a subsequent ordering of the elements in a sequence in ascending order.
//
// [ThenBy]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
func ThenBy[Source cmp.Ordered](source *OrderedEnumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return ThenByKey(source, Identity[Source])
}

// ThenByMust is like [ThenBy] but panics in case of error.
func ThenByMust[Source cmp.Ordered](source *OrderedEnumerable[Source]) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenBy(source))
}

// [ThenByLs] performs a subsequent ordering of the elements in a sequence in ascending order using a specified lesser.
//
// [ThenByLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
func ThenByLs[Source any](source *OrderedEnumerable[Source], lesser collate.Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return ThenByKeyLs(source, Identity[Source], lesser)
}

// ThenByLsMust is like [ThenByLs] but panics in case of error.
func ThenByLsMust[Source any](source *OrderedEnumerable[Source], lesser collate.Lesser[Source]) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByLs(source, lesser))
}

// [ThenByKey] performs a subsequent ordering of the elements in a sequence in ascending order according to a key.
//
// [ThenByKey]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
func ThenByKey[Source any, Key cmp.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ThenByKeyLs(source, keySelector, collate.Order[Key]{})
}

// ThenByKeyMust is like [ThenByKey] but panics in case of error.
func ThenByKeyMust[Source any, Key cmp.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByKey(source, keySelector))
}

// [ThenByKeyLs] performs a subsequent ordering of the elements in a sequence
// in ascending order according to a key using a specified lesser.
//
// [ThenByKeyLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenby
func ThenByKeyLs[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser collate.Lesser[Key]) (*OrderedEnumerable[Source], error) {
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

// ThenByKeyLsMust is like [ThenByKeyLs] but panics in case of error.
func ThenByKeyLsMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser collate.Lesser[Key]) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByKeyLs(source, keySelector, lesser))
}

// [ThenByDesc] performs a subsequent ordering of the elements in a sequence in descending order.
//
// [ThenByDesc]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending
func ThenByDesc[Source cmp.Ordered](source *OrderedEnumerable[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return ThenByDescKey(source, Identity[Source])
}

// ThenByDescMust is like [ThenByDesc] but panics in case of error.
func ThenByDescMust[Source cmp.Ordered](source *OrderedEnumerable[Source]) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByDesc(source))
}

// [ThenByDescLs] performs a subsequent ordering of the elements in a sequence in descending order using a specified lesser.
//
// [ThenByDescLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending
func ThenByDescLs[Source any](source *OrderedEnumerable[Source], lesser collate.Lesser[Source]) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if lesser == nil {
		return nil, ErrNilLesser
	}
	return ThenByDescKeyLs(source, Identity[Source], lesser)
}

// ThenByDescLsMust is like [ThenByDescLs] but panics in case of error.
func ThenByDescLsMust[Source any](source *OrderedEnumerable[Source], lesser collate.Lesser[Source]) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByDescLs(source, lesser))
}

// [ThenByDescKey] performs a subsequent ordering of the elements in a sequence in descending order according to a key.
//
// [ThenByDescKey]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending
func ThenByDescKey[Source any, Key cmp.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) (*OrderedEnumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return ThenByDescKeyLs(source, keySelector, collate.Order[Key]{})
}

// ThenByDescKeyMust is like [ThenByDescKey] but panics in case of error.
func ThenByDescKeyMust[Source any, Key cmp.Ordered](source *OrderedEnumerable[Source],
	keySelector func(Source) Key) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByDescKey(source, keySelector))
}

// [ThenByDescKeyLs] performs a subsequent ordering of the elements in a sequence
// in descending order according to a key using a specified lesser.
//
// [ThenByDescKeyLs]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending
func ThenByDescKeyLs[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser collate.Lesser[Key]) (*OrderedEnumerable[Source], error) {
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

// ThenByDescKeyLsMust is like [ThenByDescKeyLs] but panics in case of error.
func ThenByDescKeyLsMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser collate.Lesser[Key]) *OrderedEnumerable[Source] {
	return errorhelper.Must(ThenByDescKeyLs(source, keySelector, lesser))
}
