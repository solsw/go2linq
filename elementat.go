package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [ElementAt] returns the element at a specified index in a sequence.
//
// [ElementAt]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.elementat
func ElementAt[Source any](source iter.Seq[Source], index int) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	if index < 0 {
		return generichelper.ZeroValue[Source](), ErrIndexOutOfRange
	}
	i := 0
	for s := range source {
		if i == index {
			return s, nil
		}
		i++
	}
	return generichelper.ZeroValue[Source](), ErrIndexOutOfRange
}

// [ElementAtOrDefault] returns the element at a specified index in a sequence or a [zero value] if the index is out of range.
//
// [ElementAtOrDefault]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.elementatordefault
// [zero value]: https://go.dev/ref/spec#The_zero_value
func ElementAtOrDefault[Source any](source iter.Seq[Source], index int) (Source, error) {
	if source == nil {
		return generichelper.ZeroValue[Source](), ErrNilSource
	}
	s, err := ElementAt(source, index)
	if err == ErrIndexOutOfRange {
		return generichelper.ZeroValue[Source](), nil
	}
	return s, nil
}
