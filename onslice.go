//go:build go1.18

package go2linq

// OnSlice is an Enumerator implementation based on a slice of T.
type OnSlice[T any] struct {
	// indx-1 - index of the current element in elel
	//     ^^ because initially enumerator is positioned before the first element in the collection
	indx int
	// elements of the OnSlice instance
	elel []T
}

// NewOnSlice creates a new OnSlice with the specified contents.
func NewOnSlice[T any](ee ...T) *OnSlice[T] {
	var onSlice OnSlice[T]
	onSlice.elel = make([]T, len(ee))
	if len(ee) > 0 {
		copy(onSlice.elel, ee)
	}
	return &onSlice
}

// NewOnSliceEn creates a new Enumerator based on the corresponding OnSlice.
func NewOnSliceEn[T any](ee ...T) Enumerator[T] {
	return NewOnSlice(ee...)
}

// MoveNext implements the Enumerator.MoveNext method.
func (en *OnSlice[T]) MoveNext() bool {
	if en.indx >= len(en.elel) {
		return false
	}
	en.indx++
	return true
}

// Current implements the Enumerator.Current method.
func (en *OnSlice[T]) Current() T {
	return en.Item(en.indx - 1)
}

// Reset implements the Enumerator.Reset method.
func (en *OnSlice[T]) Reset() {
	en.indx = 0
}

// Count implements the Counter interface.
func (en *OnSlice[T]) Count() int {
	return len(en.elel)
}

// Item implements the Itemer interface.
func (en *OnSlice[T]) Item(i int) T {
	// https://docs.microsoft.com/dotnet/api/system.collections.ienumerator.current#remarks
	// https://docs.microsoft.com/dotnet/api/system.collections.generic.list-1.item#exceptions
	if !(0 <= i && i < len(en.elel)) {
		var t0 T
		return t0
	}
	return en.elel[i]
}

// Slice implements the Slicer interface.
func (en *OnSlice[T]) Slice() []T {
	return en.elel
}
