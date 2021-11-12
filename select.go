//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 3 – "Select" (and a rename...)
// https://codeblog.jonskeet.uk/2010/12/23/reimplementing-linq-to-objects-part-3-quot-select-quot-and-a-rename/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.select

// Select projects each element of a sequence into a new form.
// Select panics if 'source' or 'selector' is nil.
func Select[Source, Result any](source Enumerator[Source], selector func(Source) Result) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	return OnFunc[Result]{
		MvNxt: func() bool { return source.MoveNext() },
		Crrnt: func() Result { return selector(source.Current()) },
		Rst: func() { source.Reset() },
	}
}

// SelectErr is like Select but returns an error instead of panicking.
func SelectErr[Source, Result any](source Enumerator[Source], selector func(Source) Result) (res Enumerator[Result], err error) {
	defer func() {
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return Select(source, selector), nil
}

// SelectIdx projects each element of a sequence into a new form by incorporating the element's index.
// SelectIdx panics if 'source' or 'selector' is nil.
func SelectIdx[Source, Result any](source Enumerator[Source], selector func(Source, int) Result) Enumerator[Result] {
	if source == nil {
		panic(ErrNilSource)
	}
	if selector == nil {
		panic(ErrNilSelector)
	}
	var i int = -1 // position before the first element
	return OnFunc[Result]{
		MvNxt: func() bool { i++; return source.MoveNext() },
		Crrnt: func() Result { return selector(source.Current(), i) },
		Rst: func() { i = -1; source.Reset() },
	}
}

// SelectIdxErr is like SelectIdx but returns an error instead of panicking.
func SelectIdxErr[Source, Result any](source Enumerator[Source], selector func(Source, int) Result) (res Enumerator[Result], err error) {
	defer func() { 
		catchPanic[Enumerator[Result]](recover(), &res, &err)
	}()
	return SelectIdx(source, selector), nil
}
