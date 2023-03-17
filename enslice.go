package go2linq

import (
	"github.com/solsw/generichelper"
)

// EnSlice is an [Enumerable] implementation based on a [slice] of T.
//
// [slice]: https://go.dev/ref/spec#Slice_types
type EnSlice[T any] []T

// NewEnSlice creates a new [EnSlice] with a specified slice as contents.
func NewEnSlice[T any](slice ...T) Enumerable[T] {
	en := EnSlice[T](slice)
	return &en
}

// GetEnumerator implements the [Enumerable] interface.
func (en *EnSlice[T]) GetEnumerator() Enumerator[T] {
	return newEnrSlice(*en...)
}

// Count implements the [Counter] interface.
func (en *EnSlice[T]) Count() int {
	return len(*en)
}

// Item implements the [Itemer] interface.
func (en *EnSlice[T]) Item(i int) T {
	// https://learn.microsoft.com/dotnet/api/system.collections.ienumerator.current#remarks
	// https://learn.microsoft.com/dotnet/api/system.collections.generic.list-1.item#exceptions
	// if i < 0 {
	// 	panic("Enumeration has not started. Call MoveNext.")
	// }
	// if i >= len(*en) {
	// 	panic("Enumeration already finished.")
	// }

	if !(0 <= i && i < len(*en)) {
		return generichelper.ZeroValue[T]()
	}
	return (*en)[i]
}

// Slice implements the [Slicer] interface.
func (en *EnSlice[T]) Slice() []T {
	return *en
}
