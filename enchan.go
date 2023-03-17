package go2linq

// EnChan is an [Enumerable] implementation based on a [channel].
//
// [channel]: https://go.dev/ref/spec#Channel_types
type EnChan[T any] <-chan T

// NewEnChan creates a new [EnChan] with a specified channel as contents.
func NewEnChan[T any](ch <-chan T) Enumerable[T] {
	en := EnChan[T](ch)
	return &en
}

// GetEnumerator implements the [Enumerable] interface.
func (en *EnChan[T]) GetEnumerator() Enumerator[T] {
	return newEnrChan(*en)
}
