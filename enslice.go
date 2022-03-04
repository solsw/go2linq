//go:build go1.18

package go2linq

// EnSlice is an Enumerable implementation based on a slice of T.
type EnSlice[T any] struct {
	sl []T
}

// NewEnSlice creates a new EnSlice with the specified contents.
func NewEnSlice[T any](slice ...T) Enumerable[T] {
	return &EnSlice[T]{sl: slice}
}

// GetEnumerator implements the Enumerable interface.
func (en *EnSlice[T]) GetEnumerator() Enumerator[T] {
	return newEnrSlice(en.sl...)
}

// Count implements the Counter interface.
func (en *EnSlice[T]) Count() int {
	return en.GetEnumerator().(*enrSlice[T]).Count()
}

// Item implements the Itemer interface.
func (en *EnSlice[T]) Item(i int) T {
	return en.GetEnumerator().(*enrSlice[T]).Item(i)
}

// Slice implements the Slicer interface.
func (en *EnSlice[T]) Slice() []T {
	return en.GetEnumerator().(*enrSlice[T]).Slice()
}
