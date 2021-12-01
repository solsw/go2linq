//go:build go1.18

package go2linq

// OnChan is an Enumerator implementation based on channel.
type OnChan[T any] struct {
	chn   <-chan T
	crrnt T
}

// NewOnChan creates a new OnChan based on the provided functions.
func NewOnChan[T any](chn <-chan T) Enumerator[T] {
	return &OnChan[T]{chn: chn}
}

// MoveNext implements the Enumerator.MoveNext method.
func (en *OnChan[T]) MoveNext() bool {
	if en.chn == nil {
		return false
	}
	var open bool
	en.crrnt, open = <-en.chn
	return open
}

// Current implements the Enumerator.Current method.
func (en *OnChan[T]) Current() T {
	return en.crrnt
}

// Reset implements the Enumerator.Reset method.
//
// OnChan.Reset method does nothing. Hence OnChan cannot be used in functions
// that require an Enumerator with real Reset method (see ConcatSelf and the like).
func (en *OnChan[T]) Reset() {}
