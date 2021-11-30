//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 9 - SelectMany
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-9-selectmany/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

// SelectMany projects each element of a sequence to an Enumerator
// and flattens the resulting sequences into one sequence.
func SelectMany[Source, Result any](source Enumerator[Source], selector func(Source) Enumerator[Result]) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	t := Empty[Result]()
	return OnFunc[Result]{
			mvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !source.MoveNext() {
						return false
					}
					t = selector(source.Current())
				}
			},
			crrnt: func() Result { return t.Current() },
			//		crrnt: t.Current, // yields wrong results
			rst: func() { t = Empty[Result](); source.Reset() },
		},
		nil
}

// SelectManyMust is like SelectMany but panics in case of error.
func SelectManyMust[Source, Result any](source Enumerator[Source], selector func(Source) Enumerator[Result]) Enumerator[Result] {
	r, err := SelectMany(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyIdx projects each element of a sequence and its index to an Enumerator
// and flattens the resulting sequences into one sequence.
func SelectManyIdx[Source, Result any](source Enumerator[Source], selector func(Source, int) Enumerator[Result]) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	i := -1
	t := Empty[Result]()
	return OnFunc[Result]{
			mvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !source.MoveNext() {
						return false
					}
					i++
					t = selector(source.Current(), i)
				}
			},
			crrnt: func() Result { return t.Current() },
			rst:   func() { i = -1; t = Empty[Result](); source.Reset() },
		},
		nil
}

// SelectManyIdxMust is like SelectManyIdx but panics in case of error.
func SelectManyIdxMust[Source, Result any](source Enumerator[Source], selector func(Source, int) Enumerator[Result]) Enumerator[Result] {
	r, err := SelectManyIdx(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyColl projects each element of a sequence to an Enumerator,
// flattens the resulting sequences into one sequence,
// and invokes a result selector function on each element therein.
func SelectManyColl[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source) Enumerator[Collection], resultSelector func(Source, Collection) Result) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if collectionSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	var e1 Source
	t := Empty[Collection]()
	return OnFunc[Result]{
			mvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !source.MoveNext() {
						return false
					}
					e1 = source.Current()
					t = collectionSelector(e1)
				}
			},
			crrnt: func() Result { return resultSelector(e1, t.Current()) },
			rst:   func() { t = Empty[Collection](); source.Reset() },
		},
		nil
}

// SelectManyCollMust is like SelectManyColl but panics in case of error.
func SelectManyCollMust[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source) Enumerator[Collection], resultSelector func(Source, Collection) Result) Enumerator[Result] {
	r, err := SelectManyColl(source, collectionSelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectManyCollIdx projects each element of a sequence and its index to an Enumerator,
// flattens the resulting sequences into one sequence,
// and invokes a result selector function on each element therein.
func SelectManyCollIdx[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source, int) Enumerator[Collection], resultSelector func(Source, Collection) Result) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if collectionSelector == nil || resultSelector == nil {
		return nil, ErrNilSelector
	}
	var e1 Source
	i := -1
	t := Empty[Collection]()
	return OnFunc[Result]{
			mvNxt: func() bool {
				for {
					if t.MoveNext() {
						return true
					}
					if !source.MoveNext() {
						return false
					}
					e1 = source.Current()
					i++
					t = collectionSelector(e1, i)
				}
			},
			crrnt: func() Result { return resultSelector(e1, t.Current()) },
			rst:   func() { i = -1; t = Empty[Collection](); source.Reset() },
		},
		nil
}

// SelectManyCollIdxMust is like SelectManyCollIdx but panics in case of error.
func SelectManyCollIdxMust[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source, int) Enumerator[Collection], resultSelector func(Source, Collection) Result) Enumerator[Result] {
	r, err := SelectManyCollIdx(source, collectionSelector, resultSelector)
	if err != nil {
		panic(err)
	}
	return r
}
