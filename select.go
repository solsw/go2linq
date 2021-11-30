//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 3 - "Select" (and a rename...)
// https://codeblog.jonskeet.uk/2010/12/23/reimplementing-linq-to-objects-part-3-quot-select-quot-and-a-rename/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

// Select projects each element of a sequence into a new form.
func Select[Source, Result any](source Enumerator[Source], selector func(Source) Result) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	return OnFunc[Result]{
			mvNxt: func() bool { return source.MoveNext() },
			crrnt: func() Result { return selector(source.Current()) },
			rst:   func() { source.Reset() },
		},
		nil
}

// SelectMust is like Select but panics in case of error.
func SelectMust[Source, Result any](source Enumerator[Source], selector func(Source) Result) Enumerator[Result] {
	r, err := Select(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}

// SelectIdx projects each element of a sequence into a new form by incorporating the element's index.
func SelectIdx[Source, Result any](source Enumerator[Source], selector func(Source, int) Result) (Enumerator[Result], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	if selector == nil {
		return nil, ErrNilSelector
	}
	var i int = -1 // position before the first element
	return OnFunc[Result]{
			mvNxt: func() bool { i++; return source.MoveNext() },
			crrnt: func() Result { return selector(source.Current(), i) },
			rst:   func() { i = -1; source.Reset() },
		},
		nil
}

// SelectIdxMust is like SelectIdx but panics in case of error.
func SelectIdxMust[Source, Result any](source Enumerator[Source], selector func(Source, int) Result) Enumerator[Result] {
	r, err := SelectIdx(source, selector)
	if err != nil {
		panic(err)
	}
	return r
}
