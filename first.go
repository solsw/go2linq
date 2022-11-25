package go2linq

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault

// First returns the first element of a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first)
func First[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			return ZeroValue[Source](), ErrEmptySource
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0), nil
		}
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return ZeroValue[Source](), ErrEmptySource
	}
	return enr.Current(), nil
}

// FirstMust is like First but panics in case of error.
func FirstMust[Source any](source Enumerable[Source]) Source {
	r, err := First(source)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstPred returns the first element in a sequence that satisfies a specified condition.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first)
func FirstPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return ZeroValue[Source](), ErrNilPredicate
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return ZeroValue[Source](), ErrEmptySource
	}
	r := enr.Current()
	if predicate(r) {
		return r, nil
	}
	for enr.MoveNext() {
		r = enr.Current()
		if predicate(r) {
			return r, nil
		}
	}
	return ZeroValue[Source](), ErrNoMatch
}

// FirstPredMust is like FirstPred but panics in case of error.
func FirstPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	r, err := FirstPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault)
func FirstOrDefault[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	r, err := First(source)
	if err != nil {
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// FirstOrDefaultMust is like FirstOrDefault but panics in case of error.
func FirstOrDefaultMust[Source any](source Enumerable[Source]) Source {
	r, err := FirstOrDefault(source)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstOrDefaultPred returns the first element of the sequence that satisfies a condition
// or a default value if no such element is found.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault)
func FirstOrDefaultPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return ZeroValue[Source](), ErrNilPredicate
	}
	r, err := FirstPred(source, predicate)
	if err != nil {
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// FirstOrDefaultPredMust is like FirstOrDefaultPred but panics in case of error.
func FirstOrDefaultPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	r, err := FirstOrDefaultPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
