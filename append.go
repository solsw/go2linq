package go2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// [Append] appends a value to the end of the sequence.
//
// [Append]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.append
func Append[Source any](source iter.Seq[Source], element Source) (iter.Seq[Source], error) {
	if source == nil {
		return nil, errorhelper.CallerError(ErrNilSource)
	}
	repeat1, _ := Repeat(element, 1)
	r, err := Concat(source, repeat1)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
