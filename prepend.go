package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Prepend] adds a value to the beginning of the sequence.
//
// [Prepend]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.prepend
func Prepend[Source any](source iter.Seq[Source], element Source) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	repeat1, _ := Repeat(element, 1)
	return Concat(repeat1, source)
}
