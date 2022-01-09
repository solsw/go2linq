//go:build go1.18

package go2linq

// OnFunc is an Enumerator implementation based on fields-functions.
type OnFunc[T any] struct {
	mvNxt func() bool
	crrnt func() T
	rst   func()
}

// NewOnFunc creates a new OnFunc based on the provided functions.
func NewOnFunc[T any](mvNxt func() bool, crrnt func() T, rst func()) OnFunc[T] {
	return OnFunc[T]{mvNxt, crrnt, rst}
}

// NewOnFuncEn creates a new Enumerator based on the corresponding OnFunc.
func NewOnFuncEn[T any](mvNxt func() bool, crrnt func() T, rst func()) Enumerator[T] {
	return NewOnFunc(mvNxt, crrnt, rst)
}

// MoveNext implements the Enumerator.MoveNext method.
func (en OnFunc[T]) MoveNext() bool {
	if en.mvNxt == nil {
		return false
	}
	return en.mvNxt()
}

// Current implements the Enumerator.Current method.
func (en OnFunc[T]) Current() T {
	if en.crrnt == nil {
		return Default[T]()
	}
	return en.crrnt()
}

// Reset implements the Enumerator.Reset method.
func (en OnFunc[T]) Reset() {
	if en.rst == nil {
		return
	}
	en.rst()
}
