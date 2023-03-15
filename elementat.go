package go2linq

// Reimplementing LINQ to Objects: Part 31 – ElementAt / ElementAtOrDefault
// https://codeblog.jonskeet.uk/2011/01/11/reimplementing-linq-to-objects-part-31-elementat-elementatordefault/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault

// ElementAt returns the element at a specified index in a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat)
func ElementAt[Source any](source Enumerable[Source], index int) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if index < 0 {
		return ZeroValue[Source](), ErrIndexOutOfRange
	}
	if counter, ok := source.(Counter); ok {
		if index >= counter.Count() {
			return ZeroValue[Source](), ErrIndexOutOfRange
		}
		if itemer, ok := source.(Itemer[Source]); ok {
			return itemer.Item(index), nil
		}
	}
	i := 0
	enr := source.GetEnumerator()
	for enr.MoveNext() {
		if i == index {
			return enr.Current(), nil
		}
		i++
	}
	return ZeroValue[Source](), ErrIndexOutOfRange
}

// ElementAtMust is like [ElementAt] but panics in case of error.
func ElementAtMust[Source any](source Enumerable[Source], index int) Source {
	r, err := ElementAt(source, index)
	if err != nil {
		panic(err)
	}
	return r
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault)
func ElementAtOrDefault[Source any](source Enumerable[Source], index int) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	r, err := ElementAt(source, index)
	if err == ErrIndexOutOfRange {
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// ElementAtOrDefaultMust is like [ElementAtOrDefault] but panics in case of error.
func ElementAtOrDefaultMust[Source any](source Enumerable[Source], index int) Source {
	r, err := ElementAtOrDefault(source, index)
	if err != nil {
		panic(err)
	}
	return r
}
