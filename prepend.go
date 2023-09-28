package go2linq

import (
	"github.com/solsw/errorhelper"
)

// https://learn.microsoft.com/dotnet/api/system.linq.enumerable.prepend

// [Prepend] adds a value to the beginning of the sequence.
//
// [Prepend]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.prepend
func Prepend[Source any](source Enumerable[Source], element Source) (Enumerable[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return Concat(RepeatMust(element, 1), source)
}

// PrependMust is like [Prepend] but panics in case of error.
func PrependMust[Source any](source Enumerable[Source], element Source) Enumerable[Source] {
	return errorhelper.Must(Prepend(source, element))
}
