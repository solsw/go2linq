package go2linq

import (
	"iter"
)

// [ToSlice] creates a [slice] from a sequence.
//
// [slice]: https://go.dev/ref/spec#Slice_types
// [ToSlice]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.tolist
func ToSlice[Source any](source iter.Seq[Source]) ([]Source, error) {
	if source == nil {
		return nil, ErrNilSource
	}
	var ss []Source
	for s := range source {
		ss = append(ss, s)
	}
	return ss, nil
}
