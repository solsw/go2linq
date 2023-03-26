package go2linq

// enrSlice is an [Enumerator] implementation based on a [slice] of T.
//
// [slice]: https://go.dev/ref/spec#Slice_types
type enrSlice[T any] struct {
	// indx-1 - index of the current element in slc
	//     ^^ because initially enumerator is positioned before the first element in the collection
	indx int
	// elements of the enrSlice instance
	slc []T
}

// newEnrSlice creates a new [enrSlice] with a specified contents.
func newEnrSlice[T any](ee ...T) *enrSlice[T] {
	var enr enrSlice[T]
	enr.slc = make([]T, len(ee))
	if len(ee) > 0 {
		copy(enr.slc, ee)
	}
	return &enr
}

// enrSlice_moveNext is used for testing
func enrSlice_moveNext[T any](enr *enrSlice[T]) bool {
	if enr.indx >= len(enr.slc) {
		return false
	}
	enr.indx++
	return true
}

// MoveNext implements the [Enumerator.MoveNext] method.
func (enr *enrSlice[T]) MoveNext() bool {
	return enrSlice_moveNext(enr)
}

// enrSlice_current is used for testing
func enrSlice_current[T any](enr *enrSlice[T]) T {
	return enr.slc[enr.indx-1]
}

// Current implements the [Enumerator.Current] method.
func (enr *enrSlice[T]) Current() T {
	return enrSlice_current(enr)
}

// Reset implements the [Enumerator.Reset] method.
func (enr *enrSlice[T]) Reset() {
	enr.indx = 0
}
