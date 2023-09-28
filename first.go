package go2linq

import (
	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 11 - First/Single/Last and the …OrDefault versions
// https://codeblog.jonskeet.uk/2010/12/29/reimplementing-linq-to-objects-part-11-first-single-last-and-the-ordefault-versions/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault

// [First] returns the first element of a sequence.
//
// [First]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
func First[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if counter, cok := source.(Counter); cok {
		if counter.Count() == 0 {
			return generichelper.ZeroValue[Source](), ErrEmptySource
		}
		if itemer, iok := source.(Itemer[Source]); iok {
			return itemer.Item(0), nil
		}
	}
	enr := source.GetEnumerator()
	if !enr.MoveNext() {
		return generichelper.ZeroValue[Source](), ErrEmptySource
	}
	return enr.Current(), nil
}

// FirstMust is like [First] but panics in case of error.
func FirstMust[Source any](source Enumerable[Source]) Source {
	return errorhelper.Must(First(source))
}

// [FirstPred] returns the first element in a sequence that satisfies a specified condition.
//
// [FirstPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.first
func FirstPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
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
	return generichelper.ZeroValue[Source](), ErrNoMatch
}

// FirstPredMust is like [FirstPred] but panics in case of error.
func FirstPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	return errorhelper.Must(FirstPred(source, predicate))
}

// [FirstOrDefault] returns the first element of a sequence, or a [zero value] if the sequence contains no elements.
//
// [FirstOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func FirstOrDefault[Source any](source Enumerable[Source]) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	r, err := First(source)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}

// FirstOrDefaultMust is like [FirstOrDefault] but panics in case of error.
func FirstOrDefaultMust[Source any](source Enumerable[Source]) Source {
	return errorhelper.Must(FirstOrDefault(source))
}

// [FirstOrDefaultPred] returns the first element of the sequence that satisfies a condition
// or a [zero value] if no such element is found.
//
// [FirstOrDefaultPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.firstordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func FirstOrDefaultPred[Source any](source Enumerable[Source], predicate func(Source) bool) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if predicate == nil {
		return generichelper.ZeroValue[Source](), ErrNilPredicate
	}
	r, err := FirstPred(source, predicate)
	if err != nil {
		return generichelper.ZeroValue[Source](), nil
	}
	return r, nil
}

// FirstOrDefaultPredMust is like [FirstOrDefaultPred] but panics in case of error.
func FirstOrDefaultPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) Source {
	return errorhelper.Must(FirstOrDefaultPred(source, predicate))
}
