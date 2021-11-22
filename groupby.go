//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 21 - GroupByErr
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-21-groupby/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

// GroupBySelEq groups the elements of a sequence according to a key selector function.
// The keys are compared by using an Equaler
// and each group's elements are projected by using a specified function.
// If 'eq' is nil reflect.DeepEqual is used. 'source' is enumerated immediately.
// GroupBySelEq panics if 'source' or 'keySelector' or 'elementSelector' is nil.
func GroupBySelEq[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, eq Equaler[Key]) Enumerator[Grouping[Key, Element]] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		panic(ErrNilSelector)
	}
	if eq == nil {
		eq = EqualerFunc[Key](DeepEqual[Key])
	}
	lk := ToLookupSelEq(source, keySelector, elementSelector, eq)
	return lk.GetEnumerator()
}

// GroupBySelEqErr is like GroupBySelEq but returns an error instead of panicking.
func GroupBySelEqErr[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, eq Equaler[Key]) (res Enumerator[Grouping[Key, Element]], err error) {
	defer func() {
		catchPanic[Enumerator[Grouping[Key, Element]]](recover(), &res, &err)
	}()
	return GroupBySelEq(source, keySelector, elementSelector, eq), nil
}

// GroupBySelResEq groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// Key values are compared by using a specified Equaler,
// and the elements of each group are projected by using a specified function.
// If 'eq' is nil reflect.DeepEqual is used. 'source' is enumerated immediately.
// GroupBySelResEq panics if 'source' or 'keySelector' or 'elementSelector' or 'resultSelector' is nil.
func GroupBySelResEq[Source, Key, Element, Result any](source Enumerator[Source], keySelector func(Source) Key,
  elementSelector func(Source) Element, resultSelector func(Key, Enumerator[Element]) Result, eq Equaler[Key]) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	grgr := GroupBySelEq(source, keySelector, elementSelector, eq)
	return Select(grgr, func(gr Grouping[Key,Element]) Result {
		return resultSelector(gr.key, gr.GetEnumerator())
	})
}

// GroupBySelResEqErr is like GroupBySelResEq but returns an error instead of panicking.
func GroupBySelResEqErr[Source, Key, Element, Result any](source Enumerator[Source], keySelector func(Source) Key,
  elementSelector func(Source) Element, resultSelector func(Key, Enumerator[Element]) Result, eq Equaler[Key]) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupBySelResEq(source, keySelector, elementSelector, resultSelector, eq), nil
}

// GroupBySelRes groups the elements of a sequence according to a specified
// key selector function and creates a result value from each group and its key.
// The elements of each group are projected by using a specified function.
// Key values are compared by using reflect.DeepEqual. 'source' is enumerated immediately.
// GroupBySelRes panics if 'source' or 'keySelector' or 'elementSelector' or 'resultSelector' is nil.
func GroupBySelRes[Source, Key, Element, Result any](source Enumerator[Source], keySelector func(Source) Key,
  elementSelector func(Source) Element, resultSelector func(Key, Enumerator[Element]) Result) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	return GroupBySelResEq(source, keySelector, elementSelector, resultSelector, nil)
}

// GroupBySelResErr is like GroupBySelRes but returns an error instead of panicking.
func GroupBySelResErr[Source, Key, Element, Result any](source Enumerator[Source], keySelector func(Source) Key,
  elementSelector func(Source) Element, resultSelector func(Key, Enumerator[Element]) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupBySelRes(source, keySelector, elementSelector, resultSelector), nil
}

// GroupBySel groups the elements of a sequence according to a specified key selector function
// and projects the elements for each group by using a specified function.
// The keys are compared by using reflect.DeepEqual. 'source' is enumerated immediately.
// GroupBySel panics if 'source' or 'keySelector' or 'elementSelector' is nil.
func GroupBySel[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) Enumerator[Grouping[Key, Element]] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || elementSelector == nil {
		panic(ErrNilSelector)
	}
	return GroupBySelEq(source, keySelector, elementSelector, nil)
}

// GroupBySelErr is like GroupBySel but returns an error instead of panicking.
func GroupBySelErr[Source, Key, Element any](source Enumerator[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (res Enumerator[Grouping[Key, Element]], err error) {
	defer func() {
		catchPanic[Enumerator[Grouping[Key, Element]]](recover(), &res, &err)
	}()
	return GroupBySel(source, keySelector, elementSelector), nil
}

// GroupByRes groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared by using reflect.DeepEqual. 'source' is enumerated immediately.
// GroupByRes panics if 'source' or 'keySelector' or 'resultSelector' is nil.
func GroupByRes[Source, Key, Result any](source Enumerator[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerator[Source]) Result) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	return GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, nil)
}

// GroupByResErr is like GroupByRes but returns an error instead of panicking.
func GroupByResErr[Source, Key, Result any](source Enumerator[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerator[Source]) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupByRes(source, keySelector, resultSelector), nil
}

// GroupByResEq groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared by using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used. 'source' is enumerated immediately.
// GroupByResEq panics if 'source' or 'keySelector' or 'resultSelector' is nil.
func GroupByResEq[Source, Key, Result any](source Enumerator[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerator[Source]) Result, eq Equaler[Key]) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	return GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, eq)
}

// GroupByResEqErr is like GroupByResEq but returns an error instead of panicking.
func GroupByResEqErr[Source, Key, Result any](source Enumerator[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerator[Source]) Result, eq Equaler[Key]) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return GroupByResEq(source, keySelector, resultSelector, eq), nil
}

// GroupBy groups the elements of a sequence according to a specified key selector function.
// The keys are compared by using reflect.DeepEqual. 'source' is enumerated immediately.
// GroupBy panics if 'source' or 'keySelector' is nil.
func GroupBy[Source, Key any](source Enumerator[Source], keySelector func(Source) Key) Enumerator[Grouping[Key, Source]] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	return GroupBySelEq(source, keySelector, Identity[Source], nil)
}

// GroupByErr is like GroupBy but returns an error instead of panicking.
func GroupByErr[Source, Key any](source Enumerator[Source], keySelector func(Source) Key) (res Enumerator[Grouping[Key, Source]], err error) {
	defer func() {
		catchPanic[Enumerator[Grouping[Key, Source]]](recover(), &res, &err)
	}()
	return GroupBy(source, keySelector), nil
}

// GroupByEq groups the elements of a sequence according to a specified key selector function
// and compares the keys by using a specified Equaler.
// If 'eq' is nil reflect.DeepEqual is used. 'source' is enumerated immediately.
// GroupByEq panics if 'source' or 'keySelector' is nil.
func GroupByEq[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, eq Equaler[Key]) Enumerator[Grouping[Key, Source]] {
	if source == nil {
		panic(ErrNilSource)
	}
	if keySelector == nil {
		panic(ErrNilSelector)
	}
	return GroupBySelEq(source, keySelector, Identity[Source], eq)
}

// GroupByEqErr is like GroupByEq but returns an error instead of panicking.
func GroupByEqErr[Source, Key any](source Enumerator[Source],
	keySelector func(Source) Key, eq Equaler[Key]) (res Enumerator[Grouping[Key, Source]], err error) {
	defer func() {
		catchPanic[Enumerator[Grouping[Key, Source]]](recover(), &res, &err)
	}()
	return GroupByEq(source, keySelector, eq), nil
}
