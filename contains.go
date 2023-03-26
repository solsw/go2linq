package go2linq

import (
	"github.com/solsw/collate"
	"github.com/solsw/generichelper"
)

// Reimplementing LINQ to Objects: Part 32 – Contains
// https://codeblog.jonskeet.uk/2011/01/12/reimplementing-linq-to-objects-part-32-contains/
// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains

// [Contains] determines whether a sequence contains a specified element using [collate.DeepEqualer].
//
// [Contains]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func Contains[Source any](source Enumerable[Source], value Source) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	return ContainsEq(source, value, nil)
}

// ContainsMust is like [Contains] but panics in case of error.
func ContainsMust[Source any](source Enumerable[Source], value Source) bool {
	return generichelper.Must(Contains(source, value))
}

// [ContainsEq] determines whether a sequence contains a specified element using a specified equaler.
// If 'equaler' is nil [collate.DeepEqualer] is used.
//
// [ContainsEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func ContainsEq[Source any](source Enumerable[Source], value Source, equaler collate.Equaler[Source]) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if equaler == nil {
		equaler = collate.DeepEqualer[Source]{}
	}
	enr := source.GetEnumerator()
	for enr.MoveNext() {
		if equaler.Equal(value, enr.Current()) {
			return true, nil
		}
	}
	return false, nil
}

// ContainsEqMust is like [ContainsEq] but panics in case of error.
func ContainsEqMust[Source any](source Enumerable[Source], value Source, equaler collate.Equaler[Source]) bool {
	return generichelper.Must(ContainsEq(source, value, equaler))
}
