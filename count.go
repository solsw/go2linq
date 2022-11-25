package go2linq

// Reimplementing LINQ to Objects: Part 7 - Count and LongCount
// https://codeblog.jonskeet.uk/2010/12/26/reimplementing-linq-to-objects-part-7-count-and-longcount/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount

// Count returns the number of elements in a sequence.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count,
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount)
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

// CountMust is like Count but panics in case of error.
func CountMust[Source any](source Enumerable[Source]) int {
	r, err := Count(source)
	if err != nil {
		panic(err)
	}
	return r
}

// CountPred returns a number that represents how many elements in the specified sequence satisfy a condition.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count,
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount)
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

// CountPredMust is like CountPred but panics in case of error.
func CountPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) int {
	r, err := CountPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
