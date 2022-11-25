package go2linq

// enrFunc is an Enumerator implementation based on fields-functions.
type enrFunc[T any] struct {
	mvNxt func() bool
	crrnt func() T
	rst   func()
}

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
