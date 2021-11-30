//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 31 – ElementAt / ElementAtOrDefault
// https://codeblog.jonskeet.uk/2011/01/11/reimplementing-linq-to-objects-part-31-elementat-elementatordefault/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementat
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault

// ElementAt returns the element at a specified index in a sequence.
func ElementAt[Source any](source Enumerator[Source], index int) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if index < 0 {
		var s0 Source
		return s0, ErrIndexOutOfRange
	}
	if counter, ok := source.(Counter); ok {
		if index >= counter.Count() {
			var s0 Source
			return s0, ErrIndexOutOfRange
		}
		if itemer, ok := source.(Itemer[Source]); ok {
			return itemer.Item(index), nil
		}
	}
	i := 0
	for source.MoveNext() {
		if i == index {
			return source.Current(), nil
		}
		i++
	}
	var s0 Source
	return s0, ErrIndexOutOfRange
}

// ElementAtMust is like ElementAt but panics in case of error.
func ElementAtMust[Source any](source Enumerator[Source], index int) Source {
	r, err := ElementAt(source, index)
	if err != nil {
		panic(err)
	}
	return r
}

// ElementAtOrDefault returns the element at a specified index in a sequence or a default value if the index is out of range.
func ElementAtOrDefault[Source any](source Enumerator[Source], index int) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	r, err := ElementAt(source, index)
	if err == ErrIndexOutOfRange {
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// ElementAtOrDefaultMust is like ElementAtOrDefault but panics in case of error.
func ElementAtOrDefaultMust[Source any](source Enumerator[Source], index int) Source {
	r, err := ElementAtOrDefault(source, index)
	if err != nil {
		panic(err)
	}
	return r
}
