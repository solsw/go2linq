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
// First panics if 'source' is nil or the source sequence is empty.
func First[Source any](source Enumerator[Source]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			panic(ErrEmptySource)
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0)
		}
	}
	if !source.MoveNext() {
		panic(ErrEmptySource)
	}
	return source.Current()
}

// FirstErr is like First but returns an error instead of panicking.
func FirstErr[Source any](source Enumerator[Source]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return First(source), nil
}

// FirstPred returns the first element in a sequence that satisfies a specified condition.
// FirstPred panics if 'source' or 'predicate' is nil, no element satisfies the condition in 'predicate' or the source sequence is empty.
func FirstPred[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	if !source.MoveNext() {
		panic(ErrEmptySource)
	}
	r := source.Current()
	if predicate(r) {
		return r
	}
	for source.MoveNext() {
		r = source.Current()
		if predicate(r) {
			return r
		}
	}
	panic(ErrNoMatch)
}

// FirstPredErr is like FirstPred but returns an error instead of panicking.
func FirstPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return FirstPred(source, predicate), nil
}

// FirstOrDefault returns the first element of a sequence, or a default value if the sequence contains no elements.
// FirstOrDefault panics if 'source' is nil.
func FirstOrDefault[Source any](source Enumerator[Source]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	r, err := FirstErr(source)
	if err != nil {
		var s0 Source
		return s0
	}
	return r
}

// FirstOrDefaultErr is like FirstOrDefault but returns an error instead of panicking.
func FirstOrDefaultErr[Source any](source Enumerator[Source]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return FirstOrDefault(source), nil
}

// FirstOrDefaultPred returns the first element of the sequence that satisfies a condition or a default value if no such element is found.
// FirstOrDefaultPred panics if 'source' or 'predicate' is nil.
func FirstOrDefaultPred[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	r, err := FirstPredErr(source, predicate)
	if err != nil {
		var s0 Source
		return s0
	}
	return r
}

// FirstOrDefaultPredErr is like FirstOrDefaultPred but returns an error instead of panicking.
func FirstOrDefaultPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return FirstOrDefaultPred(source, predicate), nil
}

// Single returns the only element of a sequence, and panics if there is not exactly one element in the sequence.
// Single panics if 'source' is nil, the input sequence contains more than one element or the input sequence is empty.
func Single[Source any](source Enumerator[Source]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			panic(ErrEmptySource)
		}
		if counter.Count() > 1 {
			panic(ErrMultipleElements)
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0)
		}
	}
	if !source.MoveNext() {
		panic(ErrEmptySource)
	}
	if source.MoveNext() {
		panic(ErrMultipleElements)
	}
	return source.Current()
}

// SingleErr is like Single but returns an error instead of panicking.
func SingleErr[Source any](source Enumerator[Source]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return Single(source), nil
}

// SinglePred returns the only element of a sequence that satisfies a specified condition.
// SinglePred panics if 'source' or 'predicate' is nil, no element satisfies the condition in 'predicate',
// more than one element satisfies the condition in 'predicate' or the source sequence is empty.
func SinglePred[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	empty := true
	found := false
	var r Source
	for source.MoveNext() {
		empty = false
		c := source.Current()
		if predicate(c) {
			if found {
				panic(ErrMultipleMatch)
			}
			found = true
			r = c
		}
	}
	if empty {
		panic(ErrEmptySource)
	}
	if !found {
		panic(ErrNoMatch)
	}
	return r
}

// SinglePredErr is like SinglePred but returns an error instead of panicking.
func SinglePredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return SinglePred(source, predicate), nil
}

// SingleOrDefault returns the only element of a sequence, or a default value if the sequence is empty.
// SingleOrDefault panics if 'source' is nil or the input sequence contains more than one element.
func SingleOrDefault[Source any](source Enumerator[Source]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	r, err := SingleErr(source)
	if err != nil {
		if err == ErrMultipleElements {
			panic(ErrMultipleElements)
		}
		var s0 Source
		return s0
	}
	return r
}

// SingleOrDefaultErr is like SingleOrDefault but returns an error instead of panicking.
func SingleOrDefaultErr[Source any](source Enumerator[Source]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return SingleOrDefault(source), nil
}

// SingleOrDefaultPred returns the only element of a sequence that satisfies a specified condition
// or a default value if no such element exists.
// SingleOrDefaultPred panics if 'source' or 'predicate' is nil
// or more than one element satisfies the condition in 'predicate'.
func SingleOrDefaultPred[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	r, err := SinglePredErr(source, predicate)
	if err != nil {
		if err == ErrMultipleMatch {
			panic(ErrMultipleMatch)
		}
		var s0 Source
		return s0
	}
	return r
}

// SingleOrDefaultPredErr is like SingleOrDefaultPred but returns an error instead of panicking.
func SingleOrDefaultPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return SingleOrDefaultPred(source, predicate), nil
}

// Last returns the last element of a sequence.
// Last panics if 'source' is nil or the source sequence is empty.
func Last[Source any](source Enumerator[Source]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if counter, cok := source.(Counter); cok {
		len := counter.Count()
		if len == 0 {
			panic(ErrEmptySource)
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(len - 1)
		}
	}
	if !source.MoveNext() {
		panic(ErrEmptySource)
	}
	r := source.Current()
	for source.MoveNext() {
		r = source.Current()
	}
	return r
}

// LastErr is like Last but returns an error instead of panicking.
func LastErr[Source any](source Enumerator[Source]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return Last(source), nil
}

// LastPred returns the last element of a sequence that satisfies a specified condition.
// LastPred panics if 'source' or 'predicate' is nil, no element satisfies the condition in 'predicate' or the source sequence is empty.
func LastPred[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	if !source.MoveNext() {
		panic(ErrEmptySource)
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
		panic(ErrNoMatch)
	}
	return r
}

// LastPredErr is like LastPred but returns an error instead of panicking.
func LastPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return LastPred(source, predicate), nil
}

// LastOrDefault returns the last element of a sequence, or a default value if the sequence contains no elements.
// LastOrDefault panics if 'source' is nil.
func LastOrDefault[Source any](source Enumerator[Source]) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	r, err := LastErr(source)
	if err != nil {
		var s0 Source
		return s0
	}
	return r
}

// LastOrDefaultErr is like LastOrDefault but returns an error instead of panicking.
func LastOrDefaultErr[Source any](source Enumerator[Source]) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return LastOrDefault(source), nil
}

// LastOrDefaultPred returns the last element of a sequence that satisfies a condition
// or a default value if no such element is found.
// LastOrDefaultPred panics if 'source' or 'predicate' is nil.
func LastOrDefaultPred[Source any](source Enumerator[Source], predicate func(Source) bool) Source {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	r, err := LastPredErr(source, predicate)
	if err != nil {
		var s0 Source
		return s0
	}
	return r
}

// LastOrDefaultPredErr is like LastOrDefaultPred but returns an error instead of panicking.
func LastOrDefaultPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res Source, err error) {
	defer func() {
		catchErrPanic[Source](recover(), &res, &err)
	}()
	return LastOrDefaultPred(source, predicate), nil
}
