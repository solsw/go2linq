package go2linq

import (
	"github.com/solsw/errorhelper"
)

// Reimplementing LINQ to Objects: Part 10 - Any and All
// https://codeblog.jonskeet.uk/2010/12/28/reimplementing-linq-to-objects-part-10-any-and-all/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any

// [Any] determines whether a sequence contains any elements.
//
// [Any]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func Any[Source any](source Enumerable[Source]) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if counter, ok := source.(Counter); ok {
		return counter.Count() > 0, nil
	}
	return source.GetEnumerator().MoveNext(), nil
}

// AnyMust is like [Any] but panics in case of error.
func AnyMust[Source any](source Enumerable[Source]) bool {
	return errorhelper.Must(Any(source))
}

// [AnyPred] determines whether any element of a sequence satisfies a condition.
//
// [AnyPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.any
func AnyPred[Source any](source Enumerable[Source], predicate func(Source) bool) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if predicate == nil {
		return false, ErrNilPredicate
	}
	enr := source.GetEnumerator()
	for enr.MoveNext() {
		if predicate(enr.Current()) {
			return true, nil
		}
	}
	return false, nil
}

// AnyPredMust is like [AnyPred] but panics in case of error.
func AnyPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) bool {
	return errorhelper.Must(AnyPred(source, predicate))
}
