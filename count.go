//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 7 - Count and LongCount
// https://codeblog.jonskeet.uk/2010/12/26/reimplementing-linq-to-objects-part-7-count-and-longcount/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount

// Count returns the number of elements in a sequence.
func Count[Source any](source Enumerator[Source]) (int, error) {
	if source == nil {
		return -1, ErrNilSource
	}
	var c int
	if TryGetNonEnumeratedCountMust(source, &c) {
		return c, nil
	}
	r := 0
	for source.MoveNext() {
		r++
	}
	return r, nil
}

// CountMust is like Count but panics in case of error.
func CountMust[Source any](source Enumerator[Source]) int {
	r, err := Count(source)
	if err != nil {
		panic(err)
	}
	return r
}

// CountPred returns a number that represents how many elements in the specified sequence satisfy a condition.
func CountPred[Source any](source Enumerator[Source], predicate func(Source) bool) (int, error) {
	if source == nil {
		return -1, ErrNilSource
	}
	if predicate == nil {
		return -1, ErrNilPredicate
	}
	r := 0
	for source.MoveNext() {
		if predicate(source.Current()) {
			r++
		}
	}
	return r, nil
}

// CountPredMust is like CountPred but panics in case of error.
func CountPredMust[Source any](source Enumerator[Source], predicate func(Source) bool) int {
	r, err := CountPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
