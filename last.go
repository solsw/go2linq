//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

// Last returns the last element of a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last)
func Last[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		len := counter.Count()
		if len == 0 {
			return ZeroValue[Source](), ErrEmptySource
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(len - 1), nil
		}
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return ZeroValue[Source](), ErrEmptySource
	}
	r := enr.Current()
	for enr.MoveNext() {
		r = enr.Current()
	}
	return r, nil
}

// LastMust is like Last but panics in case of error.
func LastMust[Source any](source Enumerable[Source]) Source {
	r, err := Last(source)
	if err != nil {
		panic(err)
	}
	return r
}

// LastPred returns the last element of a sequence that satisfies a specified condition.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last)
func LastPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
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
	found := false
	var r Source
	c := enr.Current()
	if predicate(c) {
		found = true
		r = c
	}
	for enr.MoveNext() {
		c = enr.Current()
		if predicate(c) {
			found = true
			r = c
		}
	}
	if !found {
		return ZeroValue[Source](), ErrNoMatch
	}
	return r, nil
}

// LastPredMust is like LastPred but panics in case of error.
func LastPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	r, err := LastPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// LastOrDefault returns the last element of a sequence, or a default value if the sequence contains no elements.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault)
func LastOrDefault[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	r, err := Last(source)
	if err != nil {
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// LastOrDefaultMust is like LastOrDefault but panics in case of error.
func LastOrDefaultMust[Source any](source Enumerable[Source]) Source {
	r, err := LastOrDefault(source)
	if err != nil {
		panic(err)
	}
	return r
}

// LastOrDefaultPred returns the last element of a sequence that satisfies a condition
// or a default value if no such element is found.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault)
func LastOrDefaultPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return ZeroValue[Source](), ErrNilPredicate
	}
	r, err := LastPred(source, predicate)
	if err != nil {
		return ZeroValue[Source](), nil
	}
	return r, nil
}

// LastOrDefaultPredMust is like LastOrDefaultPred but panics in case of error.
func LastOrDefaultPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	r, err := LastOrDefaultPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
