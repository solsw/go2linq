package go2linq

import (
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault

// Single returns the only element of a sequence, and returns an error if there is not exactly one element in the sequence.
// (https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single)
func Single[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			return generichelper.ZeroValue[Source](), ErrEmptySource
		}
		if counter.Count() > 1 {
			return generichelper.ZeroValue[Source](), ErrMultipleElements
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0), nil
		}
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	if enr.MoveNext() {
		return generichelper.ZeroValue[Source](), ErrMultipleElements
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
// (https://learn.microsoft.com/dotnet/api/system.linq.enumerable.single)
func SinglePred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
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
				return generichelper.ZeroValue[Source](), ErrMultipleMatch
			}
			found = true
			r = c
		}
	}
	if empty {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	if !found {
		return generichelper.ZeroValue[Source](), ErrNoMatch
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

// [SingleOrDefault] returns the only element of a sequence, or a [zero value] if the sequence is empty.
//
// [SingleOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func SingleOrDefault[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	r, err := Single(source)
	if err != nil {
		if err == ErrMultipleElements {
			return generichelper.ZeroValue[Source](), ErrMultipleElements
		}
		return generichelper.ZeroValue[Source](), nil
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

// [SingleOrDefaultPred] returns the only element of a sequence that satisfies a specified condition
// or a [zero value] if no such element exists.
//
// [SingleOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.singleordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func SingleOrDefaultPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	r, err := SinglePred(source, predicate)
	if err != nil {
		if err == ErrMultipleMatch {
			return generichelper.ZeroValue[Source](), ErrMultipleMatch
		}
		return generichelper.ZeroValue[Source](), nil
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
