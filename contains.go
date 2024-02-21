package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [Contains] determines whether a sequence contains a specified element using [generichelper.DeepEqual].
//
// [Contains]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func Contains[Source any](source iter.Seq[Source], value Source) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	return ContainsEq(source, value, generichelper.DeepEqual[Source])
}

// [ContainsEq] determines whether a sequence contains a specified element using a specified 'equal'.
//
// [ContainsEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.contains
func ContainsEq[Source any](source iter.Seq[Source], value Source, equal func(Source, Source) bool) (bool, error) {
	if source == nil {
		return false, ErrNilSource
	}
	if equal == nil {
		return false, ErrNilEqual
	}
	for s := range source {
		if equal(s, value) {
			return true, nil
		}
	}
	return false, nil
}
