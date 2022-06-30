//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.trygetnonenumeratedcount

// TryGetNonEnumeratedCount attempts to determine the number of elements in a sequence without forcing an enumeration.
// (https://docs.microsoft.com/dotnet/api/system.linq.enumerable.trygetnonenumeratedcount)
func TryGetNonEnumeratedCount[Source any](source Enumerable[Source], count *int) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if counter, ok := source.GetEnumerator().(Counter); ok {
		*count = counter.Count()
		return true, nil
	}
	return false, nil
}

// TryGetNonEnumeratedCountMust is like TryGetNonEnumeratedCount but panics in case of an error.
func TryGetNonEnumeratedCountMust[Source any](source Enumerable[Source], count *int) bool {
	r, err := TryGetNonEnumeratedCount(source, count)
	if err != nil {
		panic(err)
	}
	return r
}
