package go2linq

import (
	"iter"
)

// [Append] appends a value to the end of the sequence.
//
// [Append]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.append
func Append[Source any](source iter.Seq[Source], element Source) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	repeat1, _ := Repeat(element, 1)
	return Concat(source, repeat1)
}
