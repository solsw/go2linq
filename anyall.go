//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 10 – Any and All
// https://codeblog.jonskeet.uk/2010/12/28/reimplementing-linq-to-objects-part-10-any-and-all/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.any
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all

// Any determines whether a sequence contains any elements.
// Any panics if 'source' is nil.
func Any[Source any](source Enumerator[Source]) bool {
	if source == nil {
		panic(ErrNilSource)
	}
	if c, ok := source.(Counter); ok {
		return c.Count() > 0
	}
	return source.MoveNext()
}

// AnyErr is like Any but returns an error instead of panicking.
func AnyErr[Source any](source Enumerator[Source]) (res bool, err error) {
	defer func() {
		catchPanic[bool](recover(), &res, &err)
	}()
	return Any(source), nil
}

// AnyPred determines whether any element of a sequence satisfies a condition.
// AnyPred panics if 'source' or 'predicate' is nil.
func AnyPred[Source any](source Enumerator[Source], predicate func(Source) bool) bool {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	for source.MoveNext() {
		if predicate(source.Current()) {
			return true
		}
	}
	return false
}

// AnyPredErr is like AnyPred but returns an error instead of panicking.
func AnyPredErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res bool, err error) {
	defer func() {
		catchPanic[bool](recover(), &res, &err)
	}()
	return AnyPred(source, predicate), nil
}

// All determines whether all elements of a sequence satisfy a condition.
// All panics if 'source' or 'predicate' is nil.
func All[Source any](source Enumerator[Source], predicate func(Source) bool) bool {
	if source == nil {
		panic(ErrNilSource)
	}
	if predicate == nil {
		panic(ErrNilPredicate)
	}
	for source.MoveNext() {
		if !predicate(source.Current()) {
			return false
		}
	}
	return true
}

// AllErr is like All but returns an error instead of panicking.
func AllErr[Source any](source Enumerator[Source], predicate func(Source) bool) (res bool, err error) {
	defer func() {
		catchPanic[bool](recover(), &res, &err)
	}()
	return All(source, predicate), nil
}
