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
