//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 31 – ElementAt / ElementAtOrDefault
// https://codeblog.jonskeet.uk/2011/01/11/reimplementing-linq-to-objects-part-31-elementat-elementatordefault/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault

// ElementAt returns the element at a specified index in a sequence.
// ElementAt panics if 'source' is nil or 'index' is out of range.
func ElementAt[Source any](source Enumerator[Source], index int) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if index < 0 {
		panic(ErrIndexOutOfRange)
	}
	if counter, ok := source.(Counter); ok {
		if index >= counter.Count() {
			panic(ErrIndexOutOfRange)
		}
		if itemer, ok := source.(Itemer[Source]); ok {
			return itemer.Item(index)
		}
	}
	i := 0
	for source.MoveNext() {
		if i == index {
			return source.Current()
		}
		i++
	}
	panic(ErrIndexOutOfRange)
}

// ElementAtErr is like ElementAt but returns an error instead of panicking.
func ElementAtErr[Source any](source Enumerator[Source], index int) (res Source, err error) {
	defer func() {
		catchPanic[Source](recover(), &res, &err)
	}()
	return ElementAt(source, index), nil
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
// ElementAtOrDefault panics if 'source' is nil.
func ElementAtOrDefault[Source any](source Enumerator[Source], index int) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	r, err := ElementAtErr(source, index)
	var s0 Source
	if err == ErrIndexOutOfRange {
		return s0
	}
	return r
}

// ElementAtOrDefaultErr is like ElementAtOrDefault but returns an error instead of panicking.
func ElementAtOrDefaultErr[Source any](source Enumerator[Source], index int) (res Source, err error) {
	defer func() {
		catchPanic[Source](recover(), &res, &err)
	}()
	return ElementAtOrDefault(source, index), nil
}
