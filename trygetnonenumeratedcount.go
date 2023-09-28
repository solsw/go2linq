package go2linq

import (
	"github.com/solsw/errorhelper"
)

// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.trygetnonenumeratedcount

// [TryGetNonEnumeratedCount] attempts to determine the number of elements in a sequence without forcing an enumeration.
//
// [TryGetNonEnumeratedCount]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.trygetnonenumeratedcount
func TryGetNonEnumeratedCount[Source any](source Enumerable[Source], count *int) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if counter, ok := source.(Counter); ok {
		*count = counter.Count()
		return true, nil
	}
	return false, nil
}

// TryGetNonEnumeratedCountMust is like [TryGetNonEnumeratedCount] but panics in case of error.
func TryGetNonEnumeratedCountMust[Source any](source Enumerable[Source], count *int) bool {
	return errorhelper.Must(TryGetNonEnumeratedCount(source, count))
}
