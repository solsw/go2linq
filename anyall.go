//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 10 - Any and All
// https://codeblog.jonskeet.uk/2010/12/28/reimplementing-linq-to-objects-part-10-any-and-all/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all

// Any determines whether a sequence contains any elements.
func Any[Source any](source Enumerable[Source]) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if c, ok := source.(Counter); ok {
		return c.Count() > 0, nil
	}
	return source.GetEnumerator().MoveNext(), nil
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

// All determines whether all elements of a sequence satisfy a condition.
func All[Source any](source Enumerable[Source], predicate func(Source) bool) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if predicate == nil {
		return false, ErrNilPredicate
	}
	enr := source.GetEnumerator()
	for enr.MoveNext() {
		if !predicate(enr.Current()) {
			return false, nil
		}
	}
	return true, nil
}

// AllMust is like All but panics in case of error.
func AllMust[Source any](source Enumerable[Source], predicate func(Source) bool) bool {
	r, err := All(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
