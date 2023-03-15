package go2linq

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

// Single returns the only element of a sequence, and returns an error if there is not exactly one element in the sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single)
func Single[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			return ZeroValue[Source](), ErrEmptySource
		}
		if counter.Count() > 1 {
			return ZeroValue[Source](), ErrMultipleElements
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0), nil
		}
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return ZeroValue[Source](), ErrEmptySource
	}
	if enr.MoveNext() {
		return ZeroValue[Source](), ErrMultipleElements
	}
	return enr.Current(), nil
}

// SingleMust is like [Single] but panics in case of error.
func SingleMust[Source any](source Enumerable[Source]) Source {
	r, err := Single(source)
	if err != nil {
		panic(err)
	}
	return r
}

// SinglePred returns the only element of a sequence that satisfies a specified condition.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single)
func SinglePred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return ZeroValue[Source](), ErrNilPredicate
	}
	enr := source.GetEnumerator()
	empty := true
	found := false
	var r Source
	for enr.MoveNext() {
		empty = false
		c := enr.Current()
		if predicate(c) {
			if found {
				return ZeroValue[Source](), ErrMultipleMatch
			}
			found = true
			r = c
		}
	}
	if empty {
		return ZeroValue[Source](), ErrEmptySource
	}
	if !found {
		return ZeroValue[Source](), ErrNoMatch
	}
	return r, nil
}

// SinglePredMust is like [SinglePred] but panics in case of error.
func SinglePredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	r, err := SinglePred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// SingleOrDefault returns the only element of a sequence, or a default value if the sequence is empty.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault)
func SingleOrDefault[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	r, err := Single(source)
	if err != nil {
		if err == ErrMultipleElements {
			return ZeroValue[Source](), ErrMultipleElements
		}
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// SingleOrDefaultMust is like [SingleOrDefault] but panics in case of error.
func SingleOrDefaultMust[Source any](source Enumerable[Source]) Source {
	r, err := SingleOrDefault(source)
	if err != nil {
		panic(err)
	}
	return r
}

// SingleOrDefaultPred returns the only element of a sequence that satisfies a specified condition
// or a default value if no such element exists.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault)
func SingleOrDefaultPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return ZeroValue[Source](), ErrNilPredicate
	}
	r, err := SinglePred(source, predicate)
	if err != nil {
		if err == ErrMultipleMatch {
			return ZeroValue[Source](), ErrMultipleMatch
		}
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// SingleOrDefaultPredMust is like [SingleOrDefaultPred] but panics in case of error.
func SingleOrDefaultPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	r, err := SingleOrDefaultPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
