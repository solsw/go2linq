package go2linq

// EnChan is an [Enumerable] implementation based on a [channel].
//
// [channel]: https://go.dev/ref/spec#Channel_types
type EnChan[T any] <-chan T

// NewEnChan creates a new [EnChan] with a specified channel as contents.
func NewEnChan[T any](ch <-chan T) *EnChan[T] {
	en := EnChan[T](ch)
	return &en
}

// NewEnChanEn creates a new [EnChan] with a specified channel as contents
// and returns it as [Enumerable].
func NewEnChanEn[T any](ch <-chan T) Enumerable[T] {
	return NewEnChan[T](ch)
}

// GetEnumerator implements the [Enumerable] interface.
func (en *EnChan[T]) GetEnumerator() Enumerator[T] {
	return newEnrChan[T](*en)
}
