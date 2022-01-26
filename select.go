//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 3 - "Select" (and a rename...)
// https://codeblog.jonskeet.uk/2010/12/23/reimplementing-linq-to-objects-part-3-quot-select-quot-and-a-rename/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

func enrSelect[Source, Result any](source Enumerable[Source], selector func(Source) Result) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enr := source.GetEnumerator()
		return enrFunc[Result]{
			mvNxt: func() bool { return enr.MoveNext() },
			crrnt: func() Result { return selector(enr.Current()) },
			rst:   func() { enr.Reset() },
		}
	}
}

// Select projects each element of a sequence into a new form.
func Select[Source, Result any](source Enumerable[Source], selector func(Source) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(enrSelect(source, selector)), nil
}

// SelectMust is like Select but panics in case of error.
func SelectMust[Source, Result any](source Enumerable[Source], selector func(Source) Result) Enumerable[Result] {
	r, err := Select(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

func enrSelectIdx[Source, Result any](source Enumerable[Source], selector func(Source, int) Result) func() Enumerator[Result] {
	return func() Enumerator[Result] {
		enr := source.GetEnumerator()
		i := -1 // position before the first element
		return enrFunc[Result]{
			mvNxt: func() bool { i++; return enr.MoveNext() },
			crrnt: func() Result { return selector(enr.Current(), i) },
			rst:   func() { i = -1; enr.Reset() },
		}
	}
}

// SelectIdx projects each element of a sequence into a new form by incorporating the element's index.
func SelectIdx[Source, Result any](source Enumerable[Source], selector func(Source, int) Result) (Enumerable[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return OnFactory(enrSelectIdx(source, selector)), nil
}

// SelectIdxMust is like SelectIdx but panics in case of error.
func SelectIdxMust[Source, Result any](source Enumerable[Source], selector func(Source, int) Result) Enumerable[Result] {
	r, err := SelectIdx(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}
