//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.first
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.single
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.last
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

// First returns the first element of a sequence.
func First[Source any](source Enumerator[Source]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			var s0 Source
			return s0, ErrEmptySource
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0), nil
		}
	}
	if !source.MoveNext() {
		var s0 Source
		return s0, ErrEmptySource
	}
	return source.Current(), nil
}

// FirstMust is like First but panics in case of error.
func FirstMust[Source any](source Enumerator[Source]) Source {
	r, err := First(source)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstPred returns the first element in a sequence that satisfies a specified condition.
func FirstPred[Source any](source Enumerator[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if predicate == nil {
		var s0 Source
		return s0, ErrNilPredicate
	}
	if !source.MoveNext() {
		var s0 Source
		return s0, ErrEmptySource
	}
	r := source.Current()
	if predicate(r) {
		return r, nil
	}
	for source.MoveNext() {
		r = source.Current()
		if predicate(r) {
			return r, nil
		}
	}
	var s0 Source
	return s0, ErrNoMatch
}

// FirstPredMust is like FirstPred but panics in case of error.
func FirstPredMust[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	r, err := FirstPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
func FirstOrDefault[Source any](source Enumerator[Source]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	r, err := First(source)
	if err != nil {
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// FirstOrDefaultMust is like FirstOrDefault but panics in case of error.
func FirstOrDefaultMust[Source any](source Enumerator[Source]) Source {
	r, err := FirstOrDefault(source)
	if err != nil {
		panic(err)
	}
	return r
}

// FirstOrDefaultPred returns the first element of the sequence that satisfies a condition or a default value if no such element is found.
func FirstOrDefaultPred[Source any](source Enumerator[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if predicate == nil {
		var s0 Source
		return s0, ErrNilPredicate
	}
	r, err := FirstPred(source, predicate)
	if err != nil {
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// FirstOrDefaultPredMust is like FirstOrDefaultPred but panics in case of error.
func FirstOrDefaultPredMust[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	r, err := FirstOrDefaultPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// Single returns the only element of a sequence, and panics if there is not exactly one element in the sequence.
func Single[Source any](source Enumerator[Source]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			var s0 Source
			return s0, ErrEmptySource
		}
		if counter.Count() > 1 {
			var s0 Source
			return s0, ErrMultipleElements
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0), nil
		}
	}
	if !source.MoveNext() {
		var s0 Source
		return s0, ErrEmptySource
	}
	if source.MoveNext() {
		var s0 Source
		return s0, ErrMultipleElements
	}
	return source.Current(), nil
}

// SingleMust is like Single but panics in case of error.
func SingleMust[Source any](source Enumerator[Source]) Source {
	r, err := Single(source)
	if err != nil {
		panic(err)
	}
	return r
}

// SinglePred returns the only element of a sequence that satisfies a specified condition.
func SinglePred[Source any](source Enumerator[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if predicate == nil {
		var s0 Source
		return s0, ErrNilPredicate
	}
	empty := true
	found := false
	var r Source
	for source.MoveNext() {
		empty = false
		c := source.Current()
		if predicate(c) {
			if found {
				var s0 Source
				return s0, ErrMultipleMatch
			}
			found = true
			r = c
		}
	}
	if empty {
		var s0 Source
		return s0, ErrEmptySource
	}
	if !found {
		var s0 Source
		return s0, ErrNoMatch
	}
	return r, nil
}

// SinglePredMust is like SinglePred but panics in case of error.
func SinglePredMust[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	r, err := SinglePred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// SingleOrDefault returns the only element of a sequence, or a default value if the sequence is empty.
func SingleOrDefault[Source any](source Enumerator[Source]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	r, err := Single(source)
	if err != nil {
		if err == ErrMultipleElements {
			var s0 Source
			return s0, ErrMultipleElements
		}
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// SingleOrDefaultMust is like SingleOrDefault but panics in case of error.
func SingleOrDefaultMust[Source any](source Enumerator[Source]) Source {
	r, err := SingleOrDefault(source)
	if err != nil {
		panic(err)
	}
	return r
}

// SingleOrDefaultPred returns the only element of a sequence that satisfies a specified condition
// or a default value if no such element exists.
func SingleOrDefaultPred[Source any](source Enumerator[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if predicate == nil {
		var s0 Source
		return s0, ErrNilPredicate
	}
	r, err := SinglePred(source, predicate)
	if err != nil {
		if err == ErrMultipleMatch {
			var s0 Source
			return s0, ErrMultipleMatch
		}
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// SingleOrDefaultPredMust is like SingleOrDefaultPred but panics in case of error.
func SingleOrDefaultPredMust[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	r, err := SingleOrDefaultPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// Last returns the last element of a sequence.
func Last[Source any](source Enumerator[Source]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		len := counter.Count()
		if len == 0 {
			var s0 Source
			return s0, ErrEmptySource
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(len - 1), nil
		}
	}
	if !source.MoveNext() {
		var s0 Source
		return s0, ErrEmptySource
	}
	r := source.Current()
	for source.MoveNext() {
		r = source.Current()
	}
	return r, nil
}

// LastMust is like Last but panics in case of error.
func LastMust[Source any](source Enumerator[Source]) Source {
	r, err := Last(source)
	if err != nil {
		panic(err)
	}
	return r
}

// LastPred returns the last element of a sequence that satisfies a specified condition.
func LastPred[Source any](source Enumerator[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if predicate == nil {
		var s0 Source
		return s0, ErrNilPredicate
	}
	if !source.MoveNext() {
		var s0 Source
		return s0, ErrEmptySource
	}
	found := false
	var r Source
	c := source.Current()
	if predicate(c) {
		found = true
		r = c
	}
	for source.MoveNext() {
		c = source.Current()
		if predicate(c) {
			found = true
			r = c
		}
	}
	if !found {
		var s0 Source
		return s0, ErrNoMatch
	}
	return r, nil
}

// LastPredMust is like LastPred but panics in case of error.
func LastPredMust[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	r, err := LastPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}

// LastOrDefault returns the last element of a sequence, or a default value if the sequence contains no elements.
func LastOrDefault[Source any](source Enumerator[Source]) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	r, err := Last(source)
	if err != nil {
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// LastOrDefaultMust is like LastOrDefault but panics in case of error.
func LastOrDefaultMust[Source any](source Enumerator[Source]) Source {
	r, err := LastOrDefault(source)
	if err != nil {
		panic(err)
	}
	return r
}

// LastOrDefaultPred returns the last element of a sequence that satisfies a condition
// or a default value if no such element is found.
func LastOrDefaultPred[Source any](source Enumerator[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		var s0 Source
		return s0, ErrNilSource
	}
	if predicate == nil {
		var s0 Source
		return s0, ErrNilPredicate
	}
	r, err := LastPred(source, predicate)
	if err != nil {
		var s0 Source
		return s0, nil
	}
	return r, nil
}

// LastOrDefaultPredMust is like LastOrDefaultPred but panics in case of error.
func LastOrDefaultPredMust[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	r, err := LastOrDefaultPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
