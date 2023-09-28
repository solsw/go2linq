package go2linq

import (
	"github.com/solsw/errorhelper"
)

// Reimplementing LINQ to Objects: Part 7 - Count and LongCount
// https://codeblog.jonskeet.uk/2010/12/26/reimplementing-linq-to-objects-part-7-count-and-longcount/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.longcount

// [Count] returns the number of elements in a sequence.
//
// [Count]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
func Count[Source any](source Enumerable[Source]) (int, error) {
	if source == nil {
		return -1, ErrNilSource
	}
	var c int
	if TryGetNonEnumeratedCountMust(source, &c) {
		return c, nil
	}
	enr := source.GetEnumerator()
	r := 0
	for enr.MoveNext() {
		r++
	}
	return r, nil
}

// CountMust is like [Count] but panics in case of error.
func CountMust[Source any](source Enumerable[Source]) int {
	return errorhelper.Must(Count(source))
}

// [CountPred] returns a number that represents how many elements in a specified sequence satisfy a condition.
//
// [CountPred]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.count
func CountPred[Source any](source Enumerable[Source], predicate func(Source) bool) (int, error) {
	if source == nil {
		return -1, ErrNilSource
	}
	if predicate == nil {
		return -1, ErrNilPredicate
	}
	enr := source.GetEnumerator()
	r := 0
	for enr.MoveNext() {
		if predicate(enr.Current()) {
			r++
		}
	}
	return r, nil
}

// CountPredMust is like [CountPred] but panics in case of error.
func CountPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) int {
	return errorhelper.Must(CountPred(source, predicate))
}
