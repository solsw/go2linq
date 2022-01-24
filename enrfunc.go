//go:build go1.18

package go2linq

// enrFunc is an Enumerator implementation based on fields-functions.
type enrFunc[T any] struct {
	mvNxt func() bool
	crrnt func() T
	rst   func()
}

// // newEnrFunc creates a new enrFunc based on the provided functions.
// func newEnrFunc[T any](mvNxt func() bool, crrnt func() T, rst func()) enrFunc[T] {
// 	return enrFunc[T]{mvNxt, crrnt, rst}
// }

// MoveNext implements the Enumerator.MoveNext method.
func (enr enrFunc[T]) MoveNext() bool {
	if enr.mvNxt == nil {
		return false
	}
	return enr.mvNxt()
}

// Current implements the Enumerator.Current method.
func (enr enrFunc[T]) Current() T {
	if enr.crrnt == nil {
		return ZeroValue[T]()
	}
	return enr.crrnt()
}

// Reset implements the Enumerator.Reset method.
func (enr enrFunc[T]) Reset() {
	if enr.rst == nil {
		return
	}
	enr.rst()
}
