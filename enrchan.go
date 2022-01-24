//go:build go1.18

package go2linq

// enrChan is an Enumerator implementation based on a channel.
type enrChan[T any] struct {
	chn   <-chan T
	crrnt T
}

// newEnrChan creates a new enrChan based on the provided channel.
func newEnrChan[T any](ch <-chan T) *enrChan[T] {
	return &enrChan[T]{chn: ch}
}

// MoveNext implements the Enumerator.MoveNext method.
func (en *enrChan[T]) MoveNext() bool {
	if en.chn == nil {
		return false
	}
	var open bool
	en.crrnt, open = <-en.chn
	return open
}

// Current implements the Enumerator.Current method.
func (en *enrChan[T]) Current() T {
	return en.crrnt
}

// Reset implements the Enumerator.Reset method.
//
// enrChan.Reset method does nothing. Hence enrChan cannot be used in functions that require an Enumerator with a real Reset method.
func (*enrChan[T]) Reset() {}
