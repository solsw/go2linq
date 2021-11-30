//go:build go1.18

package go2linq

// https://docs.microsoft.com/dotnet/api/system.linq.enumerable.trygetnonenumeratedcount

// TryGetNonEnumeratedCount attempts to determine the number of elements in a sequence without forcing an enumeration.
func TryGetNonEnumeratedCount[Source any](source Enumerator[Source], count *int) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if c, ok := source.(Counter); ok {
		*count = c.Count()
		return true, nil
	}
	return false, nil
}

// TryGetNonEnumeratedCountMust is like TryGetNonEnumeratedCount but panics in case of error.
func TryGetNonEnumeratedCountMust[Source any](source Enumerator[Source], count *int) bool {
	r, err := TryGetNonEnumeratedCount(source, count)
	if err != nil {
		panic(err)
	}
	return r
}
