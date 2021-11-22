//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 26b – OrderBy{,Descending}/ThenBy{,Descending}
// https://codeblog.jonskeet.uk/2011/01/05/reimplementing-linq-to-objects-part-26b-orderby-descending-thenby-descending/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.orderbydescending
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenby
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.thenbydescending

// OrderByLs sorts the elements of a sequence in ascending order by using a specified lesser.
// OrderByLs panics if 'source' or 'keySelector' or 'lesser' is nil.
func OrderByLs[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if lesser == nil {
		panic(ErrNilLesser)
	}
	return &OrderedEnumerable[Source]{
		source,
		projectionLesser(lesser, keySelector),
	}
}

// OrderByLsErr is like OrderByLs but returns an error instead of panicking.
func OrderByLsErr[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return OrderByLs(source, keySelector, lesser), nil
}

// OrderByCmp sorts the elements of a sequence in ascending order by using a specified comparer.
// OrderByCmp panics if 'source' or 'keySelector' or 'comparer' is nil.
func OrderByCmp[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if comparer == nil {
		panic(ErrNilComparer)
	}
	var ls Lesser[Key] = ComparerFunc[Key](comparer.Compare)
	return OrderByLs(source, keySelector, ls)
}

// OrderByCmpErr is like OrderByCmp but returns an error instead of panicking.
func OrderByCmpErr[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return OrderByCmp(source, keySelector, comparer), nil
}

// OrderByDescendingLs sorts the elements of a sequence in descending order by using a specified lesser.
// OrderByDescendingLs panics if 'source' or 'keySelector' or 'lesser' is nil.
func OrderByDescendingLs[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if lesser == nil {
		panic(ErrNilLesser)
	}
	return &OrderedEnumerable[Source]{
		source,
		reverseLesser(projectionLesser(lesser, keySelector)),
	}
}

// OrderByDescendingLsErr is like OrderByDescendingLs but returns an error instead of panicking.
func OrderByDescendingLsErr[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return OrderByDescendingLs(source, keySelector, lesser), nil
}

// OrderByDescendingCmp sorts the elements of a sequence in descending order by using a specified comparer.
// OrderByDescendingCmp panics if 'source' or 'keySelector' or 'comparer' is nil.
func OrderByDescendingCmp[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if comparer == nil {
		panic(ErrNilComparer)
	}
	var ls Lesser[Key] = ComparerFunc[Key](comparer.Compare)
	return OrderByDescendingLs(source, keySelector, ls)
}

// OrderByDescendingCmpErr is like OrderByDescendingCmp but returns an error instead of panicking.
func OrderByDescendingCmpErr[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return OrderByDescendingCmp(source, keySelector, comparer), nil
}

// ThenByLs performs a subsequent ordering of the elements in a sequence in ascending order by using a specified lesser.
// ThenByLs panics if 'source' or 'keySelector' or 'lesser' is nil.
func ThenByLs[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if source.ls == nil || lesser == nil {
		panic(ErrNilLesser)
	}
	return &OrderedEnumerable[Source]{
		source.en,
		compoundLesser(source.ls, projectionLesser(lesser, keySelector)),
	}
}

// ThenByLsErr is like ThenByLs but returns an error instead of panicking.
func ThenByLsErr[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return ThenByLs(source, keySelector, lesser), nil
}

// ThenByCmp performs a subsequent ordering of the elements in a sequence in ascending order by using a specified comparer.
// ThenByCmp panics if 'source' or 'keySelector' or 'comparer' is nil.
func ThenByCmp[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if comparer == nil {
		panic(ErrNilComparer)
	}
	var ls Lesser[Key] = ComparerFunc[Key](comparer.Compare)
	return ThenByLs(source, keySelector, ls)
}

// ThenByCmpErr is like ThenByCmp but returns an error instead of panicking.
func ThenByCmpErr[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return ThenByCmp(source, keySelector, comparer), nil
}

// ThenByDescendingLs performs a subsequent ordering of the elements in a sequence in descending order by using a specified lesser.
// ThenByDescendingLs panics if 'source' or 'keySelector' or 'lesser' is nil.
func ThenByDescendingLs[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if source.ls == nil || lesser == nil {
		panic(ErrNilLesser)
	}
	return &OrderedEnumerable[Source]{
		source.en,
		compoundLesser(source.ls, reverseLesser(projectionLesser(lesser, keySelector))),
	}
}

// ThenByDescendingLsErr is like ThenByDescendingLs but returns an error instead of panicking.
func ThenByDescendingLsErr[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, lesser Lesser[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return ThenByDescendingLs(source, keySelector, lesser), nil
}

// ThenByDescendingCmp performs a subsequent ordering of the elements in a sequence in descending order by using a specified comparer.
// ThenByDescendingCmp panics if 'source' or 'keySelector' or 'comparer' is nil.
func ThenByDescendingCmp[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) *OrderedEnumerable[Source] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	if comparer == nil {
		panic(ErrNilComparer)
	}
	var ls Lesser[Key] = ComparerFunc[Key](comparer.Compare)
	return ThenByDescendingLs(source, keySelector, ls)
}

// ThenByDescendingCmpErr is like ThenByDescendingCmp but returns an error instead of panicking.
func ThenByDescendingCmpErr[Source, Key any](source *OrderedEnumerable[Source],
	keySelector func(Source) Key, comparer Comparer[Key]) (res *OrderedEnumerable[Source], err error) {
	defer func() {
		catchPanic[*OrderedEnumerable[Source]](recover(), &res, &err)
	}()
	return ThenByDescendingCmp(source, keySelector, comparer), nil
}
