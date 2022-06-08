//go:build go1.18

package go2linq

// EnSlice is an Enumerable implementation based on a slice of T.
type EnSlice[T any] struct {
	sl []T
}

// NewEnSlice creates a new EnSlice with the specified slice as contents.
func NewEnSlice[T any](slice ...T) Enumerable[T] {
	return &EnSlice[T]{sl: slice}
}

// GetEnumerator implements the Enumerable interface.
func (en *EnSlice[T]) GetEnumerator() Enumerator[T] {
	return newEnrSlice(en.sl...)
}

func (en *EnSlice[T]) enrSlice() *enrSlice[T] {
	return en.GetEnumerator().(*enrSlice[T])
}

// Count implements the Counter interface.
func (en *EnSlice[T]) Count() int {
	return len(en.enrSlice().elel)
}

// Item implements the Itemer interface.
func (en *EnSlice[T]) Item(i int) T {
	return en.enrSlice().item(i)
}

// Slice implements the Slicer interface.
func (en *EnSlice[T]) Slice() []T {
	return en.enrSlice().elel
}
