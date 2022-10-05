//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 10 - Any and All
// https://codeblog.jonskeet.uk/2010/12/28/reimplementing-linq-to-objects-part-10-any-and-all/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any

// Any determines whether a sequence contains any elements.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any)
func Any[Source any](source Enumerable[Source]) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	enr := source.GetEnumerator()
	if c, ok := enr.(Counter); ok {
		return c.Count() > 0, nil
	}
	return enr.MoveNext(), nil
}

// AnyMust is like Any but panics in case of error.
func AnyMust[Source any](source Enumerable[Source]) bool {
	r, err := Any(source)
	if err != nil {
		panic(err)
	}
	return r
}

// AnyPred determines whether any element of a sequence satisfies a condition.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any)
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

// AnyPredMust is like AnyPred but panics in case of error.
func AnyPredMust[Source any](source Enumerable[Source], predicate func(Source) bool) bool {
	r, err := AnyPred(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
