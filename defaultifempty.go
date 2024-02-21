package go2linq

import (
	"iter"

	"github.com/solsw/generichelper"
)

// [DefaultIfEmpty] returns the elements of a specified sequence
// or the type parameter's [zero value] in a singleton collection if the sequence is empty.
//
// [DefaultIfEmpty]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
// [zero value]: https://go.dev/ref/spec#The_zero_value
func DefaultIfEmpty[Source any](source iter.Seq[Source]) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return DefaultIfEmptyDef(source, generichelper.ZeroValue[Source]())
}

// [DefaultIfEmptyDef] returns the elements of a specified sequence
// or a specified value in a singleton collection if the sequence is empty.
//
// [DefaultIfEmptyDef]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.defaultifempty
func DefaultIfEmptyDef[Source any](source iter.Seq[Source], defaultValue Source) (iter.Seq[Source], error) {
	if source == nil {
		return nil, ErrNilSource
	}
	return func(yield func(Source) bool) {
			empty := true
			for s := range source {
				empty = false
				if !yield(s) {
					return
				}
			}
			if empty {
				if !yield(defaultValue) {
					return
				}
			}
		},
		nil
}
