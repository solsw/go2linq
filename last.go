package go2linq

import (
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault

// [Last] returns the last element of a sequence.
//
// [Last]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func Last[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		len := counter.Count()
		if len == 0 {
			return generichelper.ZeroValue[Source](), ErrEmptySource
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(len - 1), nil
		}
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	r := enr.Current()
	for enr.MoveNext() {
		r = enr.Current()
	}
	return r, nil
}

// LastMust is like [Last] but panics in case of error.
func LastMust[Source any](source Enumerable[Source]) Source {
	return generichelper.Must(Last(source))
}

// [LastPred] returns the last element of a sequence that satisfies a specified condition.
//
// [LastPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.last
func LastPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return generichelper.ZeroValue[Source](), ErrEmptySource
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
		return generichelper.ZeroValue[Source](), ErrNoMatch
	}
	return r, nil
}

// LastPredMust is like [LastPred] but panics in case of error.
func LastPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	return generichelper.Must(LastPred(source, predicate))
}

// [LastOrDefault] returns the last element of a sequence or a [zero value] if the sequence contains no elements.
//
// [LastOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func LastOrDefault[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	r, err := Last(source)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}

// LastOrDefaultMust is like [LastOrDefault] but panics in case of error.
func LastOrDefaultMust[Source any](source Enumerable[Source]) Source {
	return generichelper.Must(LastOrDefault(source))
}

// [LastOrDefaultPred] returns the last element of a sequence that satisfies a condition
// or a [zero value] if no such element is found.
//
// [LastOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.lastordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func LastOrDefaultPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	r, err := LastPred(source, predicate)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}

// LastOrDefaultPredMust is like [LastOrDefaultPred] but panics in case of error.
func LastOrDefaultPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	return generichelper.Must(LastOrDefaultPred(source, predicate))
}
