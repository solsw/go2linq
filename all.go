//go:build go1.18

package go2linq

// Reimplementing LINQ to Objects: Part 10 - Any and All
// https://codeblog.jonskeet.uk/2010/12/28/reimplementing-linq-to-objects-part-10-any-and-all/
// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all

// All determines whether all elements of a sequence satisfy a condition.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.all)
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

// AllMust is like All but panics in case of an error.
func AllMust[Source any](source Enumerable[Source], predicate func(Source) bool) bool {
	r, err := All(source, predicate)
	if err != nil {
		panic(err)
	}
	return r
}
