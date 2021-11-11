//go:build go1.18

package go2linq

// OnFunc is an Enumerator implementation based on fields-functions.
type OnFunc[T any] struct {
	MvNxt func() bool
	Crrnt func() T
	Rst func()
}

// MoveNext implements the Enumerator.MoveNext method.
func (en OnFunc[T]) MoveNext() bool {
	if en.MvNxt == nil {
		return false
	}
	return en.MvNxt()
}

// Current implements the Enumerator.Current method.
func (en OnFunc[T]) Current() T {
	if en.Crrnt == nil {
		var t0 T
		return t0
	}
	return en.Crrnt()
}

// Reset implements the Enumerator.Reset method.
func (en OnFunc[T]) Reset() {
	if en.Rst == nil {
		return
	}
	en.Rst()
}
