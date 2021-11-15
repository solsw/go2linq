//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 7 - Count and LongCount
// https://codeblog.jonskeet.uk/2010/12/26/reimplementing-linq-to-objects-part-7-count-and-longcount/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.count
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.longcount

// Count returns the number of elements in a sequence.
// Count panics if 'source' is nil.
func Count[Source any](source Enumerator[Source]) int {
	if source == nil {
		panic(ErrNilSource)
	}
	if c, ok := source.(Counter); ok {
		return c.Count()
	}
	r := 0
	for source.MoveNext() {
		r++
	}
	return r
}

// CountErr is like Count but returns an error instead of panicking.
func CountErr[Source any](source Enumerator[Source]) (res int, err error) {
	defer func() {
		catchPanic[int](recover(), &res, &err)
	}()
	return Count(source), nil
}

// CountPred returns a number that represents how many elements in the specified sequence satisfy a condition.
// CountPred panics if 'source' or 'predicate' is nil.
func CountPred[Source any](source Enumerator[Source], predicate func(Source) bool) int {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	r := 0
	for source.MoveNext() {
		if predicate(source.Current()) {
			r++
		}
	}
	return r
}

// CountPredErr is like CountPred but returns an error instead of panicking.
func CountPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res int, err error) {
	defer func() {
		catchPanic[int](recover(), &res, &err)
	}()
	return CountPred(source, predicate), nil
}
