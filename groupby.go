//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 21 - GroupByErr
// https://codeblog.jonskeet.uk/2011/01/01/reimplementing-linq-to-objects-part-21-groupby/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.groupby

// GroupBySelEq groups the elements of a sequence according to a key selector function.
// The keys are compared using an Equaler
// and each group's elements are projected using a specified function.
// If 'equaler' is nil DeepEqual is used. 'source' is enumerated immediately.
func GroupBySelEq[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler Equaler[Key]) (Enumerable[*Grouping[Key, Element]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	if equaler == nil {
		equaler = DeepEqual[Key]{}
	}
	lk := ToLookupSelEqMust(source, keySelector, elementSelector, equaler)
	return lk, nil
}

// GroupBySelEqMust is like GroupBySelEq but panics in case of error.
func GroupBySelEqMust[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element, equaler Equaler[Key]) Enumerable[*Grouping[Key, Element]] {
	r, err := GroupBySelEq(source, keySelector, elementSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelResEq groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// Key values are compared using a specified Equaler,
// and the elements of each group are projected using a specified function.
// If 'equaler' is nil DeepEqual is used. 'source' is enumerated immediately.
func GroupBySelResEq[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result, equaler Equaler[Key]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	grgr := GroupBySelEqMust(source, keySelector, elementSelector, equaler)
	return Select(grgr, func(gr *Grouping[Key, Element]) Result {
		return resultSelector(gr.key, gr)
	})
}

// GroupBySelResEqMust is like GroupBySelResEq but panics in case of error.
func GroupBySelResEqMust[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result, equaler Equaler[Key]) Enumerable[Result] {
	r, err := GroupBySelResEq(source, keySelector, elementSelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySelRes groups the elements of a sequence according to a specified
// key selector function and creates a result value from each group and its key.
// The elements of each group are projected using a specified function.
// Key values are compared using reflect.DeepEqual. 'source' is enumerated immediately.
func GroupBySelRes[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelResEq(source, keySelector, elementSelector, resultSelector, nil)
}

// GroupBySelResMust is like GroupBySelRes but panics in case of error.
func GroupBySelResMust[Source, Key, Element, Result any](source Enumerable[Source], keySelector func(Source) Key,
	elementSelector func(Source) Element, resultSelector func(Key, Enumerable[Element]) Result) Enumerable[Result] {
	r, err := GroupBySelRes(source, keySelector, elementSelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBySel groups the elements of a sequence according to a specified key selector function
// and projects the elements for each group using a specified function.
// The keys are compared using reflect.DeepEqual. 'source' is enumerated immediately.
func GroupBySel[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) (Enumerable[*Grouping[Key, Element]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || elementSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelEq(source, keySelector, elementSelector, nil)
}

// GroupBySelMust is like GroupBySel but panics in case of error.
func GroupBySelMust[Source, Key, Element any](source Enumerable[Source],
	keySelector func(Source) Key, elementSelector func(Source) Element) Enumerable[*Grouping[Key, Element]] {
	r, err := GroupBySel(source, keySelector, elementSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByRes groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using reflect.DeepEqual. 'source' is enumerated immediately.
func GroupByRes[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, nil)
}

// GroupByResMust is like GroupByRes but panics in case of error.
func GroupByResMust[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result) Enumerable[Result] {
	r, err := GroupByRes(source, keySelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByResEq groups the elements of a sequence according to a specified key selector function
// and creates a result value from each group and its key.
// The keys are compared using a specified Equaler.
// If 'equaler' is nil DeepEqual is used. 'source' is enumerated immediately.
func GroupByResEq[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result, equaler Equaler[Key]) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelResEq(source, keySelector, Identity[Source], resultSelector, equaler)
}

// GroupByResEqMust is like GroupByResEq but panics in case of error.
func GroupByResEqMust[Source, Key, Result any](source Enumerable[Source],
	keySelector func(Source) Key, resultSelector func(Key, Enumerable[Source]) Result, equaler Equaler[Key]) Enumerable[Result] {
	r, err := GroupByResEq(source, keySelector, resultSelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupBy groups the elements of a sequence according to a specified key selector function.
// The keys are compared using reflect.DeepEqual. 'source' is enumerated immediately.
func GroupBy[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) (Enumerable[*Grouping[Key, Source]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelEq(source, keySelector, Identity[Source], nil)
}

// GroupByMust is like GroupBy but panics in case of error.
func GroupByMust[Source, Key any](source Enumerable[Source], keySelector func(Source) Key) Enumerable[*Grouping[Key, Source]] {
	r, err := GroupBy(source, keySelector)
	if err != nil {
		panic(err)
	}
	return r
}

// GroupByEq groups the elements of a sequence according to a specified key selector function
// and compares the keys using a specified Equaler.
// If 'equaler' is nil DeepEqual is used. 'source' is enumerated immediately.
func GroupByEq[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, equaler Equaler[Key]) (Enumerable[*Grouping[Key, Source]], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if keySelector == nil {
		return nil, ErrNilSelector
	}
	return GroupBySelEq(source, keySelector, Identity[Source], equaler)
}

// GroupByEqMust is like GroupByEq but panics in case of error.
func GroupByEqMust[Source, Key any](source Enumerable[Source],
	keySelector func(Source) Key, equaler Equaler[Key]) Enumerable[*Grouping[Key, Source]] {
	r, err := GroupByEq(source, keySelector, equaler)
	if err != nil {
		panic(err)
	}
	return r
}
