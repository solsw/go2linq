//go:build go1.18

package go2linq

// OnChan is an Enumerator implementation based on a channel.
type OnChan[T any] struct {
	chn   <-chan T
	crrnt T
}

// NewOnChan creates a new OnChan based on the provided channel.
func NewOnChan[T any](ch <-chan T) *OnChan[T] {
	return &OnChan[T]{chn: ch}
}

// NewOnChanEn creates a new Enumerator based on the corresponding OnChan.
func NewOnChanEn[T any](ch <-chan T) Enumerator[T] {
	return NewOnChan(ch)
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
// that require an Enumerator with a real Reset method (see ConcatSelf and the like).
func (*OnChan[T]) Reset() {}
