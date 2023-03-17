package go2linq

// EnFunc is an [Enumerable] implementation
// based on functions corresponding to the [Enumerable]'s methods.
type EnFunc[T any] struct {
	mvNxt func() bool
	crrnt func() T
	rst   func()
}

// NewEnFunc creates a new [EnFunc] with a specified functions as the [Enumerable]'s methods.
func NewEnFunc[T any](mvNxt func() bool, crrnt func() T, rst func()) Enumerable[T] {
	return &EnFunc[T]{mvNxt: mvNxt, crrnt: crrnt, rst: rst}
}

// GetEnumerator implements the [Enumerable] interface.
func (en *EnFunc[T]) GetEnumerator() Enumerator[T] {
	return enrFunc[T]{
		mvNxt: en.mvNxt,
		crrnt: en.crrnt,
		rst:   en.rst,
	}
}
