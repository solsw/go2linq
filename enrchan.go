package go2linq

// enrChan is an [Enumerator] implementation based on a [channel].
//
// [channel]: https://go.dev/ref/spec#Channel_types
type enrChan[T any] struct {
	chn   <-chan T
	crrnt T
}

// newEnrChan creates a new [enrChan] based on the provided channel.
func newEnrChan[T any](ch <-chan T) *enrChan[T] {
	return &enrChan[T]{chn: ch}
}

// enrChan_moveNext is used for testing
func enrChan_moveNext[T any](enr *enrChan[T]) bool {
	if enr.chn == nil {
		return false
	}
	var open bool
	enr.crrnt, open = <-enr.chn
	return open
}

// MoveNext implements the [Enumerator.MoveNext] method.
func (enr *enrChan[T]) MoveNext() bool {
	return enrChan_moveNext(enr)
}

// enrChan_current is used for testing
func enrChan_current[T any](enr *enrChan[T]) T {
	return enr.crrnt
}

// Current implements the [Enumerator.Current] method.
func (enr *enrChan[T]) Current() T {
	return enrChan_current(enr)
}

// Reset implements the [Enumerator.Reset] method.
// The method panics.
func (*enrChan[T]) Reset() {
	panic("Reset() not supported")
}
