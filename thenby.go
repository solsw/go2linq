//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

// ThenByLs performs a subsequent ordering of the elements in a sequence in ascending order using a specified lesser.
func ThenByLs[Source, Key any](source *OrderedEnumerable[Source],
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

// ThenByLsMust is like ThenByLs but panics in case of error.
func ThenByLsMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := ThenByLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByCmp performs a subsequent ordering of the elements in a sequence in ascending order using a specified comparer.
func ThenByCmp[Source, Key any](source *OrderedEnumerable[Source],
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
	return ThenByLs(source, keySelector, ls)
}

// ThenByCmpMust is like ThenByCmp but panics in case of error.
func ThenByCmpMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	r, err := ThenByCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByDescendingLs performs a subsequent ordering of the elements in a sequence in descending order using a specified lesser.
func ThenByDescendingLs[Source, Key any](source *OrderedEnumerable[Source],
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

// ThenByDescendingLsMust is like ThenByDescendingLs but panics in case of error.
func ThenByDescendingLsMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	r, err := ThenByDescendingLs(source, keySelector, lesser)
	if err != nil {
		panic(err)
	}
	return r
}

// ThenByDescendingCmp performs a subsequent ordering of the elements in a sequence in descending order using a specified comparer.
func ThenByDescendingCmp[Source, Key any](source *OrderedEnumerable[Source],
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
	return ThenByDescendingLs(source, keySelector, ls)
}

// ThenByDescendingCmpMust is like ThenByDescendingCmp but panics in case of error.
func ThenByDescendingCmpMust[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	r, err := ThenByDescendingCmp(source, keySelector, comparer)
	if err != nil {
		panic(err)
	}
	return r
}
