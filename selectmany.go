//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 9 – SelectMany
// https://codeblog.jonskeet.uk/2010/12/27/reimplementing-linq-to-objects-part-9-selectmany/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.selectmany

// SelectMany projects each element of a sequence to an Enumerator
// and flattens the resulting sequences into one sequence.
// SelectMany panics if 'source' or 'selector' is nil.
func SelectMany[Source, Result any](source Enumerator[Source],
	selector func(Source) Enumerator[Result]) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	t := Empty[Result]()
	return OnFunc[Result]{
		MvNxt: func() bool {
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
		Crrnt: func() Result { return t.Current() },
//		Crrnt: t.Current, // yields wrong results
		Rst:   func() { t = Empty[Result](); source.Reset() },
	}
}

// SelectManyErr is like SelectMany but returns an error instead of panicking.
func SelectManyErr[Source, Result any](source Enumerator[Source],
	selector func(Source) Enumerator[Result]) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return SelectMany(source, selector), nil
}

// SelectManyIdx projects each element of a sequence and its index to an Enumerator
// and flattens the resulting sequences into one sequence.
// SelectManyIdx panics if 'source' or 'selector' is nil.
func SelectManyIdx[Source, Result any](source Enumerator[Source],
	selector func(Source, int) Enumerator[Result]) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	i := -1
	t := Empty[Result]()
	return OnFunc[Result]{
		MvNxt: func() bool {
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
		Crrnt: func() Result { return t.Current() },
		Rst:   func() { i = -1; t = Empty[Result](); source.Reset() },
	}
}

// SelectManyIdxErr is like SelectManyIdx but returns an error instead of panicking.
func SelectManyIdxErr[Source, Result any](source Enumerator[Source],
	selector func(Source, int) Enumerator[Result]) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return SelectManyIdx(source, selector), nil
}

// SelectManyColl projects each element of a sequence to an Enumerator,
// flattens the resulting sequences into one sequence,
// and invokes a result selector function on each element therein.
// SelectManyColl panics if 'source' or 'collectionSelector' or 'resultSelector' is nil.
func SelectManyColl[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source) Enumerator[Collection],
	resultSelector func(Source, Collection) Result) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if collectionSelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	var e1 Source
	t := Empty[Collection]()
	return OnFunc[Result]{
		MvNxt: func() bool {
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
		Crrnt: func() Result { return resultSelector(e1, t.Current()) },
		Rst:   func() { t = Empty[Collection](); source.Reset() },
	}
}

// SelectManyCollErr is like SelectManyColl but returns an error instead of panicking.
func SelectManyCollErr[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source) Enumerator[Collection],
	resultSelector func(Source, Collection) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return SelectManyColl(source, collectionSelector, resultSelector), nil
}

// SelectManyCollIdx projects each element of a sequence and its index to an Enumerator,
// flattens the resulting sequences into one sequence,
// and invokes a result selector function on each element therein.
// SelectManyCollIdx panics if 'source' or 'collectionSelector' or 'resultSelector' is nil.
func SelectManyCollIdx[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source, int) Enumerator[Collection],
	resultSelector func(Source, Collection) Result) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if collectionSelector == nil || resultSelector == nil {
		panic(ErrNilSelector)
	}
	var e1 Source
	i := -1
	t := Empty[Collection]()
	return OnFunc[Result]{
		MvNxt: func() bool {
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
		Crrnt: func() Result { return resultSelector(e1, t.Current()) },
		Rst:   func() { i = -1; t = Empty[Collection](); source.Reset() },
	}
}

// SelectManyCollIdxErr is like SelectManyCollIdx but returns an error instead of panicking.
func SelectManyCollIdxErr[Source, Collection, Result any](source Enumerator[Source],
	collectionSelector func(Source, int) Enumerator[Collection],
	resultSelector func(Source, Collection) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return SelectManyCollIdx(source, collectionSelector, resultSelector), nil
}
